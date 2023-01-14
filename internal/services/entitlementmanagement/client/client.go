package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	AccessPackageCatalogClient *msgraph.AccessPackageCatalogClient
}

func NewClient(o *common.ClientOptions) *Client {
	msClient := msgraph.NewAccessPackageCatalogClient(o.TenantID)
	o.ConfigureClient(&msClient.BaseClient)

	return &Client{
		AccessPackageCatalogClient: msClient,
	}
}
