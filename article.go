package ftapi

type Annotation struct {
	APIURL     string   `json:"apiUrl"`
	DirectType string   `json:"directType"`
	ID         string   `json:"id"`
	Predicate  string   `json:"predicate"`
	PrefLabel  string   `json:"prefLabel"`
	Type       string   `json:"type"`
	Types      []string `json:"types"`
}

type Article struct {
	Annotations []Annotation `json:"annotations"`
	BodyXML     string       `json:"bodyXML"`
	Brands      []string     `json:"brands"`
	Byline      string       `json:"byline"`
	ID          string       `json:"id"`
	MainImage   struct {
		ID string `json:"id"`
	} `json:"mainImage"`
	PublishedDate string `json:"publishedDate"`
	RequestURL    string `json:"requestUrl"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	WebURL        string `json:"webUrl"`
	Comments      struct {
		Enabled bool `json:"enabled"`
	} `json:"comments"`
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
	err = c.FromURL(url, result)
	return result, err
}

func (c *Client) MainImageSet(article *Article) (result *ImageSet, err error) {
	if article.MainImage.ID == "" {
		return nil, nil
	}
	return c.ImageSet(article.MainImage.ID)
}
