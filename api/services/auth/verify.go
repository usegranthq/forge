package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sibiraj-s/unique-names-generator"
	"github.com/usegranthq/backend/config"
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

func StartUserVerification(c *gin.Context, user *ent.User) {
	attemptExpiry := time.Now().Add(30 * time.Minute)
	attemptId := uuid.New()

	codeOptions := unique.Options{
		Dictionaries: [][]string{
			unique.Adjectives,
			unique.Animals,
			unique.Colors,
			unique.Countries,
			unique.StarWars,
			unique.Names,
		},
		Separator: &[]string{"-"}[0],
		Length:    3,
	}
	code := unique.New(codeOptions)

	claims := jwt.MapClaims{
		"email":      user.Email,
		"attempt_id": attemptId.String(),
		"exp":        attemptExpiry.Unix(),
	}

	token, err := utils.Jwt.SignToken(claims)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	_, err = db.Client.UserVerification.Create().
		SetUserID(user.ID).
		SetAttemptID(attemptId).
		SetCode(code).
		SetExpiresAt(attemptExpiry).
		Save(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	if err := external.Postman.SendLoginEmail(c, user.Email, code); err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	cookieExpiry := int(attemptExpiry.Sub(time.Now()).Seconds())
	secure := config.Get("NODE_ENV") == "production"
	c.SetCookie("verify", token, cookieExpiry, "/", "/", secure, true)

	c.JSON(http.StatusOK, gin.H{})
}

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
	verifyCookie, err := c.Cookie("verify")
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

	loginExpiry := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"sub": user.ID.String(),
		"exp": loginExpiry.Unix(),
	}

	token, err := utils.Jwt.SignToken(claims)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	if _, err := db.Client.UserSession.Create().
		SetUser(user).
		SetToken(token).
		SetExpiresAt(loginExpiry).
		Save(c); err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	cookieExpiry := int(loginExpiry.Sub(time.Now()).Seconds())
	secure := config.Get("NODE_ENV") == "production"
	c.SetCookie("auth", token, cookieExpiry, "/", "/", secure, true)

	c.JSON(http.StatusOK, gin.H{})
}
