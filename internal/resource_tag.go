package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceTagCreate,
		Read:   resourceTagRead,
		Delete: resourceTagDelete,
		Update: nil, // No update API – tag must be recreated

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceTagCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	payload := map[string]interface{}{
		"name": d.Get("name").(string),
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tags", client.Endpoint), bytes.NewBuffer(jsonBody))
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

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create tag: %s", string(data))
	}

	var result struct {
		ID int `json:"Id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	d.SetId(strconv.Itoa(result.ID))
	return resourceTagRead(d, meta)
}

func resourceTagRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	// 1. Try primary request: GET /tags/{id}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/tags/%s", client.Endpoint, d.Id()), nil)
	req.Header.Set("X-API-Key", client.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	} else if resp.StatusCode == 200 {
		// read response body
		data, _ := io.ReadAll(resp.Body)
		fmt.Printf("[DEBUG] Response from GET /tags/%s: %s\n", d.Id(), string(data))

		// try to decode
		var tag struct {
			Name string `json:"Name"`
		}
		if err := json.Unmarshal(data, &tag); err == nil && tag.Name != "" {
			d.Set("name", tag.Name)
			return nil
		}
		fmt.Println("[DEBUG] Failed to parse tag by ID, falling back to full list.")
	}

	// 2. Fallback: GET /tags and search by ID
	reqList, _ := http.NewRequest("GET", fmt.Sprintf("%s/tags", client.Endpoint), nil)
	reqList.Header.Set("X-API-Key", client.APIKey)

	respList, err := http.DefaultClient.Do(reqList)
	if err != nil {
		return err
	}
	defer respList.Body.Close()

	if respList.StatusCode != 200 {
		return fmt.Errorf("failed to fallback to GET /tags list")
	}

	var tags []struct {
		ID   int    `json:"Id"`
		Name string `json:"Name"`
	}
	if err := json.NewDecoder(respList.Body).Decode(&tags); err != nil {
		return fmt.Errorf("failed to decode fallback tag list: %s", err)
	}

	for _, tag := range tags {
		if strconv.Itoa(tag.ID) == d.Id() {
			d.Set("name", tag.Name)
			return nil
		}
	}

	// If still not found
	d.SetId("")
	return nil
}

func resourceTagDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/tags/%s", client.Endpoint, d.Id()), nil)
	req.Header.Set("X-API-Key", client.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 || resp.StatusCode == 404 {
		return nil
	}

	data, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("failed to delete tag: %s", string(data))
}
