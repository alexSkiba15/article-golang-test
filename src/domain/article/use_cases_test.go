package article

import (
	"github.com/stretchr/testify/assert"
	"rest-project/src/adapters"
	"rest-project/src/fakers"
	"testing"
)

func TestUseCases(t *testing.T) {
	newArticleRepository := fakers.NewFakeMemoryRepository[adapters.Article]()
	useCases := NewArticleUseCasesImpl(newArticleRepository, nil)
	t.Run("CreateArticle", func(t *testing.T) {
		result, _ := useCases.Create(Input{Text: "test", Title: "test"})
		assert.Equal(t, result.Title, "test")
	})
	t.Run("UpdateArticle", func(t *testing.T) {
		result, _ := useCases.Create(Input{Text: "test", Title: "test"})
		newText := "test123123"
		article, _ := useCases.UpdateArticle(result.ID, Input{Text: newText, Title: "test"})
		assert.Equal(t, article.Text, newText)
	})
	t.Run("GetAllArticles", func(t *testing.T) {
		_, _ = useCases.Create(Input{Text: "test", Title: "test"})
		_, _ = useCases.Create(Input{Text: "test2", Title: "test2"})
		articles, _ := useCases.GetAllArticleData()
		articleLen := len(*articles)
		assert.Equal(t, articleLen, 4)
	})
	t.Run("Delete", func(t *testing.T) {
		result, _ := useCases.Create(Input{Text: "test", Title: "test"})
		article, err := useCases.GetArticle(result.ID)
		assert.Equal(t, err, nil)
		err = useCases.Delete(article.ID)
		assert.Equal(t, err, nil)
		_, err = useCases.GetArticle(article.ID)
		assert.NotEqual(t, err, nil)
	})
}
