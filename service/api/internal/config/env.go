package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	Env              = os.Getenv("ENVIRONMENT")
	IsEnvDevelopment = os.Getenv("ENVIRONMENT") == "development"
	API_PROTOCOL     = os.Getenv("API_PROTOCOL")
	API_HOST         = os.Getenv("API_HOST")
	API_PORT         = os.Getenv("API_PORT")
	API_URL          = API_PROTOCOL + API_HOST + API_PORT
	FRONT_PROTOCOL   = os.Getenv("FRONT_PROTOCOL")
	FRONT_HOST       = os.Getenv("FRONT_HOST")
	FRONT_PORT       = os.Getenv("FRONT_PORT")
	FRONT_URL        = FRONT_PROTOCOL + FRONT_HOST + FRONT_PORT

	POSTGRES_URL                     = os.Getenv("POSTGRES_URL")
	POSTGRES_MAX_CONNECTIONS, _      = strconv.Atoi(os.Getenv("POSTGRES_MAX_CONNECTIONS"))
	POSTGRES_MAX_IDLE_CONNECTIONS, _ = strconv.Atoi(os.Getenv("POSTGRES_MAX_IDLE_CONNECTIONS"))

	COGNITO_KEYS_URL = fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", os.Getenv("COGNITO_REGION"), os.Getenv("COGNITO_USERPOOL_ID"))
	COGNITO_SECRET   = os.Getenv("COGNITO_SECRET")

	REDIS_ADDR = os.Getenv("REDIS_ADDRESS")
	REDIS_PASS = os.Getenv("REDIS_PASSWORD")

	EMAIL_HOST      = os.Getenv("EMAIL_HOST")
	EMAIL_HOST_USER = os.Getenv("EMAIL_HOST_USER")
	EMAIL_HOST_PASS = os.Getenv("EMAIL_HOST_PASS")
	EMAIL_PORT      = os.Getenv("EMAIL_PORT")
)
