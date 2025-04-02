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

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Delete: resourceUserDelete,
		Update: resourceUserUpdate,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"role": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"ldap_user": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"team_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Optional Portainer team ID. Only applicable for standard users (role = 2).",
			},
			"generate_api_key": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"api_key_description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "terraform-generated-api-key",
			},
			"api_key_raw": {
				Type:     schema.TypeString,
				Computed: true,
				Sensitive: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	role := d.Get("role").(int)
	ldapUser := d.Get("ldap_user").(bool)

	if ldapUser && password != "" {
		return fmt.Errorf("cannot set password for LDAP user")
	}
	if !ldapUser && password == "" {
		return fmt.Errorf("password is required for non-LDAP user")
	}

	body := map[string]interface{}{
		"Username": username,
		"Role":     role,
	}
	if !ldapUser {
		body["Password"] = password
	}

	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users", client.Endpoint), bytes.NewBuffer(jsonBody))
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
		return fmt.Errorf("failed to create user: %s", string(data))
	}

	var result struct {
		ID int `json:"Id"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&result)

	if result.ID == 0 {
		return resourceUserReadByUsername(d, meta)
	}
	d.SetId(strconv.Itoa(result.ID))

	teamID, ok := d.GetOk("team_id")
	if ok {
		if role != 2 {
			return fmt.Errorf("team_id can only be used with standard users (role = 2)")
		}

		teamMembership := map[string]interface{}{
			"UserID": result.ID,
			"TeamID": teamID.(int),
			"Role":   2,
		}
		jsonMembership, _ := json.Marshal(teamMembership)

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/team_memberships", client.Endpoint), bytes.NewBuffer(jsonMembership))
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
			return fmt.Errorf("failed to assign user to team: %s", string(data))
		}
	}

	if d.Get("generate_api_key").(bool) {
		description := d.Get("api_key_description").(string)
		if password == "" {
			return fmt.Errorf("password must be set to generate API key")
		}
		apiPayload := map[string]interface{}{
			"description": description,
			"password":    password,
		}
		jsonTokenBody, _ := json.Marshal(apiPayload)
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/users/%d/tokens", client.Endpoint, result.ID), bytes.NewBuffer(jsonTokenBody))
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

		if resp.StatusCode != 200 {
			data, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("failed to generate API key: %s", string(data))
		}

		var tokenResp struct {
			RawAPIKey string `json:"rawAPIKey"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
			return err
		}
		d.Set("api_key_raw", tokenResp.RawAPIKey)
	}

	return resourceUserRead(d, meta)
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/users/%s", client.Endpoint, d.Id()), nil)
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
		return fmt.Errorf("failed to read user")
	}

	var user struct {
		Username string `json:"Username"`
		Role     int    `json:"Role"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return err
	}

	d.Set("username", user.Username)
	d.Set("role", user.Role)
	return nil
}

func resourceUserReadByUsername(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/users", client.Endpoint), nil)
	req.Header.Set("X-API-Key", client.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to list users for lookup")
	}

	var users []struct {
		ID       int    `json:"Id"`
		Username string `json:"Username"`
		Role     int    `json:"Role"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return err
	}

	username := d.Get("username").(string)
	for _, u := range users {
		if u.Username == username {
			d.SetId(strconv.Itoa(u.ID))
			d.Set("role", u.Role)
			return nil
		}
	}

	return fmt.Errorf("user created but not found in user list")
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	id := d.Id()

	if d.HasChange("password") {
		oldPw, newPw := d.GetChange("password")
		payload := map[string]string{
			"password":    oldPw.(string),
			"newPassword": newPw.(string),
		}
		jsonBody, _ := json.Marshal(payload)
		req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/%s/passwd", client.Endpoint, id), bytes.NewBuffer(jsonBody))
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

		if resp.StatusCode != 204 {
			data, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("failed to update password: %s", string(data))
		}
	}

	body := map[string]interface{}{
		"username": d.Get("username").(string),
		"role":     d.Get("role").(int),
		"useCache": true,
	}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/%s", client.Endpoint, id), bytes.NewBuffer(jsonBody))
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
		return fmt.Errorf("failed to update user: %s", data)
	}

	return resourceUserRead(d, meta)
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	id := d.Id()

	if keyID, ok := d.Get("api_key_id").(int); ok && keyID > 0 {
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/users/%s/tokens/%d", client.Endpoint, id, keyID), nil)
		req.Header.Set("X-API-Key", client.APIKey)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			defer resp.Body.Close()
		}
	}

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/users/%s", client.Endpoint, id), nil)
	req.Header.Set("X-API-Key", client.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 || resp.StatusCode == 204 {
		return nil
	}

	data, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("failed to delete user: %s", string(data))
}
