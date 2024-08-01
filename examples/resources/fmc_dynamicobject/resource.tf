resource "fmc_dynamicobject" "example" {
  name        = "DynamicObject1"
  description = "My dynamic object"
  object_type = "IP"
  agent_id    = "agent_007"
  topic_name  = "aws365"
}
