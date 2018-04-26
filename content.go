package ftapi

import "time"
import "fmt"

type ContentRef struct {
    APIURL string `json:"apiUrl"`
    ID     string `json:"id"`
}

func (c *Client) ContentRefsAnnotatedByUUID(uuid string, limit int) (result []ContentRef, err error) {
    limitQuery := ""
    if limit != 0 {
        limitQuery = fmt.Sprintf("&limit=%d", limit)
    }
    url := fmt.Sprintf("/content?isAnnotatedBy=%s%s",uuid,limitQuery)
    result = []ContentRef{}
    _, err = c.FromURL(url, &result)
    return result, err
}

func (c *Client) ContentRefsAnnotatedByUUIDWithTimeWindow(uuid string, limit int, fromDate time.Time, toDate time.Time) (result []ContentRef, err error) {
    limitQuery := ""
    if limit != 0 {
        limitQuery = fmt.Sprintf("&limit=%d", limit)
    }
    url := fmt.Sprintf("/content?isAnnotatedBy=%s&fromDate=%s&toDate=%s%s", uuid, fromDate.Format("2006-01-02"), toDate.Format("2006-01-02"), limitQuery)
    result = []ContentRef{}
    _, err = c.FromURL(url, &result)
    return result, err
}

