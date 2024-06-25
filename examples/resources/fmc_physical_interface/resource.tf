resource "fmc_physical_interface" "example" {
  device_id    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  mode         = "NONE"
  name         = "GigabitEthernet0/1"
  logical_name = "myinterface-0-1"
}
