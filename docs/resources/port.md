---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_port Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource can manage a Port.
---

# fmc_port (Resource)

This resource can manage a Port.

## Example Usage

```terraform
resource "fmc_port" "example" {
  port        = "443"
  name        = "tcp443"
  protocol    = "TCP"
  description = "Port TCP/443 (HTTPS)"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) User-created name of the resource.
- `protocol` (String) IANA protocol number or Ethertype. This is handled differently for Transport and Network layer protocols. Transport layer protocols are identified by the IANA protocol number (e.g. 6 means TCP, and 17 means UDP). Network layer protocols are identified by the decimal form of the IEEE Registration Authority Ethertype (e.g. 2048 means IP).

### Optional

- `description` (String) Optional user-created description.
- `domain` (String) The name of the FMC domain
- `overridable` (Boolean) Indicates whether object values can be overridden.
- `port` (String) Port number in decimal for TCP or UDP. Otherwise a protocol-specific value.

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import fmc_port.example "<id>"
```
