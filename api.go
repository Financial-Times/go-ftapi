package ftapi

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "fmt"
)


type Client struct {
    Key string
}

func (c *Client) jsonAtURL(url string, obj interface{}) (error) {
    return c.jsonAtURLWithCookie(url, obj, nil)
}

func (c *Client) jsonAtURLWithCookie(url string, obj interface{}, cookie *http.Cookie) (error) {

    log.Println("Getting",url)

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Println("Failed to build a request for ",url)
        return err
    }

    req.Header.Add("X-API-Key", c.Key)

    if cookie != nil {
        req.AddCookie(cookie)
    }

    resp, err := client.Do(req)
    if err != nil {
        log.Println("Failed to execute request for ",url)
        return err
    }

    defer resp.Body.Close()

    if err != nil {
        log.Println("Failed to get ",url)
        return err
    }

    if (resp.StatusCode != 200) {
        return fmt.Errorf("%s %s", resp.Status, http.StatusText(resp.StatusCode))
    }

    if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
        read, err2 := ioutil.ReadAll(resp.Body)
        if err2 != nil {
            log.Println("Failed to decode JSON and encountered an error trying to read the body")
        } else {
            log.Println("Failed to decode JSON from",read)
        }
        return err
    }

    log.Println("Got",obj)
    return nil
}

