package basket

import (
	"github.com/gin-gonic/gin"
)

func errorResponse(err error) gin.H {
	return gin.H{"BasketItemError": err.Error()}
}
