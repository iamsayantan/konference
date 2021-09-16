package middlewares

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/iamsayantan/konference/config"
	"github.com/iamsayantan/konference/server/dto"
	"github.com/iamsayantan/konference/server/rendering"
	"net/http"
)

type contextKey int

const (
	keyAuthUserID = 0
)

func AuthChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			accessTokenCookie, err := r.Cookie("konference-auth")
			if err != nil {
				rendering.RenderError(w, r, "Unauthenticated Access", "auth.no_access_token", http.StatusUnauthorized)
				return
			}

			accessToken = accessTokenCookie.Value
		}

		if accessToken == "" {
			rendering.RenderError(w, r, "Unauthenticated Access", "auth.no_access_token", http.StatusUnauthorized)
			return
		}

		jwtTokenClaims := &dto.UserClaims{}
		token, err := jwt.ParseWithClaims(accessToken, jwtTokenClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppSecret), nil
		})

		if err != nil {
			rendering.RenderError(w, r, "Unauthenticated Access", "auth.invalid_signature", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			rendering.RenderError(w, r, "Unauthenticated Access", "auth.invalid_access_token", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, keyAuthUserID, jwtTokenClaims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetAuthenticatedUserIdFromContext is used to get the id of the authenticated user from the context. This method
// should only be called from handlers which is protected by the AuthChecker middleware.
func GetAuthenticatedUserIdFromContext(ctx context.Context) uint {
	return ctx.Value(keyAuthUserID).(uint)
}
