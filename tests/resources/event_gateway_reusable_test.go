package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestEventGatewayReusable(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("EGW Control plane", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "my_event_gateway" {
  				name = "my_test_event_gateway"
			}
		`)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.my_event_gateway", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway.my_event_gateway", "name", "my_test_event_gateway"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp.AddAttribute("labels", `{key = "value"}`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway.my_event_gateway", "labels.%", "1"),
						resource.TestCheckResourceAttr("konnect_event_gateway.my_event_gateway", "labels.key", "value"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})

	})

	t.Run("EGW Backend Cluster", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf_test_event_gateway" {
  				name = "my_test_event_gateway"
			}
		`)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(`
		resource "konnect_event_gateway_backend_cluster" "tf_test_egw_backend_cluster" {
			name = "tf_test_egw_backend_cluster"
			authentication = {
    			anonymous = {}
  			}
			bootstrap_servers = [
				"10.0.0.1:8080"
			]
			tls = {
				enabled = false
				insecure_skip_verify = true
			}
		}
		`)
		require.NoError(t, err)

		// Set up dependencies
		egwBackendCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.tf_test_event_gateway", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster", "name", "tf_test_egw_backend_cluster"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster.AddAttribute("labels", `{key = "value"}`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster", "labels.%", "1"),
						resource.TestCheckResourceAttr("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster", "labels.key", "value"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("EGW Virtual Cluster", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.KonnectBeta,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.KonnectBeta

		// Event Gateway (Control Plane)
		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf_test_event_gateway" {
			  name = "my_test_event_gateway"
			}
		`)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_backend_cluster" "tf_test_egw_backend_cluster" {
			  name = "tf_test_egw_backend_cluster"

			  authentication = {
				anonymous = {}
			  }

			  bootstrap_servers = [
				"10.0.0.1:8080"
			  ]

			  tls = {
				enabled = false
				insecure_skip_verify = true
			  }
			}
		`)
		require.NoError(t, err)
		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		// Virtual Cluster
		egwVirtualCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster" "tf_test_egw_virtual_cluster" {
			  name      = "tf_test_egw_virtual_cluster"
			  dns_label = "vcluster-1"
			  acl_mode  = "passthrough"

			  authentication = [
				{
				  anonymous = {}
				}
			  ]

			  destination = {
				id = konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster.id
			  }
			}
		`)
		require.NoError(t, err)

		// Wire gateway → virtual cluster
		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(egwBackendCluster).
						Upsert(egwVirtualCluster).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway.tf_test_event_gateway",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_virtual_cluster.tf_test_egw_virtual_cluster",
								plancheck.ResourceActionCreate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster.tf_test_egw_virtual_cluster",
							"name",
							"tf_test_egw_virtual_cluster",
						),
					),
				},
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(egwBackendCluster).
						Upsert(egwVirtualCluster).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(egwBackendCluster).
						Upsert(
							egwVirtualCluster.AddAttribute("labels", `{ key = "value" }`),
						).
						Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster.tf_test_egw_virtual_cluster",
							"labels.%",
							"1",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_virtual_cluster.tf_test_egw_virtual_cluster",
							"labels.key",
							"value",
						),
					),
				},
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(egwBackendCluster).
						Upsert(egwVirtualCluster).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("EGW_Consume_Policy_Decrypt", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.KonnectBeta,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.KonnectBeta

		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf-test-egw-decrypt" {
			  name = "tf-test-egw-decrypt"
			}
		`)
		require.NoError(t, err)

		backend, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_backend_cluster" "tf-test-backend" {
			  name = "tf-test-backend"

			  authentication = {
				anonymous = {}
			  }

			  bootstrap_servers = ["127.0.0.1:9092"]

			  tls = {
				enabled = false
				insecure_skip_verify = true
			  }
			}
		`)
		require.NoError(t, err)

		backend.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		virtualCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster" "tf-test-virtual" {
			  name      = "tf-test-virtual"
			  dns_label = "tf-test-vc"
			  acl_mode  = "passthrough"

			  authentication = [
				{
				  anonymous = {}
				}
			  ]

			  destination = {
				id = konnect_event_gateway_backend_cluster.tf-test-backend.id
			  }
			}
		`)
		require.NoError(t, err)

		virtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		staticKey, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_static_key" "tf-test-static-key" {
			  name  = "tf-test-static-key"
			  value = "$${vault.env[\"MY_ENV_VAR\"]}"
			}
		`)
		require.NoError(t, err)

		staticKey.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		decryptPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_decrypt" "tf-test-decrypt-policy" {
			  name        = "tf-test-decrypt-policy"
			  description = "initial decrypt policy"
			  enabled     = false
			  condition   = "context.topic.name.endsWith('my_suffix')"

			  config = {
				failure_mode = "passthrough"

				key_sources = [
				  {
					static = {
					  name = "tf-test-static-key"
					}
				  }
				]
				part_of_record = ["key"]
			  }
			}
		`)
		require.NoError(t, err)

		decryptPolicy.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)
		decryptPolicy.AddAttribute(
			"virtual_cluster_id",
			virtualCluster.ResourcePath()+".id",
		)
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(backend).Upsert(virtualCluster).Upsert(staticKey).Upsert(decryptPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.tf-test-egw-decrypt", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_backend_cluster.tf-test-backend", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_virtual_cluster.tf-test-virtual", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_static_key.tf-test-static-key", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_consume_policy_decrypt.tf-test-decrypt-policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt.tf-test-decrypt-policy", "enabled", "false"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_decrypt.tf-test-decrypt-policy", "config.failure_mode", "passthrough"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(backend).Upsert(virtualCluster).Upsert(staticKey).Upsert(decryptPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(backend).Upsert(virtualCluster).Upsert(staticKey).Upsert(decryptPolicy.AddAttribute("enabled", "true")).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway_consume_policy_decrypt.tf-test-decrypt-policy", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("EGW Consume Policy Modify Headers", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.KonnectBeta,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.KonnectBeta

		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf_test_event_gateway_modify_headers" {
				name = "test-gateway-modify-headers"
			}
		`)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_backend_cluster" "tf_test_egw_backend_cluster_modify_headers" {
				name = "test-backend-cluster-modify-headers"
				authentication = {
					anonymous = {}
				}
				bootstrap_servers = [
					"10.0.0.1:9092"
				]
				tls = {
					enabled = false
					insecure_skip_verify = true
				}
			}
		`)
		require.NoError(t, err)

		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwVirtualCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster" "tf_test_egw_virtual_cluster_modify_headers" {
				name      = "test-vcluster-modify-headers"
				dns_label = "vcluster-modify-headers"
				acl_mode  = "passthrough"

				authentication = [
					{
						anonymous = {}
					}
				]

				destination = {
					id = konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster_modify_headers.id
				}
			}
		`)
		require.NoError(t, err)

		// Wire gateway → virtual cluster
		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		consumePolicyModifyHeaders, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_modify_headers" "test_modify_headers_policy" {
				name        = "test-consume-modify-headers-policy"
				description = "Test consume policy for modifying headers"
				condition   = "context.topic.name == 'header-topic'"
				enabled     = true

				config = {
					actions = [
						{
							set = {
								key   = "X-Custom-Header"
								value = "custom-value"
							}
						},
						{
							remove = {
								key = "X-Remove-Header"
							}
						}
					]
				}
			}
		`)
		require.NoError(t, err)

		// Wire policy to gateway and virtual cluster
		consumePolicyModifyHeaders.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)
		consumePolicyModifyHeaders.AddAttribute(
			"virtual_cluster_id",
			egwVirtualCluster.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				// Create all resources
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(egwBackendCluster).
						Upsert(egwVirtualCluster).
						Upsert(consumePolicyModifyHeaders).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway.tf_test_event_gateway_modify_headers",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster_modify_headers",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_virtual_cluster.tf_test_egw_virtual_cluster_modify_headers",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
								plancheck.ResourceActionCreate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"name",
							"test-consume-modify-headers-policy",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"description",
							"Test consume policy for modifying headers",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"enabled",
							"true",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"condition",
							"context.topic.name == 'header-topic'",
						),
						resource.TestCheckResourceAttrSet(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"id",
						),
					),
				},
				// Update policy with labels
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(egwBackendCluster).
						Upsert(egwVirtualCluster).
						Upsert(
							consumePolicyModifyHeaders.AddAttribute("labels", `{ env = "dev", service = "headers" }`),
						).
						Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"labels.%",
							"2",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"labels.env",
							"dev",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"labels.service",
							"headers",
						),
					),
				},
				// Update enabled flag to false
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(egwBackendCluster).
						Upsert(egwVirtualCluster).
						Upsert(
							consumePolicyModifyHeaders.AddAttribute("enabled", "false"),
						).
						Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"enabled",
							"false",
						),
					),
				},
				// Revert to enabled and remove labels
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(egwBackendCluster).
						Upsert(egwVirtualCluster).
						Upsert(
							consumePolicyModifyHeaders.AddAttribute("enabled", "true"),
						).
						Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_consume_policy_modify_headers.test_modify_headers_policy",
							"enabled",
							"true",
						),
					),
				},
			},
		})
	})

	t.Run("EGW Consume Policy Schema Validation", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.KonnectBeta,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.KonnectBeta

		// Event Gateway (Control Plane) - Parent resource
		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf_test_event_gateway_schema_validation" {
				name = "test-gateway-schema-validation"
			}
		`)
		require.NoError(t, err)

		// Event Gateway Backend Cluster - Parent resource
		egwBackendCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_backend_cluster" "tf_test_egw_backend_cluster_schema_validation" {
				name = "test-backend-cluster-schema-validation"
				authentication = {
					anonymous = {}
				}
				bootstrap_servers = [
					"10.0.0.1:9092"
				]
				tls = {
					enabled = false
					insecure_skip_verify = true
				}
			}
		`)
		require.NoError(t, err)

		// Wire gateway → backend cluster
		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		// Event Gateway Virtual Cluster - Parent resource
		egwVirtualCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster" "tf_test_egw_virtual_cluster_schema_validation" {
				name      = "test-vcluster-schema-validation"
				dns_label = "vcluster-schema-validation"
				acl_mode  = "passthrough"

				authentication = [
					{
						anonymous = {}
					}
				]

				destination = {
					id = konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster_schema_validation.id
				}
			}
		`)
		require.NoError(t, err)

		// Wire gateway → virtual cluster
		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		// Event Gateway Consume Policy Schema Validation - Main resource
		consumePolicySchemaValidation, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_schema_validation" "test_schema_validation_policy" {
				name        = "test-consume-schema-validation-policy"
				description = "Test consume policy for schema validation"
				condition   = "context.topic.name == 'validated-topic'"
				enabled     = true

				config = {
					type                     = "json"
					key_validation_action    = "mark"
					value_validation_action  = "skip"
				}
			}
		`)
		require.NoError(t, err)

		consumePolicySchemaValidation.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)
		consumePolicySchemaValidation.AddAttribute(
			"virtual_cluster_id",
			egwVirtualCluster.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySchemaValidation).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.tf_test_event_gateway_schema_validation", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster_schema_validation", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_virtual_cluster.tf_test_egw_virtual_cluster_schema_validation", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "name", "test-consume-schema-validation-policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "description", "Test consume policy for schema validation"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "enabled", "true"),
						resource.TestCheckResourceAttrSet("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "id"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySchemaValidation).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySchemaValidation.AddAttribute("labels", `{ env = "test", validation = "strict" }`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "labels.%", "2"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "labels.env", "test"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_schema_validation.test_schema_validation_policy", "labels.validation", "strict"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("EGW Consume Policy Skip Record", func(t *testing.T) {

		builder := hclbuilder.NewWithProvider(
			hclbuilder.KonnectBeta,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.KonnectBeta

		// Event Gateway (Control Plane) - Parent resource
		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf_test_event_gateway_skip_record" {
				name = "tf_test_event_gateway_skip_record"
			}
		`)
		require.NoError(t, err)

		// Event Gateway Backend Cluster - Parent resource
		egwBackendCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_backend_cluster" "tf_test_egw_backend_cluster_skip_record" {
				name = "tf_test_egw_backend_cluster_skip_record"
				authentication = {
					anonymous = {}
				}
				bootstrap_servers = [
					"10.0.0.1:9092"
				]
				tls = {
					enabled = false
					insecure_skip_verify = true
				}
			}
		`)
		require.NoError(t, err)

		// Wire gateway → backend cluster
		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		// Event Gateway Virtual Cluster - Parent resource
		egwVirtualCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster" "tf_test_egw_virtual_cluster_skip_record" {
				name      = "tf_test_egw_virtual_cluster_skip_record"
				dns_label = "vcluster-skip-record"
				acl_mode  = "passthrough"

				authentication = [
					{
						anonymous = {}
					}
				]

				destination = {
					id = konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster_skip_record.id
				}
			}
		`)
		require.NoError(t, err)

		// Wire gateway → virtual cluster
		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		// Event Gateway Consume Policy Skip Record - Main resource
		consumePolicySkipRecord, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_consume_policy_skip_record" "test_skip_record_policy" {
				name        = "test_skip_record_policy"
				description = "Test consume policy for skipping records"
				condition   = "context.topic.name == 'skip-topic'"
				enabled     = true
			}
		`)
		require.NoError(t, err)

		consumePolicySkipRecord.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)
		consumePolicySkipRecord.AddAttribute(
			"virtual_cluster_id",
			egwVirtualCluster.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySkipRecord).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.tf_test_event_gateway_skip_record", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster_skip_record", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_virtual_cluster.tf_test_egw_virtual_cluster_skip_record", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", "name", "test_skip_record_policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", "description", "Test consume policy for skipping records"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", "condition", "context.topic.name == 'skip-topic'"),
						resource.TestCheckResourceAttrSet("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", "id"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySkipRecord).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(consumePolicySkipRecord.AddAttribute("condition", `context.topic.name == 'new-skip-topic'`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_skip_record.test_skip_record_policy", "condition", "context.topic.name == 'new-skip-topic'"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	//	t.Run("EGW Data Plane Certificate", func(t *testing.T) {
	//		builder := hclbuilder.NewWithProvider(
	//			hclbuilder.KonnectBeta,
	//			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
	//		)
	//		builder.ProviderProperty = hclbuilder.KonnectBeta
	//
	//		// Event Gateway (Control Plane) - Parent resource
	//		egwCp, err := hclbuilder.FromString(`
	//			resource "konnect_event_gateway" "tf_test_event_gateway_certificate" {
	//				name = "test-gateway-certificate"
	//			}
	//		`)
	//		require.NoError(t, err)
	//
	//		// Event Gateway Data Plane Certificate - Main resource
	//		// Using a valid self-signed test certificate (PEM format)
	//		dataplaneCertificate, err := hclbuilder.FromString(`
	//resource "konnect_event_gateway_data_plane_certificate" "test_certificate" {
	//  name        = "test-data-plane-certificate"
	//  description = "Test certificate for data plane"
	//
	//  certificate = <<EOF
	//-----BEGIN CERTIFICATE-----
	//MIIBtjCCAVugAwIBAgIUH7xT2dL5z5wWqvZkYqQ1yYbG5cUwCgYIKoZIzj0EAwIw
	//EzERMA8GA1UEAwwIdGVzdC1jZXJ0MB4XDTI0MDEwMTAwMDAwMFoXDTM0MDEwMTAw
	//MDAwMFowEzERMA8GA1UEAwwIdGVzdC1jZXJ0MFkwEwYHKoZIzj0CAQYIKoZIzj0D
	//AQcDQgAEkG+8k5X2nN3H+V5+0U7aV4j5xO0lF5sJ2M3PqZrV4s0cVtW1Lr0F6d7C
	//mHnQ6B+9m0PZb+7E9B8F1z6kWqNTMFEwHQYDVR0OBBYEFJm7x9Z2l5+f4E0p5xVZ
	//xC9W0M0JMB8GA1UdIwQYMBaAFJm7x9Z2l5+f4E0p5xVZxC9W0M0JMA8GA1UdEwEB
	///wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIhAPb9xYl3KJm3E+J6uVZkV5c9Wv3D
	//d5+1m+7zZ5pBv9kRAiAQf5gWz8g7J9c8C6QJY0U6hJ4qvZ3qZ7nZxW4Y9qzQ==
	//-----END CERTIFICATE-----
	//EOF
	//}
	//`)
	//		require.NoError(t, err)
	//
	//		// Wire certificate to gateway
	//		dataplaneCertificate.AddAttribute(
	//			"gateway_id",
	//			egwCp.ResourcePath()+".id",
	//		)
	//
	//		resource.Test(t, resource.TestCase{
	//			ProtoV6ProviderFactories: providerFactory,
	//			Steps: []resource.TestStep{
	//				// Create all resources
	//				{
	//					Config: builder.
	//						Upsert(egwCp).
	//						Upsert(dataplaneCertificate).
	//						Build(),
	//
	//					ConfigPlanChecks: resource.ConfigPlanChecks{
	//						PreApply: []plancheck.PlanCheck{
	//							plancheck.ExpectResourceAction(
	//								"konnect_event_gateway.tf_test_event_gateway_certificate",
	//								plancheck.ResourceActionCreate,
	//							),
	//							plancheck.ExpectResourceAction(
	//								"konnect_event_gateway_data_plane_certificate.test_certificate",
	//								plancheck.ResourceActionCreate,
	//							),
	//						},
	//					},
	//
	//					Check: resource.ComposeAggregateTestCheckFunc(
	//						resource.TestCheckResourceAttr(
	//							"konnect_event_gateway_data_plane_certificate.test_certificate",
	//							"name",
	//							"test-data-plane-certificate",
	//						),
	//						resource.TestCheckResourceAttr(
	//							"konnect_event_gateway_data_plane_certificate.test_certificate",
	//							"description",
	//							"Test certificate for data plane",
	//						),
	//						resource.TestCheckResourceAttrSet(
	//							"konnect_event_gateway_data_plane_certificate.test_certificate",
	//							"id",
	//						),
	//						resource.TestCheckResourceAttrSet(
	//							"konnect_event_gateway_data_plane_certificate.test_certificate",
	//							"metadata.sha256_fingerprint",
	//						),
	//						resource.TestCheckResourceAttrSet(
	//							"konnect_event_gateway_data_plane_certificate.test_certificate",
	//							"metadata.subject",
	//						),
	//						resource.TestCheckResourceAttrSet(
	//							"konnect_event_gateway_data_plane_certificate.test_certificate",
	//							"metadata.issuer",
	//						),
	//					),
	//				},
	//				// Update certificate description
	//				{
	//					Config: builder.
	//						Upsert(egwCp).
	//						Upsert(
	//							dataplaneCertificate.AddAttribute("description", `Updated certificate for data plane`),
	//						).
	//						Build(),
	//
	//					Check: resource.ComposeAggregateTestCheckFunc(
	//						resource.TestCheckResourceAttr(
	//							"konnect_event_gateway_data_plane_certificate.test_certificate",
	//							"description",
	//							"Updated certificate for data plane",
	//						),
	//					),
	//				},
	//				// Revert description
	//				{
	//					Config: builder.
	//						Upsert(egwCp).
	//						Upsert(
	//							dataplaneCertificate.AddAttribute("description", `Test certificate for data plane`),
	//						).
	//						Build(),
	//
	//					Check: resource.ComposeAggregateTestCheckFunc(
	//						resource.TestCheckResourceAttr(
	//							"konnect_event_gateway_data_plane_certificate.test_certificate",
	//							"description",
	//							"Test certificate for data plane",
	//						),
	//					),
	//				},
	//			},
	//		})
	//	}) // invalid certificate.

	t.Run("EGW_Listener", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.KonnectBeta,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.KonnectBeta
		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf_test_event_gateway_listener" {
				name = "test-gateway-listener"
			}
		`)
		require.NoError(t, err)
		egwListener, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener" "test_listener" {
				name        = "test-listener"
				description = "Test listener configuration"
				addresses   = ["0.0.0.0"]
				ports       = ["9092"]
			}
		`)
		require.NoError(t, err)
		egwListener.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(egwListener).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_event_gateway.tf_test_event_gateway_listener",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectResourceAction(
								"konnect_event_gateway_listener.test_listener",
								plancheck.ResourceActionCreate,
							),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"name",
							"test-listener",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"description",
							"Test listener configuration",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"addresses.0",
							"0.0.0.0",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"ports.0",
							"9092",
						),
						resource.TestCheckResourceAttrSet(
							"konnect_event_gateway_listener.test_listener",
							"id",
						),
						resource.TestCheckResourceAttrSet(
							"konnect_event_gateway_listener.test_listener",
							"created_at",
						),
						resource.TestCheckResourceAttrSet(
							"konnect_event_gateway_listener.test_listener",
							"updated_at",
						),
					),
				},
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(egwListener).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(
							egwListener.AddAttribute("description", `Test listener configuration`),
						).
						Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"description",
							"Test listener configuration",
						),
					),
				},
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(
							egwListener.AddAttribute("ports", `["9092", "9093", "9094"]`),
						).
						Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"ports.#",
							"3",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"ports.0",
							"9092",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"ports.1",
							"9093",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"ports.2",
							"9094",
						),
					),
				},
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(
							egwListener.AddAttribute("labels", `{env = "test", team = "infrastructure"}`),
						).
						Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"labels.%",
							"2",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"labels.env",
							"test",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"labels.team",
							"infrastructure",
						),
					),
				},
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(
							egwListener.AddAttribute("addresses", `["localhost"]`),
						).
						Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"addresses.0",
							"localhost",
						),
					),
				},
				// Revert to original configuration
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(egwListener).
						Build(),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"name",
							"test-listener",
						),
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_listener.test_listener",
							"description",
							"Test listener configuration",
						),
					),
				},
			},
		})
	})

	/**We are getting unexpected error after running plan twice*/
	t.Run("EGW_Listener_Policy_Forward_To_Virtual_Cluster", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf_test_event_gateway_listener_policy" {
				name = "test-gateway-listener-policy"
			}
		`)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_backend_cluster" "tf_test_egw_backend_cluster" {
				name = "tf_test_egw_backend_cluster"
				authentication = {
					anonymous = {}
				}
				bootstrap_servers = [
					"10.0.0.1:9092"
				]
				tls = {
					enabled = false
					insecure_skip_verify = true
				}
			}
		`)
		require.NoError(t, err)

		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		egwVirtualCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster" "tf_test_egw_virtual_cluster" {
				name      = "tf_test_egw_virtual_cluster"
				dns_label = "vcluster-listener-policy"
				acl_mode  = "passthrough"

				authentication = [
					{
						anonymous = {}
					}
				]

				destination = {
					id = konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster.id
				}
			}
		`)
		require.NoError(t, err)

		// Wire virtual cluster to gateway
		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		// Event Gateway Listener - Parent resource
		egwListener, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener" "tf_test_egw_listener" {
				name      = "tf_test_egw_listener"
				addresses = ["0.0.0.0"]
				ports     = ["9092", "9093", "9094", "9095"]
				labels    = {}

				lifecycle {
					ignore_changes = [labels]
				}
				description = "Description for listener"
			}
		`)
		require.NoError(t, err)

		// Wire listener to gateway
		egwListener.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		// Event Gateway Listener Policy Forward To Virtual Cluster - Main resource
		listenerPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener_policy_forward_to_virtual_cluster" "tf_test_egw_forward_policy" {
				name        = "tf_test_egw_forward_policy"
				description = "Test forward to virtual cluster policy"
				enabled     = true

				config = {
					port_mapping = {
						advertised_host = "broker.example.com"
						bootstrap_port  = "at_start"
						min_broker_id   = 0
						destination = {
							virtual_cluster_reference_by_name = {
								name = konnect_event_gateway_virtual_cluster.tf_test_egw_virtual_cluster.name
							}
						}
					}
				}
			}
		`)
		require.NoError(t, err)

		// Wire policy to listener and gateway
		listenerPolicy.AddAttribute(
			"event_gateway_listener_id",
			egwListener.ResourcePath()+".id",
		)
		listenerPolicy.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(egwListener).Upsert(listenerPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.tf_test_event_gateway_listener_policy", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_virtual_cluster.tf_test_egw_virtual_cluster", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_listener.tf_test_egw_listener", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "name", "tf_test_egw_forward_policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "description", "Test forward to virtual cluster policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "enabled", "true"),
						resource.TestCheckResourceAttrSet("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.tf_test_egw_forward_policy", "id"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(egwListener).Upsert(listenerPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("EGW_Listener_Policy_TLS_Server", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.KonnectBeta,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.KonnectBeta

		// Event Gateway (Control Plane) - Parent resource
		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf_test_event_gateway_tls" {
				name = "test-gateway-tls-server"
			}
		`)
		require.NoError(t, err)

		// Event Gateway Listener - Parent resource
		egwListener, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener" "test_listener_tls" {
				name      = "test-listener-tls"
				addresses = ["0.0.0.0"]
				ports     = ["9092", "9093"]
				labels    = {}
				description = "Listener for TLS server policy"

				lifecycle {
					ignore_changes = [labels]
				}
			}
		`)
		require.NoError(t, err)

		// Wire listener to gateway
		egwListener.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		// Event Gateway Listener Policy TLS Server - Main resource
		tlsServerPolicy, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_listener_policy_tls_server" "test_tls_policy" {
				name        = "test-tls-server-policy"
				description = "Test TLS server policy for encryption"
				enabled     = true

				config = {
					allow_plaintext = false
					certificates = [
						{
							certificate = "-----BEGIN CERTIFICATE-----\nMIIC7jCCAd agAwIBAgIUWcN6i/K3PaY5E9O+V1YQUAECE4YwDQYJKoZIhvcNAQEL\nBQAwEzERMA8GA1UEAwwIdGVzdC1jZXJ0MB4XDTI0MDEwMTAwMDAwMFoXDTM0MDEw\nMTAwMDAwMFowEzERMA8GA1UEAwwIdGVzdC1jZXJ0MIIBIjANBgkqhkiG9w0BAQEF\nAAOCAQ8AMIIBCgKCAQEAyL2kR4u5nE5+9Z0qQKpZ9l+XKZcE8GmU5lZ5vCjYV+Zp\nY5dQy3FZ5Qqz1zv7F0s8g5mC3Y8d+RZy5P8FJ1h3oN6m8F+9y5HqYxKp5nJq9yR\np6Zy1x3Z5yq7H8x5p0Z9YkJ1+2p5ZK3ZJr9X5yH8E8R5Y8ZK8y5+Z9QIDAQAB\no1MwUTAdBgNVHQ4EFgQUy0v7uG7Nn6Jw+uXxvR4g6Y7K6H4wHwYDVR0jBBgwFoAU\ny0v7uG7Nn6Jw+uXxvR4g6Y7K6H4wDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0B\nAQsFAAOCAQEAQq9x5X8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5\nZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8\nZ5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z\n9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5\nY8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5Y8=\n-----END CERTIFICATE-----"
							key = "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDIvaRHi7mcTn71\nnSpAqln2X5cplwTwaZTmVnm8KNhX5mljl1DLcVnlCrPXO/sXSzyDmYLdjx35FnLk\n/wUnWHeg3qbwX73HkepjEqnmcmr3JGnpnLXHdnnKrsfzHmnRn1iQnX7anlkrdk\nm/1fnIdwTxHljxkrzLn5n1D1AgMBAAECggEBAKzZ5c8V9PX8e7Z5Y+Jp5+qZ8x\n5ZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK\n9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8\nZ5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5\n+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z\n9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5Y8E7Z5Y+Jp5+qZ8x5ZK9Y8Z5+Z9x5\nY8ECgYEA6L2mR4u5nE5+9Z0qQKpZ9l+XKZcE8GmU5lZ5vCjYV+ZpY5dQy3FZ5Qq\nz1zv7F0s8g5mC3Y8d+RZy5P8FJ1h3oN6m8F+9y5HqYxKp5nJq9yRp6Zy1x3Z5y\nq7H8x5p0Z9YkJ1+2p5ZK3ZJr9X5yH8E8R5Y8ZK8y5+Z9QIDAQAB\n-----END PRIVATE KEY-----"
						}
					]
					versions = {
						min = "TLSv1.2"
						max = "TLSv1.3"
					}
				}
			}
		`)
		require.NoError(t, err)

		// Wire policy to listener and gateway
		tlsServerPolicy.AddAttribute(
			"event_gateway_listener_id",
			egwListener.ResourcePath()+".id",
		)
		tlsServerPolicy.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwListener).Upsert(tlsServerPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.tf_test_event_gateway_tls", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_listener.test_listener_tls", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_listener_policy_tls_server.test_tls_policy", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_tls_server.test_tls_policy", "name", "test-tls-server-policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_tls_server.test_tls_policy", "description", "Test TLS server policy for encryption"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_tls_server.test_tls_policy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_tls_server.test_tls_policy", "config.allow_plaintext", "false"),
						resource.TestCheckResourceAttrSet("konnect_event_gateway_listener_policy_tls_server.test_tls_policy", "id"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwListener).Upsert(tlsServerPolicy).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwListener.AddAttribute("description", "Update description")).Build(),
					Check:  resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttr("konnect_event_gateway_listener.test_listener_tls", "description", "Update description")),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("EGW_Produce_Policy_Encrypt", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.KonnectBeta,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.KonnectBeta

		// Event Gateway (Control Plane) - Parent resource
		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf_test_event_gateway_produce_encrypt" {
				name = "tf_test_event_gateway_produce_encrypt"
			}
		`)
		require.NoError(t, err)

		// Event Gateway Backend Cluster - Parent resource
		egwBackendCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_backend_cluster" "tf_test_egw_backend_cluster_produce_encrypt" {
				name = "tf_test_egw_backend_cluster_produce_encrypt"
				authentication = {
					anonymous = {}
				}
				bootstrap_servers = [
					"10.0.0.1:9092"
				]
				tls = {
					enabled = false
					insecure_skip_verify = true
				}
			}
		`)
		require.NoError(t, err)

		// Wire gateway → backend cluster
		egwBackendCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		// Event Gateway Virtual Cluster - Parent resource
		egwVirtualCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster" "tf_test_egw_virtual_cluster_produce_encrypt" {
				name      = "tf_test_egw_virtual_cluster_produce_encrypt"
				dns_label = "vcluster-produce-encrypt"
				acl_mode  = "passthrough"

				authentication = [
					{
						anonymous = {}
					}
				]

				destination = {
					id = konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster_produce_encrypt.id
				}
			}
		`)
		require.NoError(t, err)

		// Wire gateway → virtual cluster
		egwVirtualCluster.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		// Event Gateway Static Key - Parent resource for encryption
		staticKey, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_static_key" "test_encryption_key" {
				name  = "test-encryption-key"
				value = "$${vault.env[\"MY_ENV_VAR\"]}"
			}
		`)
		require.NoError(t, err)

		staticKey.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		producePolicyEncrypt, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_encrypt" "test_produce_encrypt_policy" {
				name        = "test-produce-encrypt-policy"
				description = "Test produce policy for encryption"
				condition   = "context.topic.name == 'encrypted-topic'"
				enabled     = true

				config = {
					failure_mode = "error"
					encryption_key = {
						static = {
							key = {
								reference_by_name = {
									name = konnect_event_gateway_static_key.test_encryption_key.name
								}
							}
						}
					}
					part_of_record = ["value"]
				}
			}
		`)
		require.NoError(t, err)

		// Wire policy to gateway and virtual cluster
		producePolicyEncrypt.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)
		producePolicyEncrypt.AddAttribute(
			"virtual_cluster_id",
			egwVirtualCluster.ResourcePath()+".id",
		)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				// Create all resources
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(staticKey).Upsert(producePolicyEncrypt).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("konnect_event_gateway.tf_test_event_gateway_produce_encrypt", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster_produce_encrypt", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_virtual_cluster.tf_test_egw_virtual_cluster_produce_encrypt", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_static_key.test_encryption_key", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_produce_policy_encrypt.test_produce_encrypt_policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt.test_produce_encrypt_policy", "name", "test-produce-encrypt-policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt.test_produce_encrypt_policy", "description", "Test produce policy for encryption"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt.test_produce_encrypt_policy", "enabled", "true"),
						resource.TestCheckResourceAttrSet("konnect_event_gateway_produce_policy_encrypt.test_produce_encrypt_policy", "id"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(staticKey).Upsert(producePolicyEncrypt).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Upsert(egwVirtualCluster).Upsert(staticKey).Upsert(producePolicyEncrypt.AddAttribute("description", "Updated Test produce policy for encryption")).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_encrypt.test_produce_encrypt_policy", "description", "Updated Test produce policy for encryption")),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("EGW_Produce_Policy_Modify_Headers", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.KonnectBeta,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.KonnectBeta

		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "test-gateway-produce-headers" {
			  name = "test-gateway-produce-headers"
			}
		`)
		require.NoError(t, err)

		backendCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_backend_cluster" "test-backend-cluster-produce-headers" {
			  name = "test-backend-cluster-produce-headers"
			  authentication = {
				anonymous = {}
			  }
			  bootstrap_servers = ["10.0.0.1:9092"]
			  tls = {
				enabled = false
				insecure_skip_verify = true
			  }
			  gateway_id = konnect_event_gateway.test-gateway-produce-headers.id
			}
		`)
		require.NoError(t, err)

		virtualCluster, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_virtual_cluster" "test-vcluster-produce-headers" {
			  name      = "test-vcluster-produce-headers"
			  dns_label = "vcluster-produce-headers"
			  acl_mode  = "passthrough"

			  authentication = [{
				anonymous = {}
			  }]

			  destination = {
				id = konnect_event_gateway_backend_cluster.test-backend-cluster-produce-headers.id
			  }

			  gateway_id = konnect_event_gateway.test-gateway-produce-headers.id
			}
		`)
		require.NoError(t, err)

		policyCreate, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_produce_policy_modify_headers" "test-produce-headers-policy" {
			  name        = "test-produce-headers-policy"
			  description = "Test produce policy for modifying headers"
			  condition   = "context.topic.name == 'header-topic'"
			  enabled     = true

			  config = {
				actions = [
				  {
					set = {
					  key   = "X-Producer-Header"
					  value = "producer-value"
					}
				  },
				  {
					remove = {
					  key = "X-Remove-Header"
					}
				  }
				]
			  }

			  gateway_id         = konnect_event_gateway.test-gateway-produce-headers.id
			  virtual_cluster_id = konnect_event_gateway_virtual_cluster.test-vcluster-produce-headers.id
			}
		`)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.
						Upsert(egwCp).
						Upsert(backendCluster).
						Upsert(virtualCluster).
						Upsert(policyCreate).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.test-gateway-produce-headers", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_backend_cluster.test-backend-cluster-produce-headers", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_virtual_cluster.test-vcluster-produce-headers", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", "name", "test-produce-headers-policy"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", "description", "Test produce policy for modifying headers"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", "condition", "context.topic.name == 'header-topic'"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", "enabled", "true"),
						resource.TestCheckResourceAttrSet("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", "id"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(backendCluster).Upsert(virtualCluster).Upsert(policyCreate).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(backendCluster).Upsert(virtualCluster).Upsert(policyCreate.AddAttribute("description", "Updated Test produce policy for modifying headers")).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.test-produce-headers-policy", "description", "Updated Test produce policy for modifying headers")),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("EGW_Static_Key", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.KonnectBeta,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.KonnectBeta

		// Parent Gateway
		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "test_static_key_gateway" {
			  name = "test-event-gateway-static-key"
			}
		`)
		require.NoError(t, err)

		// Static Key (initial)
		staticKey, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_static_key" "test_egw_static_key" {
			  name  = "test_egw_static_key"
			  value = "$${vault.env['MY_ENV_VAR']}"
			}
		`)
		require.NoError(t, err)

		staticKey.AddAttribute(
			"gateway_id",
			egwCp.ResourcePath()+".id",
		)

		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{

				{
					Config: builder.
						Upsert(egwCp).
						Upsert(staticKey).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.test_static_key_gateway", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_static_key.test_egw_static_key", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_static_key.test_egw_static_key", "name", "test_egw_static_key"),
						resource.TestCheckResourceAttrSet("konnect_event_gateway_static_key.test_egw_static_key", "id"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(staticKey).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(staticKey.AddAttribute("labels", `{key = "value"}`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_static_key.test_egw_static_key", "labels.%", "1"),
						resource.TestCheckResourceAttr("konnect_event_gateway_static_key.test_egw_static_key", "labels.key", "value"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("EGW_Schema_Registry_Confluent", func(t *testing.T) {

		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		egw, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf-test-egw-schema-registry" {
			  name = "tf-test-egw-schema-registry"
			}
		`)
		require.NoError(t, err)

		schemaRegistry, err := hclbuilder.FromString(`
			resource "konnect_event_gateway_schema_registry" "test_schema_registry_confluent" {
			  gateway_id = konnect_event_gateway.tf-test-egw-schema-registry.id

			  confluent = {
				name        = "test_schema_registry_confluent"
				description = "test_schema_registry_confluent_description"

				labels = {
				  env = "test"
				}

				config = {
				  endpoint        = "https://key-hovercraft.com"
				  schema_type     = "avro"
				  timeout_seconds = 8

				  authentication = {
					basic = {
					  username = "test-user"
					  password = "$${vault.env['MY_ENV_VAR']}"
					}
				  }
				}
			  }
			}
		`)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{

				{
					Config: builder.
						Upsert(egw).
						Upsert(schemaRegistry).
						Build(),

					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.tf-test-egw-schema-registry", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_schema_registry.test_schema_registry_confluent", plancheck.ResourceActionCreate),
						},
					},

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_schema_registry.test_schema_registry_confluent", "confluent.name", "test_schema_registry_confluent"),
						resource.TestCheckResourceAttr("konnect_event_gateway_schema_registry.test_schema_registry_confluent", "confluent.description", "test_schema_registry_confluent_description"),
						resource.TestCheckResourceAttr("konnect_event_gateway_schema_registry.test_schema_registry_confluent", "confluent.config.schema_type", "avro"),
						resource.TestCheckResourceAttr("konnect_event_gateway_schema_registry.test_schema_registry_confluent", "confluent.config.timeout_seconds", "8"),
						resource.TestCheckResourceAttrSet("konnect_event_gateway_schema_registry.test_schema_registry_confluent", "id"),
					),
				},
				{
					Config: builder.Upsert(egw).Upsert(schemaRegistry).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egw).Upsert(schemaRegistry.AddAttribute("confluent.description", "test_schema_registry_confluent_description")).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_event_gateway_schema_registry.test_schema_registry_confluent",
							"confluent.description",
							"test_schema_registry_confluent_description",
						),
					),
				},
				{
					Config: builder.Upsert(egw).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

}
