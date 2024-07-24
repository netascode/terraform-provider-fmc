---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_file_rules Data Source - terraform-provider-fmc"
subcategory: "Policy"
description: |-
  This data source can read the File Rules.
---

# fmc_file_rules (Data Source)

This data source can read the File Rules.

## Example Usage

```terraform
data "fmc_file_rules" "example" {
  id             = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  file_policy_id = ""
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `file_policy_id` (String) ID of the file policy(container).

### Optional

- `domain` (String) The name of the FMC domain
- `id` (String) The id of the object
- `name` (String) The name of file rule (dummy)

### Read-Only

- `action` (String) The name of file rule
- `direction` (String) The name of file rule
- `file_categories` (Attributes List) The name of file rule (see [below for nested schema](#nestedatt--file_categories))
- `file_types` (Attributes List) The name of file rule (see [below for nested schema](#nestedatt--file_types))
- `protocol` (String) The name of file rule

<a id="nestedatt--file_categories"></a>
### Nested Schema for `file_categories`

Read-Only:

- `id` (String) The name of file rule
- `name` (String) The name of file rule


<a id="nestedatt--file_types"></a>
### Nested Schema for `file_types`

Read-Only:

- `id` (String) The name of file rule
- `name` (String)