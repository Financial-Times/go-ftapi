package ftapi

type Thing struct {
    RawJSON    *[]byte
    APIURL     string   `json:"apiUrl"`
    DirectType string   `json:"directType"`
    ID         string   `json:"id"`
    PrefLabel  string   `json:"prefLabel"`
    Types      []string `json:"types"`
}

type Concept struct {
    Thing
    Aliases	[]string `json:"aliases"`
    BroaderConcepts []Annotation `json:"broaderConcepts"`
    NarrowerConcepts []Annotation `json:"narrowerConcepts"`
    RelatedConcepts []Annotation `json:"relatedConcepts"`
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

func (c *Client) ConceptByUUID(uuid string) (result *Concept, err error) {
    url := "https://api.ft.com/things/"+uuid
    return c.Concept(url)
}

func (c *Client) Concept(url string) (result *Concept, err error) {
    result = &Concept{}
    raw, err := c.FromURL(url, result)
    result.RawJSON = raw
    return result, err
}
