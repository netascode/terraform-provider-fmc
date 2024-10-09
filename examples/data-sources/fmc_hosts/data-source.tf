// Comments made manually, which will be removed at next go-generate

data "fmc_hosts" "example" {
//  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"  // This ID is for TF purposes only
  items = {
    "hosts_1" = {
//      id = fmc_hosts.test.items["hosts_1"].id  // We want to search by name, hence ID here is not needed
    }
  }
}
