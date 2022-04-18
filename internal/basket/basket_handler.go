package basket

import (
	"github.com/AlaraEfe/BasketService/internal/api"
	config "github.com/AlaraEfe/BasketService/packages/config"
	"github.com/AlaraEfe/BasketService/packages/mw"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BasketItemHandler struct {
	repo *basketItemRepository
}

func NewBasketItemHandler(r *gin.RouterGroup, repo *basketItemRepository, config *config.Config) {
	h := &BasketItemHandler{repo: repo}

	r.POST("/add", h.addToCart, mw.AuthMiddlewareUser(config.JWTConfig.SecretKey))
	r.GET("/list/:id", h.listCart, mw.AuthMiddlewareUser(config.JWTConfig.SecretKey))
	r.DELETE("/deleteBasketItemByID/:id", h.deleteBasketItemByID, mw.AuthMiddlewareUser(config.JWTConfig.SecretKey))
	r.PUT("/updateBasketItemByID/:id", h.UpdateBasketItemByID, mw.AuthMiddlewareUser(config.JWTConfig.SecretKey))

}

func (b *BasketItemHandler) addToCart(c *gin.Context) {
	basketItem := &api.BasketItem{}
	if err := c.ShouldBindJSON(&basketItem); err != nil {

		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	u, err := b.repo.addToCart(ResponseToBasketItem(basketItem))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, BasketItemToResponse(u))
}

func (b *BasketItemHandler) listCart(c *gin.Context) {

	value, _ := strconv.Atoi(c.Param("id"))

	items, err := b.repo.listCart(uint(value))

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, basketItemsToResponse(items))

}

func (b *BasketItemHandler) deleteBasketItemByID(c *gin.Context) {
	value, _ := strconv.Atoi(c.Param("id"))

	err := b.repo.deleteBasketItemByID(uint(value))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusNoContent, nil)

}

func (b *BasketItemHandler) UpdateBasketItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	basketItemBody := &api.BasketItem{ID: uint(id)}
	if err := c.Bind(&basketItemBody); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	item, err := b.repo.UpdateBasketItemByID(ResponseToBasketItem(basketItemBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, BasketItemToResponse(item))
}
