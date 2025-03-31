package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type WebhookPayload struct {
	EndpointID  int    `json:"endpointID"`
	RegistryID  int    `json:"registryID,omitempty"`
	ResourceID  string `json:"resourceID"`
	WebhookType int    `json:"webhookType"`
}

type WebhookResponse struct {
	ID         int    `json:"Id"`
	EndpointID int    `json:"EndpointId"`
	RegistryID int    `json:"RegistryId"`
	ResourceID string `json:"ResourceId"`
	Token      string `json:"Token"`
	Type       int    `json:"Type"`
}

func resourceWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebhookCreate,
		Read:   resourceWebhookRead,
		Delete: resourceWebhookDelete,
		Schema: map[string]*schema.Schema{
			"endpoint_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"registry_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"webhook_type": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"token": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	payload := WebhookPayload{
		EndpointID:  d.Get("endpoint_id").(int),
		RegistryID:  d.Get("registry_id").(int),
		ResourceID:  d.Get("resource_id").(string),
		WebhookType: d.Get("webhook_type").(int),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/webhooks", client.Endpoint), bytes.NewBuffer(jsonPayload))
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
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create webhook: %s", string(body))
	}

	var result WebhookResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%d", result.ID))
	d.Set("token", result.Token)
	return nil
}

func resourceWebhookRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	webhookID := d.Id()

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/webhooks/%s", client.Endpoint, webhookID), nil)
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
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete webhook: %s", string(body))
	}

	d.SetId("")
	return nil
}
