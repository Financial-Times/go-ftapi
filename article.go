package ftapi

type Article struct {
    BodyXML   string     `json:"bodyXML"`
    Brands    []string   `json:"brands"`
    Byline    string     `json:"byline"`
    ID        string     `json:"id"`
    MainImage struct {
        ID string        `json:"id"`
    }                    `json:"mainImage"`
    PublishedDate string `json:"publishedDate"`
    RequestUrl    string `json:"requestUrl"`
    Title         string `json:"title"`
    Type          string `json:"type"`
    WebUrl        string `json:"webUrl"`
    Comments struct {
        Enabled bool     `json:"enabled"`
    }                    `json:"comments"`
}

func (c *Client) GetArticleByUuid(uuid string) (result *Article, err error) {
    url := "https://api.ft.com/content/"+uuid
    return c.GetArticle(url)
}

func (c *Client) GetArticle(url string) (result *Article, err error) {
    result = &Article{}
    err = c.getJsonAtUrl(url, result)
    return result, err
}

func (c *Client) GetMainImageSet(article *Article) (result *ImageSet, err error) {
    if article.MainImage.ID == "" {
        return nil, nil
    }
    return c.GetImageSet(article.MainImage.ID)
}
