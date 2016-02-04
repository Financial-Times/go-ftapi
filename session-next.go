package ftapi

import "net/http"

type NextSession struct {
    EmailAddress   string `json:"emailAddress"`
    ErightsId      string `json:"erightsId"`
    FirstName      string `json:"firstName"`
    Groups         string `json:"groups"`
    Industry       string `json:"industry"`
    LastName       string `json:"lastName"`
    PassportId     string `json:"passportId"`
    Position       string `json:"position"`
    Products       string `json:"products"`
    Responsibility string `json:"responsibility"`
    SessionToken   string `json:"sessionToken"`
    Title          string `json:"title"`
    Uuid           string `json:"uuid"`
}

func (c *Client) GetNextSessionFromCookie(cookie *http.Cookie) (*NextSession, error) {

	url := "https://session-next.ft.com/"
	result := &NextSession{}
	err := c.getJsonAtUrlWithCookie(url, result, cookie)
	return result, err

}
