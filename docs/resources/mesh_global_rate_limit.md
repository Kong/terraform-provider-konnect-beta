---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_global_rate_limit Resource - terraform-provider-konnect-beta"
subcategory: ""
description: |-
  MeshGlobalRateLimit Resource
---

# konnect_mesh_global_rate_limit (Resource)

MeshGlobalRateLimit Resource

## Example Usage

```terraform
resource "konnect_mesh_global_rate_limit" "my_meshglobalratelimit" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    from = [
      {
        default = {
          backend = {
            rate_limit_service = {
              limit_on_service_fail = true
              timeout               = "...my_timeout..."
              url                   = "...my_url..."
            }
          }
          http = {
            disabled = false
            on_rate_limit = {
              headers = {
                add = [
                  {
                    name  = "...my_name..."
                    value = "...my_value..."
                  }
                ]
                set = [
                  {
                    name  = "...my_name..."
                    value = "...my_value..."
                  }
                ]
              }
              status = 10
            }
            ratelimit_on_request = [
              {
                kind = "OnHeader"
                limits = [
                  {
                    request_rate = {
                      interval = "...my_interval..."
                      num      = 7
                    }
                    value = "...my_value..."
                  }
                ]
                name = "...my_name..."
              }
            ]
            request_rate = {
              interval = "...my_interval..."
              num      = 3
            }
          }
          mode = "Limit"
        }
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
      }
    ]
    target_ref = {
      kind = "MeshService"
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
        default = {
          backend = {
            rate_limit_service = {
              limit_on_service_fail = true
              timeout               = "...my_timeout..."
              url                   = "...my_url..."
            }
          }
          http = {
            disabled = false
            on_rate_limit = {
              headers = {
                add = [
                  {
                    name  = "...my_name..."
                    value = "...my_value..."
                  }
                ]
                set = [
                  {
                    name  = "...my_name..."
                    value = "...my_value..."
                  }
                ]
              }
              status = 1
            }
            ratelimit_on_request = [
              {
                kind = "OnHeader"
                limits = [
                  {
                    request_rate = {
                      interval = "...my_interval..."
                      num      = 8
                    }
                    value = "...my_value..."
                  }
                ]
                name = "...my_name..."
              }
            ]
            request_rate = {
              interval = "...my_interval..."
              num      = 8
            }
          }
          mode = "Shadow"
        }
        target_ref = {
          kind = "MeshGateway"
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
      }
    ]
  }
  type = "MeshGlobalRateLimit"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cp_id` (String) Id of the Konnect resource. Requires replacement if changed.
- `mesh` (String) name of the mesh. Requires replacement if changed.
- `name` (String) name of the MeshGlobalRateLimit. Requires replacement if changed.
- `spec` (Attributes) Spec is the specification of the Kuma MeshGlobalRateLimit resource. (see [below for nested schema](#nestedatt--spec))
- `type` (String) the type of the resource. must be "MeshGlobalRateLimit"

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

- `from` (Attributes List) From list makes a match between clients and corresponding configurations (see [below for nested schema](#nestedatt--spec--from))
- `target_ref` (Attributes) TargetRef is a reference to the resource the policy takes an effect on.
The resource could be either a real store object or virtual resource
defined inplace. (see [below for nested schema](#nestedatt--spec--target_ref))
- `to` (Attributes List) To list makes a match between clients and corresponding configurations (see [below for nested schema](#nestedatt--spec--to))

<a id="nestedatt--spec--from"></a>
### Nested Schema for `spec.from`

Optional:

- `default` (Attributes) Default is a configuration specific to the group of clients referenced in
'targetRef' (see [below for nested schema](#nestedatt--spec--from--default))
- `target_ref` (Attributes) TargetRef is a reference to the resource that represents a group of
clients.
Not Null (see [below for nested schema](#nestedatt--spec--from--target_ref))

<a id="nestedatt--spec--from--default"></a>
### Nested Schema for `spec.from.default`

Optional:

- `backend` (Attributes) Backend defines location of rate limit backend service. Not Null (see [below for nested schema](#nestedatt--spec--from--default--backend))
- `http` (Attributes) Not Null (see [below for nested schema](#nestedatt--spec--from--default--http))
- `mode` (String) Mode defines rate limit behavior when limits are reached. Possible options: Limit and Shadow. Setting Shadow will
not block over the limit requests but will update metrics. This is useful for testing rate limit configuration.
must be one of ["Limit", "Shadow"]

<a id="nestedatt--spec--from--default--backend"></a>
### Nested Schema for `spec.from.default.backend`

Optional:

- `rate_limit_service` (Attributes) Not Null (see [below for nested schema](#nestedatt--spec--from--default--backend--rate_limit_service))

<a id="nestedatt--spec--from--default--backend--rate_limit_service"></a>
### Nested Schema for `spec.from.default.backend.rate_limit_service`

Optional:

- `limit_on_service_fail` (Boolean) LimitOnServiceFail will pass limit requests if ratelimit service is not reachable.
- `timeout` (String) Timeout for rate limit request made form Data Plane Proxy to rate limit service.
- `url` (String) Url defines address of rate limit service.



<a id="nestedatt--spec--from--default--http"></a>
### Nested Schema for `spec.from.default.http`

Optional:

- `disabled` (Boolean) Define if rate limiting should be disabled.
- `on_rate_limit` (Attributes) Describes the actions to take on a rate limit event (see [below for nested schema](#nestedatt--spec--from--default--http--on_rate_limit))
- `ratelimit_on_request` (Attributes List) Defines rate limit based on request content (see [below for nested schema](#nestedatt--spec--from--default--http--ratelimit_on_request))
- `request_rate` (Attributes) Defines how many requests are allowed per interval. (see [below for nested schema](#nestedatt--spec--from--default--http--request_rate))

<a id="nestedatt--spec--from--default--http--on_rate_limit"></a>
### Nested Schema for `spec.from.default.http.on_rate_limit`

Optional:

- `headers` (Attributes) The Headers to be added to the HTTP response on a rate limit event (see [below for nested schema](#nestedatt--spec--from--default--http--on_rate_limit--headers))
- `status` (Number) The HTTP status code to be set on a rate limit event

<a id="nestedatt--spec--from--default--http--on_rate_limit--headers"></a>
### Nested Schema for `spec.from.default.http.on_rate_limit.headers`

Optional:

- `add` (Attributes List) (see [below for nested schema](#nestedatt--spec--from--default--http--on_rate_limit--headers--add))
- `set` (Attributes List) (see [below for nested schema](#nestedatt--spec--from--default--http--on_rate_limit--headers--set))

<a id="nestedatt--spec--from--default--http--on_rate_limit--headers--add"></a>
### Nested Schema for `spec.from.default.http.on_rate_limit.headers.add`

Optional:

- `name` (String) Not Null
- `value` (String) Not Null


<a id="nestedatt--spec--from--default--http--on_rate_limit--headers--set"></a>
### Nested Schema for `spec.from.default.http.on_rate_limit.headers.set`

Optional:

- `name` (String) Not Null
- `value` (String) Not Null




<a id="nestedatt--spec--from--default--http--ratelimit_on_request"></a>
### Nested Schema for `spec.from.default.http.ratelimit_on_request`

Optional:

- `kind` (String) Kind defines type of rate limit config. Possible options: OnHeader. Not Null; must be "OnHeader"
- `limits` (Attributes List) Limits defines limit configuration. Not Null (see [below for nested schema](#nestedatt--spec--from--default--http--ratelimit_on_request--limits))
- `name` (String) Name of the request element on which rate limit should apply. E.g. header name. Not Null

<a id="nestedatt--spec--from--default--http--ratelimit_on_request--limits"></a>
### Nested Schema for `spec.from.default.http.ratelimit_on_request.limits`

Optional:

- `request_rate` (Attributes) Defines how many requests are allowed per interval. (see [below for nested schema](#nestedatt--spec--from--default--http--ratelimit_on_request--limits--request_rate))
- `value` (String) Value of the request element on which rate limit should apply. E.g. header value. Not Null

<a id="nestedatt--spec--from--default--http--ratelimit_on_request--limits--request_rate"></a>
### Nested Schema for `spec.from.default.http.ratelimit_on_request.limits.request_rate`

Optional:

- `interval` (String) The interval the number of units is accounted for. Only 1s, 1m, 1h or 24h can be configured. Not Null
- `num` (Number) Number of units per interval (depending on usage it can be a number of requests,
or a number of connections).
Not Null




<a id="nestedatt--spec--from--default--http--request_rate"></a>
### Nested Schema for `spec.from.default.http.request_rate`

Optional:

- `interval` (String) The interval the number of units is accounted for. Only 1s, 1m, 1h or 24h can be configured. Not Null
- `num` (Number) Number of units per interval (depending on usage it can be a number of requests,
or a number of connections).
Not Null




<a id="nestedatt--spec--from--target_ref"></a>
### Nested Schema for `spec.from.target_ref`

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

- `default` (Attributes) Default is a configuration specific to the group of clients referenced in
'targetRef' (see [below for nested schema](#nestedatt--spec--to--default))
- `target_ref` (Attributes) TargetRef is a reference to the resource that represents a group of
clients.
Not Null (see [below for nested schema](#nestedatt--spec--to--target_ref))

<a id="nestedatt--spec--to--default"></a>
### Nested Schema for `spec.to.default`

Optional:

- `backend` (Attributes) Backend defines location of rate limit backend service. Not Null (see [below for nested schema](#nestedatt--spec--to--default--backend))
- `http` (Attributes) Not Null (see [below for nested schema](#nestedatt--spec--to--default--http))
- `mode` (String) Mode defines rate limit behavior when limits are reached. Possible options: Limit and Shadow. Setting Shadow will
not block over the limit requests but will update metrics. This is useful for testing rate limit configuration.
must be one of ["Limit", "Shadow"]

<a id="nestedatt--spec--to--default--backend"></a>
### Nested Schema for `spec.to.default.backend`

Optional:

- `rate_limit_service` (Attributes) Not Null (see [below for nested schema](#nestedatt--spec--to--default--backend--rate_limit_service))

<a id="nestedatt--spec--to--default--backend--rate_limit_service"></a>
### Nested Schema for `spec.to.default.backend.rate_limit_service`

Optional:

- `limit_on_service_fail` (Boolean) LimitOnServiceFail will pass limit requests if ratelimit service is not reachable.
- `timeout` (String) Timeout for rate limit request made form Data Plane Proxy to rate limit service.
- `url` (String) Url defines address of rate limit service.



<a id="nestedatt--spec--to--default--http"></a>
### Nested Schema for `spec.to.default.http`

Optional:

- `disabled` (Boolean) Define if rate limiting should be disabled.
- `on_rate_limit` (Attributes) Describes the actions to take on a rate limit event (see [below for nested schema](#nestedatt--spec--to--default--http--on_rate_limit))
- `ratelimit_on_request` (Attributes List) Defines rate limit based on request content (see [below for nested schema](#nestedatt--spec--to--default--http--ratelimit_on_request))
- `request_rate` (Attributes) Defines how many requests are allowed per interval. (see [below for nested schema](#nestedatt--spec--to--default--http--request_rate))

<a id="nestedatt--spec--to--default--http--on_rate_limit"></a>
### Nested Schema for `spec.to.default.http.on_rate_limit`

Optional:

- `headers` (Attributes) The Headers to be added to the HTTP response on a rate limit event (see [below for nested schema](#nestedatt--spec--to--default--http--on_rate_limit--headers))
- `status` (Number) The HTTP status code to be set on a rate limit event

<a id="nestedatt--spec--to--default--http--on_rate_limit--headers"></a>
### Nested Schema for `spec.to.default.http.on_rate_limit.headers`

Optional:

- `add` (Attributes List) (see [below for nested schema](#nestedatt--spec--to--default--http--on_rate_limit--headers--add))
- `set` (Attributes List) (see [below for nested schema](#nestedatt--spec--to--default--http--on_rate_limit--headers--set))

<a id="nestedatt--spec--to--default--http--on_rate_limit--headers--add"></a>
### Nested Schema for `spec.to.default.http.on_rate_limit.headers.add`

Optional:

- `name` (String) Not Null
- `value` (String) Not Null


<a id="nestedatt--spec--to--default--http--on_rate_limit--headers--set"></a>
### Nested Schema for `spec.to.default.http.on_rate_limit.headers.set`

Optional:

- `name` (String) Not Null
- `value` (String) Not Null




<a id="nestedatt--spec--to--default--http--ratelimit_on_request"></a>
### Nested Schema for `spec.to.default.http.ratelimit_on_request`

Optional:

- `kind` (String) Kind defines type of rate limit config. Possible options: OnHeader. Not Null; must be "OnHeader"
- `limits` (Attributes List) Limits defines limit configuration. Not Null (see [below for nested schema](#nestedatt--spec--to--default--http--ratelimit_on_request--limits))
- `name` (String) Name of the request element on which rate limit should apply. E.g. header name. Not Null

<a id="nestedatt--spec--to--default--http--ratelimit_on_request--limits"></a>
### Nested Schema for `spec.to.default.http.ratelimit_on_request.limits`

Optional:

- `request_rate` (Attributes) Defines how many requests are allowed per interval. (see [below for nested schema](#nestedatt--spec--to--default--http--ratelimit_on_request--limits--request_rate))
- `value` (String) Value of the request element on which rate limit should apply. E.g. header value. Not Null

<a id="nestedatt--spec--to--default--http--ratelimit_on_request--limits--request_rate"></a>
### Nested Schema for `spec.to.default.http.ratelimit_on_request.limits.request_rate`

Optional:

- `interval` (String) The interval the number of units is accounted for. Only 1s, 1m, 1h or 24h can be configured. Not Null
- `num` (Number) Number of units per interval (depending on usage it can be a number of requests,
or a number of connections).
Not Null




<a id="nestedatt--spec--to--default--http--request_rate"></a>
### Nested Schema for `spec.to.default.http.request_rate`

Optional:

- `interval` (String) The interval the number of units is accounted for. Only 1s, 1m, 1h or 24h can be configured. Not Null
- `num` (Number) Number of units per interval (depending on usage it can be a number of requests,
or a number of connections).
Not Null




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
terraform import konnect_mesh_global_rate_limit.my_konnect_mesh_global_rate_limit "{ \"cp_id\": \"bf138ba2-c9b1-4229-b268-04d9d8a6410b\",  \"mesh\": \"\",  \"name\": \"\"}"
```
