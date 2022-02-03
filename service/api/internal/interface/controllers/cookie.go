package controllers

import (
	"net/http"

	"github.com/maru44/enva/service/api/pkg/config"
)

func setCookie(w http.ResponseWriter, key string, value string, age int) {
	sameSite := http.SameSiteNoneMode
	secure := true
	if config.IsEnvDevelopment {
		sameSite = http.SameSiteLaxMode
		secure = false
	}
	http.SetCookie(w, &http.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		Domain:   config.FRONT_HOST,
		MaxAge:   age,
		SameSite: sameSite,
		Secure:   secure,
		HttpOnly: true,
	})
}

func destroyCookie(w http.ResponseWriter, key string) {
	setCookie(w, key, "", -1)
}
