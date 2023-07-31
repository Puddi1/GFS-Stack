package handlers

// To clean and riarrange

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Puddi1/GFS-Stack/env"
	"github.com/Puddi1/GFS-Stack/utils"
)

// Struct that represent the JSON format of the body of the update user api request
type UpdateUserBody struct {
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Data     UpdateUserData `json:"data"`
}

// Nested data type, simplle key to value type
type UpdateUserData struct {
	KeyValue map[string]string
}

// OAuth providers
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

// HandleSignUpUserWithEmail allows you to sign up users with email and password via supabase Auth
func HandleSignUpUserWithEmail(email string, password string) error {
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
		log.Fatalf("Error during User signup request: %s", err)
	}

	resBody := HandleResponseBodyToString(res)

	fmt.Printf("\nUser signed up, response: %s", resBody)

	return nil
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
		log.Fatalf("Error during User login request: %s", err)
	}

	resBody := HandleResponseBodyToString(res)

	fmt.Printf("\nUser logged in, response: %s", resBody)
	fmt.Println(res)

	return res, nil
}

// HandleLoginUserWithPhone allows you to login users with phone and password via supabase Auth
func HandleLoginUserWithPhone(phone string, password string) error {
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
		log.Fatalf("Error during User login with phone request: %s", err)
	}

	resBody := HandleResponseBodyToString(res)

	fmt.Printf("\nUser logged in with phone, response: %s", resBody)

	return nil
}

// HandleSendRecoveryEmail allows you to send users a recovery email via supabase Auth
func HandleSendRecoveryEmail(email string) error {
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
		log.Fatalf("Error during Recovery email request: %s", err)
	}

	resBody := HandleResponseBodyToString(res)

	fmt.Printf("\nRecovery email sent, response: %s", resBody)

	return nil
}

// HandleUpdateUser allows you to update users' email and password via supabase Auth
func HandleUpdateUser(u *UpdateUserBody) error {
	var UserAccessToken string

	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/user"
	body := utils.StructToJSON(u)
	headers := [][2]string{
		{"apikey", apiKey},
		{"Authorization", "Bearer " + UserAccessToken},
		{"Content-Type", "application/json"},
	}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPut,
		Url:        url,
		Body:       body,
		Headers:    headers,
	})
	if err != nil {
		log.Fatalf("Error during User login: %s", err)
	}

	resBody := HandleResponseBodyToString(res)

	fmt.Printf("\nUser logged in, response: %s", resBody)

	return nil
}

// HandleLogOutUser allows you to logout users via supabase Auth
func HandleLogOutUser(email string, password string) error {
	var UserToken string

	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/logout"
	body := []byte{}
	headers := [][2]string{
		{"Content-Type", "application/json"},
		{"apikey", apiKey},
		{"Authorization", "Bearer " + UserToken},
	}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        url,
		Body:       body,
		Headers:    headers,
	})
	if err != nil {
		log.Fatalf("Error during User login: %s", err)
	}

	resBody := HandleResponseBodyToString(res)

	fmt.Printf("\nUser logged in, response: %s", resBody)

	return nil
}

// HandleSendUserInvite allows you to send users invites via email via supabase Auth
func HandleSendUserInvite(email string) error {
	var SUPABASE_KEY string

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
		log.Fatalf("Error during User login: %s", err)
	}

	resBody := HandleResponseBodyToString(res)

	fmt.Printf("\nUser logged in, response: %s", resBody)

	return nil
}

// HandleLoginWithMagicLinkViaEmail allows you to login users with magic link sent via email via supabase Auth
func HandleLoginWithMagicLinkViaEmail(email string) error {
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
		log.Fatalf("Error during User login with magic link via email: %s", err)
	}

	resBody := HandleResponseBodyToString(res)

	fmt.Printf("\nUser logged in with magic link via email, response: %s", resBody)

	return nil
}

// HandleLoginViaSMSOTP allows you to login users with SMS OTP via supabase Auth
// Note: VerifyViaSMSOTP needed, twilio credentials needed
func HandleLoginViaSMSOTP(phone string) error {
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
		log.Fatalf("Error during User login with magic link via email: %s", err)
	}

	resBody := HandleResponseBodyToString(res)

	fmt.Printf("\nUser logged in with magic link via email, response: %s", resBody)

	return nil
}

// HandleVerifyViaSMSOTP allows you to very users' SMS OTP via supabase Auth
// Note: LoginViaSMSOTP needed, twilio credentials needed
func HandleVerifyViaSMSOTP(phone string, token string) error {
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
		log.Fatalf("Error during User Verification via SMS OTP: %s", err)
	}

	resBody := HandleResponseBodyToString(res)

	fmt.Printf("\nUser logged in with SMS OTP, response: %s", resBody)

	return nil
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

// HandleGetCurrentUser allows you to get the current user infos via supabase Auth
func HandleGetCurrentUser() (*http.Response, error) {
	var userToken string

	apiKey := env.ENVs["SUPABASE_API_PRIVATE_KEY"]
	url := env.ENVs["SUPABASE_URL"] + "/auth/v1/user"
	body := []byte{}
	headers := [][2]string{
		{"Content-Type", "application/json"},
		{"apikey", apiKey},
		{"Authorization", "Bearer" + userToken},
	}

	res, err := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodGet,
		Url:        url,
		Body:       body,
		Headers:    headers,
	})
	if err != nil {
		log.Fatalf("Error during User fetch: %s", err)
	}

	return res, nil
}

// Login with third party auth
// Check error handling and return statements of functions here
// check standalone variables to complete and where to find them
