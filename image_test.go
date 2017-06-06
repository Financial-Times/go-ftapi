package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    ftapi "github.com/Financial-Times/go-ftapi"
)

var test_image *ftapi.Image
var test_image_client *ftapi.Client

func TestImage(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("FT_API_KEY")
    log.Println("Using API key: ",key)
    test_image_client = &ftapi.Client{key, ""}

    result, err := test_image_client.ImageByUUID("ccca5db0-a7da-11e5-9700-2b669a5aeb83")

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal("http://www.ft.com/thing/ccca5db0-a7da-11e5-9700-2b669a5aeb83", result.ID)
    }

    test_image = result
}

