---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_load_balancing_strategy Resource - terraform-provider-konnect-beta"
subcategory: ""
description: |-
  MeshLoadBalancingStrategy Resource
---

# konnect_mesh_load_balancing_strategy (Resource)

MeshLoadBalancingStrategy Resource

## Example Usage

```terraform
resource "konnect_mesh_load_balancing_strategy" "my_meshloadbalancingstrategy" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    target_ref = {
      kind = "Dataplane"
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
          load_balancer = {
            least_request = {
              active_request_bias = {
                integer = 10
              }
              choice_count = 4
            }
            maglev = {
              hash_policies = [
                {
                  connection = {
                    source_ip = false
                  }
                  cookie = {
                    name = "...my_name..."
                    path = "...my_path..."
                    ttl  = "...my_ttl..."
                  }
                  filter_state = {
                    key = "...my_key..."
                  }
                  header = {
                    name = "...my_name..."
                  }
                  query_parameter = {
                    name = "...my_name..."
                  }
                  terminal = false
                  type     = "Connection"
                }
              ]
              table_size = 26413
            }
            random = {
              # ...
            }
            ring_hash = {
              hash_function = "XXHash"
              hash_policies = [
                {
                  connection = {
                    source_ip = false
                  }
                  cookie = {
                    name = "...my_name..."
                    path = "...my_path..."
                    ttl  = "...my_ttl..."
                  }
                  filter_state = {
                    key = "...my_key..."
                  }
                  header = {
                    name = "...my_name..."
                  }
                  query_parameter = {
                    name = "...my_name..."
                  }
                  terminal = false
                  type     = "QueryParameter"
                }
              ]
              max_ring_size = 5614666
              min_ring_size = 623920
            }
            round_robin = {
              # ...
            }
            type = "Maglev"
          }
          locality_awareness = {
            cross_zone = {
              failover = [
                {
                  from = {
                    zones = [
                      "..."
                    ]
                  }
                  to = {
                    type = "Any"
                    zones = [
                      "..."
                    ]
                  }
                }
              ]
              failover_threshold = {
                percentage = {
                  str = "...my_str..."
                }
              }
            }
            disabled = true
            local_zone = {
              affinity_tags = [
                {
                  key    = "...my_key..."
                  weight = 7
                }
              ]
            }
          }
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
  type = "MeshLoadBalancingStrategy"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cp_id` (String) Id of the Konnect resource. Requires replacement if changed.
- `mesh` (String) name of the mesh. Requires replacement if changed.
- `name` (String) name of the MeshLoadBalancingStrategy. Requires replacement if changed.
- `spec` (Attributes) Spec is the specification of the Kuma MeshLoadBalancingStrategy resource. (see [below for nested schema](#nestedatt--spec))
- `type` (String) the type of the resource. must be "MeshLoadBalancingStrategy"

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
defined inplace. (see [below for nested schema](#nestedatt--spec--target_ref))
- `to` (Attributes List) To list makes a match between the consumed services and corresponding configurations (see [below for nested schema](#nestedatt--spec--to))

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

- `default` (Attributes) Default is a configuration specific to the group of destinations referenced in
'targetRef' (see [below for nested schema](#nestedatt--spec--to--default))
- `target_ref` (Attributes) TargetRef is a reference to the resource that represents a group of
destinations.
Not Null (see [below for nested schema](#nestedatt--spec--to--target_ref))

<a id="nestedatt--spec--to--default"></a>
### Nested Schema for `spec.to.default`

Optional:

- `load_balancer` (Attributes) LoadBalancer allows to specify load balancing algorithm. (see [below for nested schema](#nestedatt--spec--to--default--load_balancer))
- `locality_awareness` (Attributes) LocalityAwareness contains configuration for locality aware load balancing. (see [below for nested schema](#nestedatt--spec--to--default--locality_awareness))

<a id="nestedatt--spec--to--default--load_balancer"></a>
### Nested Schema for `spec.to.default.load_balancer`

Optional:

- `least_request` (Attributes) LeastRequest selects N random available hosts as specified in 'choiceCount' (2 by default)
and picks the host which has the fewest active requests (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--least_request))
- `maglev` (Attributes) Maglev implements consistent hashing to upstream hosts. Maglev can be used as
a drop in replacement for the ring hash load balancer any place in which
consistent hashing is desired. (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--maglev))
- `random` (Attributes) Random selects a random available host. The random load balancer generally
performs better than round-robin if no health checking policy is configured.
Random selection avoids bias towards the host in the set that comes after a failed host. (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--random))
- `ring_hash` (Attributes) RingHash  implements consistent hashing to upstream hosts. Each host is mapped
onto a circle (the “ring”) by hashing its address; each request is then routed
to a host by hashing some property of the request, and finding the nearest
corresponding host clockwise around the ring. (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--ring_hash))
- `round_robin` (Attributes) RoundRobin is a load balancing algorithm that distributes requests
across available upstream hosts in round-robin order. (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--round_robin))
- `type` (String) Not Null; must be one of ["RoundRobin", "LeastRequest", "RingHash", "Random", "Maglev"]

<a id="nestedatt--spec--to--default--load_balancer--least_request"></a>
### Nested Schema for `spec.to.default.load_balancer.least_request`

Optional:

- `active_request_bias` (Attributes) ActiveRequestBias refers to dynamic weights applied when hosts have varying load
balancing weights. A higher value here aggressively reduces the weight of endpoints
that are currently handling active requests. In essence, the higher the ActiveRequestBias
value, the more forcefully it reduces the load balancing weight of endpoints that are
actively serving requests. (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--least_request--active_request_bias))
- `choice_count` (Number) ChoiceCount is the number of random healthy hosts from which the host with
the fewest active requests will be chosen. Defaults to 2 so that Envoy performs
two-choice selection if the field is not set.

<a id="nestedatt--spec--to--default--load_balancer--least_request--active_request_bias"></a>
### Nested Schema for `spec.to.default.load_balancer.least_request.active_request_bias`

Optional:

- `integer` (Number)
- `str` (String)



<a id="nestedatt--spec--to--default--load_balancer--maglev"></a>
### Nested Schema for `spec.to.default.load_balancer.maglev`

Optional:

- `hash_policies` (Attributes List) HashPolicies specify a list of request/connection properties that are used to calculate a hash.
These hash policies are executed in the specified order. If a hash policy has the “terminal” attribute
set to true, and there is already a hash generated, the hash is returned immediately,
ignoring the rest of the hash policy list. (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--maglev--hash_policies))
- `table_size` (Number) The table size for Maglev hashing. Maglev aims for “minimal disruption”
rather than an absolute guarantee. Minimal disruption means that when
the set of upstream hosts change, a connection will likely be sent
to the same upstream as it was before. Increasing the table size reduces
the amount of disruption. The table size must be prime number limited to 5000011.
If it is not specified, the default is 65537.

<a id="nestedatt--spec--to--default--load_balancer--maglev--hash_policies"></a>
### Nested Schema for `spec.to.default.load_balancer.maglev.hash_policies`

Optional:

- `connection` (Attributes) (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--maglev--hash_policies--connection))
- `cookie` (Attributes) (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--maglev--hash_policies--cookie))
- `filter_state` (Attributes) (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--maglev--hash_policies--filter_state))
- `header` (Attributes) (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--maglev--hash_policies--header))
- `query_parameter` (Attributes) (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--maglev--hash_policies--query_parameter))
- `terminal` (Boolean) Terminal is a flag that short-circuits the hash computing. This field provides
a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback
to rest of the policy list”, it saves time when the terminal policy works.
If true, and there is already a hash computed, ignore rest of the list of hash polices.
- `type` (String) Not Null; must be one of ["Header", "Cookie", "Connection", "SourceIP", "QueryParameter", "FilterState"]

<a id="nestedatt--spec--to--default--load_balancer--maglev--hash_policies--connection"></a>
### Nested Schema for `spec.to.default.load_balancer.maglev.hash_policies.connection`

Optional:

- `source_ip` (Boolean) Hash on source IP address.


<a id="nestedatt--spec--to--default--load_balancer--maglev--hash_policies--cookie"></a>
### Nested Schema for `spec.to.default.load_balancer.maglev.hash_policies.cookie`

Optional:

- `name` (String) The name of the cookie that will be used to obtain the hash key. Not Null
- `path` (String) The name of the path for the cookie.
- `ttl` (String) If specified, a cookie with the TTL will be generated if the cookie is not present.


<a id="nestedatt--spec--to--default--load_balancer--maglev--hash_policies--filter_state"></a>
### Nested Schema for `spec.to.default.load_balancer.maglev.hash_policies.filter_state`

Optional:

- `key` (String) The name of the Object in the per-request filterState, which is
an Envoy::Hashable object. If there is no data associated with the key,
or the stored object is not Envoy::Hashable, no hash will be produced.
Not Null


<a id="nestedatt--spec--to--default--load_balancer--maglev--hash_policies--header"></a>
### Nested Schema for `spec.to.default.load_balancer.maglev.hash_policies.header`

Optional:

- `name` (String) The name of the request header that will be used to obtain the hash key. Not Null


<a id="nestedatt--spec--to--default--load_balancer--maglev--hash_policies--query_parameter"></a>
### Nested Schema for `spec.to.default.load_balancer.maglev.hash_policies.query_parameter`

Optional:

- `name` (String) The name of the URL query parameter that will be used to obtain the hash key.
If the parameter is not present, no hash will be produced. Query parameter names
are case-sensitive.
Not Null




<a id="nestedatt--spec--to--default--load_balancer--random"></a>
### Nested Schema for `spec.to.default.load_balancer.random`


<a id="nestedatt--spec--to--default--load_balancer--ring_hash"></a>
### Nested Schema for `spec.to.default.load_balancer.ring_hash`

Optional:

- `hash_function` (String) HashFunction is a function used to hash hosts onto the ketama ring.
The value defaults to XX_HASH. Available values – XX_HASH, MURMUR_HASH_2.
must be one of ["XXHash", "MurmurHash2"]
- `hash_policies` (Attributes List) HashPolicies specify a list of request/connection properties that are used to calculate a hash.
These hash policies are executed in the specified order. If a hash policy has the “terminal” attribute
set to true, and there is already a hash generated, the hash is returned immediately,
ignoring the rest of the hash policy list. (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies))
- `max_ring_size` (Number) Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries,
but can be lowered to further constrain resource use.
- `min_ring_size` (Number) Minimum hash ring size. The larger the ring is (that is,
the more hashes there are for each provided host) the better the request distribution
will reflect the desired weights. Defaults to 1024 entries, and limited to 8M entries.

<a id="nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies"></a>
### Nested Schema for `spec.to.default.load_balancer.ring_hash.hash_policies`

Optional:

- `connection` (Attributes) (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies--connection))
- `cookie` (Attributes) (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies--cookie))
- `filter_state` (Attributes) (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies--filter_state))
- `header` (Attributes) (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies--header))
- `query_parameter` (Attributes) (see [below for nested schema](#nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies--query_parameter))
- `terminal` (Boolean) Terminal is a flag that short-circuits the hash computing. This field provides
a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback
to rest of the policy list”, it saves time when the terminal policy works.
If true, and there is already a hash computed, ignore rest of the list of hash polices.
- `type` (String) Not Null; must be one of ["Header", "Cookie", "Connection", "SourceIP", "QueryParameter", "FilterState"]

<a id="nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies--connection"></a>
### Nested Schema for `spec.to.default.load_balancer.ring_hash.hash_policies.connection`

Optional:

- `source_ip` (Boolean) Hash on source IP address.


<a id="nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies--cookie"></a>
### Nested Schema for `spec.to.default.load_balancer.ring_hash.hash_policies.cookie`

Optional:

- `name` (String) The name of the cookie that will be used to obtain the hash key. Not Null
- `path` (String) The name of the path for the cookie.
- `ttl` (String) If specified, a cookie with the TTL will be generated if the cookie is not present.


<a id="nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies--filter_state"></a>
### Nested Schema for `spec.to.default.load_balancer.ring_hash.hash_policies.filter_state`

Optional:

- `key` (String) The name of the Object in the per-request filterState, which is
an Envoy::Hashable object. If there is no data associated with the key,
or the stored object is not Envoy::Hashable, no hash will be produced.
Not Null


<a id="nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies--header"></a>
### Nested Schema for `spec.to.default.load_balancer.ring_hash.hash_policies.header`

Optional:

- `name` (String) The name of the request header that will be used to obtain the hash key. Not Null


<a id="nestedatt--spec--to--default--load_balancer--ring_hash--hash_policies--query_parameter"></a>
### Nested Schema for `spec.to.default.load_balancer.ring_hash.hash_policies.query_parameter`

Optional:

- `name` (String) The name of the URL query parameter that will be used to obtain the hash key.
If the parameter is not present, no hash will be produced. Query parameter names
are case-sensitive.
Not Null




<a id="nestedatt--spec--to--default--load_balancer--round_robin"></a>
### Nested Schema for `spec.to.default.load_balancer.round_robin`



<a id="nestedatt--spec--to--default--locality_awareness"></a>
### Nested Schema for `spec.to.default.locality_awareness`

Optional:

- `cross_zone` (Attributes) CrossZone defines locality aware load balancing priorities when dataplane proxies inside local zone
are unavailable (see [below for nested schema](#nestedatt--spec--to--default--locality_awareness--cross_zone))
- `disabled` (Boolean) Disabled allows to disable locality-aware load balancing.
When disabled requests are distributed across all endpoints regardless of locality.
- `local_zone` (Attributes) LocalZone defines locality aware load balancing priorities between dataplane proxies inside a zone (see [below for nested schema](#nestedatt--spec--to--default--locality_awareness--local_zone))

<a id="nestedatt--spec--to--default--locality_awareness--cross_zone"></a>
### Nested Schema for `spec.to.default.locality_awareness.cross_zone`

Optional:

- `failover` (Attributes List) Failover defines list of load balancing rules in order of priority (see [below for nested schema](#nestedatt--spec--to--default--locality_awareness--cross_zone--failover))
- `failover_threshold` (Attributes) FailoverThreshold defines the percentage of live destination dataplane proxies below which load balancing to the
next priority starts.
Example: If you configure failoverThreshold to 70, and you have deployed 10 destination dataplane proxies.
Load balancing to next priority will start when number of live destination dataplane proxies drops below 7.
Default 50 (see [below for nested schema](#nestedatt--spec--to--default--locality_awareness--cross_zone--failover_threshold))

<a id="nestedatt--spec--to--default--locality_awareness--cross_zone--failover"></a>
### Nested Schema for `spec.to.default.locality_awareness.cross_zone.failover`

Optional:

- `from` (Attributes) From defines the list of zones to which the rule applies (see [below for nested schema](#nestedatt--spec--to--default--locality_awareness--cross_zone--failover--from))
- `to` (Attributes) To defines to which zones the traffic should be load balanced. Not Null (see [below for nested schema](#nestedatt--spec--to--default--locality_awareness--cross_zone--failover--to))

<a id="nestedatt--spec--to--default--locality_awareness--cross_zone--failover--from"></a>
### Nested Schema for `spec.to.default.locality_awareness.cross_zone.failover.from`

Optional:

- `zones` (List of String) Not Null


<a id="nestedatt--spec--to--default--locality_awareness--cross_zone--failover--to"></a>
### Nested Schema for `spec.to.default.locality_awareness.cross_zone.failover.to`

Optional:

- `type` (String) Type defines how target zones will be picked from available zones. Not Null; must be one of ["None", "Only", "Any", "AnyExcept"]
- `zones` (List of String)



<a id="nestedatt--spec--to--default--locality_awareness--cross_zone--failover_threshold"></a>
### Nested Schema for `spec.to.default.locality_awareness.cross_zone.failover_threshold`

Optional:

- `percentage` (Attributes) Not Null (see [below for nested schema](#nestedatt--spec--to--default--locality_awareness--cross_zone--failover_threshold--percentage))

<a id="nestedatt--spec--to--default--locality_awareness--cross_zone--failover_threshold--percentage"></a>
### Nested Schema for `spec.to.default.locality_awareness.cross_zone.failover_threshold.percentage`

Optional:

- `integer` (Number)
- `str` (String)




<a id="nestedatt--spec--to--default--locality_awareness--local_zone"></a>
### Nested Schema for `spec.to.default.locality_awareness.local_zone`

Optional:

- `affinity_tags` (Attributes List) AffinityTags list of tags for local zone load balancing. (see [below for nested schema](#nestedatt--spec--to--default--locality_awareness--local_zone--affinity_tags))

<a id="nestedatt--spec--to--default--locality_awareness--local_zone--affinity_tags"></a>
### Nested Schema for `spec.to.default.locality_awareness.local_zone.affinity_tags`

Optional:

- `key` (String) Key defines tag for which affinity is configured. Not Null
- `weight` (Number) Weight of the tag used for load balancing. The bigger the weight the bigger the priority.
Percentage of local traffic load balanced to tag is computed by dividing weight by sum of weights from all tags.
For example with two affinity tags first with weight 80 and second with weight 20,
then 80% of traffic will be redirected to the first tag, and 20% of traffic will be redirected to second one.
Setting weights is not mandatory. When weights are not set control plane will compute default weight based on list order.
Default: If you do not specify weight we will adjust them so that 90% traffic goes to first tag, 9% to next, and 1% to third and so on.





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
terraform import konnect_mesh_load_balancing_strategy.my_konnect_mesh_load_balancing_strategy "{ \"cp_id\": \"bf138ba2-c9b1-4229-b268-04d9d8a6410b\",  \"mesh\": \"\",  \"name\": \"\"}"
```
