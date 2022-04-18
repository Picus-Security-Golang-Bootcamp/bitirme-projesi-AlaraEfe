package category

import (
	"github.com/AlaraEfe/BasketService/internal/models"
	config "github.com/AlaraEfe/BasketService/packages/config"
	"github.com/AlaraEfe/BasketService/packages/mw"
	"strconv"
	//"fmt"
	"encoding/csv"
	"github.com/AlaraEfe/BasketService/internal/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type categoryHandler struct {
	repo *CategoryRepository
}

func NewCategoryHandler(r *gin.RouterGroup, repo *CategoryRepository, config *config.Config) {
	h := &categoryHandler{repo: repo}

	r.GET("/list", h.listCategories)
	r.POST("/upload", h.uploadCategoryFile, mw.AuthMiddlewareAdmin(config.JWTConfig.SecretKey))
}

func (ch *categoryHandler) uploadCategoryFile(c *gin.Context) {
	csvFile, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var categories api.CategorySlice

	for _, line := range csvLines[1:] {
		category := &models.Category{}
		ID, _ := strconv.Atoi(line[0])
		category.ID = uint(ID)
		category.CategoryName = line[1]

		category, err := ch.repo.create(category)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		categories = append(categories, CategoryToResponse(category))

	}

	c.JSON(http.StatusOK, categories)

}

func (ch *categoryHandler) listCategories(c *gin.Context) {
	categories, _ := ch.repo.listCategories()

	categoriesResponses := categoriesToResponse(categories)

	c.JSON(http.StatusOK, categoriesResponses)

}
