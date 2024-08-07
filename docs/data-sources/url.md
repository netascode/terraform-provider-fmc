---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_url Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source can read the URL.
---

# fmc_url (Data Source)

This data source can read the URL.

## Example Usage

```terraform
data "fmc_url" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) The name of the FMC domain
- `id` (String) The id of the object
- `name` (String) User-created name of the resource.

### Read-Only

- `description` (String) Optional user-created description.
- `overridable` (Boolean) Indicates whether object values can be overridden.
- `url` (String) URL value.
