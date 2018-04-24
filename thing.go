package ftapi

import (
    "net/url"
)

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

type ConceptSearchResponse struct {
    RawJSON *[]byte
    Concepts []struct {
        APIURL     string   `json:"apiUrl"`
        DirectType string   `json:"type"`
        ID         string   `json:"id"`
        PrefLabel  string   `json:"prefLabel"`
    } `json:"concepts"`
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

func (c *Client) ConceptSearch(s string) (result []Thing, err error) {
    return c.ConceptSearchByTypes(s, []string{
        NewOntology.Topic,
        Ontology.Organisation,
        NewOntology.Location,
        Ontology.Person,
    })
}

func (c *Client) ConceptSearchByTypes(s string, types []string) (result []Thing, err error) {
    result = []Thing{}

    if len(types) == 0 {
        return result, nil
    }

    u, _ := url.Parse("https://api.ft.com/concepts")
    q := u.Query()
    q.Set("mode", "search")
    q.Set("q", s)
    for _, t := range types {
        q.Add("type", t)
    }
    u.RawQuery = q.Encode()
    resp := &ConceptSearchResponse{}

    _, err = c.FromURL(u.String(), resp)

    for _, t := range resp.Concepts {
        result = append(result, Thing{
            APIURL: t.APIURL,
            DirectType: t.DirectType,
            ID: t.ID,
            PrefLabel: t.PrefLabel,
        })
    }
    return result, err
}

