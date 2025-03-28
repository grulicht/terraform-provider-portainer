package internal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type dockerImageAuth struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	ServerAddress string `json:"serveraddress"`
}

func resourceDockerImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceDockerImageCreate,
		Read:   resourceDockerImageRead,
		Delete: resourceDockerImageDelete,
		Update: nil,
		Schema: map[string]*schema.Schema{
			"endpoint_id":   {Type: schema.TypeInt, Required: true, ForceNew: true},
			"image":         {Type: schema.TypeString, Required: true, ForceNew: true},
			"registry_auth": {Type: schema.TypeString, Optional: true, Sensitive: true, ForceNew: true},
		},
	}
}

func resourceDockerImageCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	image := d.Get("image").(string)
	endpointID := d.Get("endpoint_id").(int)
	auth := d.Get("registry_auth").(string)

	params := url.Values{}
	params.Add("fromImage", image)
	fullURL := fmt.Sprintf("%s/endpoints/%d/docker/images/create?%s", client.Endpoint, endpointID, params.Encode())

	req, err := http.NewRequest("POST", fullURL, strings.NewReader(""))
	if err != nil {
		return err
	}
	req.Header.Set("X-API-Key", client.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", "0")

	if auth != "" {
		split := strings.SplitN(auth, ":", 2)
		if len(split) != 2 {
			return fmt.Errorf("invalid registry_auth format (expected username:password)")
		}
		payload := dockerImageAuth{
			Username:      split[0],
			Password:      split[1],
			Email:         "",
			ServerAddress: strings.Split(image, "/")[0],
		}
		jsonData, _ := json.Marshal(payload)
		encoded := base64.StdEncoding.EncodeToString(jsonData)
		req.Header.Set("X-Registry-Auth", encoded)
	} else {
		req.Header.Set("X-Registry-Auth", base64.StdEncoding.EncodeToString([]byte(`{}`)))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to pull image, status code: %d, body: %s", resp.StatusCode, string(body))
	}
	fmt.Printf("[DEBUG] Docker image pull result: %s\n", string(body))

	d.SetId(fmt.Sprintf("%d-%s", endpointID, image))
	return nil
}

func resourceDockerImageRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDockerImageDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	endpointID := d.Get("endpoint_id").(int)
	image := d.Get("image").(string)

	encodedImage := url.PathEscape(image)
	deleteURL := fmt.Sprintf("%s/endpoints/%d/docker/images/%s", client.Endpoint, endpointID, encodedImage)

	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-API-Key", client.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to delete image, status code: %d, body: %s", resp.StatusCode, string(body))
	}
	fmt.Printf("[DEBUG] Docker image delete result: %s\n", string(body))

	d.SetId("")
	return nil
}
