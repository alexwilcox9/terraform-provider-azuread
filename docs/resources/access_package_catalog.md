---
subcategory: "Entitlement Management"
---

# Resource: azuread_access_package_catalog

Manages a Access Package Catalog within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application roles: `EntitlementManagement.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Identity Governance Administrator`

## Example Usage

```terraform
resource "azuread_access_package_catalog" "example" {
  display_name       = "Example Catalog"
  description        = "An Access Package Catalog to manage resource access"
  catalog_type       = "UserManaged"
  state              = "Published"
  externally_visible = false
}
```

## Argument Reference

The following arguments are supported:

* `catalog_type` - (Required) A catalog type. Possible values are `UserManaged`, `ServiceDefault`,
* `description` - The description of the catalog.
* `display_name` - (Required) The friendly name for this named location.
* `externally_visible` - (Optional) A catalog type. Possible values are `published`, `unpublished`,
* `state` - (Required) The state of the catalog. Possible values are `UserManaged`, `ServiceDefault`,

---


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the named location.

## Import

Named Locations can be imported using the `id`, e.g.

```shell
terraform import azuread_access_package_catalog.example 00000000-0000-0000-0000-000000000000
```
