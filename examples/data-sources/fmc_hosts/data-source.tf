data "fmc_hosts" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  items = {
    "hosts_1" = {
      id = fmc_hosts.test.items["hosts_1"].id
    }
  }
}
