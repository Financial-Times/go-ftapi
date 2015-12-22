package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    "local/ftapi"
)

var test_image *ftapi.Image
var test_image_client *ftapi.Client

func TestGetImage(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("FT_API_KEY")
    log.Println("Using API key: ",key)
    test_image_client = &ftapi.Client{key}

    result, err := test_image_client.GetImageByUuid("ccca5db0-a7da-11e5-9700-2b669a5aeb83")

    a.Nil(err)
    a.Equal("http://www.ft.com/thing/ccca5db0-a7da-11e5-9700-2b669a5aeb83", result.ID)

    test_image = result
}

