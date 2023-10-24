package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rest-project/src/domain/article"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ArticleHandler struct {
	ArticleUseCases article.UseCases
}

func (h *ArticleHandler) GetAll(c *gin.Context) {
	//data, _ := h.ArticleUseCases.GetAllArticleData()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "asdadasdasdasda:}"})
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var newArticleInput article.Input

	if err := c.BindJSON(&newArticleInput); err != nil {
		return
	}
	data, _ := h.ArticleUseCases.Create(newArticleInput)
	c.JSON(http.StatusCreated, data)
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	articleID := c.Param("articleID")
	parsedUUID, err := uuid.Parse(articleID)
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return
	}

	err = h.ArticleUseCases.Delete(parsedUUID)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *ArticleHandler) GetArticleById(c *gin.Context) {
	articleID := c.Param("articleID")
	parsedUUID, err := uuid.Parse(articleID)
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return
	}

	responseArticle, err := h.ArticleUseCases.GetArticle(parsedUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(http.StatusOK, responseArticle)
	}
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	articleID := c.Param("articleID")
	parsedUUID, err := uuid.Parse(articleID)
	var newArticleInput article.Input

	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return
	}
	if err := c.BindJSON(&newArticleInput); err != nil {
		return
	}

	responseArticle, err := h.ArticleUseCases.UpdateArticle(parsedUUID, newArticleInput)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(http.StatusOK, responseArticle)
	}
}

func NewArticleHandler(articleUC article.UseCases) ArticleHandler {
	handler := &ArticleHandler{
		ArticleUseCases: articleUC,
	}
	return *handler
}
