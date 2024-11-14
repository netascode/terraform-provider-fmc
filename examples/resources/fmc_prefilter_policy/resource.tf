resource "fmc_prefilter_policy" "example" {
  name                              = "POLICY1"
  description                       = "My prefilter policy"
  default_action                    = "BLOCK_TUNNELS"
  default_action_log_begin          = true
  default_action_log_end            = false
  default_action_send_events_to_fmc = true
  default_action_syslog_config_id   = "35e197ca-33a8-11ef-b2d1-d98ae17766e7"
  default_action_snmp_config_id     = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
