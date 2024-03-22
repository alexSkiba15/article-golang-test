package controllers

import (
	"fmt"
	"rest-project/domain/article"

	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ArticleHandler struct {
	ArticleUseCases article.UseCases
}

func (h *ArticleHandler) GetAll(c *fiber.Ctx) error {
	data, err := h.ArticleUseCases.GetAllArticleData(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *ArticleHandler) Create(c *fiber.Ctx) error {
	var newArticleInput article.Input

	if err := c.BodyParser(&newArticleInput); err != nil {
		return nil
	}
	data, _ := h.ArticleUseCases.Create(c.Context(), newArticleInput)
	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *ArticleHandler) Delete(c *fiber.Ctx) error {
	articleID := c.Params("articleID")
	parsedUUID, err := uuid.Parse(articleID)
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return nil
	}

	err = h.ArticleUseCases.Delete(c.Context(), parsedUUID)
	if err != nil {
		fmt.Println(err)
	}
	return c.Status(fiber.StatusNoContent).JSON(nil)
}

func (h *ArticleHandler) GetArticleById(c *fiber.Ctx) error {
	articleID := c.Params("articleID")
	parsedUUID, err := uuid.Parse(articleID)
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return nil
	}

	responseArticle, err := h.ArticleUseCases.GetArticle(c.Context(), parsedUUID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(nil)
	} else {
		return c.Status(fiber.StatusOK).JSON(responseArticle)
	}
}

func (h *ArticleHandler) UpdateArticle(c *fiber.Ctx) error {
	articleID := c.Params("articleID")
	parsedUUID, errUUID := uuid.Parse(articleID)
	var newArticleInput article.Input

	if errUUID != nil {
		fmt.Printf("Error parsing UUID: %v\n", errUUID)
		return nil
	}
	if errBind := c.BodyParser(&newArticleInput); errBind != nil {
		return nil
	}

	responseArticle, errArticle := h.ArticleUseCases.UpdateArticle(c.Context(), parsedUUID, newArticleInput)
	if errArticle != nil {
		return c.Status(fiber.StatusNotFound).JSON(nil)
	} else {
		return c.Status(fiber.StatusOK).JSON(responseArticle)
	}
}

func NewArticleHandler(articleUC article.UseCases) ArticleHandler {
	handler := &ArticleHandler{
		ArticleUseCases: articleUC,
	}
	return *handler
}
