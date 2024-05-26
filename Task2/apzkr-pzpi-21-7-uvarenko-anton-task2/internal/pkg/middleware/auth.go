package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func TokenVerifier(jwtHandler jwt.JWT, userTypesAllowed []core.UsersUserType) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rawToken := strings.Split(ctx.Request.Header.Get("Authorization"), " ")
		if len(rawToken) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		claims, err := jwtHandler.VerifyToken(rawToken[1], userTypesAllowed)
		if err != nil {
			fmt.Println(err)
			if errors.Is(err, pkg.ErrForbiden) {
				ctx.AbortWithStatusJSON(http.StatusForbidden, nil)
				return
			}

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		ctx.Set("user_id", claims.ID)
		ctx.Set("user_type", string(claims.UserType))

		ctx.Next()
	}
}
