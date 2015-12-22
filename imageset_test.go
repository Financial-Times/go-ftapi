package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    "local/ftapi"
)

var test_imageset *ftapi.ImageSet
var test_imageset_client *ftapi.Client

func TestGetImageSet(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("FT_API_KEY")
    log.Println("Using API key: ",key)
    test_imageset_client = &ftapi.Client{key}

    result, err := test_imageset_client.GetImageSetByUuid("e1ee1f3a-8fbe-11e5-1582-a29c65546762")

    a.Nil(err)
    a.Equal("http://www.ft.com/thing/e1ee1f3a-8fbe-11e5-1582-a29c65546762", result.ID)

    test_imageset = result
}

func TestGetImageSetMembers(t *testing.T) {
    a := assert.New(t)

    images, err := test_imageset_client.GetImageSetMembers(test_imageset, 0)

    a.Nil(err)
    a.Equal(len(images), 1)
    a.Equal("http://www.ft.com/thing/e1ee1f3a-8fbe-11e5-8be4-3506bf20cc2b", images[0].ID)

}

