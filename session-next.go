package ftapi

import "net/http"

type NextSession struct {
    EMailAddress   string `json:"emailAddress"`
    ERightsID      string `json:"erightsId"`
    FirstName      string `json:"firstName"`
    Groups         string `json:"groups"`
    Industry       string `json:"industry"`
    LastName       string `json:"lastName"`
    PassportID     string `json:"passportId"`
    Position       string `json:"position"`
    Products       string `json:"products"`
    Responsibility string `json:"responsibility"`
    SessionToken   string `json:"sessionToken"`
    Title          string `json:"title"`
    UUID           string `json:"uuid"`
}

func (c *Client) NextSessionFromCookie(cookie *http.Cookie) (*NextSession, error) {

	url := "https://session-next.ft.com/"
	result := &NextSession{}
	err := c.jsonAtURLWithCookie(url, result, cookie)
	return result, err

}
