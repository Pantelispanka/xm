package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"xm-challenge/internal/domain"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"gopkg.in/square/go-jose.v2/jwt"
)

var Verifier *oidc.IDTokenVerifier
var oauth2Config oauth2.Config
var keyset oidc.KeySet

type IAuth interface {
	AuthMiddleware(h http.HandlerFunc) http.HandlerFunc
}

type AuthImplementation struct {
}

func InitVerifier() {
	log.Println("Starting oidc configuration")
	oidcProvider := os.Getenv("OIDC_PROVIDER")
	if oidcProvider == "" {
		oidcProvider = "*"
	}

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, oidcProvider)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	clientID := "accounts-api"

	oauth2Config = oauth2.Config{
		ClientID: clientID,
		Endpoint: provider.Endpoint(),
		Scopes:   []string{oidc.ScopeOpenID, "profile", "email", "openid"},
	}

	oidcConfig := &oidc.Config{
		ClientID:          clientID,
		SkipClientIDCheck: true,
		SkipIssuerCheck:   true,
	}
	keyset = oidc.NewRemoteKeySet(ctx, oidcProvider)

	Verifier = provider.Verifier(oidcConfig)
}

func (a *AuthImplementation) AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiAll := r.Header.Get("Authorization")
		if apiAll == "" {
			err := domain.ErrorReport{
				Error:  "Not authorized",
				Status: http.StatusUnauthorized,
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}
		apiKeyAr := strings.Split(apiAll, " ")
		authType := apiKeyAr[0]
		if authType == "Bearer" {
			apiKey := apiKeyAr[1]
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			_, err := Verifier.Verify(ctx, apiKey)
			if err != nil {
				err := domain.ErrorReport{
					Error:  err.Error(),
					Status: http.StatusUnauthorized,
				}
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(err)
				return
			}
		} else {
			err := domain.ErrorReport{
				Error:  "Not authorized",
				Status: http.StatusUnauthorized,
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}

		h.ServeHTTP(w, r)
	})
}

type JwtClaims struct {
	*jwt.Claims
	Name             string   `json:"name,omitEmpty"`
	PreferedUsername string   `json:"preferred_username"`
	GivenName        string   `json:"given_name"`
	FamilyName       string   `json:"family_name,omitEmpty"`
	Email            string   `json:"email"`
	Groups           []string `json:"groups"`
}

func GetClaims(token string) (claims JwtClaims, err error) {
	resultCl := JwtClaims{}
	tokenVer, err := Verifier.Verify(context.TODO(), token)
	erro := tokenVer.Claims(&resultCl)
	if erro != nil {
		log.Println("failed to parse Claims: ", erro.Error())
	}
	return resultCl, err
}

func Verify(ctx context.Context, apiKey string)
