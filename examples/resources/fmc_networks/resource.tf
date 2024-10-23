resource "fmc_networks" "example" {
  items = {
    networks_1 = {
      description = "My Network 1"
      prefix      = "10.1.1.1"
    }
  }
}
