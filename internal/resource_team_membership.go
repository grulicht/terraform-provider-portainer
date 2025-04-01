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

		Importer: &schema.ResourceImporter{
			State: resourceTeamMembershipImport,
		},

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
	return resourceTeamMembershipRead(d, meta)
}

func resourceTeamMembershipRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/team_memberships", client.Endpoint), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-API-Key", client.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to fetch team memberships list: %s", resp.Status)
	}

	var memberships []TeamMembershipResponse
	if err := json.NewDecoder(resp.Body).Decode(&memberships); err != nil {
		return err
	}

	for _, m := range memberships {
		if strconv.Itoa(m.ID) == d.Id() {
			d.Set("role", m.Role)
			d.Set("team_id", m.TeamID)
			d.Set("user_id", m.UserID)
			return nil
		}
	}

	d.SetId("")
	return nil
}

func resourceTeamMembershipImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceTeamMembershipRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
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
