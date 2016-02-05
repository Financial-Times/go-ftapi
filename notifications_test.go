package ftapi_test

import (
    "testing"
    "log"
    "os"
    "github.com/stretchr/testify/assert"
    "time"
    ftapi "github.com/Financial-Times/go-ftapi"
)

var test_notifications *ftapi.Notifications
var test_notifications_client *ftapi.Client

func TestRawNotifications(t *testing.T) {
    a := assert.New(t)

    key := os.Getenv("FT_API_KEY")
    log.Println("Using API key: ",key)
    test_notifications_client = &ftapi.Client{key}

    result, err := test_notifications_client.RawNotificationsSince(time.Now().Add(-96 * time.Hour))

    // assuming >200 things have been published in the last 96 hours //

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal(200, len(result.Notifications))
    }

    test_notifications = result
}

func TestNextRawNotifications(t *testing.T) {
    a := assert.New(t)

    result, err := test_notifications_client.NextRawNotifications(test_notifications)

    // assuming >400 things have been published in the last 96 hours //

    a.Nil(err)

    if a.NotNil(result) {
        a.Equal(200, len(result.Notifications))
    }
}

func TestAllNotifications(t *testing.T) {
    a := assert.New(t)

    result, err := test_notifications_client.AllNotifications(96 * time.Hour)

    // assuming >400 things have been published in the last 96 hours //

    a.Nil(err)

    if a.NotNil(result) {
        log.Printf("%d notifications in the last 96 hours.", len(result))
        a.True(len(result)>400, "Should be more than 400 notifications.")
    }
}
