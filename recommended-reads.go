package ftapi

import (
        "net/url"
)

type Recommendation struct {
    ID         string  `json:"id"`
    Popularity float64 `json:"popularity"`
    Published  string  `json:"published"`
    Score      float64 `json:"score"`
    Title      string  `json:"title"`
}

type Recommendations struct {
    Articles []Recommendation `json:"articles"`
    Status  string `json:"status"`
    Type    string `json:"type"`
    Version string `json:"version"`
}

func (c *Client) GetContextualRecommendationsByUuid(uuid string) (result *Recommendations, err error) {
    u, err := url.Parse("https://api.ft.com/recommended-reads-api/recommend/contextual")
    q := u.Query()
    q.Set("contentid", uuid)
    q.Set("count", "10")
    q.Set("sort", "rel")
    q.Set("recency", "7")
    u.RawQuery = q.Encode()
    result = &Recommendations{}
    err = c.getJsonAtUrl(u.String(), result)
    return result, err
}

func (c *Client) GetBehaviouralRecommendationsByUuid(uuid string, userid string) (result *Recommendations, err error) {
    u, err := url.Parse("https://api.ft.com/recommended-reads-api/recommend/behavioural")
    q := u.Query()
    q.Set("contentid", uuid)
    q.Set("userid", userid)
    q.Set("count", "10")
    q.Set("sort", "rel")
    q.Set("recency", "7")
    u.RawQuery = q.Encode()
    result = &Recommendations{}
    err = c.getJsonAtUrl(u.String(), result)
    return result, err
}

func (c *Client) GetBehaviouralRecommendations(userid string) (result *Recommendations, err error) {
    u, err := url.Parse("https://api.ft.com/recommended-reads-api/recommend/behavioural")
    q := u.Query()
    q.Set("userid", userid)
    q.Set("count", "10")
    q.Set("sort", "rel")
    q.Set("recency", "7")
    u.RawQuery = q.Encode()
    result = &Recommendations{}
    err = c.getJsonAtUrl(u.String(), result)
    return result, err
}

func (c *Client) GetPopularRecommendations() (result *Recommendations, err error) {
    u, err := url.Parse("https://api.ft.com/recommended-reads-api/recommend/popular")
    q := u.Query()
    q.Set("count", "10")
    q.Set("recency", "7")
    u.RawQuery = q.Encode()
    result = &Recommendations{}
    err = c.getJsonAtUrl(u.String(), result)
    return result, err
}
