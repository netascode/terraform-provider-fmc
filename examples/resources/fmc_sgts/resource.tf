resource "fmc_sgts" "example" {
  items = {
    sgt_1 = {
      name        = "SGT1"
      type        = ""
      description = "My SGT object"
      tag         = "11"
    }
  }
}
