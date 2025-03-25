package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAuth() *schema.Resource {
	return &schema.Resource{
		Create: resourceAuthCreate,
		Read:   schema.Noop,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				Sensitive: true,
				ForceNew: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
				Sensitive: true,
				ForceNew: true,
			},
			"jwt": {
				Type:     schema.TypeString,
				Computed: true,
				Sensitive: true,
			},
		},
	}
}

func resourceAuthCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	creds := map[string]string{
		"username": d.Get("username").(string),
		"password": d.Get("password").(string),
	}

	jsonBody, _ := json.Marshal(creds)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth", client.Endpoint), bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to authenticate: %s", string(data))
	}

	var response struct {
		JWT string `json:"jwt"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	d.SetId("auth-result")
	d.Set("jwt", response.JWT)

	return nil
}
