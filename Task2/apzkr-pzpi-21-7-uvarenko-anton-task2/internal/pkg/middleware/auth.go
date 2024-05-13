package middleware

import (
	"net/http"
	"strings"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func TokenVerifier(authClient *auth.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rawToken := strings.Split(ctx.Request.Header.Get("Authorization"), " ")
		if len(rawToken) < 2 {
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}

		token, err := authClient.VerifyIDToken(ctx, rawToken[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}

		if token.Expires < time.Now().Unix() {
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}

		ctx.Next()
	}
}
