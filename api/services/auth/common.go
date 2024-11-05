package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sibiraj-s/unique-names-generator"
	"github.com/usegranthq/backend/constants"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/ent/verification"
	"github.com/usegranthq/backend/external"
	"github.com/usegranthq/backend/utils"
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

	_, err = db.Client.Verification.Create().
		SetUserID(user.ID).
		SetAttemptID(attemptId).
		SetType(verification.TypeSIGNUP).
		SetCode(code).
		SetExpiresAt(attemptExpiry).
		Save(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	utils.SafeRoutine(func() {
		if err := external.Postman.SendLoginEmail(c, user.Email, code); err != nil {
			utils.HttpError.InternalServerError(c)
			return
		}
	})

	cookieExpiry := int(attemptExpiry.Sub(time.Now()).Seconds())
	utils.Http.SetCookie(c, constants.VerifyCookie, token, cookieExpiry)
}

func CreateUserSession(c *gin.Context, user *ent.User) {
	sessionId, err := utils.GenerateRandom(32)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	sessionExpiry := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"sub": sessionId,
		"exp": sessionExpiry.Unix(),
	}

	sessionCookie, err := utils.Jwt.SignToken(claims)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	if _, err := db.Client.UserSession.Create().
		SetUser(user).
		SetToken(sessionId).
		SetExpiresAt(sessionExpiry).
		Save(c); err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	_, err = db.Client.User.UpdateOneID(user.ID).SetLastLogin(time.Now()).Save(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	cookieExpiry := int(sessionExpiry.Sub(time.Now()).Seconds())
	utils.Http.SetCookie(c, constants.AuthCookie, sessionCookie, cookieExpiry)
	utils.Http.DeleteCookie(c, constants.VerifyCookie)
}

func VerifyCaptcha(c *gin.Context, token string) error {
	turnstileResponse, err := external.Turnstile.Verify(token)

	if err != nil {
		utils.HttpError.InternalServerError(c, "Failed to verify captcha")
		return err
	}

	if !turnstileResponse.Success {
		errorCodes := strings.Join(turnstileResponse.ErrorCodes, ", ")
		errorMessage := fmt.Sprintf("Captcha verification failed: %s", errorCodes)
		utils.HttpError.BadRequest(c, errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}

func DoEmailSignup(c *gin.Context, email string) error {
	if utils.Emails.IsDisposableEmail(email) {
		utils.HttpError.BadRequest(c, "Disposable emails are not allowed. Note: when you delete your account, we delete everything without a trace.")
		return errors.New("disposable email not allowed")
	}

	_, err := db.Client.User.
		Query().
		Where(user.Email(email)).
		Only(c)

	if err == nil {
		utils.HttpError.Conflict(c, "User already exists")
		return err
	} else if !ent.IsNotFound(err) {
		utils.HttpError.InternalServerError(c)
		return err
	}

	newUser, err := db.Client.User.
		Create().
		SetEmail(email).
		SetProvider(user.ProviderEMAIL).
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return err
	}

	StartUserVerification(c, newUser)
	return nil
}

func DoOauthSignup(c *gin.Context, email string, provider user.Provider) error {
	err := db.Client.User.
		Create().
		SetEmail(email).
		SetProvider(provider).
		SetVerifiedAt(time.Now()).
		OnConflictColumns(user.FieldEmail).
		Ignore().
		Exec(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return err
	}

	user, err := db.Client.User.
		Query().
		Where(user.Email(email)).
		Only(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return err
	}

	CreateUserSession(c, user)
	return nil
}
