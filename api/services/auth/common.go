package auth

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sibiraj-s/unique-names-generator"
	"github.com/usegranthq/backend/config"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/user"
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
	c.SetCookie("_ug_verify", token, cookieExpiry, "/", "/", secure, true)
}

func CreateUserSession(c *gin.Context, user *ent.User) {
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

	_, err = db.Client.User.UpdateOneID(user.ID).SetLastLogin(time.Now()).Save(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	cookieExpiry := int(loginExpiry.Sub(time.Now()).Seconds())
	secure := config.Get("NODE_ENV") == "production"
	c.SetCookie("_ug_auth", token, cookieExpiry, "/", "/", secure, true)
}

func DoSignup(c *gin.Context, email string) error {
	if utils.IsDisposableEmail(email) {
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
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return err
	}

	StartUserVerification(c, newUser)
	return nil
}

func DoOauthSignup(c *gin.Context, email string) error {
	err := db.Client.User.
		Create().
		SetEmail(email).
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
