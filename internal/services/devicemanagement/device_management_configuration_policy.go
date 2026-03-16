// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package devicemanagement

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicy"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
)

var _ sdk.ResourceWithUpdate = DeviceManagementConfigurationPolicy{}

func settingSchema(depth int) *pluginsdk.Resource {
	result := map[string]*pluginsdk.Schema{
		"setting_definition_id": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},
		"setting_instance_template_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},
		"value": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},
		"setting_value_template_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},
	}

	if depth < 10 {
		result["children"] = pointer.To(pluginsdk.Schema{
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem:     settingSchema(depth + 1),
		})
	}

	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"choice_setting": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				Elem: &pluginsdk.Resource{
					Schema: result,
				},
			},
		},
	}
}

type DeviceManagementConfigurationPolicy struct{}

func (r DeviceManagementConfigurationPolicy) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return beta.ValidateDeviceManagementConfigurationPolicyID
}

func (r DeviceManagementConfigurationPolicy) ResourceType() string {
	return "azuread_device_management_configuration_policy"
}

func (r DeviceManagementConfigurationPolicy) ModelObject() interface{} {
	return &DeviceManagementConfigurationPolicyModel{}
}

func (r DeviceManagementConfigurationPolicy) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"description": {
			Description: "The description of the configuration policy",
			Type:        pluginsdk.TypeString,
			Required:    true,
		},
		"name": {
			Description: "The name of the configuration policy",
			Type:        pluginsdk.TypeString,
			Required:    true,
		},
		"platforms": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"technologies": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice(beta.PossibleValuesForDeviceManagementConfigurationTechnologies(), false),
		},
		"role_scope_tag_ids": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
		"settings": {
			Type:     pluginsdk.TypeList,
			Required: true,
			MaxItems: 1,
			Elem:     settingSchema(0),
		},
	}
}

func (r DeviceManagementConfigurationPolicy) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

type DeviceManagementConfigurationPolicyModel struct {
	Description     string                                        `tfschema:"description"`
	Name            string                                        `tfschema:"name"`
	Platforms       string                                        `tfschema:"platforms"`
	Technologies    string                                        `tfschema:"technologies"`
	RoleScopeTagIds []string                                      `tfschema:"role_scope_tag_ids"`
	Settings        []DeviceManagementConfigurationPolicySettings `tfschema:"settings"`
}

type DeviceManagementConfigurationPolicySettings struct {
	ChoiceSettings []DeviceManagementConfigurationChoiceSetting `tfschema:"choice_setting"`
}

type DeviceManagementConfigurationChoiceSetting struct {
	SettingDefinitionId       string                                         `tfschema:"setting_definition_id"`
	SettingInstanceTemplateId string                                         `tfschema:"setting_instance_template_id"`
	Value                     string                                         `tfschema:"value"`
	SettingValueTemplateId    string                                         `tfschema:"setting_value_template_id"`
	Children                  []*DeviceManagementConfigurationPolicySettings `tfschema:"children"`
}

func (r DeviceManagementConfigurationPolicy) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DeviceManagement.ConfigurationPolicyClient

			var model DeviceManagementConfigurationPolicyModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			// a := make([]beta.DeviceManagementConfigurationSetting, 0)

			// b := beta.DeviceManagementConfigurationChoiceSettingInstance{
			// 	SettingDefinitionId: pointer.To("device_vendor_msft_policy_config_secguidev22h2~policy~cat_secguide_pol_secguide_a001_block_flash"),
			// 	SettingInstanceTemplateReference: pointer.To(beta.DeviceManagementConfigurationSettingInstanceTemplateReference{
			// 		SettingInstanceTemplateId: pointer.To("bf7fb31d-c639-489b-8226-94ac86cb7b03"),
			// 	}),
			// 	ChoiceSettingValue: pointer.To(beta.DeviceManagementConfigurationChoiceSettingValue{
			// 		Value: nullable.Value("device_vendor_msft_policy_config_secguidev22h2~policy~cat_secguide_pol_secguide_a001_block_flash_1"),
			// 		SettingValueTemplateReference: pointer.To(beta.DeviceManagementConfigurationSettingValueTemplateReference{
			// 			SettingValueTemplateId: pointer.To("955247be-9bcd-470b-af8a-9c7fa1ea70d9"),
			// 		}),
			// 		Children: pointer.To([]beta.DeviceManagementConfigurationSettingInstance{
			// 			beta.DeviceManagementConfigurationChoiceSettingInstance{
			// 				SettingDefinitionId: pointer.To("device_vendor_msft_policy_config_secguidev22h2~policy~cat_secguide_pol_secguide_a001_block_flash_pol_secguide_block_flash"),
			// 				SettingInstanceTemplateReference: pointer.To(beta.DeviceManagementConfigurationSettingInstanceTemplateReference{
			// 					SettingInstanceTemplateId: pointer.To("adf0ffe0-1a26-4ad4-9016-016599eb0251"),
			// 				}),
			// 				ChoiceSettingValue: pointer.To(beta.DeviceManagementConfigurationChoiceSettingValue{
			// 					SettingValueTemplateReference: pointer.To(beta.DeviceManagementConfigurationSettingValueTemplateReference{
			// 						SettingValueTemplateId: pointer.To("a4c5edad-e134-43a9-95b3-7b629bb9c256"),
			// 					}),
			// 					Value: nullable.Value("device_vendor_msft_policy_config_secguidev22h2~policy~cat_secguide_pol_secguide_a001_block_flash_pol_secguide_block_flash_block embedded flash activation only"),
			// 				}),
			// 			},
			// 		}),
			// 	}),
			// }

			// a = append(a, beta.DeviceManagementConfigurationSetting{
			// 	SettingInstance: b,
			// })

			properties := beta.DeviceManagementConfigurationPolicy{
				Description:     nullable.Value(model.Description),
				Name:            nullable.Value(model.Name),
				Platforms:       pointer.To(beta.DeviceManagementConfigurationPlatforms(model.Platforms)),
				Technologies:    pointer.To(beta.DeviceManagementConfigurationTechnologies(model.Technologies)),
				RoleScopeTagIds: pointer.To(model.RoleScopeTagIds),
				Settings:        expandDeviceManagementConfigurationPolicySettings(model.Settings),
			}

			resp, err := client.CreateConfigurationPolicy(ctx, properties, configurationpolicy.DefaultCreateConfigurationPolicyOperationOptions())
			if err != nil {
				return fmt.Errorf("creating assignment schedule request: %v", err)
			}

			request := resp.Model
			if request == nil {
				return fmt.Errorf("creating assignment schedule request: model was nil")
			}
			if request.Id == nil || *request.Id == "" {
				return fmt.Errorf("creating assignment schedule request: ID returned for request is nil/empty")
			}

			id := beta.NewDeviceManagementConfigurationPolicyID(pointer.From(request.Id))

			metadata.SetID(id)

			return nil
		},
	}
}

func (r DeviceManagementConfigurationPolicy) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DeviceManagement.ConfigurationPolicyClient

			id, err := beta.ParseDeviceManagementConfigurationPolicyID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("unable to parse ID: %v", err)
			}

			var model DeviceManagementConfigurationPolicyModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			resp, err := client.GetConfigurationPolicy(ctx, *id, configurationpolicy.DefaultGetConfigurationPolicyOperationOptions())
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			response := resp.Model

			if response != nil {
				model.Description = response.Description.GetOrZero()
				model.Name = response.Name.GetOrZero()
				if response.Platforms != nil {
					model.Platforms = string(*response.Platforms)
				}
				if response.Technologies != nil {
					model.Technologies = string(*response.Technologies)
				}
				if response.RoleScopeTagIds != nil {
					model.RoleScopeTagIds = *response.RoleScopeTagIds
				}
			}

			return metadata.Encode(&model)
		},
	}
}

func (r DeviceManagementConfigurationPolicy) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DeviceManagement.ConfigurationPolicyClient

			id, err := beta.ParseDeviceManagementConfigurationPolicyID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("unable to parse ID: %v", err)
			}

			var model DeviceManagementConfigurationPolicyModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			properties := beta.DeviceManagementConfigurationPolicy{
				Description:     nullable.Value(model.Description),
				Name:            nullable.Value(model.Name),
				RoleScopeTagIds: pointer.To(model.RoleScopeTagIds),
			}

			_, err = client.UpdateConfigurationPolicy(ctx, *id, properties, configurationpolicy.DefaultUpdateConfigurationPolicyOperationOptions())
			if err != nil {
				return fmt.Errorf("creating updated assignment schedule request: %v", err)
			}

			return nil
		},
	}
}

func (r DeviceManagementConfigurationPolicy) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DeviceManagement.ConfigurationPolicyClient

			id, err := beta.ParseDeviceManagementConfigurationPolicyID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("unable to parse ID: %v", err)
			}

			_, err = client.DeleteConfigurationPolicy(ctx, *id, configurationpolicy.DefaultDeleteConfigurationPolicyOperationOptions())
			if err != nil {
				return fmt.Errorf("deleting device management configuration policy %q: %+v", id, err)
			}

			return nil
		},
	}
}

func expandDeviceManagementConfigurationPolicySettings(input []DeviceManagementConfigurationPolicySettings) *[]beta.DeviceManagementConfigurationSetting {
	result := make([]beta.DeviceManagementConfigurationSetting, 0)

	for _, setting := range input {
		for _, choiceSetting := range setting.ChoiceSettings {
			choiceSettingresult := beta.DeviceManagementConfigurationChoiceSettingInstance{
				ChoiceSettingValue: pointer.To(beta.DeviceManagementConfigurationChoiceSettingValue{
					Value: nullable.Value(choiceSetting.Value),
				}),
			}
			if choiceSetting.SettingDefinitionId != "" {
				choiceSettingresult.SettingDefinitionId = pointer.To(choiceSetting.SettingDefinitionId)
			}
			if choiceSetting.SettingValueTemplateId != "" {
				choiceSettingresult.ChoiceSettingValue.SettingValueTemplateReference = pointer.To(beta.DeviceManagementConfigurationSettingValueTemplateReference{
					SettingValueTemplateId: pointer.To(choiceSetting.SettingValueTemplateId),
				})
			}
			children := make([]DeviceManagementConfigurationPolicySettings, 0)
			for _, child := range choiceSetting.Children {
				children = append(children, pointer.From(child))
			}
			if len(children) > 0 {
				choiceSettingresult.ChoiceSettingValue.Children = expandDeviceManagementConfigurationSettingInstance(children)
			}
			result = append(result, beta.DeviceManagementConfigurationSetting{
				SettingInstance: choiceSettingresult,
			})
		}
	}

	return &result
}

func expandDeviceManagementConfigurationSettingInstance(input []DeviceManagementConfigurationPolicySettings) *[]beta.DeviceManagementConfigurationSettingInstance {
	result := make([]beta.DeviceManagementConfigurationSettingInstance, 0)

	for _, setting := range input {
		for _, choiceSetting := range setting.ChoiceSettings {
			choiceSettingresult := beta.DeviceManagementConfigurationChoiceSettingInstance{
				ChoiceSettingValue: pointer.To(beta.DeviceManagementConfigurationChoiceSettingValue{
					Value: nullable.Value(choiceSetting.Value),
				}),
			}
			if choiceSetting.SettingDefinitionId != "" {
				choiceSettingresult.SettingDefinitionId = pointer.To(choiceSetting.SettingDefinitionId)
			}
			if choiceSetting.SettingValueTemplateId != "" {
				choiceSettingresult.ChoiceSettingValue.SettingValueTemplateReference = pointer.To(beta.DeviceManagementConfigurationSettingValueTemplateReference{
					SettingValueTemplateId: pointer.To(choiceSetting.SettingValueTemplateId),
				})
			}
			children := make([]DeviceManagementConfigurationPolicySettings, 0)
			for _, child := range choiceSetting.Children {
				children = append(children, pointer.From(child))
			}
			if len(children) > 0 {
				choiceSettingresult.ChoiceSettingValue.Children = expandDeviceManagementConfigurationSettingInstance(children)
			}
			result = append(result, choiceSettingresult)
		}
	}
	return &result
}
