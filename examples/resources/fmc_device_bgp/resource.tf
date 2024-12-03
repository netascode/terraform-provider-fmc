resource "fmc_device_bgp" "example" {
  device_id                 = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ipv4_learned_route_map_id = ""
  ipv4_neighbors = [
    {
      ipv4_neighbor_address     = "10.1.1.1"
      pv4_neighbor_romote_as    = "65534"
      ipv4_neighbor_description = "My BGP Peer"
      ipv4_neighbor_filter_access_lists = [
        {
        }
      ]
      ipv4_neighbor_filter_route_map_lists = [
        {
        }
      ]
      ipv4_neighbor_filter_prefix_lists = [
        {
        }
      ]
      ipv4_neighbor_filter_as_path_lists = [
        {
        }
      ]
      ipv4_neighbor_generate_default_route_map          = ""
      ipv4_neighbor_routes_advertise_map                = ""
      ipv4_neighbor_routes_advertise_exist_nonexist_map = ""
      ipv4_neighbor_authentication_password             = ""
    }
  ]
  ipv4_aggregate_addresses = [
    {
    }
  ]
  ipv4_filterings = [
    {
    }
  ]
  ipv4_networks = [
    {
    }
  ]
  ipv4_redistributions = [
    {
    }
  ]
  ipv4_route_injections = [
    {
    }
  ]
}
