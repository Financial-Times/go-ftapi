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

func (c *Client) getJsonAtUrl(url string, obj interface{}) (error) {
    log.Println("Getting",url)

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Println("Failed to build a request for ",url)
        return err
    }

    req.Header.Add("X-Api-Key", c.Key)
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Failed to execute request for ",url)
        return err
    }

    defer resp.Body.Close()

    if err != nil {
        log.Println("Failed to get ",url)
        return err
    } else {
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
        read, err2 := ioutil.ReadAll(resp.Body)
        if err2 != nil {
            log.Println("Decoded JSON but encountered an error trying to read the body")
        } else {
            log.Println("Decoded JSON from",read)
        }
        log.Println("Got",obj)
        return nil
    }
}

