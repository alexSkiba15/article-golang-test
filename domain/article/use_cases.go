package article

import (
	"context"
	"fmt"
	"time"

	"github.com/alexSkiba15/article-golang-test/adapters"
	_ "github.com/alexSkiba15/article-golang-test/adapters/models"
	"github.com/alexSkiba15/article-golang-test/domain/entities"

	"github.com/google/uuid"
)

type UseCasesImpl struct {
	articleRepository adapters.ArticleRepository
}

func NewArticleUseCasesImpl(rep adapters.ArticleRepository) *UseCasesImpl {
	return &UseCasesImpl{articleRepository: rep}
}

func (a *UseCasesImpl) GetAllArticleData(ctx context.Context) (*[]entities.Article, error) {
	articles, err := a.articleRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	resultArticles := make([]entities.Article, len(*articles))

	for i, article := range *articles {
		resultArticles[i] = entities.Article{
			Base:  *entities.NewBase(article.ID, article.CreatedAt, article.UpdatedAt),
			Title: article.Title,
			Text:  article.Text,
		}
	}
	return &resultArticles, nil
}

func (a *UseCasesImpl) Create(ctx context.Context, i Input) (*entities.Article, error) {
	fmt.Printf("Creating article with title: %v and text: %v\n", i.Title, i.Text)
	result := &entities.Article{
		Base:  *entities.NewBase(adapters.UUIDGenerator(), time.Now().UTC(), time.Now().UTC()),
		Title: i.Title,
		Text:  i.Text,
	}
	fmt.Println(result)
	err := a.articleRepository.Create(ctx, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *UseCasesImpl) Delete(ctx context.Context, id uuid.UUID) error {
	err := a.articleRepository.DeleteById(ctx, id)
	return err
}

func (a *UseCasesImpl) GetArticle(ctx context.Context, articleId uuid.UUID) (*entities.Article, error) {
	art, err := a.articleRepository.GetById(ctx, articleId)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	return art, nil
}

func (a *UseCasesImpl) UpdateArticle(ctx context.Context, articleId uuid.UUID, i Input) (*entities.Article, error) {
	currentArticle, err := a.articleRepository.GetById(ctx, articleId)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	currentArticle.Title = i.Title
	currentArticle.Text = i.Text
	err = a.articleRepository.Update(ctx, currentArticle)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	return &entities.Article{
		Base:  *entities.NewBase(currentArticle.ID, currentArticle.CreatedAt, currentArticle.UpdatedAt),
		Title: currentArticle.Title,
		Text:  currentArticle.Text,
	}, nil
}
