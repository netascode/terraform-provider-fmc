resource "fmc_file_rules" "example" {
  file_policy_id = ""
  action         = "DETECT"
  protocol       = "ANY"
  direction      = "ANY"
  file_categories = [
    {
      id   = "2"
      name = "Archive"
    }
  ]
  file_types = [
    {
      id   = "19"
      name = "7Z"
    }
  ]
}
