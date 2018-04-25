package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    ftapi "github.com/Financial-Times/go-ftapi"
)

var test_suggest_article *ftapi.Article
var test_suggest_client *ftapi.Client

func TestSuggest(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("FT_API_KEY")
    log.Println("Using API key: ",key)
    test_suggest_client = ftapi.NewClient(key)

    art := &ftapi.Article{
        BodyXML: "Test suggestion body which mentions the Financial Times",
    }

    result, err := test_suggest_client.Suggest(art)

    a.Nil(err)

    if a.NotNil(result) {
        var found bool
        for _, s := range result.Suggestions {
            if s.ID == "http://www.ft.com/thing/73cc33b5-d0cb-3815-8347-bc49e1ddbd5c" {
                found = true;
                break;
            }
        }
        a.Equal(found, true)
    }
}
