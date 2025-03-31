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

type TeamMembershipPayload struct {
	Role   int `json:"role"`
	TeamID int `json:"teamID"`
	UserID int `json:"userID"`
}

type TeamMembershipResponse struct {
	ID     int `json:"Id"`
	Role   int `json:"Role"`
	TeamID int `json:"TeamID"`
	UserID int `json:"UserID"`
}

func resourceTeamMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceTeamMembershipCreate,
		Read:   resourceTeamMembershipRead,
		Delete: resourceTeamMembershipDelete,
		Schema: map[string]*schema.Schema{
			"role": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Membership role: 1 = team leader, 2 = regular member",
			},
			"team_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the team",
			},
			"user_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the user",
			},
		},
	}
}

func resourceTeamMembershipCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	payload := TeamMembershipPayload{
		Role:   d.Get("role").(int),
		TeamID: d.Get("team_id").(int),
		UserID: d.Get("user_id").(int),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/team_memberships", client.Endpoint), bytes.NewBuffer(data))
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
		return fmt.Errorf("failed to create team membership: %s", body)
	}

	var result TeamMembershipResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	d.SetId(strconv.Itoa(result.ID))
	return nil
}

func resourceTeamMembershipRead(d *schema.ResourceData, meta interface{}) error {
	// Not supported via Portainer API, remove from state to avoid drift
	return nil
}

func resourceTeamMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	id := d.Id()

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/team_memberships/%s", client.Endpoint, id), nil)
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
		return fmt.Errorf("failed to delete team membership: %s", body)
	}

	d.SetId("")
	return nil
}
