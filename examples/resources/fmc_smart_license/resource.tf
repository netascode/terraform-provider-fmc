// Enable Evaluation Mode
resource "fmc_smart_license" "license" {
  registration_type = "EVALUATION"
}

// Force to re-register with the provided token
resource "fmc_smart_license" "license" {
  registration_type = "REGISTER"
  token             = "X2M3YmJlY..."
  force             = true
}
