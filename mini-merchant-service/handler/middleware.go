package handler

import (
	"mini-merchant-service/service"
	"mini-merchant-service/auth"
	"mini-merchant-service/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//middleware
func Middleware(userService service.UserService, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || len(authHeader) == 0 {
			errorResponse := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "unauthorize"})

			c.AbortWithStatusJSON(401, errorResponse)
			return
		}

		token, err := authService.ValidateToken(authHeader)

		if err != nil {
			errorResponse := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": err.Error()})

			c.AbortWithStatusJSON(401, errorResponse)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			errorResponse := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "unauthorize"})

			c.AbortWithStatusJSON(401, errorResponse)
			return
		}

		userID := ""

		if claim["user_id"] != nil {
			userID = claim["user_id"].(string)
		}

		c.Set("currentUser", gin.H{
			"user_id": userID,
		})
	}
}

//corsmiddleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
