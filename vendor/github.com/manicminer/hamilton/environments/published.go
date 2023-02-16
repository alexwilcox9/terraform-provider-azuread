package environments

type ApiAppId = string

// PublishedApis is a map containing Application IDs for well known APIs published by Microsoft.
// They can be used to acquire access tokens, but are primarily described here for easy inclusion in
// application manifests and service principal assignments.
// These are collected from various sources, including the Azure Portal, CLI, PowerShell and the following doc:
// https://learn.microsoft.com/en-us/troubleshoot/azure/active-directory/verify-first-party-apps-sign-in#application-ids-of-commonly-used-microsoft-applications
var PublishedApis = map[string]ApiAppId{
	"ApplicationInsights":               "f5c26e74-f226-4ae8-85f0-b4af0080ac9e",
	"AttestationService":                "c61423b7-1d1f-430d-b444-0eee53298103",
	"AzureActiveDirectoryGraph":         "00000002-0000-0000-c000-000000000000",
	"AzureAdIdentityGovernanceInsights": "58c746b0-a0b0-4647-a8f6-12dde5981638",
	"AzureAdIntegratedApp":              "af47b99c-8954-4b45-ab68-8121157418ef",
	"AzureAdNotification":               "fc03f97a-9db0-4627-a216-ec98ce54e018",
	"AzureAnalysisServices":             "4ac7d521-0382-477b-b0f8-7e1d95f85ca2",
	"AzureAppConfiguration":             "35ffadb3-7fc1-497e-b61b-381d28e744cc",
	"AzureAppService":                   "abfa0a7c-a6b6-4736-8310-5855508787cd",
	"AzureBatch":                        "ddbf3205-c6bd-46ae-8127-60eb93363864",
	"AzureContainerRegistry":            "6a0ec4d3-30cb-4a83-91c0-ae56bc0e3d26",
	"AzureCosmosDb":                     "a232010e-820c-4083-83bb-3ace5fc29d0b",
	"AzureDataBricks":                   "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d",
	"AzureDataCatalog":                  "9d3e55ba-79e0-4b7c-af50-dc460b81dca1",
	"AzureDataLake":                     "e9f49c6b-5ce5-44c8-925d-015017e9f7ad",
	"AzureDevOps":                       "499b84ac-1321-427f-aa17-267ca6975798",
	"AzureDigitalTwins":                 "0b07f429-9f4b-4714-9392-cc5e8e80c8b0",
	"AzureEventHubs":                    "80369ed6-5f11-4dd9-bef3-692475845e77",
	"AzureHdInsightCluster":             "7865c1d2-f040-46cc-875f-831a1ef6a28a",
	"AzureHealthcare":                   "4f6778d8-5aef-43dc-a1ff-b073724b9495",
	"AzureIamSupportability":            "a57aca87-cbc0-4f3c-8b9e-dc095fdc8978",
	"AzureImportExport":                 "7de4d5c5-5b32-4235-b8a9-33b34d6bcd2a",
	"AzureIotCentral":                   "9edfcdd9-0bc5-4bd4-b287-c3afc716aac7",
	"AzureIotHubDeviceProvisioning":     "0cd79364-7a90-4354-9984-6e36c841418d",
	"AzureKeyVault":                     "cfa8b339-82a2-471a-a3c9-0fc0be7a4093",
	"AzureKubernetesServiceAadServer":   "6dae42f8-4368-4678-94ff-3960e28e3630",
	"AzureMaps":                         "ba1ea022-5807-41d5-bbeb-292c7e1cf5f6",
	"AzureMediaServices":                "374b2a64-3b6b-436b-934c-b820eacca870",
	"AzurePortal":                       "c44b4083-3bb0-49c1-b47d-974e53cbdf3c",
	"AzureSecurityInsights":             "98785600-1bb7-4fb9-b9fa-19afe2c8a360",
	"AzureServiceBus":                   "80a10ef9-8168-493d-abf9-3297c4ef6e3c",
	"AzureServiceDeploy":                "5b306cba-9c71-49db-96c3-d17ca2379c4d",
	"AzureServiceManagement":            "797f4846-ba00-4fd7-ba43-dac1f8f63013",
	"AzureSqlDatabase":                  "022907d3-0f1b-48f7-badc-1ba6abab6d66",
	"AzureStackHciService":              "1322e676-dee7-41ee-a874-ac923822781c",
	"AzureStorage":                      "e406a681-f3d4-42a8-90b6-c2b029497af1",
	"AzureStreamAnalytics":              "66f1e791-7bfb-4e18-aed8-1720056421c7",
	"AzureSynapseGateway":               "1ac05c7e-12d2-4605-bf9d-549d7041c6b3",
	"AzureSynapseStudio":                "ec52d13d-2e85-410e-a89a-8c79fb6a32ac",
	"AzureTimeSeriesInsights":           "120d688d-1518-4cf7-bd38-182f158850b6",
	"AzureVPN":                          "41b23e61-6c1e-4545-b367-cd054e0ed4b4",
	"Bing":                              "9ea1ad79-fdb6-4f9a-8bc3-2b70f96e34c7",
	"BotFrameworkDevPortal":             "f3723d34-6ff5-4ceb-a148-d99dcd2511fc",
	"BranchConnectWebService":           "57084ef3-d413-4087-a28f-f6f3b1ad7786",
	"CognitiveServices":                 "7d312290-28c8-473c-a0ed-8e53749b6d6d",
	"ComputeRecommendationService":      "b9a92e36-2cf8-4f4e-bcb3-9d99e00e14ab",
	"ConnectionsService":                "b7912db9-aa33-4820-9d4f-709830fdd78f",
	"CortanaAtWorkBingServices":         "22d7579f-06c2-4baa-89d2-e844486adb9d",
	"CortanaAtWorkService":              "2a486b53-dbd2-49c0-a2bc-278bdfc30833",
	"CortanaRuntimeService":             "81473081-50b9-469a-b9d8-303109583ecb",
	"CustomerInsights":                  "38c77d00-5fcb-4cce-9d93-af4738258e3c",
	"DataMigrationService":              "a4bad4aa-bf02-4631-9f78-a64ffdba8150",
	"DomainControllerServices":          "2565bd9d-da50-47d4-8b85-4c97f669dc36",
	"Dynamic365BusinessCentral":         "996def3d-b36c-4153-8607-a6fd3c01b89f",
	"Dynamics365DataExportService":      "b861dbcc-a7ef-4219-a005-0e4de4ea7dcf",
	"DynamicsCrm":                       "00000007-0000-0000-c000-000000000000",
	"DynamicsErp":                       "00000015-0000-0000-c000-000000000000",
	"FlowService":                       "7df0a125-d3be-4c96-aa54-591f83ff541c",
	"GraphConnectorService":             "56c1da01-2129-48f7-9355-af6d59d42766",
	"InformationProtectionSyncService":  "870c4f2e-85b6-4d43-bdda-6ed9a579b725",
	"InTune":                            "c161e42e-d4df-4a3d-9b42-e7a3c31f59d4",
	"KustoService":                      "2746ea77-4702-4b45-80ca-3c97e680e8b7",
	"KustoServiceMFA":                   "725d0e77-e1fd-48f1-a295-2115457f7609",
	"LogAnalytics":                      "ca7f3f0b-7d91-482c-8e09-c5d840d0eac5",
	"MileIqAdminCenter":                 "de096ee1-dae7-4ee1-8dd5-d88ccc473815",
	"MileIqDashboard":                   "f7069a8d-9edc-4300-b365-ae53c9627fc4",
	"MileIqRestService":                 "b692184e-b47f-4706-b352-84b288d2d9ee",
	"MixedReality":                      "c7ddd9b4-5172-4e28-bd29-1e0792947d18",
	"MicrosoftAzureCli":                 "04b07795-8ddb-461a-bbee-02f9e1bf7b46",
	"MicrosoftAzureFrontDoor":           "ad0e1c7e-6d38-4ba4-9efd-0bc77ba9f037",
	"MicrosoftAzureFrontDoorCdn":        "205478c0-bd83-4e1b-a9d6-db63a3e1e1c8",
	"Microsoft365DataAtRestEncryption":  "c066d759-24ae-40e7-a56f-027002b5d3e4",
	"MicrosoftGraph":                    "00000003-0000-0000-c000-000000000000",
	"MicrosoftInvoicing":                "b6b84568-6c01-4981-a80f-09da9a20bbed",
	"MicrosoftOffice":                   "d3590ed6-52b3-4102-aeff-aad2292ab01c",
	"Microsoft.StorageSync":             "9469b9f5-6722-4481-a2b2-14ed560b706f",
	"MicrosoftTeams":                    "1fec8e78-bce4-4aaf-ab1b-5451cc387264",
	"MicrosoftTeamsWebClient":           "5e3ce6c0-2b1f-4285-8d4b-75ee78787346",
	"Office365Connectors":               "48af08dc-f6d2-435f-b2a7-069abd99c086",
	"Office365Demeter":                  "982bda36-4632-4165-a46a-9863b1bbcf7d",
	"Office365DwEngineV2":               "441509e5-a165-4363-8ee7-bcf0b7d26739",
	"Office365ExchangeOnline":           "00000002-0000-0ff1-ce00-000000000000",
	"Office365ExchangeOnlineProtection": "00000007-0000-0ff1-ce00-000000000000",
	"Office365InformationProtection":    "2f3f02c9-5679-4a5c-a605-0de55b07d135",
	"Office365Management":               "c5393580-f805-4401-95e8-94b7a6ef2fc2",
	"Office365Portal":                   "00000006-0000-0ff1-ce00-000000000000",
	"Office365SharePointOnline":         "00000003-0000-0ff1-ce00-000000000000",
	"Office365SuiteUx":                  "4345a7b9-9a63-4910-a426-35363201d503",
	"Office365Zoom":                     "0d38933a-0bbd-41ca-9ebd-28c4b5ba7cb7",
	"OfficeHome":                        "4765445b-32c6-49b0-83e6-1d93765276ca",
	"OfficeUwpPwa":                      "0ec893e0-5785-4de6-99da-4ed124e5296c",
	"OneNote":                           "2d4d3d8e-2be3-4bef-9f87-7875a61c29de",
	"OneProfileService":                 "b2cc270f-563e-4d8a-af47-f00963a71dcd",
	"OssRdbms":                          "123cd850-d9df-40bd-94d5-c9f07b7fa203",
	"OssRdbmsPostgreSqlFlexibleServerAadAuthentication": "5657e26c-cc92-45d9-bc47-9da6cfdb4ed9",
	"PeopleCardsService":          "394866fc-eedb-4f01-8536-3ff84b16be2a",
	"PolicyAdministrationService": "0469d4cd-df37-4d93-8a61-f8c75b809164",
	"PowerAppsRuntimeService":     "82f77645-8a66-4745-bcdf-9706824f9ad0",
	"PowerBiService":              "00000009-0000-0000-c000-000000000000",
	"Purview":                     "73c2949e-da2d-457a-9607-fcc665198967",
	"RightsManagementServices":    "00000012-0000-0000-c000-000000000000",
	"ServiceTrust":                "d6fdaa33-e821-4211-83d0-cf74736489e1",
	"Signup":                      "b4bddae8-ab25-483e-8670-df09b9f1d0ea",
	"SkypeForBusinessOnline":      "00000004-0000-0ff1-ce00-000000000000",
	"SpeechRecognition":           "1a6fcee6-0816-469b-acac-fe7ef2e87b83",
	"TargetedMessagingService":    "4c4f550b-42b2-4a16-93f9-fdb9e01bb6ed",
	"TeamsServices":               "cc15fd57-2c6c-4117-a88c-83b1d56b4bbe",
	"ThreatProtection":            "8ee8fdad-f234-4243-8f3b-15c294843740",
	"UniversalPrint":              "da9b70f6-5323-4ce6-ae5c-88dcc5082966",
	"WindowsDefenderAtp":          "fc780465-2017-40d4-a0c5-307022471b92",
	"WindowsVirtualDesktop":       "9cdead84-a844-4324-93f2-b2e6bb768d07",
	"Yammer":                      "00000005-0000-0ff1-ce00-000000000000",
}
