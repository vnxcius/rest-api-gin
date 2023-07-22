package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vnxcius/gin-api/controllers"
	"github.com/vnxcius/gin-api/database"
	"github.com/vnxcius/gin-api/models"
)

var ID int

func RoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func CriaAlunoMock() {
	a := models.Aluno{
		Nome: "Aluno Teste",
		CPF:  "12345678910",
		RG:   "123456789",
	}

	database.DB.Create(&a)
	ID = int(a.ID)
}

func DeletaAlunoMock() {
	var a models.Aluno
	database.DB.Delete(&a, ID)

}

func TestStatusCodeParam(t *testing.T) {
	r := RoutesSetup()
	r.GET("/:nome", controllers.Greetings)

	req, _ := http.NewRequest("GET", "/vini", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	resMock := `{"API diz:":"E ai vini!"}` // Testar o body da requisição por mock
	resBody, _ := io.ReadAll(res.Body)

	assert.Equal(t, resMock, string(resBody))
}

func TestShowAlunosHandler(t *testing.T) {
	database.Connection()
	CriaAlunoMock() //* Garantir que haja ao menos 1 registro neste teste, cria e deleta um aluno do banco de dados
	defer DeletaAlunoMock()

	r := RoutesSetup()
	r.GET("/alunos", controllers.ShowAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	
	// fmt.Println(res.Body imprimir o corpo da resposta
	}
	
func TestSearch(t *testing.T) {
	database.Connection()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	
	r := RoutesSetup()
	r.GET("/alunos/search/:cpf", controllers.Search)
	
	req, _ := http.NewRequest("GET", "/alunos/search/12345678910", nil)
	res := httptest.NewRecorder()
	
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestSearchID(t *testing.T) {
	database.Connection()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	
	r := RoutesSetup()

	r.GET("/alunos/:id", controllers.GetAluno)

	path := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var alunoMock models.Aluno

	json.Unmarshal(res.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678910", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
}

func TestDelete(t *testing.T) {
	database.Connection()
	CriaAlunoMock()
	
	r := RoutesSetup()

	r.DELETE("alunos/:id", controllers.DeleteAluno)

	path := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

}

func TestUpdate(t *testing.T) {
	database.Connection()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	
	r := RoutesSetup()

	r.PATCH("/alunos/:id", controllers.EditarAluno)

	
	a := models.Aluno{
		Nome: "Aluno Teste",
		CPF:  "12345678911",
		RG:   "123456788",
	}
	
	alunoJson, _ := json.Marshal(a)
	
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(alunoJson))
	res := httptest.NewRecorder()
	
	r.ServeHTTP(res, req)

	var alunoMockAtt models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMockAtt)

	fmt.Println(alunoMockAtt.CPF)
}