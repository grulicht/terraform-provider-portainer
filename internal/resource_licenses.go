package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type LicensePayload struct {
	Key string `json:"key"`
}

type LicenseResponse struct {
	ConflictingKeys []string `json:"conflictingKeys"`
}

func resourceLicenses() *schema.Resource {
	return &schema.Resource{
		Create: resourceLicensesCreate,
		Read:   resourceLicensesRead,
		Delete: resourceLicensesDelete,
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "License key to be attached",
				Sensitive:   true,
				ForceNew:    true,
			},
			"force": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Force attach even if there are conflicting licenses",
			},
			"conflicting_keys": {
				Type:        schema.TypeList,
				Computed:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of conflicting license keys, if any",
			},
		},
	}
}

func resourceLicensesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	licenseKey := d.Get("key").(string)
	force := d.Get("force").(bool)

	payload := LicensePayload{
		Key: licenseKey,
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/licenses/add", client.Endpoint)
	if force {
		url += "?force=true"
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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
		return fmt.Errorf("failed to attach license: %s", string(body))
	}

	var response LicenseResponse
	_ = json.NewDecoder(resp.Body).Decode(&response)

	_ = d.Set("conflicting_keys", response.ConflictingKeys)
	
	d.SetId(licenseKey)
	return nil
}

func resourceLicensesRead(d *schema.ResourceData, meta interface{}) error {
	return nil // Not supported by Portainer API
}

func resourceLicensesDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
