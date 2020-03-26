package ftapi

import "time"

type Annotation struct {
	Thing
	Predicate string `json:"predicate"`
	Type      string `json:"type"`
}

type Standout struct {
	EditorsChoice bool `json:"editorsChoice"`
	Exclusive     bool `json:"exclusive"`
	Scoop         bool `json:"scoop"`
}

type Article struct {
	Thing
	Annotations           []Annotation `json:"annotations"`
	BodyXML               string       `json:"bodyXML"`
	Brands                []string     `json:"brands"`
	Byline                string       `json:"byline"`
	CanBeSyndicated       string       `json:"canBeSyndicated"`
	CanBeDistributed      string       `json:"canBeDistributed"`
	EditorialDesk         string       `json:"editorialDesk"`
	Embeds                []ImageSet   `json:"embeds"`
	MainImage             ImageSet     `json:"mainImage"`
	PublishedDate         time.Time
	RawPublishedDate      string `json:"publishedDate"`
	FirstPublishedDate    time.Time
	RawFirstPublishedDate string   `json:"firstPublishedDate"`
	RequestURL            string   `json:"requestUrl"`
	Standfirst            string   `json:"standfirst"`
	Standout              Standout `json:"standout"`
	Title                 string   `json:"title"`
	Type                  string   `json:"type"`
	WebURL                string   `json:"webUrl"`
	Comments              struct {
		Enabled bool `json:"enabled"`
	} `json:"comments"`
}

func (a Article) UUID() string {
	return FinalUUID(a.ID)
}

func (c *Client) ArticleByUUID(uuid string) (result *Article, err error) {
	url := "/content/" + uuid
	return c.Article(url)
}

func (c *Client) EnrichedArticleByUUID(uuid string) (result *Article, err error) {
	url := "/enrichedcontent/" + uuid
	return c.Article(url)
}

func (c *Client) EnrichedArticle(url string) (result *Article, err error) {
	return c.EnrichedArticleByUUID(FinalUUID(url))
}

func (c *Client) InternalArticleByUUID(uuid string) (result *Article, err error) {
	url := "/internalcontent/" + uuid
	return c.Article(url)
}

func (c *Client) InternalArticle(url string) (result *Article, err error) {
	return c.InternalArticleByUUID(FinalUUID(url))
}

func (c *Client) Article(url string) (result *Article, err error) {
	result = &Article{}
	raw, err := c.FromURL(url, result)
	result.RawJSON = raw
	if err != nil {
		return result, err
	}
	result.PublishedDate, err = time.Parse("2006-01-02T15:04:05.000Z", result.RawPublishedDate)
	if result.RawFirstPublishedDate != "" {
		result.FirstPublishedDate, err = time.Parse("2006-01-02T15:04:05.000Z", result.RawFirstPublishedDate)
	}
	return result, err
}

func (c *Client) MainImageSet(article *Article) (result *ImageSet, err error) {
	if article.MainImage.ID == "" {
		return nil, nil
	}
	return c.ImageSet(article.MainImage.ID)
}
