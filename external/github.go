package external

import (
	"context"
	"strings"

	"github.com/google/go-github/v66/github"
	"github.com/usegranthq/backend/config"
	"github.com/usegranthq/backend/utils"
	"golang.org/x/oauth2"
	ghOauth2 "golang.org/x/oauth2/github"
)

type githubExternal struct {
	HmacSecretKey string
}

var Github = githubExternal{}
var githubOauthConfig *oauth2.Config

func init() {
	githubOauthConfig = &oauth2.Config{
		ClientID:     config.Get("GITHUB_OAUTH_CLIENT_ID"),                 // Replace with your GitHub client ID
		ClientSecret: config.Get("GITHUB_OAUTH_CLIENT_SECRET"),             // Replace with your GitHub client secret
		Scopes:       []string{"user:email"},                               // Requesting both profile and email scopes
		RedirectURL:  config.Get("CLIENT_URL") + "/callbacks/github/oauth", // Callback URL
		Endpoint:     ghOauth2.Endpoint,                                    // GitHub's OAuth 2.0 endpoint
	}
	Github.HmacSecretKey = config.Get("GITHUB_HMAC_SECRET_KEY")
}

func (g *githubExternal) getClient(code string) (*github.Client, error) {
	// exchange the authorization code for an access token
	token, err := githubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	// create a GitHub client with the access token
	client := github.NewClient(githubOauthConfig.Client(context.Background(), token))
	return client, nil
}

func (g *githubExternal) getUserPrimaryEmail(client *github.Client) (string, error) {
	emails, _, err := client.Users.ListEmails(context.Background(), nil)
	if err != nil {
		return "", err
	}

	primaryEmail := ""
	for _, email := range emails {
		if email.GetPrimary() && email.GetVerified() {
			primaryEmail = email.GetEmail()
			break
		}
	}

	return strings.ToLower(primaryEmail), nil
}

func (g *githubExternal) getUserName(client *github.Client) (string, error) {
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		return "", err
	}

	return user.GetName(), nil
}

func (g *githubExternal) GenerateOauthUrl() (string, error) {
	state, err := utils.Hmac.GenerateHMACState(g.HmacSecretKey)
	if err != nil {
		return "", err
	}

	url := githubOauthConfig.AuthCodeURL(state)
	return url, nil
}

func (g *githubExternal) GetGithubUser(code string) (string, string, error) {
	client, err := g.getClient(code)

	if err != nil {
		return "", "", err
	}

	primaryEmail, err := g.getUserPrimaryEmail(client)
	if err != nil {
		return "", "", err
	}

	userName, err := g.getUserName(client)
	if err != nil {
		return "", "", err
	}

	return primaryEmail, userName, nil
}
