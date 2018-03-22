package ftapi

import "encoding/json"

type SuggestResponse struct {
    RawJSON *[]byte
    Suggestions []Annotation `json:"suggestions"`
}

// Suggest takes an article object (which may have come from the API or have been partially constructed by the client) and returns suggested annotations.
func (c *Client) Suggest(article *Article) (result *SuggestResponse, err error) {
    articleJSON, err := json.Marshal(article)
    if err != nil {
        return nil, err
    }

    result = &SuggestResponse{}
    raw, err := c.FromURLWithBody("https://api.ft.com/content/suggest", articleJSON, result)
    result.RawJSON = raw
    return result, err
}
