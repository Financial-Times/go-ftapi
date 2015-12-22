package ftapi

type Image struct {
    BinaryUrl     string  `json:"binaryUrl"`
    Description   string  `json:"description"`
    ID            string  `json:"id"`
    PixelHeight   int     `json:"pixelHeight"`
    PixelWidth    int     `json:"pixelWidth"`
    PublishedDate string  `json:"publishedDate"`
    RequestUrl    string  `json:"requestUrl"`
    Title         string  `json:"title"`
    Type          string  `json:"type"`
}

func (c *Client) GetImageByUuid(uuid string) (result *Image, err error) {
    url := "https://api.ft.com/content/"+uuid
    return c.GetImage(url)
}

func (c *Client) GetImage(url string) (result *Image, err error) {
    result = &Image{}
    err = c.getJsonAtUrl(url, result)
    return result, err
}
