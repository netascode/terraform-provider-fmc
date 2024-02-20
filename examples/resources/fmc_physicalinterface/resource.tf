resource "fmc_physicalinterface" "example" {
  device_id = "<Device_ID>"
  enabled   = true
  mode      = "NONE"
  name      = "GigabitEthernet0/1"
  ifname    = "interface_name"
}
