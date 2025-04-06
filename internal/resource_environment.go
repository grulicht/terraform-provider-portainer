package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnvironmentCreate,
		Read:   resourceEnvironmentRead,
		Delete: resourceEnvironmentDelete,
		Update: resourceEnvironmentUpdate,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Environment type: 1 = Docker, 2 = Agent, 3 = Azure, 4 = Edge Agent, 5 = Kubernetes",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					t := val.(int)
					if t < 1 || t > 5 {
						errs = append(errs, fmt.Errorf("%q must be between 1 and 5", key))
					}
					return
				},
			},
			"group_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "ID of the Portainer endpoint group. Default is 1 (Unassigned).",
			},
			"tag_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "List of tag IDs to assign to the environment.",
			},
		},
	}
}

func resourceEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("Name", d.Get("name").(string))
	_ = writer.WriteField("URL", d.Get("environment_address").(string))
	_ = writer.WriteField("EndpointCreationType", strconv.Itoa(d.Get("type").(int)))
	_ = writer.WriteField("GroupID", strconv.Itoa(d.Get("group_id").(int)))

	_ = writer.WriteField("TLS", "true")
	_ = writer.WriteField("TLSSkipVerify", "true")
	_ = writer.WriteField("TLSSkipClientVerify", "true")

	if v, ok := d.GetOk("tag_ids"); ok {
		tagIds := v.([]interface{})
		formatted := "["
		for i, id := range tagIds {
			if i > 0 {
				formatted += ","
			}
			formatted += fmt.Sprintf("%d", id.(int))
		}
		formatted += "]"
		_ = writer.WriteField("TagIds", formatted)
	}

	writer.Close()

	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}

	var result struct {
		ID int `json:"Id"`
	}

	url := fmt.Sprintf("%s/endpoints", client.Endpoint)
	if err := client.DoMultipartRequest("POST", url, &body, headers, &result); err != nil {
		return fmt.Errorf("failed to create environment: %w", err)
	}

	d.SetId(strconv.Itoa(result.ID))
	return resourceEnvironmentRead(d, meta)
}

func resourceEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	url := fmt.Sprintf("%s/endpoints/%s", client.Endpoint, d.Id())

	var env struct {
		Name      string `json:"Name"`
		Type      int    `json:"Type"`
		URL       string `json:"URL"`
		PublicURL string `json:"PublicURL"`
		GroupID   int    `json:"GroupId"`
		TagIds    []int  `json:"TagIds"`
	}

	resp, err := client.DoRequest("GET", url, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("failed to read environment")
	}

	if err := json.NewDecoder(resp.Body).Decode(&env); err != nil {
		return err
	}

	d.Set("name", env.Name)
	d.Set("type", env.Type)
	d.Set("group_id", env.GroupID)
	d.Set("tag_ids", env.TagIds)

	if env.Type == 1 {
		d.Set("environment_address", env.URL)
	} else {
		d.Set("environment_address", env.PublicURL)
	}

	return nil
}

func resourceEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	id := d.Id()
	url := fmt.Sprintf("%s/endpoints/%s", client.Endpoint, id)

	payload := map[string]interface{}{
		"name":      d.Get("name").(string),
		"url":       d.Get("environment_address").(string),
		"publicURL": d.Get("environment_address").(string),
		"groupID":   d.Get("group_id").(int),
		"tagIDs":    d.Get("tag_ids").([]interface{}),
	}

	resp, err := client.DoRequest("PUT", url, nil, payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to update environment: %s", string(body))
	}

	return resourceEnvironmentRead(d, meta)
}

func resourceEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	url := fmt.Sprintf("%s/endpoints/%s", client.Endpoint, d.Id())

	resp, err := client.DoRequest("DELETE", url, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 || resp.StatusCode == 204 {
		return nil
	}

	data, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("failed to delete environment: %s", string(data))
}
