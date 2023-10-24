package article

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"rest-project/src/adapters"
	"time"
)

type UseCasesImpl struct {
	articleRepo adapters.Repo[adapters.Article]
	context     context.Context
}

func NewArticleUseCasesImpl(articleRepo adapters.Repo[adapters.Article], context context.Context) *UseCasesImpl {
	return &UseCasesImpl{articleRepo: articleRepo, context: context}
}

func (a *UseCasesImpl) GetAllArticleData() (*[]Article, error) {
	articles, err := a.articleRepo.GetAll(a.context)
	if err != nil {
		return nil, err
	}
	resultArticles := make([]Article, len(*articles))

	for i, article := range *articles {
		resultArticles[i] = Article{
			ID:        article.ID,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
			Title:     article.Title,
			Text:      article.Text,
		}
	}
	return &resultArticles, nil
}

func (a *UseCasesImpl) Create(i Input) (*Article, error) {
	result := &Article{
		ID:        adapters.UUIDGenerator(),
		Title:     i.Title,
		Text:      i.Text,
		CreatedAt: time.Now().UTC(),
	}

	err := a.articleRepo.Create(&adapters.Article{
		Title: result.Title,
		Text:  result.Text,
		ArticlePK: adapters.ArticlePK{
			ID: result.ID,
		},
		BaseModel: adapters.BaseModel{
			CreatedAt: time.Now().UTC(),
		},
	}, a.context)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *UseCasesImpl) Delete(id uuid.UUID) error {
	err := a.articleRepo.DeleteById(id, a.context)
	return err
}

func (a *UseCasesImpl) GetArticle(articleId uuid.UUID) (*Article, error) {
	art, err := a.articleRepo.GetById(articleId, a.context)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	return &Article{
		ID:        art.ID,
		CreatedAt: art.CreatedAt,
		UpdatedAt: art.UpdatedAt,
		Title:     art.Title,
		Text:      art.Text,
	}, nil
}

func (a *UseCasesImpl) UpdateArticle(articleId uuid.UUID, i Input) (*Article, error) {
	currentArticle, err := a.articleRepo.GetById(articleId, a.context)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	currentArticle.Title = i.Title
	currentArticle.Text = i.Text
	err = a.articleRepo.Update(currentArticle, a.context)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return &Article{
		ID:        currentArticle.ID,
		CreatedAt: currentArticle.CreatedAt,
		UpdatedAt: currentArticle.UpdatedAt,
		Title:     currentArticle.Title,
		Text:      currentArticle.Text,
	}, nil
}
