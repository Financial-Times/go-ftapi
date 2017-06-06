package ftapi

type FinancialInstrument struct {
    Thing
    FIGI string `json:"FIGI"`
}

type Organisation struct {
    Thing
    Labels     []string `json:"labels"`
    LEICode    string `json:"leiCode"`
    IndustryClassification Thing `json:"industryClassification"`
    ParentOrganisation Thing `json:"parentOrganisation"`
    Subsidiaries []Thing `json:"subsidiaries"`
    Memberships []Membership `json:"memberships"`
    FinancialInstrument Thing `json:"financialInstrument"`
}

func (c *Client) OrganisationByUUID(uuid string) (result *Organisation, err error) {
    url := "https://api.ft.com/organisations/"+uuid
    return c.Organisation(url)
}

func (c *Client) Organisation(url string) (result *Organisation, err error) {
    result = &Organisation{}
    raw, err := c.FromURL(url, result)
    result.RawJSON = raw
    return result, err
}
