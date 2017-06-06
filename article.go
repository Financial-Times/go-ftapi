package ftapi

import "time"

type Annotation struct {
    	Thing
    	Predicate  string   `json:"predicate"`
	Type       string   `json:"type"`
}

type Article struct {
	Thing
	Annotations []Annotation `json:"annotations"`
	BodyXML     string       `json:"bodyXML"`
	Brands      []string     `json:"brands"`
	Byline      string       `json:"byline"`
	CanBeSyndicated string	 `json:"canBeSyndicated"`
	MainImage   struct {
		ID string `json:"id"`
	} `json:"mainImage"`
	PublishedDate time.Time
	RawPublishedDate string `json:"publishedDate"`
	RequestURL    string `json:"requestUrl"`
	Standfirst    string `json:"standfirst"`
	Standout      struct {
		EditorsChoice bool `json:"editorsChoice"`
		Exclusive bool `json:"exclusive"`
		Scoop bool `json:"scoop"`
	} `json:"standout"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	WebURL        string `json:"webUrl"`
	Comments      struct {
		Enabled bool `json:"enabled"`
	} `json:"comments"`
}

type ArticleRef struct {
    APIURL string `json:"apiUrl"`
    ID     string `json:"id"`
}

func (a Article) UUID() string {
	return FinalUUID(a.ID)
}

func (c *Client) ArticleByUUID(uuid string) (result *Article, err error) {
	url := "https://api.ft.com/content/" + uuid
	return c.Article(url)
}

func (c *Client) EnrichedArticleByUUID(uuid string) (result *Article, err error) {
	url := "https://api.ft.com/enrichedcontent/" + uuid
	return c.Article(url)
}

func (c *Client) EnrichedArticle(url string) (result *Article, err error) {
	return c.EnrichedArticleByUUID(FinalUUID(url))
}

func (c *Client) Article(url string) (result *Article, err error) {
	result = &Article{}
	raw, err := c.FromURL(url, result)
    result.RawJSON = raw
    if err == nil {
        result.PublishedDate, err = time.Parse("2006-01-02T15:04:05.000Z", result.RawPublishedDate)
    }
	return result, err
}

func (c *Client) ArticleRefsAnnotatedByUUID(uuid string) (result []ArticleRef, err error) {
	url := "https://api.ft.com/content?isAnnotatedBy=" + uuid
	result = []ArticleRef{}
	_, err = c.FromURL(url, &result)
	return result, err
}

func (c *Client) MainImageSet(article *Article) (result *ImageSet, err error) {
	if article.MainImage.ID == "" {
		return nil, nil
	}
	return c.ImageSet(article.MainImage.ID)
}
