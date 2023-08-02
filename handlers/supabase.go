package handlers

// To clean and riarrange

import (
	"fmt"
	"net/http"

	"github.com/Puddi1/GFS-Stack/env"
	"github.com/Puddi1/GFS-Stack/utils"
)

//**********************//
//*****    AUTH    *****//
//**********************//

// UpdateUserBody represent the JSON format of the body of the update user api request
type UpdateUserBody struct {
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Data     UpdateUserData `json:"data"`
}

// UpdateUserData is a nested UpdateUserBody data type, simple key to value type
type UpdateUserData struct {
	KeyValue map[string]string
}

// OAuth providers map, be sure to handle correctly the key in the frontend
var OAuth = map[string]string{
	"Apple":     "apple",
	"Azure":     "azure",
	"Bitbucket": "bitbucket",
	"Discord":   "discord",
	"Facebook":  "facebook",
	"Figma":     "figma",
	"Github":    "github",
	"Gitlab":    "gitlab",
	"Google":    "google",
	"Kexcloak":  "keycloak",
	"Linkedin":  "linkedin",
	"Notion":    "notion",
	"Slack":     "slack",
	"Spotify":   "spotify",
	"Twitch":    "twitch",
	"Twitter":   "twitter",
	"Workos":    "workos",
}

// GetUser allows you to get the current user infos via supabase Auth
func GetUser(userAccessToken string) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/user"
	body := []byte{}
	headers := [][2]string{
		{"Content-Type", "application/json"},
		{"apikey", apiKey},
		{"Authorization", "Bearer" + userAccessToken},
	}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodGet,
		Url:        url,
		Body:       body,
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleSignUpUserWithEmail allows you to sign up users with email and password via supabase Auth
func HandleSignUpUserWithEmail(email string, password string) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/signup"
	body := map[string]string{"email": email, "password": password}
	headers := [][2]string{{"Content-Type", "application/json"}, {"apikey", apiKey}}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        url,
		Body:       utils.MapToJSON(body),
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleLoginUserWithEmail allows you to login users with email and password via supabase Auth
func HandleLoginUserWithEmail(email string, password string) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/token?grant_type=password"
	body := map[string]string{"email": email, "password": password}
	headers := [][2]string{{"Content-Type", "application/json"}, {"apikey", apiKey}}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        url,
		Body:       utils.MapToJSON(body),
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleLoginUserWithPhone allows you to login users with phone and password via supabase Auth
// Note: You must enter your own twilio credentials on the auth settings page to enable SMS-based Logins.
func HandleLoginUserWithPhone(phone string, password string) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/signup"
	body := map[string]string{"phone": phone, "password": password}
	headers := [][2]string{{"Content-Type", "application/json"}, {"apikey", apiKey}}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        url,
		Body:       utils.MapToJSON(body),
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleSendRecoveryEmail allows you to send users a recovery email via supabase Auth
func HandleSendRecoveryEmail(email string) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/recover"
	body := map[string]string{"email": email}
	headers := [][2]string{{"Content-Type", "application/json"}, {"apikey", apiKey}}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        url,
		Body:       utils.MapToJSON(body),
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleUpdateUser allows you to update users' email and password via supabase Auth
func HandleUpdateUser(u *UpdateUserBody, userAccessToken string) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/user"
	body := utils.StructToJSON(u)
	headers := [][2]string{
		{"apikey", apiKey},
		{"Authorization", "Bearer " + userAccessToken},
		{"Content-Type", "application/json"},
	}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPut,
		Url:        url,
		Body:       body,
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleLogOutUser allows you to logout users via supabase Auth
func HandleLogOutUser(email string, password string, userAccessToken string) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/logout"
	body := []byte{}
	headers := [][2]string{
		{"Content-Type", "application/json"},
		{"apikey", apiKey},
		{"Authorization", "Bearer " + userAccessToken},
	}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        url,
		Body:       body,
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleSendUserInvite allows you to send users invites via email via supabase Auth
func HandleSendUserInvite(email string, SUPABASE_KEY string) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/invite"
	body := map[string]string{"email": email}
	headers := [][2]string{
		{"Content-Type", "application/json"},
		{"apikey", apiKey},
		{"Authorization", "Bearer " + SUPABASE_KEY},
	}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        url,
		Body:       utils.MapToJSON(body),
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleLoginWithMagicLinkViaEmail allows you to login users with magic link sent via email via supabase Auth
func HandleLoginWithMagicLinkViaEmail(email string) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/magiclink"
	body := map[string]string{"email": email}
	headers := [][2]string{
		{"Content-Type", "application/json"},
		{"apikey", apiKey},
	}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        url,
		Body:       utils.MapToJSON(body),
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleLoginViaSMSOTP allows you to login users with SMS OTP via supabase Auth
// Note: VerifyViaSMSOTP needed, twilio credentials needed
func HandleLoginViaSMSOTP(phone string) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/otp"
	body := map[string]string{"phone": phone}
	headers := [][2]string{
		{"Content-Type", "application/json"},
		{"apikey", apiKey},
	}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        url,
		Body:       utils.MapToJSON(body),
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleVerifyViaSMSOTP allows you to very users' SMS OTP via supabase Auth
// Note: You must enter your own twilio credentials on the auth settings page to enable SMS-based Logins.
func HandleVerifyViaSMSOTP(phone string, token string) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/verify"
	body := map[string]string{
		"type":  "sms",
		"phone": phone,
		"token": token,
	}
	headers := [][2]string{
		{"Content-Type", "application/json"},
		{"apikey", apiKey},
	}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        url,
		Body:       utils.MapToJSON(body),
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleLoginWithThirdPartyOAuth allows you to login users with a third party OAuth via supabase Auth
// Note: OAuth credentials and activation needed on your supabase project settings, to redirect
// user to a precise url you need to add it on redirect URLs, if you use a default redirectUrl,
// thus "", the provider will redirect to the Site URL, which is the default URL of you app.
// Configuration: https://supabase.com/dashboard/project/<your_project>/auth/url-configuration
// Guide: https://supabase.com/docs/learn/auth-deep-dive/auth-google-oauth
func HandleLoginWithThirdPartyOAuth(provider string, redirectUrl string) (location string, code int) {
	// Compose url
	url := env.ENVs["SUPABASE_URL"] +
		"/auth/v1/authorize?provider=" + provider +
		"&redirect_to=https://" + redirectUrl
	// Send destination
	return url, http.StatusSeeOther
}

//**********************//
//*****  WEBHOOKS  *****//
//**********************//

// func i() {}

//**********************//
//*****  FUNCTIONS *****//
//**********************//

//**********************//
//*****   BUCKET   *****//
//**********************//
// Don't care much about impl. leaving at last, maybe.

// HandleSetFileCache sets the cache for a file, using "path/to/file.jpg" and version
func HandleSetFileCache(pathToFile string, version int) (*http.Response, error) {
	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/storage/v1/object/"
	body := map[string]string{}
	headers := [][2]string{}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        fmt.Sprintf("%s%s?token=%s&version=%v", url, pathToFile, apiKey, version),
		Body:       utils.MapToJSON(body),
		Headers:    headers,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
