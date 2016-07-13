package ftapi

type Thing struct {
    RawJSON    *[]byte
    APIURL     string   `json:"apiUrl"`
    DirectType string   `json:"directType"`
    ID         string   `json:"id"`
    PrefLabel  string   `json:"prefLabel"`
    Types      []string `json:"types"`
}

func (c *Client) ThingByUUID(uuid string) (result *Thing, err error) {
    url := "https://api.ft.com/things/"+uuid
    return c.Thing(url)
}

func (c *Client) Thing(url string) (result *Thing, err error) {
    result = &Thing{}
    raw, err := c.FromURL(url, result)
    result.RawJSON = raw
    return result, err
}