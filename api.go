package ftapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	Key  string
	Auth string
        BaseURL *url.URL
}

func NewClient(key string) *Client {
	return &Client{key, "", nil}
}

func NewClientSpecial(key string, auth string, baseURL string) (*Client, error) {
	if baseURL == "" {
		baseURL = "https://api.ft.com/"
	}
	b, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &Client{key, auth, b}, nil
}

func (c *Client) FromURL(relurl string, obj interface{}) (*[]byte, error) {
	r, err := url.Parse(relurl)
	if err != nil {
		return nil, err
	}
	return c.do(r, nil, nil, obj)
}

func (c *Client) FromURLWithBody(relurl string, body []byte, obj interface{}) (*[]byte, error) {
	r, err := url.Parse(relurl)
	if err != nil {
		return nil, err
	}
	return c.do(r, body, nil, obj)
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

func (c *Client) FromURLWithCookie(relurl string, obj interface{}, cookie *http.Cookie) (*[]byte, error) {
	r, err := url.Parse(relurl)
	if err != nil {
		return nil, err
	}
	return c.do(r, nil, cookie, obj)
}


func (c *Client) do(relurl *url.URL, body []byte, cookie *http.Cookie, obj interface{}) (*[]byte, error) {
	return c.doLimitedTimes(relurl, body, cookie, obj, 5)
}

func (c *Client) doLimitedTimes(relurl *url.URL, body []byte, cookie *http.Cookie, obj interface{}, times int) (*[]byte, error) {
	if times == 0 {
		return nil, fmt.Errorf("Too many redirects: %s", relurl)
	}

	client := &http.Client{
	    CheckRedirect: func(req *http.Request, via []*http.Request) error {
		// Don't let golang follow redirects, as it won't add the X-Api-Key to the second request :(
        	return http.ErrUseLastResponse
	    },
	}

	var req *http.Request
	var err error
	var absurl string

        if c.BaseURL == nil {
		if b, err := url.Parse("https://api.ft.com"); err == nil {
			c.BaseURL = b
		}
	}

	absurl = c.BaseURL.ResolveReference(relurl).String()

	if body == nil {
		req, err = http.NewRequest("GET", absurl, nil)
		if err != nil {
			log.Println("Failed to build a GET request for ", absurl)
			return nil, err
		}
	} else {
		req, err = http.NewRequest("POST", absurl, bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		log.Printf("POST %s\nbody: %s\n", absurl, string(body))
		if err != nil {
			log.Println("Failed to build a POST request for ", absurl)
			log.Println(body)
			return nil, err
		}
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
		log.Printf("Failed to execute request for %s:%s\n", absurl, err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Failed to get %s:%s\n", absurl, err.Error())
		return nil, err
	}

	log.Printf("%d %s",resp.StatusCode,absurl)

	switch resp.StatusCode {
	case 301, 302, 303, 307, 308:
		l, err := resp.Location()
		if err != nil {
			return nil, err
		}
		return c.doLimitedTimes(l, body, cookie, obj, times-1)
	case 200:
		rbody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Failed to read body from ", absurl)
		}

		if err := json.Unmarshal(rbody, obj); err != nil {
			log.Println("Failed to decode JSON from ", absurl)
			return nil, err
		}

		return &rbody, nil
	default:
		return nil, fmt.Errorf("%s %s", resp.Status, http.StatusText(resp.StatusCode))
	}
}
