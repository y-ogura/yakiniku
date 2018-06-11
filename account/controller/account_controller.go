package controller

import (
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/y-ogura/yakiniku/account" // import account package
	"github.com/y-ogura/yakiniku/account/usecase"
)

// AccountController account controller
type AccountController struct {
	AccountUsecase usecase.AccountUsecase
}

// NewAccountController mount account controller
func NewAccountController(e *echo.Echo, account usecase.AccountUsecase) {
	handler := &AccountController{
		AccountUsecase: account,
	}

	e.GET("/accounts", handler.List)
}

// List list accounts
// @Title アカウント一覧取得
// @Description アカウント一覧取得API
// @Assept json
// @Success 200 {array} account.Client true "アカウント"
// @Failure 500 {object} error true "internal server error"
// @Resource /accounts
// @Router /accounts [get]
func (c *AccountController) List(ctx echo.Context) error {
	res, err := c.AccountUsecase.List()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}
