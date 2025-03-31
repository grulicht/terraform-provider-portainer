package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CloudCredentialPayload struct {
	Provider    string                 `json:"provider"`
	Name        string                 `json:"name"`
	Credentials map[string]interface{} `json:"credentials"`
}

func resourceCloudCredentials() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudCredentialsCreate,
		Delete: resourceCloudCredentialsDelete,
		Read:   resourceCloudCredentialsRead,
		Schema: map[string]*schema.Schema{
			"provider": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"credentials": {
				Type:     schema.TypeMap,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceCloudCredentialsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	payload := CloudCredentialPayload{
		Provider:    d.Get("provider").(string),
		Name:        d.Get("name").(string),
		Credentials: mapStringInterface(d.Get("credentials").(map[string]interface{})),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/cloud/credentials", client.Endpoint), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Set("X-API-Key", client.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to create cloud credential: HTTP %d", resp.StatusCode)
	}

	var result struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	d.SetId(strconv.Itoa(result.ID))
	return nil
}

func resourceCloudCredentialsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	url := fmt.Sprintf("%s/cloud/credentials/%s", client.Endpoint, d.Id())

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-API-Key", client.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to delete cloud credential: HTTP %d", resp.StatusCode)
	}

	d.SetId("")
	return nil
}

func resourceCloudCredentialsRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func mapStringInterface(input map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = v
	}
	return output
}
