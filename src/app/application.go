package app

import (
	"github.com/atg0831/msabookstore/bookstore_oauth-api/src/domain/access_token"
	"github.com/atg0831/msabookstore/bookstore_oauth-api/src/http"
	"github.com/atg0831/msabookstore/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.Run(":8080")
}
