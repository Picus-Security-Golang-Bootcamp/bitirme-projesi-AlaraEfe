package product

import (
	"github.com/AlaraEfe/BasketService/internal/api"
	config "github.com/AlaraEfe/BasketService/packages/config"
	"github.com/AlaraEfe/BasketService/packages/mw"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	repo *ProductRepository
}

func NewProductHandler(r *gin.RouterGroup, repo *ProductRepository, config *config.Config) {
	h := &ProductHandler{repo: repo}

	r.GET("/list", h.getAll)
	r.POST("/create", h.create, mw.AuthMiddlewareAdmin(config.JWTConfig.SecretKey))
	r.GET("/:id", h.getByID)
	r.GET("/searchByProductName/:name", h.SearchByProductName)
	r.GET("/searchByProductID/:id", h.SearchByProductID)
	r.PUT("/updateProductByID/:id", h.updateProductByID, mw.AuthMiddlewareAdmin(config.JWTConfig.SecretKey))
	r.DELETE("/deleteProductByID/:id", h.deleteProductByID, mw.AuthMiddlewareAdmin(config.JWTConfig.SecretKey))
}

func (p *ProductHandler) create(c *gin.Context) {

	productBody := &api.Product{}

	if err := c.ShouldBindJSON(&productBody); err != nil {

		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := p.repo.create(responseToProduct(productBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, ProductToResponse(product))

}

func (p *ProductHandler) getAll(c *gin.Context) {
	products, _ := p.repo.getAll()

	productResponses := productsToResponse(products)

	c.JSON(http.StatusOK, productResponses)

}

func (p *ProductHandler) getByID(c *gin.Context) {

	value, _ := strconv.Atoi(c.Param("id"))

	product, err := p.repo.getByID(uint(value))

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, ProductToResponse(product))

}

func (p *ProductHandler) SearchByProductName(c *gin.Context) {

	products, err := p.repo.SearchByProductName(c.Param("name"))

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, productsToResponse(products))

}

func (p *ProductHandler) SearchByProductID(c *gin.Context) {

	value, _ := strconv.Atoi(c.Param("id"))

	product, err := p.repo.SearchByProductID(uint(value))

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, ProductToResponse(product))

}

func (p *ProductHandler) deleteProductByID(c *gin.Context) {
	value, _ := strconv.Atoi(c.Param("id"))

	err := p.repo.deleteProductByID(uint(value))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusNoContent, nil)

}

func (p *ProductHandler) updateProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	productBody := &api.Product{ID: uint(id)}
	if err := c.Bind(&productBody); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := p.repo.updateProductByID(responseToProduct(productBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, ProductToResponse(product))
}
