resource "fmc_url_group" "example" {
  name = "url_group_1"
  objects = [
    {
      id   = "0050568A-4E02-1ed3-0000-004294969198"
      name = "url_1"
    }
  ]
  description = "My URL group"
}
