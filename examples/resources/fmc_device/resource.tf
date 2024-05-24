resource "fmc_device" "example" {
  host_name        = "device1"
  license_caps     = ["BASE"]
  reg_key          = "key1"
  access_policy_id = "fmc_access_control_policy.test.id"
}
