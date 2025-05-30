---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_passthrough Resource - terraform-provider-konnect-beta"
subcategory: ""
description: |-
  MeshPassthrough Resource
---

# konnect_mesh_passthrough (Resource)

MeshPassthrough Resource

## Example Usage

```terraform
resource "konnect_mesh_passthrough" "my_meshpassthrough" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    default = {
      append_match = [
        {
          port     = 6
          protocol = "mysql"
          type     = "IP"
          value    = "...my_value..."
        }
      ]
      passthrough_mode = "Matched"
    }
    target_ref = {
      kind = "MeshMultiZoneService"
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
  type = "MeshPassthrough"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cp_id` (String) Id of the Konnect resource. Requires replacement if changed.
- `mesh` (String) name of the mesh. Requires replacement if changed.
- `name` (String) name of the MeshPassthrough. Requires replacement if changed.
- `spec` (Attributes) Spec is the specification of the Kuma MeshPassthrough resource. (see [below for nested schema](#nestedatt--spec))
- `type` (String) the type of the resource. must be "MeshPassthrough"

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

- `default` (Attributes) MeshPassthrough configuration. (see [below for nested schema](#nestedatt--spec--default))
- `target_ref` (Attributes) TargetRef is a reference to the resource the policy takes an effect on.
The resource could be either a real store object or virtual resource
defined in-place. (see [below for nested schema](#nestedatt--spec--target_ref))

<a id="nestedatt--spec--default"></a>
### Nested Schema for `spec.default`

Optional:

- `append_match` (Attributes List) AppendMatch is a list of destinations that should be allowed through the sidecar. (see [below for nested schema](#nestedatt--spec--default--append_match))
- `passthrough_mode` (String) Defines the passthrough behavior. Possible values: `All`, `None`, `Matched`
When `All` or `None` `appendMatch` has no effect.
If not specified then the default value is "Matched".
must be one of ["All", "Matched", "None"]

<a id="nestedatt--spec--default--append_match"></a>
### Nested Schema for `spec.default.append_match`

Optional:

- `port` (Number) Port defines the port to which a user makes a request.
- `protocol` (String) Protocol defines the communication protocol. Possible values: `tcp`, `tls`, `grpc`, `http`, `http2`, `mysql`. Default: "tcp"; must be one of ["tcp", "tls", "grpc", "http", "http2", "mysql"]
- `type` (String) Type of the match, one of `Domain`, `IP` or `CIDR` is available. Not Null; must be one of ["Domain", "IP", "CIDR"]
- `value` (String) Value for the specified Type. Not Null



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

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_mesh_passthrough.my_konnect_mesh_passthrough "{ \"cp_id\": \"bf138ba2-c9b1-4229-b268-04d9d8a6410b\",  \"mesh\": \"\",  \"name\": \"\"}"
```
