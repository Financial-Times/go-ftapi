package ftapi

import ( "time"
         "fmt"
         "log" )

type ImageSet struct {
    RawJSON    *[]byte
	Description string `json:"description"`
	ID          string `json:"id"`
	Members     []Image `json:"members"`
	PublishedDate string `json:"publishedDate"`
	RequestURL    string `json:"requestUrl"`
	Title         string `json:"title"`
	Type          string `json:"type"`
}

func (c *Client) ImageSetByUUID(uuid string) (result *ImageSet, err error) {
	url := "https://api.ft.com/content/" + uuid
	return c.ImageSet(url)
}

func (c *Client) ImageSet(url string) (result *ImageSet, err error) {
	result = &ImageSet{}
	raw, err := c.FromURL(url, result)
    result.RawJSON = raw
	return result, err
}

func (c *Client) ImageSetMembers(imageset *ImageSet, timeout time.Duration) (result []*Image, err error) {
	chImages := make(chan *Image)

	for _, member := range imageset.Members {
        log.Println(len(imageset.Members), "images to get:", member.ID, c)
		go func(url string) {
			image, err := c.Image(url)
			if err != nil {
				return
			}
			chImages <- image
		}(member.ID)
	}

	images := []*Image{}

	for {
		select {
		case image := <-chImages:
            log.Println("Received:", image.ID)
			images = append(images, image)
			if len(images) == len(imageset.Members) {
				return images, nil
			}
		case <-time.After(timeout):
            if timeout > time.Duration(0) {
    			return nil, fmt.Errorf("Image retrieval timed out after %sms", timeout.String())
            }
		}
	}

}
