package reqres

import (
	"bytes"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func resourceCreateUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"request_headers": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"request_body": &schema.Schema{
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

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	url := "https://reqres.in/api/users"
	headers := d.Get("request_headers").(map[string]interface{})
	jsonStr := []byte(d.Get("request_body").(string))

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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

	if resp.StatusCode != 201 {
		return fmt.Errorf("HTTP request error. Response code: %d", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error while reading response body. %s", err)
	}

	d.Set("body", string(bytes))
	d.SetId(time.Now().UTC().String())

	return nil
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
