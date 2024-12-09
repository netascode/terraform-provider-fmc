resource "fmc_device_ha_pair_monitoring" "example" {
  device_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  logical_name         = "outside"
  monitor_interface    = true
  ipv4_standby_address = "10.1.1.2"
}
