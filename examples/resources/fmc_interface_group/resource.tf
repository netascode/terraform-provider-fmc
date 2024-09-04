resource "fmc_interface_group" "example" {
  name           = "interface_group_2"
  interface_mode = "ROUTED"
  interfaces = [
    {
      id = "0050568A-4E02-1ed3-0000-004294969198"
    }
  ]
}
