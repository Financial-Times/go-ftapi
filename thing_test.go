package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    ftapi "github.com/Financial-Times/go-ftapi"
)

var test_thing *ftapi.Thing
var test_thing_client *ftapi.Client
var test_test_thing_client *ftapi.Client

func TestThing(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("FT_API_KEY")
    log.Println("Using API key: ",key)
    test_thing_client = ftapi.NewClient(key)

    result, err := test_thing_client.ThingByUUID("b4bd218e-c65c-362c-82ff-75ad0c607042")

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal("http://www.ft.com/thing/b4bd218e-c65c-362c-82ff-75ad0c607042", result.ID)
        a.Equal("Dagestan State University", result.PrefLabel)
    }

    test_thing = result
}

func TestTestThing(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("FT_TEST_API_KEY")
    log.Println("Using API key: ",key)

    test_test_thing_client, err := ftapi.NewClientSpecial(key,"","http://test.api.ft.com/")

    a.Nil(err)

    result, err := test_test_thing_client.ThingByUUID("b4bd218e-c65c-362c-82ff-75ad0c607042")

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal("http://www.ft.com/thing/b4bd218e-c65c-362c-82ff-75ad0c607042", result.ID)
        a.Equal("Dagestan State University", result.PrefLabel)
    }
}
