package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lailiseptiandi/go-store-api/config"
	"github.com/lailiseptiandi/go-store-api/helpers"
	"github.com/lailiseptiandi/go-store-api/service"
)

func AuthMiddleware(authService config.Service, globalService service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		authToken := strings.Split(authHeader, " ")
		if len(authToken) == 2 {
			tokenString = authToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {

			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get id parse
		userID := int(claim["user_id"].(float64))
		user, err := globalService.GetUserByID(userID)
		if err != nil {
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// if utils {

		// }

		c.Set("currentUser", user)

	}
}
