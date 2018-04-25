package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    ftapi "github.com/Financial-Times/go-ftapi"
)

var test_recommendations_client *ftapi.Client

func TestContextualRecommendationsByUUID(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("RR_API_KEY")
    log.Println("Using API key: ",key)
    test_recommendations_client = ftapi.NewClient(key)

    result, err := test_recommendations_client.ContextualRecommendationsByUUID("24b6f48a-c675-11e5-b3b1-7b2481276e45", 10, 7)

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal(10, len(result))
    }
}

func TestBehaviouralRecommendationsByUUID(t *testing.T) {
    a := assert.New(t)

    userid := os.Getenv("FT_USERID")

    if userid == "" {
        log.Println("WARNING: Didn't test behavioural recommendations as no user id was available.")
    } else {
        log.Println("Using user id: ",userid)

        result, err := test_recommendations_client.BehaviouralRecommendationsByUUID("24b6f48a-c675-11e5-b3b1-7b2481276e45", userid, 10, 7)

        a.Nil(err)

        if a.NotNil(result) {
            a.Equal(10, len(result))
        }
    }
}

func TestBehaviouralRecommendations(t *testing.T) {
    a := assert.New(t)

    userid := os.Getenv("FT_USERID")

    if userid == "" {
        log.Println("WARNING: Didn't test behavioural recommendations as no user id was available.")
    } else {
        log.Println("Using user id: ",userid)

        result, err := test_recommendations_client.BehaviouralRecommendations(userid, 10, 7)

        a.Nil(err)

        if a.NotNil(result) {
            a.Equal(10, len(result))
        }
    }
}

func TestPopularRecommendations(t *testing.T) {
    a := assert.New(t)

    result, err := test_recommendations_client.PopularRecommendations(10, 7)

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal(10, len(result))
    }
}
