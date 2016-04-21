package ftapi

import (
	"net/url"
	"strconv"
)

type Recommendation struct {
	ID            string  `json:"id"`
	Popularity    float64 `json:"popularity"`
	PublishedDate string  `json:"published"`
	Score         float64 `json:"score"`
	Title         string  `json:"title"`
	URL           string  `json:"url"`
}

type Recommendations struct {
	Articles []Recommendation `json:"articles"`
	Status   string           `json:"status"`
	Type     string           `json:"type"`
	Version  string           `json:"version"`
}

func (c *Client) RawContextualRecommendationsByUUID(uuid string, count int, recency int) (result *Recommendations, err error) {
	u, err := url.Parse("https://api.ft.com/recommended-reads-api/recommend/contextual")
	q := u.Query()
	q.Set("contentid", uuid)
	q.Set("count", strconv.Itoa(count))
	q.Set("sort", "rel")
	q.Set("recency", strconv.Itoa(recency))
	u.RawQuery = q.Encode()
	result = &Recommendations{}
	err = c.FromURL(u.String(), result)
	return result, err
}

func (c *Client) ContextualRecommendationsByUUID(uuid string, count int, recency int) (result []Recommendation, err error) {
	r, err := c.RawContextualRecommendationsByUUID(uuid, count, recency)
	if err != nil {
		return nil, err
	}
	return r.Articles, nil
}

func (c *Client) RawBehaviouralRecommendationsByUUID(uuid string, userid string, count int, recency int) (result *Recommendations, err error) {
	u, err := url.Parse("https://api.ft.com/recommended-reads-api/recommend/behavioural")
	q := u.Query()
	q.Set("contentid", uuid)
	q.Set("userid", userid)
	q.Set("count", strconv.Itoa(count))
	q.Set("sort", "rel")
	q.Set("recency", strconv.Itoa(recency))
	u.RawQuery = q.Encode()
	result = &Recommendations{}
	err = c.FromURL(u.String(), result)
	return result, err
}

func (c *Client) BehaviouralRecommendationsByUUID(uuid string, userid string, count int, recency int) (result []Recommendation, err error) {
	r, err := c.RawBehaviouralRecommendationsByUUID(uuid, userid, count, recency)
	if err != nil {
		return nil, err
	}
	return r.Articles, nil
}

func (c *Client) RawBehaviouralRecommendations(userid string, count int, recency int) (result *Recommendations, err error) {
	u, err := url.Parse("https://api.ft.com/recommended-reads-api/recommend/behavioural")
	q := u.Query()
	q.Set("userid", userid)
	q.Set("count", strconv.Itoa(count))
	q.Set("sort", "rel")
	q.Set("recency", strconv.Itoa(recency))
	u.RawQuery = q.Encode()
	result = &Recommendations{}
	err = c.FromURL(u.String(), result)
	return result, err
}

func (c *Client) BehaviouralRecommendations(userid string, count int, recency int) (result []Recommendation, err error) {
	r, err := c.RawBehaviouralRecommendations(userid, count, recency)
	if err != nil {
		return nil, err
	}
	return r.Articles, nil
}

func (c *Client) RawPopularRecommendations(count int, recency int) (result *Recommendations, err error) {
	u, err := url.Parse("https://api.ft.com/recommended-reads-api/recommend/popular")
	q := u.Query()
	q.Set("count", strconv.Itoa(count))
	q.Set("recency", strconv.Itoa(recency))
	u.RawQuery = q.Encode()
	result = &Recommendations{}
	err = c.FromURL(u.String(), result)
	return result, err
}

func (c *Client) PopularRecommendations(count int, recency int) (result []Recommendation, err error) {
	r, err := c.RawPopularRecommendations(count, recency)
	if err != nil {
		return nil, err
	}
	return r.Articles, nil
}
