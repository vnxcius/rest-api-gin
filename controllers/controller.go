package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vnxcius/gin-api/database"
	"github.com/vnxcius/gin-api/models"
)

/*
O não recebimento do contexto do Gin por parâmetro faz com
que não seja possível retornar o JSON. O contexto é muito
importante e possui informações úteis, como por exemplo,
cabeçalhos, cookies, etc.
*/
func ShowAlunos(c *gin.Context) {
	var a []models.Aluno

	database.DB.Find(&a)
	c.JSON(200, a)
}

func Greetings(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome +"!",
	})
}

func CriarAluno(c *gin.Context) {
	var a models.Aluno

	if err := c.ShouldBindJSON(&a); err != nil { // ShouldBindJSON vincula automaticamente uma request em json para os devidos valores 
		c.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error(),
		})

		return
	}

	if err := models.Validate(&a); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error(),
		})
		
		return
	}

	database.DB.Create(&a)
	c.JSON(http.StatusOK, a)
}

func GetAluno(c *gin.Context) {
	var a models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&a, id)

	if a.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found:": "Aluno não encontrado",
		})

		return
	}
	
	c.JSON(http.StatusOK, a)
}

func DeleteAluno(c *gin.Context) {
	var a models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&a, id)

	if a.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found:": "Aluno não encontrado",
		})

		return
	}

	database.DB.Delete(&a, id)
	c.JSON(http.StatusOK, gin.H{
		"Aviso:": "Aluno foi deletado com sucesso.",
	})
}

func EditarAluno(c *gin.Context) {
	var a models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&a, id)

	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error(),
		})

		return
	}
	
	if err := models.Validate(&a); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error(),
		})
		
		return
	}
	database.DB.Model(&a).UpdateColumns(a)

	c.JSON(http.StatusOK, gin.H{
		"Aviso:": "Aluno foi atualizado com sucesso.",
	})

}

func Search(c *gin.Context){
	var a models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&a)

	if a.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found:": "Aluno não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, a)
}