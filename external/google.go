package external

import (
	"context"

	"github.com/usegranthq/backend/config"
	"github.com/usegranthq/backend/utils"
	"golang.org/x/oauth2"
	gOauth2 "golang.org/x/oauth2/google"
	gApi "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type googleExternal struct {
	HmacSecretKey string
}

var Google = googleExternal{}
var googleOauthConfig *oauth2.Config

func init() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     config.Get("GOOGLE_OAUTH_CLIENT_ID"),     // Replace with your Google client ID
		ClientSecret: config.Get("GOOGLE_OAUTH_CLIENT_SECRET"), // Replace with your Google client secret
		Scopes: []string{ // Requesting both profile and email scopes
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		RedirectURL: config.Get("CLIENT_URL") + "/callbacks/google/oauth", // Callback URL
		Endpoint:    gOauth2.Endpoint,                                     // Google's OAuth 2.0 endpoint
	}
	Google.HmacSecretKey = config.Get("GOOGLE_HMAC_SECRET_KEY")
}

func (g *googleExternal) getGoogleClient(code string) (*gApi.Service, error) {
	// exchange the authorization code for an access token
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	// create a Google OAuth2 service with the access token
	service, err := gApi.NewService(
		context.Background(),
		option.WithTokenSource(googleOauthConfig.TokenSource(context.Background(), token)),
	)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (g *googleExternal) getGoogleUserInfo(service *gApi.Service) (string, string, error) {
	userinfo, err := service.Userinfo.Get().Do()
	if err != nil {
		return "", "", err
	}

	return userinfo.Name, userinfo.Email, nil
}

func (g *googleExternal) GenerateOauthUrl() (string, error) {
	state, err := utils.Hmac.GenerateHMACState(g.HmacSecretKey)
	if err != nil {
		return "", err
	}

	url := googleOauthConfig.AuthCodeURL(state)
	return url, nil
}

func (g *googleExternal) GetGoogleUser(code string) (string, string, error) {
	service, err := g.getGoogleClient(code)
	if err != nil {
		return "", "", err
	}

	userName, userEmail, err := g.getGoogleUserInfo(service)
	if err != nil {
		return "", "", err
	}

	return userEmail, userName, nil
}
