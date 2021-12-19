package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	Env              = os.Getenv("ENV")
	IsEnvDevelopment = os.Getenv("ENV") == "development"
	API_PROTOCOL     = os.Getenv("API_PROTCOL")
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

	CLI_HEADER_SEP = "=+=+=+="
)
