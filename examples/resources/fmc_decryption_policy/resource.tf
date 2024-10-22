resource "fmc_decryption_policy" "example" {
  name                             = "decryption_policy_test"
  decryption_errors                = "BLOCK"
  compressed_session               = "INHERIT_DEFAULT_ACTION"
  ssl_v2_session                   = "INHERIT_DEFAULT_ACTION"
  unknown_cipher_suite             = "INHERIT_DEFAULT_ACTION"
  unsupported_cipher_suite         = "INHERIT_DEFAULT_ACTION"
  session_not_cached               = "INHERIT_DEFAULT_ACTION"
  handshake_errors                 = "INHERIT_DEFAULT_ACTION"
  policy_action                    = "DO_NOT_DECRYPT"
  event_log_action                 = "LOG_NONE"
  disallow_untrusted_issuer_resign = true
  strip_http3                      = false
  tls13_decryption                 = false
  quic_decryption                  = false
  adaptive_probe                   = true
  block_extensions                 = [1]
  log_end                          = false
  send_events                      = false
  trusted_cas = [
    {
      ca_name = "Cisco-Trusted-Authorities"
      ca_id   = "d3856136-b65c-49e6-ac4c-0daf5be98bc1"
    }
  ]
}
