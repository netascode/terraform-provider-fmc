---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_extended_acl Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource can manage an Extended ACL.
---

# fmc_extended_acl (Resource)

This resource can manage an Extended ACL.

## Example Usage

```terraform
resource "fmc_extended_acl" "example" {
  name        = "extended_acl_1"
  description = "My Extended Access Control List"
  entries = [
    {
      action               = "DENY"
      log_level            = "WARNING"
      logging              = "PER_ACCESS_LIST_ENTRY"
      log_interval_seconds = 120
      source_network_literals = [
        {
          value = "10.1.1.0/24"
          type  = "Network"
        }
      ]
      destination_network_literals = [
        {
          value = "10.2.2.2"
          type  = "Host"
        }
      ]
      source_network_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      destination_network_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      source_port_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      destination_port_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `entries` (Attributes List) Ordered list of ACL's entries. (see [below for nested schema](#nestedatt--entries))
- `name` (String) User-created name of the resource.

### Optional

- `description` (String) Optional user-created description.
- `domain` (String) The name of the FMC domain

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Required:

- `action` (String) Indicates the redistribution access: PERMIT or DENY.
  - Choices: `PERMIT`, `DENY`
- `logging` (String) The logging mode.
  - Choices: `PER_ACCESS_LIST_ENTRY`, `DEFAULT`, `DISABLED`

Optional:

- `destination_network_literals` (Attributes Set) Set of objects that represent destinations of traffic (literally specified). (see [below for nested schema](#nestedatt--entries--destination_network_literals))
- `destination_network_objects` (Attributes Set) Set of objects that represent destinations of traffic (fmc_network, fmc_host, ...). (see [below for nested schema](#nestedatt--entries--destination_network_objects))
- `destination_port_objects` (Attributes Set) Set of objects representing destination ports. (see [below for nested schema](#nestedatt--entries--destination_port_objects))
- `log_interval_seconds` (Number) The logging interval in seconds. Must be left at 300 if `logging` is DEFAULT or DISABLED.
  - Default value: `300`
- `log_level` (String) The logging level. Recommended to be left at INFORMATIONAL if `logging` is DEFAULT or DISABLED.
  - Choices: `EMERGENCY`, `ALERT`, `CRITICAL`, `ERROR`, `WARNING`, `NOTIFICATION`, `INFORMATIONAL`, `DEBUGGING`
  - Default value: `INFORMATIONAL`
- `source_network_literals` (Attributes Set) Set of objects that represent sources of traffic (literally specified). (see [below for nested schema](#nestedatt--entries--source_network_literals))
- `source_network_objects` (Attributes Set) Set of objects that represent sources of traffic (fmc_network, fmc_host, ...). (see [below for nested schema](#nestedatt--entries--source_network_objects))
- `source_port_objects` (Attributes Set) Set of objects representing source ports. (see [below for nested schema](#nestedatt--entries--source_port_objects))

<a id="nestedatt--entries--destination_network_literals"></a>
### Nested Schema for `entries.destination_network_literals`

Optional:

- `type` (String)
- `value` (String)


<a id="nestedatt--entries--destination_network_objects"></a>
### Nested Schema for `entries.destination_network_objects`

Optional:

- `id` (String) UUID of the object (such as fmc_network.example.id, etc.).


<a id="nestedatt--entries--destination_port_objects"></a>
### Nested Schema for `entries.destination_port_objects`

Optional:

- `id` (String) UUID of the object (such as fmc_port.example.id).


<a id="nestedatt--entries--source_network_literals"></a>
### Nested Schema for `entries.source_network_literals`

Optional:

- `type` (String)
- `value` (String)


<a id="nestedatt--entries--source_network_objects"></a>
### Nested Schema for `entries.source_network_objects`

Optional:

- `id` (String) UUID of the object (such as fmc_network.example.id, etc.).


<a id="nestedatt--entries--source_port_objects"></a>
### Nested Schema for `entries.source_port_objects`

Optional:

- `id` (String) UUID of the object (such as fmc_port.example.id).

## Import

Import is supported using the following syntax:

```shell
terraform import fmc_extended_acl.example "<id>"
```
