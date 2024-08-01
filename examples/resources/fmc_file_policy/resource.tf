resource "fmc_file_policy" "example" {
  name                     = "file_policy_1"
  description              = "My file policy"
  archive_depth            = "3"
  archive_depth_action     = true
  block_encrypted_archives = false
  clean_list               = true
  custom_detection_list    = true
  first_time_file_analysis = true
  inspect_archives         = false
  threat_score             = "High"
}
