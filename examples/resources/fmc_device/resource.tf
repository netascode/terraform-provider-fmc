resource "fmc_device" "example" {
  name             = "device1"
  host_name        = "10.0.0.1"
  license_caps     = ["BASE"]
  reg_key          = "key1"
  access_policy_id = "fmc_access_control_policy.test.id"
}
