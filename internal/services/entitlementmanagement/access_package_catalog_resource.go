package entitlementmanagement

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

func accessPackageCatalogResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: accessPackageCatalogResourceCreate,
		ReadContext:   accessPackageCatalogResourceRead,
		UpdateContext: accessPackageCatalogResourceUpdate,
		DeleteContext: accessPackageCatalogResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{

			"display_name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"description": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"catalog_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.AccessPackageCatalogTypeUserManaged,
					msgraph.AccessPackageCatalogTypeServiceDefault,
				}, false),
			},

			"state": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.AccessPackageCatalogStatePublished,
					msgraph.AccessPackageCatalogStateUnpublished,
				}, false),
			},

			"externally_visible": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func accessPackageCatalogResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).EntitlementManagement.AccessPackageCatalogClient

	properties := msgraph.AccessPackageCatalog{
		DisplayName:         utils.String(d.Get("display_name").(string)),
		State:               d.Get("state").(string),
		CatalogType:         d.Get("catalog_type").(string),
		Description:         utils.String(d.Get("description").(string)),
		IsExternallyVisible: utils.Bool(d.Get("externally_visible").(bool)),
	}

	catalog, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create access package catalog")
	}

	if catalog.ID == nil || *catalog.ID == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for access package catalog is nil/empty")
	}

	d.SetId(*catalog.ID)

	return accessPackageCatalogResourceRead(ctx, d, meta)
}

func accessPackageCatalogResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).EntitlementManagement.AccessPackageCatalogClient

	catalog, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package catalog with Object ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving Access package catalog with object ID %q", d.Id())
	}

	tf.Set(d, "display_name", catalog.DisplayName)
	tf.Set(d, "state", catalog.State)
	tf.Set(d, "catalog_type", catalog.CatalogType)
	tf.Set(d, "description", catalog.Description)
	tf.Set(d, "externally_visible", catalog.IsExternallyVisible)

	return nil
}

func accessPackageCatalogResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).EntitlementManagement.AccessPackageCatalogClient

	properties := msgraph.AccessPackageCatalog{
		ID:                  utils.String(d.Id()),
		DisplayName:         utils.String(d.Get("display_name").(string)),
		State:               d.Get("state").(string),
		CatalogType:         d.Get("catalog_type").(string),
		Description:         utils.String(d.Get("description").(string)),
		IsExternallyVisible: utils.Bool(d.Get("externally_visible").(bool)),
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update access package package with ID: %q", d.Id())
	}

	return nil
}

func accessPackageCatalogResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).EntitlementManagement.AccessPackageCatalogClient
	catalogId := d.Id()

	_, status, err := client.Get(ctx, catalogId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package catalog with ID %q already deleted", catalogId)
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving access package catalog with ID %q", catalogId)
	}

	status, err = client.Delete(ctx, catalogId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting access package catalog with ID %q, got status %d", catalogId, status)
	}

	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, catalogId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of access package catalog with ID %q", catalogId)
	}

	return nil
}
