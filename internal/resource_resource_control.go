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

func resourceResourceControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceResourceControlCreate,
		Read:   resourceResourceControlRead,
		Update: resourceResourceControlUpdate,
		Delete: resourceResourceControlDelete,

		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sub_resource_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"administrators_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"public": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"teams": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"users": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func resourceResourceControlCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	body := map[string]interface{}{
		"resourceID":          d.Get("resource_id").(string),
		"subResourceIDs":      d.Get("sub_resource_ids"),
		"type":                d.Get("type").(int),
		"administratorsOnly": d.Get("administrators_only").(bool),
		"public":             d.Get("public").(bool),
		"teams":              d.Get("teams"),
		"users":              d.Get("users"),
	}

	jsonBody, _ := json.Marshal(body)
	url := fmt.Sprintf("%s/resource_controls", client.Endpoint)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
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
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create resource control: %s", string(data))
	}

	var result struct {
		ID int `json:"Id"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&result)

	d.SetId(strconv.Itoa(result.ID))
	return resourceResourceControlRead(d, meta)
}

func resourceResourceControlRead(d *schema.ResourceData, meta interface{}) error {
	// Optional: implement GET if supported
	return nil
}

func resourceResourceControlUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	id := d.Id()

	body := map[string]interface{}{
		"administratorsOnly": d.Get("administrators_only").(bool),
		"public":             d.Get("public").(bool),
		"teams":              d.Get("teams"),
		"users":              d.Get("users"),
	}

	jsonBody, _ := json.Marshal(body)
	url := fmt.Sprintf("%s/resource_controls/%s", client.Endpoint, id)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
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
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to update resource control: %s", string(data))
	}

	return resourceResourceControlRead(d, meta)
}

func resourceResourceControlDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	url := fmt.Sprintf("%s/resource_controls/%s", client.Endpoint, d.Id())
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

	if resp.StatusCode >= 400 && resp.StatusCode != 404 {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete resource control: %s", string(data))
	}

	d.SetId("")
	return nil
}
