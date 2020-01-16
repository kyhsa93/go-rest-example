package account

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-example/account/repository"
	"github.com/kyhsa93/gin-rest-example/account/router"
	"github.com/kyhsa93/gin-rest-example/account/service"
	"github.com/kyhsa93/gin-rest-example/config"
	"github.com/kyhsa93/gin-rest-example/util"
)

// InitializeAccount innitialize account module
func InitializeAccount(engine *gin.Engine, config *config.Config, util *util.Util) {
	repository := repository.NewRepository(config)
	service := service.NewService(repository)
	router.NewRouter(engine, service, util)
}
