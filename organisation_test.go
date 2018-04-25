package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    ftapi "github.com/Financial-Times/go-ftapi"
)

var test_organisation *ftapi.Organisation
var test_organisation_client *ftapi.Client

func TestOrganisation(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("FT_API_KEY")
    log.Println("Using API key: ",key)
    test_organisation_client = ftapi.NewClient(key)

    result, err := test_organisation_client.OrganisationByUUID("d90be6d9-e941-3da4-a28d-e5339f7f896f")

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal("http://www.ft.com/thing/d90be6d9-e941-3da4-a28d-e5339f7f896f", result.ID)
        a.Equal("Particle Programmatica, Inc.", result.PrefLabel)
    }

    test_organisation = result
}
