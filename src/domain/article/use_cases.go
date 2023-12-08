package article

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"rest-project/src/adapters"
	"rest-project/src/adapters/models"
	"rest-project/src/domain/entities"
	"time"
)

type UseCasesImpl struct {
	articleRepo adapters.ArticleRepository[adapters.Article]
	context     context.Context
}

func NewArticleUseCasesImpl(articleRepo adapters.Repo[adapters.Article], context context.Context) *UseCasesImpl {
	return &UseCasesImpl{articleRepo: articleRepo, context: context}
}

func (a *UseCasesImpl) GetAllArticleData() (*[]entities.Article, error) {
	articles, err := a.articleRepo.GetAll(a.context)
	if err != nil {
		return nil, err
	}
	resultArticles := make([]entities.Article, len(*articles))

	for i, article := range *articles {
		resultArticles[i] = entities.Article{
			Base:  entities.NewBase(article.ID, article.CreatedAt, &article.UpdatedAt),
			Title: article.Title,
			Text:  article.Text,
		}
	}
	return &resultArticles, nil
}

func (a *UseCasesImpl) Create(i Input) (*entities.Article, error) {
	result := &entities.Article{
		Base:  entities.NewBase(adapters.UUIDGenerator(), time.Now().UTC(), nil),
		Title: i.Title,
		Text:  i.Text,
	}

	err := a.articleRepo.Create(&adapters.Article{
		Title: result.Title,
		Text:  result.Text,
		ArticlePK: adapters.ArticlePK{
			ID: result.ID,
		},
		BaseModel: models.BaseModel{
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

func (a *UseCasesImpl) GetArticle(articleId uuid.UUID) (*entities.Article, error) {
	art, err := a.articleRepo.GetById(articleId, a.context)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	return &entities.Article{
		Base:  entities.NewBase(art.ID, art.CreatedAt, &art.UpdatedAt),
		Title: art.Title,
		Text:  art.Text,
	}, nil
}

func (a *UseCasesImpl) UpdateArticle(articleId uuid.UUID, i Input) (*entities.Article, error) {
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
	return &entities.Article{
		Base:  entities.NewBase(currentArticle.ID, currentArticle.CreatedAt, &currentArticle.UpdatedAt),
		Title: currentArticle.Title,
		Text:  currentArticle.Text,
	}, nil
}
