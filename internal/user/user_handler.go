package user

import (
	"github.com/AlaraEfe/BasketService/packages/config"
	jwtHelper "github.com/AlaraEfe/BasketService/packages/jwth"
	jwt "github.com/golang-jwt/jwt/v4"
	"os"
	"time"
	//"github.com/AlaraEfe/BasketService/internal/models"
	//"strconv"
	//"fmt"
	"github.com/AlaraEfe/BasketService/internal/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	repo   *UserRepository
	config *config.Config
}

func NewUserHandler(r *gin.RouterGroup, repo *UserRepository, config *config.Config) {
	h := &userHandler{repo: repo, config: config}

	r.POST("/login", h.login)
	r.POST("/signUp", h.signUp)

}

func (u *userHandler) login(c *gin.Context) {
	var req api.Login
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := u.repo.GetUser(&req.Email, &req.Password)
	if user == nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return

	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"iat":    time.Now().Unix(),
		"iss":    os.Getenv("ENV"),
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
		"role":   user.Role,
	})
	token := jwtHelper.GenerateToken(jwtClaims, u.config.JWTConfig.SecretKey)
	c.JSON(http.StatusOK, token)
}

func (u *userHandler) VerifyToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, u.config.JWTConfig.SecretKey)

	c.JSON(http.StatusOK, decodedClaims)
}

func (u *userHandler) signUp(c *gin.Context) {

	user := &api.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	us, err := u.repo.create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, UserToResponse(us))

}
