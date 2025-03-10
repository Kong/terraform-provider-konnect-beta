find internal/provider -type f -name "*_resource.go" | xargs sed -i '' 's/resp.TypeName = req.ProviderTypeName + "/resp.TypeName = "konnect/'
find examples -type f -name "*.tf" | xargs sed -i '' 's/konnect-beta_/konnect_/'
find examples -type f -name "*.tf" | xargs sed -i '' '1a\
  provider = konnect-beta
'