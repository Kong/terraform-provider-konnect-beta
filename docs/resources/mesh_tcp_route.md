---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_tcp_route Resource - terraform-provider-konnect-beta"
subcategory: ""
description: |-
  MeshTCPRoute Resource
---

# konnect_mesh_tcp_route (Resource)

MeshTCPRoute Resource

## Example Usage

```terraform
resource "konnect_mesh_tcp_route" "my_meshtcproute" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    target_ref = {
      kind = "Mesh"
      labels = {
        key = "value"
      }
      mesh      = "...my_mesh..."
      name      = "...my_name..."
      namespace = "...my_namespace..."
      proxy_types = [
        "Sidecar"
      ]
      section_name = "...my_section_name..."
      tags = {
        key = "value"
      }
    }
    to = [
      {
        rules = [
          {
            default = {
              backend_refs = [
                {
                  kind = "MeshService"
                  labels = {
                    key = "value"
                  }
                  mesh      = "...my_mesh..."
                  name      = "...my_name..."
                  namespace = "...my_namespace..."
                  port      = 6
                  proxy_types = [
                    "Gateway"
                  ]
                  section_name = "...my_section_name..."
                  tags = {
                    key = "value"
                  }
                  weight = 10
                }
              ]
            }
          }
        ]
        target_ref = {
          kind = "Mesh"
          labels = {
            key = "value"
          }
          mesh      = "...my_mesh..."
          name      = "...my_name..."
          namespace = "...my_namespace..."
          proxy_types = [
            "Gateway"
          ]
          section_name = "...my_section_name..."
          tags = {
            key = "value"
          }
        }
      }
    ]
  }
  type = "MeshTCPRoute"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cp_id` (String) Id of the Konnect resource. Requires replacement if changed.
- `mesh` (String) name of the mesh. Requires replacement if changed.
- `name` (String) name of the MeshTCPRoute. Requires replacement if changed.
- `spec` (Attributes) Spec is the specification of the Kuma MeshTCPRoute resource. (see [below for nested schema](#nestedatt--spec))
- `type` (String) the type of the resource. must be "MeshTCPRoute"

### Optional

- `labels` (Map of String) The labels to help identity resources

### Read-Only

- `creation_time` (String) Time at which the resource was created
- `modification_time` (String) Time at which the resource was updated
- `warnings` (List of String) warnings is a list of warning messages to return to the requesting Kuma API clients.
Warning messages describe a problem the client making the API request should correct or be aware of.

<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `target_ref` (Attributes) TargetRef is a reference to the resource the policy takes an effect on.
The resource could be either a real store object or virtual resource
defined in-place. (see [below for nested schema](#nestedatt--spec--target_ref))
- `to` (Attributes List) To list makes a match between the consumed services and corresponding
configurations (see [below for nested schema](#nestedatt--spec--to))

<a id="nestedatt--spec--target_ref"></a>
### Nested Schema for `spec.target_ref`

Required:

- `kind` (String) Kind of the referenced resource. must be one of ["Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute", "Dataplane"]

Optional:

- `labels` (Map of String) Labels are used to select group of MeshServices that match labels. Either Labels or
Name and Namespace can be used.
- `mesh` (String) Mesh is reserved for future use to identify cross mesh resources.
- `name` (String) Name of the referenced resource. Can only be used with kinds: `MeshService`,
`MeshServiceSubset` and `MeshGatewayRoute`
- `namespace` (String) Namespace specifies the namespace of target resource. If empty only resources in policy namespace
will be targeted.
- `proxy_types` (List of String) ProxyTypes specifies the data plane types that are subject to the policy. When not specified,
all data plane types are targeted by the policy.
- `section_name` (String) SectionName is used to target specific section of resource.
For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
- `tags` (Map of String) Tags used to select a subset of proxies by tags. Can only be used with kinds
`MeshSubset` and `MeshServiceSubset`


<a id="nestedatt--spec--to"></a>
### Nested Schema for `spec.to`

Optional:

- `rules` (Attributes List) Rules contains the routing rules applies to a combination of top-level
targetRef and the targetRef in this entry.
Not Null (see [below for nested schema](#nestedatt--spec--to--rules))
- `target_ref` (Attributes) TargetRef is a reference to the resource that represents a group of
destinations.
Not Null (see [below for nested schema](#nestedatt--spec--to--target_ref))

<a id="nestedatt--spec--to--rules"></a>
### Nested Schema for `spec.to.rules`

Optional:

- `default` (Attributes) Default holds routing rules that can be merged with rules from other
policies.
Not Null (see [below for nested schema](#nestedatt--spec--to--rules--default))

<a id="nestedatt--spec--to--rules--default"></a>
### Nested Schema for `spec.to.rules.default`

Optional:

- `backend_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--to--rules--default--backend_refs))

<a id="nestedatt--spec--to--rules--default--backend_refs"></a>
### Nested Schema for `spec.to.rules.default.backend_refs`

Optional:

- `kind` (String) Kind of the referenced resource. Not Null; must be one of ["Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute", "Dataplane"]
- `labels` (Map of String) Labels are used to select group of MeshServices that match labels. Either Labels or
Name and Namespace can be used.
- `mesh` (String) Mesh is reserved for future use to identify cross mesh resources.
- `name` (String) Name of the referenced resource. Can only be used with kinds: `MeshService`,
`MeshServiceSubset` and `MeshGatewayRoute`
- `namespace` (String) Namespace specifies the namespace of target resource. If empty only resources in policy namespace
will be targeted.
- `port` (Number) Port is only supported when this ref refers to a real MeshService object
- `proxy_types` (List of String) ProxyTypes specifies the data plane types that are subject to the policy. When not specified,
all data plane types are targeted by the policy.
- `section_name` (String) SectionName is used to target specific section of resource.
For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
- `tags` (Map of String) Tags used to select a subset of proxies by tags. Can only be used with kinds
`MeshSubset` and `MeshServiceSubset`
- `weight` (Number) Default: 1




<a id="nestedatt--spec--to--target_ref"></a>
### Nested Schema for `spec.to.target_ref`

Optional:

- `kind` (String) Kind of the referenced resource. Not Null; must be one of ["Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute", "Dataplane"]
- `labels` (Map of String) Labels are used to select group of MeshServices that match labels. Either Labels or
Name and Namespace can be used.
- `mesh` (String) Mesh is reserved for future use to identify cross mesh resources.
- `name` (String) Name of the referenced resource. Can only be used with kinds: `MeshService`,
`MeshServiceSubset` and `MeshGatewayRoute`
- `namespace` (String) Namespace specifies the namespace of target resource. If empty only resources in policy namespace
will be targeted.
- `proxy_types` (List of String) ProxyTypes specifies the data plane types that are subject to the policy. When not specified,
all data plane types are targeted by the policy.
- `section_name` (String) SectionName is used to target specific section of resource.
For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
- `tags` (Map of String) Tags used to select a subset of proxies by tags. Can only be used with kinds
`MeshSubset` and `MeshServiceSubset`

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_mesh_tcp_route.my_konnect_mesh_tcp_route "{ \"cp_id\": \"bf138ba2-c9b1-4229-b268-04d9d8a6410b\",  \"mesh\": \"\",  \"name\": \"\"}"
```
