package ftapi

import (
         "time"
         "log"
)

type Notification struct {
    APIURL string `json:"apiUrl"`
    ID     string `json:"id"`
    Type   string `json:"type"`
}

type Notifications struct {
    Links []struct {
        Href string `json:"href"`
        Rel  string `json:"rel"`
    } `json:"links"`
    Notifications []Notification `json:"notifications"`
    RequestURL string `json:"requestUrl"`
}

var ZuluTime = time.FixedZone("Z", 0)

func (c *Client) RawNotificationsSince(since time.Time) (*Notifications, error) {
    since = since.In(ZuluTime).Truncate(time.Second)

    rfcSince, err := since.MarshalText()
    if err != nil {
        return nil, err
    }

	url := "https://api.ft.com/content/notifications/?since=" + string(rfcSince)
	result := &Notifications{}
	err = c.jsonAtURL(url, result)
	return result, err
}

func (c *Client) NextRawNotifications(after *Notifications) (*Notifications, error) {
    if after.Links == nil || len(after.Links)==0 {
        return nil, nil
    }

    var sinceURL string

    for _,item := range after.Links {
        if item.Rel == "next" {
            sinceURL = item.Href
        }
    }

    if sinceURL == "" {
        return nil, nil
    }

    log.Printf("%s -> %s", after.RequestURL, sinceURL)

	result := &Notifications{}
	err := c.jsonAtURL(sinceURL, result)
	return result, err
}


func (c *Client) Notifications(duration time.Duration) ([]Notification, error) {
	since := time.Now().Add(-duration)
    return c.NotificationsSince(since)
}

func (c *Client) NotificationsSince(since time.Time) ([]Notification, error) {
    result, err := c.RawNotificationsSince(since)
    if result == nil {
        return nil, err
    }
    return result.Notifications, err
}

func (c *Client) AllNotifications(duration time.Duration) ([]Notification, error) {
	since := time.Now().Add(-duration)
    return c.AllNotificationsSince(since)
}

func (c *Client) AllNotificationsSince(since time.Time) ([]Notification, error) {
    result, err := c.RawNotificationsSince(since)
    if err != nil {
        return nil, err
    }

    combined := result.Notifications
    log.Printf("First page had %d notifications", len(combined))

    for {
        result, err = c.NextRawNotifications(result)
        if err != nil || result.Notifications == nil {
            log.Printf("Error after %d notifications", len(combined))
            return combined, err
        }
        log.Printf("Got %d more notifications", len(result.Notifications))
        combined = append(combined, result.Notifications...)
        log.Printf("Now have %d notifications", len(combined))
        if result.Notifications == nil || len(result.Notifications) < 200 {
            log.Printf("No more notifications.")
            break
        }
    }

    return combined, nil
}

func (c *Client) Listen(start time.Time, sleep time.Duration) (chan Notification) {
    ch := make(chan Notification)

    go func() {
        result, err := c.RawNotificationsSince(start)
        if err != nil {
            log.Println("Error getting notifications:",err)
        }

        for {
            if result.Notifications == nil {
                log.Printf("Notifications was nil")
            } else {
                for _, notification := range result.Notifications {
                    log.Printf("<-%s", notification.ID)
                    ch <- notification
                }
            }

            time.Sleep(sleep)

            result, err = c.NextRawNotifications(result)
            if err != nil {
                log.Println("Error getting notifications:",err)
            }
        }
    }()

    return ch
}
