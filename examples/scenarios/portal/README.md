# Portal Example

This scenario creates the following resources:

* Gateway Control Plane, Service + Route
* Developer Portal
* Application Authentication Strategy
* API, API Documentation and API Specification

It then links the API to the Gateway Service and publishes the API to the portal.

## Test it yourself

```bash
git clone https://github.com/Kong/terraform-provider-konnect-beta
cd examples/scenarios/portal
terraform init
terraform apply
```