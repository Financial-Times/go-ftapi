package ftapi

type Image struct {
    RawJSON       *[]byte
    BinaryURL     string  `json:"binaryUrl"`
    Description   string  `json:"description"`
    ID            string  `json:"id"`
    PixelHeight   int     `json:"pixelHeight"`
    PixelWidth    int     `json:"pixelWidth"`
    PublishedDate string  `json:"publishedDate"`
    RequestURL    string  `json:"requestUrl"`
    Title         string  `json:"title"`
    Type          string  `json:"type"`
}

func (c *Client) ImageByUUID(uuid string) (result *Image, err error) {
    url := "https://api.ft.com/content/"+uuid
    return c.Image(url)
}

func (c *Client) Image(url string) (result *Image, err error) {
    result = &Image{}
    raw, err := c.FromURL(url, result)
    result.RawJSON = raw
    return result, err
}
