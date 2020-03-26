package ftapi_test

import (
	"log"
	"os"
	"testing"
	"time"

	ftapi "github.com/Financial-Times/go-ftapi"
	"github.com/stretchr/testify/assert"
)

var test_article *ftapi.Article
var test_article_client *ftapi.Client

func TestArticle(t *testing.T) {
	a := assert.New(t)

	key := os.Getenv("FT_API_KEY")
	log.Println("Using API key: ", key)
	test_article_client = ftapi.NewClient(key)

	result, err := test_article_client.ArticleByUUID("98ca84ac-a7c3-11e5-955c-1e1d6de94879")

	a.Nil(err)

	if a.NotNil(result) {
		a.Equal("http://www.ft.com/thing/98ca84ac-a7c3-11e5-955c-1e1d6de94879", result.ID)
	}

	test_article = result
}

func TestEnrichedArticle(t *testing.T) {
	a := assert.New(t)

	key := os.Getenv("FT_API_KEY")
	log.Println("Using API key: ", key)
	test_article_client = ftapi.NewClient(key)

	// "Sepp Blatter and Michel Platini receive 8-year ban from football"
	result, err := test_article_client.EnrichedArticle("http://api.ft.com/enrichedcontent/98ca84ac-a7c3-11e5-955c-1e1d6de94879")

	a.Nil(err)

	if a.NotNil(result) {
		a.Equal("http://www.ft.com/thing/98ca84ac-a7c3-11e5-955c-1e1d6de94879", result.ID)
		a.Equal(result.PublishedDate.Equal(time.Date(2015, 12, 21, 17, 30, 14, 0, time.UTC)), true)
		a.Equal(result.FirstPublishedDate.Equal(time.Date(2015, 12, 21, 9, 49, 40, 0, time.UTC)), true)
	}

	test_article = result
}

func TestEnrichedArticleWithoutFirstPublishedDate(t *testing.T) {
	a := assert.New(t)

	key := os.Getenv("FT_API_KEY")
	log.Println("Using API key: ", key)
	test_article_client = ftapi.NewClient(key)

	// "Fifa: The fall of the house of Blatter"
	// was only published once so has no firstPublishedDate
	result, err := test_article_client.EnrichedArticle("http://api.ft.com/enrichedcontent/64573fc6-8df6-11e5-a549-b89a1dfede9b")

	a.Nil(err)

	if a.NotNil(result) {
		a.Equal("http://www.ft.com/thing/64573fc6-8df6-11e5-a549-b89a1dfede9b", result.ID)
		a.Equal(result.PublishedDate.Equal(time.Date(2015, 11, 22, 18, 13, 23, 0, time.UTC)), true)
		a.Equal(result.FirstPublishedDate.Equal(result.PublishedDate), true)
	}

	test_article = result
}

func TestInternalArticle(t *testing.T) {
	a := assert.New(t)

	key := os.Getenv("FT_API_KEY")
	log.Println("Using API key: ", key)
	test_article_client = ftapi.NewClient(key)

	result, err := test_article_client.InternalArticle("http://api.ft.com/interalcontent/98ca84ac-a7c3-11e5-955c-1e1d6de94879")

	a.Nil(err)

	if a.NotNil(result) {
		a.Equal("http://www.ft.com/thing/98ca84ac-a7c3-11e5-955c-1e1d6de94879", result.ID)
	}

	test_article = result
}

func TestMainImageSet(t *testing.T) {
	a := assert.New(t)

	mainImageSet, err := test_article_client.MainImageSet(test_article)

	a.Nil(err)

	if a.NotNil(mainImageSet) {
		a.Equal("http://www.ft.com/thing/ccca5db0-a7da-11e5-0966-bcfc402e40ca", mainImageSet.ID)
	}
}
