package entitlementmanagement_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/odata"
)

type AccessPackageCatalogResource struct{}

func TestAccAccessPackageCatalog_unpublished(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog", "test")
	r := AccessPackageCatalogResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.unpublished(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageCatalog_published(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog", "test")
	r := AccessPackageCatalogResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.published(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageCatalog_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog", "test")
	r := AccessPackageCatalogResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.unpublished(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.published(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.unpublished(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r AccessPackageCatalogResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	var id *string

	catalog, status, err := clients.EntitlementManagement.AccessPackageCatalogClient.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Access Package Catalog with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Access Package Catalog with object ID %q: %+v", state.ID, err)
	}
	id = catalog.ID

	return utils.Bool(id != nil && *id == state.ID), nil
}

func (AccessPackageCatalogResource) unpublished(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_access_package_catalog" "test" {
  display_name = "acctestNLIP-%[1]d"
  description        = "An unpublished Access Package Catalog"
  catalog_type       = "UserManaged"
  state              = "unpublished"
  externally_visible = false
}
`, data.RandomInteger)
}

func (AccessPackageCatalogResource) published(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_access_package_catalog" "test" {
  display_name = "acctestNLIP-%[1]d"
  description        = "A published Access Package Catalog"
  catalog_type       = "UserManaged"
  state              = "published"
  externally_visible = true
}
`, data.RandomInteger)
}
