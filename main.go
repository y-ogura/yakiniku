// @APIVersion 1.0.0
// @APITitle yakiniku blog
// @APIDescription yakiniku blog sample
// @BasePath http://localhost:8081/swagger-ui
// @SubApi account [/accounts]

package main

import (
	"github.com/labstack/echo"
	"github.com/y-ogura/yakiniku/account/controller"
	"github.com/y-ogura/yakiniku/account/repository"
	"github.com/y-ogura/yakiniku/account/usecase"
	swaggerui "github.com/y-ogura/yakiniku/swagger-ui"
)

func main() {
	e := echo.New()
	setup(e)

	accountRepo := repository.NewAccountRepository(database)

	accountUcase := usecase.NewAccountUsecase(accountRepo)

	controller.NewAccountController(e, accountUcase)

	swaggerui.NewSwaggerController(e)

	e.Logger.Fatal(e.Start(":8081"))
}
