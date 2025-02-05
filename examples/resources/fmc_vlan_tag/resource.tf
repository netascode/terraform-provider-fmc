resource "fmc_vlan_tag" "example" {
  name        = "my_vlan_tag"
  description = "My VLAN Tag"
  overridable = true
  start_tag   = "12"
  end_tag     = "15"
}
