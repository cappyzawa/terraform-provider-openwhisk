package openwhisk

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apache/openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider serves terraform provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"auth_token": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ver": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"verbose": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"debug": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"user_agent": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Terraform-Provider-OpenWhisk",
			},
			"additional_header": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"openwhisk_action": resourceAction(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"openwhisk_action": dataSourceAction(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	host := d.Get("host").(string)
	namespace := d.Get("namespace").(string)
	authToken := d.Get("auth_token").(string)
	version := d.Get("ver").(string)
	debug := d.Get("debug").(bool)
	verbose := d.Get("verbose").(bool)
	userAgent := d.Get("user_agent").(string)
	additionalHeader := d.Get("additional_header").(map[string]interface{})
	header := make(http.Header)
	for k, v := range additionalHeader {
		header.Set(k, v.(string))
	}
	config := &whisk.Config{
		Host:              host,
		Namespace:         namespace,
		AuthToken:         authToken,
		Version:           version,
		Debug:             debug,
		Verbose:           verbose,
		UserAgent:         userAgent,
		AdditionalHeaders: header,
	}
	client, err := whisk.NewClient(http.DefaultClient, config)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to create client",
			Detail:   fmt.Sprintf("error: %v", err),
		})
	}
	return client, diags
}
