package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    ftapi "github.com/Financial-Times/go-ftapi"
)

var test_content_client *ftapi.Client

func TestContentRefsAnnotatedByUUID(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("FT_API_KEY")
    log.Println("Using API key: ",key)
    test_content_client = &ftapi.Client{key, ""}

    result, err := test_content_client.ContentRefsAnnotatedByUUID("9b40e89c-e87b-3d4f-b72c-2cf7511d2146",14) // News

    a.Nil(err)
    if a.NotNil(result) {
        a.Equal(14, len(result))
    }
}
