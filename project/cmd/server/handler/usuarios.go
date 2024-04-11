package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cesar-oliveira-silva/goweb-aula-4-exec-manha.git/project/internal/usuarios"
	"github.com/cesar-oliveira-silva/goweb-aula-4-exec-manha.git/project/pkg/web"
	"github.com/gin-gonic/gin"
)

type CreateRequestDto struct {
	Nome        string `json:"nome"`
	Sobrenome   string `json:"sobrenome"`
	Email       string `json:"email"`
	Idade       int    `json:"idade"`
	Altura      int    `json:"altura"`
	Ativo       bool   `json:"ativo"`
	DataCriacao string `json:"dataCriacao"`
}

type UpdateRequestDto struct {
	Nome        string `json:"nome"`
	Sobrenome   string `json:"sobrenome"`
	Email       string `json:"email"`
	Idade       int    `json:"idade"`
	Altura      int    `json:"altura"`
	Ativo       bool   `json:"ativo"`
	DataCriacao string `json:"dataCriacao"`
}

type UpdateNameRequestDto struct {
	Nome string `json:"nome"`
}

type ServiceHandler struct {
	service usuarios.Service
}

func NewUser(p usuarios.Service) *ServiceHandler {
	return &ServiceHandler{
		service: p,
	}
}

func (c *ServiceHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		p, err := c.service.GetAll()
		if err != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{
			// 	"error": err.Error(),
			// })
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro ao recuperar usuarios"))
			return
		}

		if len(p) == 0 {
			//ctx.Status(http.StatusNoContent)
			ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, "Nenhum usuário cadastrado"))
			return
		}

		//ctx.JSON(http.StatusOK, p)
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))

	}
}

func (c *ServiceHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req CreateRequestDto
		if err := ctx.Bind(&req); err != nil {
			//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(),})
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		// quando chamamos a service, os dados já estarão tratados
		fmt.Println(req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.DataCriacao)
		p, err := c.service.Store(req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.DataCriacao)
		if err != nil {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		//ctx.JSON(http.StatusCreated, p)
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, p, ""))
	}
}
func (c *ServiceHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// forma de se fazermos uma conversão de alfa númerico para inteiro
		// strconv.Atoi(ctx.Param("id"))
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)

		if err != nil {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro: Id invalido"))
			return
		}

		var req UpdateRequestDto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		if req.Nome == "" {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": "o nome do Usuario é obrigatório"})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro: o nome do Usuario é obrigatório"))
			return
		}

		if req.Sobrenome == "" {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": "O sobrenome do usuario é obrigatório"})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro: o sobernome do Usuario é obrigatório"))
			return
		}

		if req.Email == "" {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": "o email do usuario é obrigatório"})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro: o email do Usuario é obrigatório"))
			return
		}

		if req.Idade == 0 {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": "A idade do usuário é obrigatória"})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro: A idade do Usuario é obrigatória"))
			return
		}

		if req.Altura == 0 {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": "A altura do usuário é obrigatória"})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro: A altura do Usuario é obrigatória"))
			return
		}

		if !req.Ativo {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": "O usuário deve ser ativado"})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro: o status do Usuario é obrigatório e deve ser ativo(true)"))
			return
		}

		if req.DataCriacao == "" {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": "O deve ter uma data de criacao"})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro: a Data de criacao do Usuario é obrigatória"))
			return
		}

		p, err := c.service.Update(id, req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.DataCriacao)
		if err != nil {
			//ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		//ctx.JSON(http.StatusOK, p)
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

func (c *ServiceHandler) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
		if err != nil {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID inválido"))

			return
		}
		var req UpdateNameRequestDto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}
		if req.Nome == "" {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": "O nome do usuario é obrigatório"})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "O nome do usuario é obrigatório"))
			return
		}
		p, err := c.service.UpdateName(id, req.Nome)
		if err != nil {
			//ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ServiceHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
		if err != nil {
			//ctx.JSON(400, gin.H{"error": "invalid ID"})
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID inválido"))
			return
		}
		err = c.service.Delete(id)
		if err != nil {
			//ctx.JSON(404, gin.H{"error": err.Error()})
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		// poderiamos usar o http.StatusNoContent -> 204 tbm!
		//ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("O usuario %d foi removido", id)})
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, fmt.Sprintf("O usuario %d foi removido", id), ""))
	}
}
