package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vnxcius/gin-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ShowAlunos)
	r.POST("/alunos", controllers.CriarAluno)
	r.GET("/alunos/:id", controllers.GetAluno) // :id e :nome são gin params
	r.DELETE("alunos/:id", controllers.DeleteAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/search/:cpf", controllers.Search)

	r.GET("/:nome", controllers.Greetings)

	r.Run() // Abrir o servidor na porta default (8080)
	
	//* r.NoRoute(controllers.PaginaNaoEncontrada) redirecionamento 404 através do gin
}