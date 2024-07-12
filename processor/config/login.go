package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauth = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/v1/callback/google", //"http://localhost:3000/main", //
		ClientID:     "554692719373-vv40si5k7elfa9p61vm26kbeouno9cfv.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/contacts.readonly", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	OauthStateString = "randomstatestring"
)
