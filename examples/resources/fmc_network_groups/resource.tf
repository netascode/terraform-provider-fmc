resource "fmc_network_groups" "example" {
  items = {
    net_group_1 = {
      description = "My Network Group 1"
      group_names = ["child_group_1"]
      objects = [
        {
        }
      ]
      literals = [
        {
          value = "10.1.1.0/24"
        }
      ]
    }
  }
}
