package middleware

import (
	"context"
	"net/http"

	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/httpresp"
	"github.com/pjimming/zustacm/core/utils/userauth"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
)

type JwtAuthMiddleware struct {
}

func NewJwtAuthMiddleware() *JwtAuthMiddleware {
	return &JwtAuthMiddleware{}
}

func JwtAuth(rdc *redis.Client) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// client
			token := r.Header.Get("Authorization")
			if token == "" {
				httpresp.HttpError(w, r, errorx.ErrorTokenInvalid())
				return
			}

			// server
			if getToken, _ := rdc.Get(context.Background(),
				userauth.GenTokenKey(token)).Bool(); !getToken {
				httpresp.HttpError(w, r, errorx.ErrorTokenInvalid())
				return
			}

			// parse token
			uc, err := userauth.ParseToken(token)
			if err != nil {
				httpresp.HttpError(w, r, errorx.ErrorAuth())
				return
			}

			ctx := userauth.CtxWithToken(
				userauth.CtxWithUser(r.Context(), uc),
				token,
			)

			next(w, r.WithContext(ctx))
		}
	}
}
