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

func resourceRegistry() *schema.Resource {
	return &schema.Resource{
		Create: resourceRegistryCreate,
		Read:   resourceRegistryRead,
		Delete: resourceRegistryDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"base_url": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"authentication": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"instance_url": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"aws_region": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRegistryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	registryType := d.Get("type").(int)
	name := d.Get("name").(string)
	url := d.Get("url").(string)
	baseURL := d.Get("base_url").(string)
	auth := d.Get("authentication").(bool)

	body := map[string]interface{}{
		"name": name,
		"type": registryType,
	}

	switch registryType {
	case 1: // Quay.io
		body["url"] = url
		body["authentication"] = true
		body["username"] = d.Get("username").(string)
		body["password"] = d.Get("password").(string)

	case 2: // Azure
		body["url"] = url
		body["baseURL"] = baseURL
		body["authentication"] = true
		body["username"] = d.Get("username").(string)
		body["password"] = d.Get("password").(string)

	case 3: // Custom
		body["url"] = url
		body["baseURL"] = baseURL
		body["authentication"] = auth
		if auth {
			body["username"] = d.Get("username").(string)
			body["password"] = d.Get("password").(string)
		}

	case 4: // GitLab
		body["url"] = url
		gitlab := map[string]interface{}{
			"InstanceURL": d.Get("instance_url").(string),
		}
		body["authentication"] = true
		body["username"] = d.Get("username").(string)
		body["password"] = d.Get("password").(string)
		body["gitlab"] = gitlab

	case 5: // ProGet
		body["url"] = url
		body["baseURL"] = baseURL
		body["authentication"] = true
		body["username"] = d.Get("username").(string)
		body["password"] = d.Get("password").(string)

	case 6: // DockerHub
	    body["url"] = url
		body["authentication"] = true
		body["username"] = d.Get("username").(string)
		body["password"] = d.Get("password").(string)

	case 7: // AWS ECR
		ecr := map[string]interface{}{}
		if v, ok := d.GetOk("aws_region"); ok {
			ecr["Region"] = v.(string)
		}
		body["url"] = url
		body["ecr"] = ecr
		body["authentication"] = auth
		if auth {
			body["username"] = d.Get("username").(string)
			body["password"] = d.Get("password").(string)
		}

	default:
		return fmt.Errorf("unsupported registry type: %d", registryType)
	}

	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/registries", client.Endpoint), bytes.NewBuffer(jsonBody))
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
		return fmt.Errorf("failed to create registry: %s", string(data))
	}

	var result struct {
		ID int `json:"Id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	d.SetId(strconv.Itoa(result.ID))
	return resourceRegistryRead(d, meta)
}

func resourceRegistryRead(d *schema.ResourceData, meta interface{}) error {
	// Not implemented (optional for now)
	return nil
}

func resourceRegistryDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/registries/%s", client.Endpoint, d.Id()), nil)
	req.Header.Set("X-API-Key", client.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete registry")
	}
	return nil
}
