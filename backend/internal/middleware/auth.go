package middleware

import (
	"net/http"
	"strings"

	"github.com/Zecci9/curly/backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")

		if header == "" {

			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "缺少Token",
				},
			)

			c.Abort()

			return
		}

		tokenString :=
			strings.TrimPrefix(
				header,
				"Bearer ",
			)

		claims, err :=
			auth.ParseToken(tokenString)

		if err != nil {

			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"message": err.Error(),
				},
			)

			c.Abort()
			return
		}

		c.Set(
			"user_id",
			claims["user_id"],
		)

		c.Next()

	}

}
