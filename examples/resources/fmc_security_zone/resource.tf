resource "fmc_security_zone" "example" {
  name           = "security_zone_1"
  description    = "My description"
  interface_mode = "ROUTED"
}
