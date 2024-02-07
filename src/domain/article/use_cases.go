package article

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"rest-project/src/adapters"
	_ "rest-project/src/adapters/models"
	"rest-project/src/domain/entities"
	"time"
)

type UseCasesImpl struct {
	uow     adapters.UnitOfWork
	context context.Context
}

func NewArticleUseCasesImpl(uow adapters.UnitOfWork, context context.Context) *UseCasesImpl {
	return &UseCasesImpl{uow: uow, context: context}
}

func (a *UseCasesImpl) GetAllArticleData() (*[]entities.Article, error) {
	articles, err := a.uow.ArticleRepository().GetAll(a.context)
	if err != nil {
		return nil, err
	}
	resultArticles := make([]entities.Article, len(*articles))

	for i, article := range *articles {
		resultArticles[i] = entities.Article{
			Base:  *entities.NewBase(article.ID, article.CreatedAt, &article.UpdatedAt),
			Title: article.Title,
			Text:  article.Text,
		}
	}
	return &resultArticles, nil
}

func (a *UseCasesImpl) Create(i Input) (*entities.Article, error) {
	result := &entities.Article{
		Base:  *entities.NewBase(adapters.UUIDGenerator(), time.Now().UTC(), nil),
		Title: i.Title,
		Text:  i.Text,
	}
	err := a.uow.ArticleRepository().Create(result, a.context)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *UseCasesImpl) Delete(id uuid.UUID) error {
	err := a.uow.ArticleRepository().DeleteById(id, a.context)
	return err
}

func (a *UseCasesImpl) GetArticle(articleId uuid.UUID) (*entities.Article, error) {
	art, err := a.uow.ArticleRepository().GetById(articleId, a.context)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	return art, nil
}

func (a *UseCasesImpl) UpdateArticle(articleId uuid.UUID, i Input) (*entities.Article, error) {
	currentArticle, err := a.uow.ArticleRepository().GetById(articleId, a.context)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	currentArticle.Title = i.Title
	currentArticle.Text = i.Text
	err = a.uow.ArticleRepository().Update(currentArticle, a.context)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return &entities.Article{
		Base:  *entities.NewBase(currentArticle.ID, currentArticle.CreatedAt, &currentArticle.UpdatedAt),
		Title: currentArticle.Title,
		Text:  currentArticle.Text,
	}, nil
}
