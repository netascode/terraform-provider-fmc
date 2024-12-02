resource "fmc_device_bgp" "example" {
  device_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  learned_route_map_id = ""
  neighbors = [
    {
      ipv4_address = "10.1.1.1"
      romote_as    = "65534"
      description  = "My BGP Peer"
    }
  ]
  maximum_paths = [
    {
    }
  ]
}
