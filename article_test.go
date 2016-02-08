package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    ftapi "github.com/Financial-Times/go-ftapi"
)

var test_article *ftapi.Article
var test_article_client *ftapi.Client

func TestArticle(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("FT_API_KEY")
    log.Println("Using API key: ",key)
    test_article_client = &ftapi.Client{key}

    result, err := test_article_client.ArticleByUUID("98ca84ac-a7c3-11e5-955c-1e1d6de94879")

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
