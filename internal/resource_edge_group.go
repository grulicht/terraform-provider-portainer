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

func resourceEdgeGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceEdgeGroupCreate,
		Read:   resourceEdgeGroupRead,
		Delete: resourceEdgeGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dynamic": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"partial_match": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"endpoints": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				ForceNew: true,
			},
			"tag_ids": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEdgeGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	body := map[string]interface{}{
		"name":         d.Get("name").(string),
		"dynamic":      d.Get("dynamic").(bool),
		"partialMatch": d.Get("partial_match").(bool),
	}

	if endpoints, ok := d.GetOk("endpoints"); ok {
		body["endpoints"] = endpoints
	}
	if tagIDs, ok := d.GetOk("tag_ids"); ok {
		body["tagIDs"] = tagIDs
	}

	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/edge_groups", client.Endpoint), bytes.NewBuffer(jsonBody))
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
		return fmt.Errorf("failed to create edge group: %s", string(data))
	}

	var result struct {
		ID int `json:"Id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	d.SetId(strconv.Itoa(result.ID))
	return resourceEdgeGroupRead(d, meta)
}

func resourceEdgeGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/edge_groups/%s", client.Endpoint, d.Id()), nil)
	req.Header.Set("X-API-Key", client.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("failed to read edge group")
	}

	var group struct {
		Name         string `json:"Name"`
		Dynamic      bool   `json:"Dynamic"`
		PartialMatch bool   `json:"PartialMatch"`
		TagIDs       []int  `json:"TagIds"`
		Endpoints    []int  `json:"Endpoints"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return err
	}

	d.Set("name", group.Name)
	d.Set("dynamic", group.Dynamic)
	d.Set("partial_match", group.PartialMatch)
	d.Set("tag_ids", group.TagIDs)
	d.Set("endpoints", group.Endpoints)

	return nil
}

func resourceEdgeGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/edge_groups/%s", client.Endpoint, d.Id()), nil)
	req.Header.Set("X-API-Key", client.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete edge group")
	}

	return nil
}
