package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    ftapi "github.com/Financial-Times/go-ftapi"
)

var test_person *ftapi.Person
var test_person_client *ftapi.Client

func TestPerson(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("FT_API_KEY")
    log.Println("Using API key: ",key)
    test_person_client = ftapi.NewClient(key)

    result, err := test_person_client.PersonByUUID("535c8b74-b14d-3f32-9923-4dfdb4a97f5f")

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal("http://www.ft.com/thing/535c8b74-b14d-3f32-9923-4dfdb4a97f5f", result.ID)
        a.Equal("Suleyman Kerimov", result.PrefLabel)
    }

    test_person = result
}
