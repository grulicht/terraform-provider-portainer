package internal

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider defines the Portainer Terraform provider schema and resources.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PORTAINER_ENDPOINT", nil),
				Description: "URL of the Portainer instance (e.g. https://portainer.example.com). '/api' will be appended automatically if missing.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PORTAINER_API_KEY", nil),
				Description: "API key to authenticate with Portainer. Only API keys are supported (not JWT tokens).",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"portainer_user":            resourceUser(),
	        "portainer_team":            resourceTeam(),
	        "portainer_environment":     resourceEnvironment(),
	        "portainer_endpoint_group":  resourceEndpointGroup(),
	        "portainer_tag":             resourceTag(),
			"portainer_registry":        resourceRegistry(),
			"portainer_backup":          resourceBackup(),
			"portainer_backup_s3":       resourceBackupS3(),
			"portainer_edge_group":      resourceEdgeGroup(),
			"portainer_edge_job":        resourceEdgeJob(),
			"portainer_auth":            resourceAuth(),
			"portainer_edge_stack":      resourceEdgeStack(),
			"portainer_custom_template": resourceCustomTemplate(),
			"portainer_stack":           resourcePortainerStack(),
            "portainer_container_exec":  resourceContainerExec(),
			"portainer_docker_network":  resourceDockerNetwork(),
			"portainer_docker_image":    resourceDockerImage(),
			"portainer_docker_volume":   resourceDockerVolume(),
			"portainer_open_amt":        resourceOpenAMT(),
			"portainer_settings":        resourceSettings(),
			"portainer_ssl":             resourceSSLSettings(),
			"portainer_team_membership": resourceTeamMembership(),
			"portainer_webhook":         resourceWebhook(),
			"portainer_webhook_execute": resourceWebhookExecute(),
			"portainer_licenses":        resourceLicenses(),
			"portainer_cloud_credentials": resourceCloudCredentials(),
		},
		ConfigureContextFunc: configureProvider,
	}
}

// APIClient is a simple client struct to store connection information.
type APIClient struct {
	Endpoint string
	APIKey   string
}

// configureProvider sets up the API client and appends '/api' if missing from the endpoint.
func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	endpoint := d.Get("endpoint").(string)
	apiKey := d.Get("api_key").(string)

	// Ensure endpoint ends with /api
	if !strings.HasSuffix(endpoint, "/api") {
		endpoint = strings.TrimRight(endpoint, "/") + "/api"
	}

	client := &APIClient{
		Endpoint: endpoint,
		APIKey:   apiKey,
	}

	var diags diag.Diagnostics
	return client, diags
}
