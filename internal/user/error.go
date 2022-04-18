package user

import (
	"github.com/gin-gonic/gin"
)

func errorResponse(err error) gin.H {
	return gin.H{"UserError": err.Error()}
}
