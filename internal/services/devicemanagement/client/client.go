// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicy"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

// CAUTION!
// The Conditional Access API has compatibility issues between API versions. If you create a policy using the Beta API,
// or even if you update an existing policy using the Beta API that was originally created with the Stable API, that
// policy will be irrevocably mutated and can no longer be updated, or even _read_ using the Stable API.
// For this reason, we are bound to using the Stable API here, as to use the Beta API, even to update a single property
// for a Conditional Access Policy, will break that policy for users. The only way to go back to the Stable API after
// breaking a policy in this way, is to delete and recreate it, which is wholly undesirable for a critical security resource.

type Client struct {
	ConfigurationPolicyClient *configurationpolicy.ConfigurationPolicyClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	configurationPolicyClient, err := configurationpolicy.NewConfigurationPolicyClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(configurationPolicyClient.Client)

	return &Client{
		ConfigurationPolicyClient: configurationPolicyClient,
	}, nil
}
