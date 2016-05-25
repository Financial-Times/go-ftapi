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
    Auth string
}

func (c *Client) FromURL(url string, obj interface{}) (*[]byte, error) {
    return c.FromURLWithCookie(url, obj, nil)
}

func (c *Client) FromPath(path string, obj interface{}) (*[]byte, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(data, &obj); err != nil {
        return nil, err
    }

    return &data, nil
}

func (c *Client) FromURLWithCookie(url string, obj interface{}, cookie *http.Cookie) (*[]byte, error) {
    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Println("Failed to build a request for ",url)
        return nil, err
    }

    req.Header.Add("X-API-Key", c.Key)

    if c.Auth != "" {
        req.Header.Add("Authorization", c.Auth)
    }

    if cookie != nil {
        req.AddCookie(cookie)
    }

    resp, err := client.Do(req)
    if err != nil {
        log.Println("Failed to execute request for ",url)
        return nil, err
    }

    defer resp.Body.Close()

    if err != nil {
        log.Println("Failed to get ",url)
        return nil, err
    }

    if (resp.StatusCode != 200) {
        return nil, fmt.Errorf("%s %s", resp.Status, http.StatusText(resp.StatusCode))
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Failed to read body from ",url)
    }

    if err := json.Unmarshal(body, obj); err != nil {
        log.Println("Failed to decode JSON from ",url)
        return nil, err
    }

    return &body, nil
}
