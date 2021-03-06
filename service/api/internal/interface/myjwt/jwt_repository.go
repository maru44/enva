package myjwt

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type JwtRepository struct {
	JwtParserAbstract
}

func (repo *JwtRepository) FetchJwk(ctx context.Context, url string) (jwk.Set, error) {
	return jwk.Fetch(ctx, url)
}

func (repo *JwtRepository) GetUserByJwt(ctx context.Context, idToken string) (*domain.User, error) {
	token, err := repo.evaluate(ctx, idToken)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrUnauthorized)
	} else if !token.Valid {
		return nil, perr.New("Invalid Token", perr.ErrUnauthorized)
	}

	claim := token.Claims.(jwt.MapClaims)
	claimJson, err := json.Marshal(claim)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrBadRequest)
	}

	var userClaim domain.UserFromClaim
	if err = json.Unmarshal(claimJson, &userClaim); err != nil {
		return nil, perr.Wrap(err, perr.ErrBadRequest)
	}
	user := userClaim.ToUser()

	return user, nil
}

func (repo *JwtRepository) evaluate(ctx context.Context, idToken string) (*jwt.Token, error) {
	token, err := repo.Parse(idToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, perr.New(fmt.Sprintf("Unexpected Signing method: %v", t.Header["alg"]), perr.ErrBadRequest)
		}

		// kid from token
		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, perr.New("kid header not found", perr.ErrBadRequest)
		}

		// keyset from context
		keySet, ok := ctx.Value(domain.CtxCognitoKeySetKey).(jwk.Set)
		if !ok {
			return nil, perr.New("keySet is not set in context", perr.ErrBadRequest)
		}

		keys, _ := keySet.LookupKeyID(kid)
		var pubKey interface{}
		if err := keys.Raw(&pubKey); err != nil {
			return nil, perr.Wrap(err, perr.ErrBadRequest)
		}

		return pubKey, nil
	})
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrUnauthorized)
	}

	return token.GetToken(), nil
}
