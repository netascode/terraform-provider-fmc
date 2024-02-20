resource "fmc_device" "example" {
  name               = "FTD1"
  host_name          = "192.168.0.152"
  reg_key            = "cisco"
  license_caps       = ["MALWARE"]
  access_policy_id   = "9BCD876543456789A0"
  access_policy_name = "Test-ACP"
  type               = "Device"
}
