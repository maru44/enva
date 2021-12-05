package domain

type (
	ContextKey string

	CtxAccess struct {
		Method string
		URL    string
	}
)

const (
	CtxUserKey          ContextKey = "user"
	CtxCognitoKeySetKey ContextKey = "cognito_keyset"
	CtxAccessKey        ContextKey = "access"
)
