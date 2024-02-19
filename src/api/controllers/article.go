package controllers

import (
	"fmt"
	"net/http"
	"rest-project/src/domain/article"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ArticleHandler struct {
	ArticleUseCases article.UseCases
}

func (h *ArticleHandler) GetAll(c *gin.Context) {
	data, _ := h.ArticleUseCases.GetAllArticleData(c)
	c.JSON(http.StatusOK, data)
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var newArticleInput article.Input

	if err := c.BindJSON(&newArticleInput); err != nil {
		return
	}
	data, _ := h.ArticleUseCases.Create(c, newArticleInput)
	c.JSON(http.StatusCreated, data)
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	articleID := c.Param("articleID")
	parsedUUID, err := uuid.Parse(articleID)
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return
	}

	err = h.ArticleUseCases.Delete(c, parsedUUID)
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

	responseArticle, err := h.ArticleUseCases.GetArticle(c, parsedUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(http.StatusOK, responseArticle)
	}
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	articleID := c.Param("articleID")
	parsedUUID, errUUID := uuid.Parse(articleID)
	var newArticleInput article.Input

	if errUUID != nil {
		fmt.Printf("Error parsing UUID: %v\n", errUUID)
		return
	}
	if errBind := c.BindJSON(&newArticleInput); errBind != nil {
		return
	}

	responseArticle, errArticle := h.ArticleUseCases.UpdateArticle(c, parsedUUID, newArticleInput)
	if errArticle != nil {
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
