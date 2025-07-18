# Changelog

## 0.7.2
> Released on 2025/07/11

### Bug Fixes
* fix default mesh hook that did not set features (MeshCreation, HostnameGeneratorCreation) and "skip-initial-policies" for a new Mesh Control Plane and Mesh

## 0.7.1
> Released on 2025/07/10

### BREAKING CHANGES
* `labels` field is now removed from `konnect_api_document`.

## 0.7.0
> Released on 2025/06/25

### BREAKING CHANGES
* `auth_strategy_sync_error` read only field removed from `konnect_api`.

### Features
* `current_version_summary` is now available on the `konnect_api` resource
* `attributes` can now be set on the `konnect_api` resource

### Bug Fixes
* `service` is now optional in `konnect_api_implementation`

## 0.6.1
> Released on 2025/06/12

### Bug Fixes
* Fix empty labels in Mesh resources causing an error

## 0.6.0
> Released on 2025/06/05

### BREAKING CHANGES
* The `deprecated` field has been removed from the `konnect_api` resource

### Features
* Add support for the Konnect IDP server. Adds the `konnect_auth_server`, `konnect_auth_server_claims`, `konnect_auth_server_scopes`, and `konnect_auth_server_clients` resources
* A new `konnect_api_version` resource has been added. This is a replacement for the deprecated `konnect_api_specification` resource.
* Add retries to all 404 returned from the API to combat replication lag
* Add retries to other retriable status codes (408, 429, 500, 502, 503, 504)
* Add retries to 400 returned from the API on DELETE Mesh Control Plane to wait for the zone to disconnect

## 0.5.2

> Released on 2025/06/05

### Bug Fixes
* Fix `terraform apply` for the `konnect_api` resource by ignoring the removed `deprecated` field

### Update dependencies
* Bump `shared-speakeasy/hooks/mesh_defaults` to `v0.0.3`

## 0.5.1
> Released on 2025/06/05

### Bug fixes

* Correctly define dependencies in gen.yaml under `terraform` and not `go`

## 0.5.0
> Released on 2025/05/27

### Features
* Use `kumalabels` custom type in Mesh resources to ignore autogenerated labels

## 0.4.1
> Released on 2025/05/21

### Bug fixes

* Updates generator to capture fixes for mismatched attribute names
 
## 0.4.0
> Released on 2025/05/21

### Features
* Add `MeshControlPlanes` data source to list all control planes
* Add a meaningful message when a user forgets to import a resource

## 0.3.0
> Released on 2025/05/19

### BREAKING CHANGES

* Removed `public_labels` from `konnect_api`

### Features

* Add `spec_content` to `konnect_api` to allow publishing of specifications when creating an API
* Add `force_destroy` option to `konnect_portal` to allow deletion of Portals with child resources


## 0.2.1
> Released on 2025/04/29

### Improvements

* Use common test cases for Kong Mesh and Konnect via shared-speakeasy

## 0.2.0
> Released on 2025/04/28

### Features

* Add support for OAS2 docs in Portal

### Bug fixes

* Fix plan output for non-computed entities in Mesh
* Support sending empty arrays in Mesh

## 0.1.3
> Released on 2025/04/24

### Bug fixes

* Fix more "Replace if changed" logic for Mesh entities
* Add integration tests

## 0.1.2
> Released on 2025/04/22

### Bug fixes

* Fix "Replace if changed" logic for Mesh entities

## 0.1.1
> Released on 2025/03/19

### Bug fixes

* Switch to PUT rather than PATCH when managing Mesh entities

## 0.1.0
> Released on 2025/03/11

### Features

* Beta support for Kong Mesh
* Beta support for Developer Portal v3

