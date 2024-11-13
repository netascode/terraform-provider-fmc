resource "fmc_prefilter_policy" "example" {
  name                              = "POLICY1"
  description                       = "My prefilter policy"
  default_action                    = "BLOCK_TUNNELS"
  default_action_log_begin          = true
  default_action_log_end            = false
  default_action_send_events_to_fmc = true
}
