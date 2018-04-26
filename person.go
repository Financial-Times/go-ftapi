package ftapi

type ChangeEvent struct {
    StartedAt  string `json:"startedAt"`
    EndedAt    string `jsoon:"endedAt"`
}

type Membership struct {
    Title string `json:"title"`
    Organisation  *Thing `json:"organisation"`
    ChangeEvents  []ChangeEvent `json:"changeEvents"`
}

type Role struct {
    Thing
    ChangeEvents  []ChangeEvent `json:"changeEvents"`

}

type Person struct {
    Thing
    Labels     []string `json:"labels"`
    Salutation string   `json:"salutation"`
    BirthYear  int      `json:"birthYear"`
    Memberships []Membership `json:"memberships"`
}

func (c *Client) PersonByUUID(uuid string) (result *Person, err error) {
    url := "/people/"+uuid
    return c.Person(url)
}

func (c *Client) Person(url string) (result *Person, err error) {
    result = &Person{}
    raw, err := c.FromURL(url, result)
    result.RawJSON = raw
    return result, err
}
