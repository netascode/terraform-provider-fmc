resource "fmc_device" "example" {
  name      = "MyDeviceName1"
  host_name = "10.0.0.1"
  license_capabilities = [
    {
    }
  ]
  registration_key    = "key1"
  performance_tier    = "FTDv5"
  snort_engine        = "SNORT3"
  object_group_search = true
  access_policy_id    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
