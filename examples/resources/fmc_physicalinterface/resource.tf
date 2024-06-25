resource "fmc_physicalinterface" "example" {
  device_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  mode      = "NONE"
  name      = "GigabitEthernet0/1"
  ifname    = "interface_name"
}
