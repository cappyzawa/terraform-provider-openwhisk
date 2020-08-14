package openwhisk

import (
	"context"
	"fmt"

	"github.com/apache/openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAction() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceActionCreate,
		UpdateContext: resourceActionUpdate,
		ReadContext:   resourceActionRead,
		DeleteContext: resourceActionDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceActionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceActionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceActionRead(ctx, d, m)
}

func resourceActionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*whisk.Client)
	name := d.Get("name").(string)
	action, _, err := client.Actions.Get(name, false)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to get action",
			Detail:   fmt.Sprintf("error: %v", err),
		})
		return diags
	}
	d.SetId(action.Name)
	d.Set("name", action.Name)
	return diags
}

func resourceActionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
