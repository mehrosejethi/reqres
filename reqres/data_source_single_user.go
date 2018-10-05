package reqres

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceSingleUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSingleUserRead,

		Schema: map[string]*schema.Schema{
			"request_headers": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"body": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceSingleUserRead(d *schema.ResourceData, meta interface{}) error {

	url := "https://reqres.in/api/users/" + d.Get("user_id").(string)
	headers := d.Get("request_headers").(map[string]interface{})

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("Error creating request: %s", err)
	}

	for name, value := range headers {
		req.Header.Set(name, value.(string))
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error during making a request: %s", url)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP request error. Response code: %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" || isContentTypeAllowed(contentType) == false {
		return fmt.Errorf("Content-Type is not a text type. Got: %s", contentType)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error while reading response body. %s", err)
	}

	d.Set("body", string(bytes))
	d.SetId(time.Now().UTC().String())

	return nil
}
