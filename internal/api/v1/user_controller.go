package v1

import (
	"github.com/gin-gonic/gin"
	"pokemon-be/internal/apperr"
	"pokemon-be/internal/config"
	"pokemon-be/internal/constants"
	"pokemon-be/internal/httpresp"
	"pokemon-be/internal/request"
	"pokemon-be/internal/service"
)

const (
	userBasePath = "/user"
)

type UserController struct {
	cfg         config.Config
	userService service.UserService
}

func NewUserController(cfg config.Config, userService service.UserService) *UserController {

	return &UserController{
		cfg:         cfg,
		userService: userService,
	}
}

func (h *UserController) AddRoutes(r *gin.Engine) {
	uc := r.Group(constants.V1BasePath + userBasePath)

	uc.POST("/register", h.Register)
	uc.POST("/login", h.Login)
}

func (h *UserController) Register(c *gin.Context) {

	var data *request.RegisterUserRequest

	if err := c.BindJSON(&data); err != nil {
		httpresp.HttpRespError(c, apperr.ErrBadRequest)
		return
	}

	err := h.userService.CreateUser(data)

	if err != nil {
		h.cfg.Logger().Error().Err(err).Msg("[Register] Error registering user")
		httpresp.HttpRespError(c, err)
		return
	}

	httpresp.HttpRespSuccess(c, nil, nil)
}

func (h *UserController) Login(c *gin.Context) {

	var data *request.LoginRequest

	if err := c.BindJSON(&data); err != nil {
		httpresp.HttpRespError(c, apperr.ErrBadRequest)
		return
	}

	user, err := h.userService.Login(data)

	if err != nil {
		h.cfg.Logger().Error().Err(err).Msg("[Login] Error login")
		httpresp.HttpRespError(c, err)
		return
	}

	httpresp.HttpRespSuccess(c, user, nil)
}
