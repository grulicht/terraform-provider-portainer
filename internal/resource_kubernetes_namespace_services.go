package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceKubernetesNamespaceService() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesNamespaceServiceCreate,
		Read:   resourceKubernetesNamespaceServiceRead,
		Update: resourceKubernetesNamespaceServiceUpdate,
		Delete: resourceKubernetesNamespaceServiceDelete,

		Schema: map[string]*schema.Schema{
			"environment_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ClusterIP",
			},
			"allocate_load_balancer_node_ports": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"external_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"external_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"load_balancer_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_affinity": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "None",
			},
			"ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "TCP",
						},
						"target_port": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"selector": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceKubernetesNamespaceServiceCreate(d *schema.ResourceData, meta interface{}) error {
	return createOrUpdateK8sService(d, meta, "POST")
}

func resourceKubernetesNamespaceServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	return createOrUpdateK8sService(d, meta, "PUT")
}

func createOrUpdateK8sService(d *schema.ResourceData, meta interface{}, method string) error {
	client := meta.(*APIClient)
	envID := d.Get("environment_id").(int)
	namespace := d.Get("namespace").(string)
	name := d.Get("name").(string)

	annotations := map[string]string{}
	if raw, ok := d.GetOk("annotations"); ok {
		for k, v := range raw.(map[string]interface{}) {
			annotations[k] = v.(string)
		}
	}

	labels := map[string]string{}
	if raw, ok := d.GetOk("labels"); ok {
		for k, v := range raw.(map[string]interface{}) {
			labels[k] = v.(string)
		}
	}

	selector := map[string]string{}
	if raw, ok := d.GetOk("selector"); ok {
		for k, v := range raw.(map[string]interface{}) {
			selector[k] = v.(string)
		}
	}

	externalIPs := []string{}
	if raw, ok := d.GetOk("external_ips"); ok {
		for _, v := range raw.([]interface{}) {
			externalIPs = append(externalIPs, v.(string))
		}
	}

	ports := []map[string]interface{}{}
	if raw, ok := d.GetOk("ports"); ok {
		for _, p := range raw.([]interface{}) {
			port := p.(map[string]interface{})
			targetPort := parseTargetPort(port["target_port"])
			ports = append(ports, map[string]interface{}{
				"Name":       port["name"],
				"Port":       port["port"],
				"Protocol":   port["protocol"],
				"TargetPort": targetPort,
			})
		}
	}

	body := map[string]interface{}{
		"name":                          name,
		"namespace":                     namespace,
		"type":                          d.Get("type").(string),
		"allocateLoadBalancerNodePorts": d.Get("allocate_load_balancer_node_ports").(bool),
		"annotations":                   annotations,
		"labels":                        labels,
		"selector":                      selector,
		"ports":                         ports,
		"externalIPs":                   externalIPs,
		"externalName":                  d.Get("external_name").(string),
		"loadBalancerIP":                d.Get("load_balancer_ip").(string),
		"sessionAffinity":               d.Get("session_affinity").(string),
	}

	jsonBody, _ := json.Marshal(body)
	url := fmt.Sprintf("%s/kubernetes/%d/namespaces/%s/services", client.Endpoint, envID, namespace)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
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
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to %s service: %s", strings.ToLower(method), string(data))
	}

	d.SetId(fmt.Sprintf("%d:%s:%s", envID, namespace, name))
	return nil
}

func parseTargetPort(value interface{}) string {
	var str string
	switch v := value.(type) {
	case int:
		str = strconv.Itoa(v)
	case float64:
		str = strconv.Itoa(int(v))
	case string:
		str = v
	default:
		str = fmt.Sprintf("%v", v)
	}

	if !strings.ContainsAny(str, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		str = "p" + str // prefix with a letter
	}

	return str
}

func resourceKubernetesNamespaceServiceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceKubernetesNamespaceServiceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
