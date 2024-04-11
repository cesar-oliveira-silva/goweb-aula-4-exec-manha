package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cesar-oliveira-silva/goweb-aula-4-exec-manha.git/project/cmd/server/handler"
	"github.com/cesar-oliveira-silva/goweb-aula-4-exec-manha.git/project/internal/usuarios"
	"github.com/cesar-oliveira-silva/goweb-aula-4-exec-manha.git/project/pkg/store"
	"github.com/cesar-oliveira-silva/goweb-aula-4-exec-manha.git/project/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// var dbConn *sql.DB

func TokenMiddleware(ctx *gin.Context) {
	tokenEnvironment := os.Getenv("TOKEN")
	token := ctx.GetHeader("token")
	if token != tokenEnvironment {
		// status StatusUnauthorized equivalente ao 401
		// ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		// 	"error": "token inválido",
		// })
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro: Token invalido"))
		return
	}
	ctx.Next()
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}
	db := store.NewFileStore("file", "usuarios.json")
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	userHandler := handler.NewUser(service)

	server := gin.Default()
	pr := server.Group("/usuarios")
	pr.Use(TokenMiddleware)
	pr.POST("/", userHandler.Store())
	pr.GET("/", userHandler.GetAll())
	pr.PUT("/:id", userHandler.Update())
	pr.PATCH("/:id", userHandler.UpdateName())
	pr.DELETE("/:id", userHandler.Delete())

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
