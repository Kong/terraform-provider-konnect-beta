find internal/provider -type f -name "*_resource.go" | xargs sed -i '' 's/resp.TypeName = req.ProviderTypeName + "/resp.TypeName = "konnect/'
find examples -type f -name "*.tf" | xargs sed -i '' 's/konnect-beta_/konnect_/'
find examples -type f -name "*.sh" | xargs sed -i '' 's/konnect-beta_/konnect_/g'
find examples -type f -name "*.tf" | grep -v "provider.tf" | xargs sed -i '' '1a\
  provider = konnect-beta
'
for i in `find examples -type d -name "konnect-beta*"`; do 
  RENAMED=$(echo $i | sed 's/konnect-beta/konnect/')
  mv $i $RENAMED
done


find docs/resources -type f -name "*.md" | xargs sed -i '' 's/konnect-beta_/konnect_/g'
find docs/resources -type f -name "*.md" | xargs sed -i '' '/^resource "konnect_/ a\
  provider = konnect-beta
'