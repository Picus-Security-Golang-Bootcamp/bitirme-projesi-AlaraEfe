package mw

import (
	jwtHelper "github.com/AlaraEfe/BasketService/packages/jwth"
	"github.com/gin-gonic/gin"

	"net/http"
)

func AuthMiddlewareAdmin(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				if decodedClaims.Role == "Admin" {
					c.Next()
					c.Abort()
					return
				}

			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()

	}
}

func AuthMiddlewareUser(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				if decodedClaims.Role == "User" || decodedClaims.Role == "Admin" {
					c.Next()
					c.Abort()
					return
				}

			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()

	}
}
