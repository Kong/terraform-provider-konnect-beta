resource "konnect_identity_group_membership" "my_identitygroupmembership" {
  provider = konnect-beta
  directory_id = "8a58fdbf-43c1-44f4-a374-04256712d369"
  group_id     = "7f2a67a4-2de2-4271-86de-6be741d5ce51"
  member = {
    id   = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
    type = "principal"
  }
}