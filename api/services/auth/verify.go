package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/userverification"
	"github.com/usegranthq/backend/external"
	"github.com/usegranthq/backend/utils"
)

const maxVerificationAttempts = 5

const (
	VerificationSuccess            = 0
	VerificationExpired            = 1
	VerificationMaxAttemptsReached = 2
	VerificationInvalidCode        = 3
	VerificationUnknown            = 4
)

func deleteVerificationCode(c *gin.Context, attemptID uuid.UUID) {
	_, _ = db.Client.UserVerification.Delete().
		Where(userverification.AttemptID(attemptID)).
		Exec(c)
}

func verifyToken(c *gin.Context, inputToken string, code string) (*ent.User, int) {
	token, err := utils.Jwt.DecodeToken(inputToken)
	if err != nil {
		return nil, VerificationInvalidCode
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, VerificationInvalidCode
	}

	attemptIDStr, ok := claims["attempt_id"].(string)
	if !ok {
		return nil, VerificationInvalidCode
	}

	attemptID, err := uuid.Parse(attemptIDStr)
	if err != nil {
		return nil, VerificationInvalidCode
	}

	userVerification, err := db.Client.UserVerification.Query().
		Where(userverification.AttemptID(attemptID)).
		Only(c)
	if err != nil {
		return nil, VerificationInvalidCode
	}

	if userVerification.Attempts >= maxVerificationAttempts {
		// delete the verification code
		deleteVerificationCode(c, attemptID)
		return nil, VerificationMaxAttemptsReached
	}

	_, err = userVerification.Update().
		SetAttempts(userVerification.Attempts + 1).
		Save(c)
	if err != nil {
		return nil, VerificationMaxAttemptsReached
	}

	if userVerification.Code != code {
		return nil, VerificationInvalidCode
	}

	user, err := userVerification.QueryUser().Only(c)
	if err != nil {
		return nil, VerificationUnknown
	}

	// delete the verification code
	deleteVerificationCode(c, attemptID)

	return user, VerificationSuccess
}

type verifyRequest struct {
	Code string `json:"code" binding:"required"`
}

func Verify(c *gin.Context) {
	verifyCookie, err := c.Cookie("_ug_verify")
	if err != nil {
		utils.HttpError.Unauthorized(c)
		return
	}

	var req verifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	user, errCode := verifyToken(c, verifyCookie, req.Code)
	if errCode != VerificationSuccess {
		switch errCode {
		case VerificationExpired:
		case VerificationInvalidCode:
			utils.HttpError.BadRequest(c, "Invalid verification code")
		case VerificationMaxAttemptsReached:
			utils.HttpError.BadRequest(c, "Maximum verification attempts reached")
		default:
			utils.HttpError.InternalServerError(c)
		}
		return
	}

	if user.VerifiedAt.IsZero() {
		if _, err := user.Update().SetVerifiedAt(time.Now()).Save(c); err != nil {
			utils.HttpError.InternalServerError(c)
			return
		}
	}

	CreateUserSession(c, user)
	c.JSON(http.StatusOK, gin.H{})
}

type verifyOauthRequest struct {
	Code  string `json:"code" binding:"required"`
	State string `json:"state" binding:"required"`
}

func VerifyGithub(c *gin.Context) {
	var req verifyOauthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	if !utils.Hmac.VerifySecureHMACState(external.Github.HmacSecretKey, req.State) {
		utils.HttpError.BadRequest(c, "Invalid or expired token")
		return
	}

	primaryEmail, _, err := external.Github.GetGithubUser(req.Code)
	if err != nil {
		utils.HttpError.BadRequest(c, "Unable to fetch user details")
		return
	}

	if err := DoOauthSignup(c, primaryEmail); err != nil {
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func VerifyGoogle(c *gin.Context) {
	var req verifyOauthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	if !utils.Hmac.VerifySecureHMACState(external.Google.HmacSecretKey, req.State) {
		utils.HttpError.BadRequest(c, "Invalid or expired token")
		return
	}

	primaryEmail, _, err := external.Google.GetGoogleUser(req.Code)
	if err != nil {
		utils.HttpError.BadRequest(c, "Unable to fetch user details")
		return
	}

	if err := DoOauthSignup(c, primaryEmail); err != nil {
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}
