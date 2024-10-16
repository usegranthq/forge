package settings

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/token"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/utils"
)

// expires_at is optional, if not provided, the token will never expire
// values can be "never", "1", "7", "30", "60", "90", "180", "360", "0"
type createTokenRequest struct {
	Name      string `json:"name" binding:"required"`
	ExpiresAt int    `json:"expires_at"`
}

func CreateToken(c *gin.Context) {
	user := c.MustGet("user").(*ent.User)

	var req createTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	var expiresAt *time.Time
	if req.ExpiresAt != 0 {
		parsedTime := time.Now().AddDate(0, 0, req.ExpiresAt)
		expiresAt = &parsedTime
	}

	apiToken, err := utils.GenerateToken("ug_api")
	if err != nil {
		utils.HttpError.InternalServerError(c, err.Error())
		return
	}

	tokenQuery := db.Client.Token.Create().
		SetName(req.Name).
		SetToken(apiToken).
		SetUserID(user.ID)

	if expiresAt != nil {
		tokenQuery.SetExpiresAt(*expiresAt)
	}

	token, err := tokenQuery.Save(c)
	if err != nil {
		utils.HttpError.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"name":         token.Name,
		"token":        token.Token,
		"last_used_at": token.LastUsedAt,
		"expires_at":   token.ExpiresAt,
		"created_at":   token.CreatedAt,
	})
}

type listTokensResponse struct {
	Name       string    `json:"name"`
	ID         uuid.UUID `json:"id"`
	ExpiresAt  time.Time `json:"expires_at"`
	LastUsedAt time.Time `json:"last_used_at,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ListTokens(c *gin.Context) {
	currentUser := c.MustGet("user").(*ent.User)

	tokens, err := db.Client.Token.Query().
		Where(token.HasUserWith(user.ID(currentUser.ID))).
		All(c)

	if err != nil {
		utils.HttpError.InternalServerError(c, err.Error())
		return
	}

	response := make([]listTokensResponse, len(tokens))
	for i, token := range tokens {
		response[i] = listTokensResponse{
			Name:       token.Name,
			ID:         token.ID,
			ExpiresAt:  token.ExpiresAt,
			LastUsedAt: token.LastUsedAt,
			CreatedAt:  token.CreatedAt,
			UpdatedAt:  token.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, response)
}

func DeleteToken(c *gin.Context) {
	currentUser := c.MustGet("user").(*ent.User)

	tokenID := c.Param("tokenID")
	if _, err := uuid.Parse(tokenID); err != nil {
		utils.HttpError.BadRequest(c, "invalid token id")
		return
	}

	err := db.Client.Token.
		DeleteOneID(uuid.MustParse(tokenID)).
		Where(token.HasUserWith(user.ID(currentUser.ID))).
		Exec(c)
	if err != nil {
		utils.HttpError.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
