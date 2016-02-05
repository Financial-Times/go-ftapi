package ftapi

import "time"
import "fmt"
import "log"

type ImageSet struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Members     []struct {
		ID string `json:"id"`
	} `json:"members"`
	PublishedDate string `json:"publishedDate"`
	RequestUrl    string `json:"requestUrl"`
	Title         string `json:"title"`
	Type          string `json:"type"`
}

func (c *Client) ImageSetByUuid(uuid string) (result *ImageSet, err error) {
	url := "https://api.ft.com/content/" + uuid
	return c.ImageSet(url)
}

func (c *Client) ImageSet(url string) (result *ImageSet, err error) {
	result = &ImageSet{}
	err = c.jsonAtUrl(url, result)
	return result, err
}

func (c *Client) ImageSetMembers(imageset *ImageSet, timeout_ms int64) (result []*Image, err error) {
	images_ch := make(chan *Image)

	for _, member := range imageset.Members {
        log.Println(len(imageset.Members), "images to get:", member.ID, c)
		go func(url string) {
			image, err := c.Image(url)
			if err != nil {
				return
			}
			images_ch <- image
		}(member.ID)
	}

	images := []*Image{}

	for {
		select {
		case image := <-images_ch:
            log.Println("Received:", image.ID)
			images = append(images, image)
			if len(images) == len(imageset.Members) {
				return images, nil
			}
		case <-time.After(time.Duration(timeout_ms) * time.Millisecond):
            if timeout_ms > 0 {
    			return nil, fmt.Errorf("Image retrieval timed out after %sms", timeout_ms)
            }
		}
	}

}
