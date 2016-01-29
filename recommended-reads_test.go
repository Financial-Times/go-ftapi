package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    ftapi "github.com/Financial-Times/go-ftapi"
)

var test_recommendations_client *ftapi.Client

func TestGetContextualRecommendationsByUuid(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("RR_API_KEY")
    log.Println("Using API key: ",key)
    test_recommendations_client = &ftapi.Client{key}

    result, err := test_recommendations_client.GetContextualRecommendationsByUuid("24b6f48a-c675-11e5-b3b1-7b2481276e45")

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal(10, len(result.Articles))
    }
}

func TestGetBehaviouralRecommendationsByUuid(t *testing.T) {
    a := assert.New(t)

    userid := os.Getenv("FT_USERID")
    log.Println("Using user id: ",userid)

    result, err := test_recommendations_client.GetBehaviouralRecommendationsByUuid("24b6f48a-c675-11e5-b3b1-7b2481276e45", userid)

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal(10, len(result.Articles))
    }
}

func TestGetBehaviouralRecommendations(t *testing.T) {
    a := assert.New(t)

    userid := os.Getenv("FT_USERID")
    log.Println("Using user id: ",userid)

    result, err := test_recommendations_client.GetBehaviouralRecommendations(userid)

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal(10, len(result.Articles))
    }
}

func TestGetPopularRecommendations(t *testing.T) {
    a := assert.New(t)

    result, err := test_recommendations_client.GetPopularRecommendations()

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal(10, len(result.Articles))
    }
}
