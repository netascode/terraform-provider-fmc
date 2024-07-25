resource "fmc_manual_nat_rules" "example" {
  nat_policy_id                  = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  nat_type                       = "STATIC"
  interface_ipv6                 = false
  fall_through                   = false
  enabled                        = true
  unidirectional                 = true
  original_source_id             = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  original_destination_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  original_source_port_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  original_destination_port_id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  translated_destination_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  translated_source_id           = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  translated_destination_port_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  translated_source_port_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  source_interface_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  destination_interface_id       = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  description                    = "description of nat rule"
}
