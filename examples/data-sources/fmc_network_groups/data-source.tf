data "fmc_network_groups" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  items = {
    "net_group_1" = {
      id = fmc_network_groups.test.items["net_group_1"].id
    }
  }
}
