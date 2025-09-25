find internal/provider -type f -name "*_resource.go" | xargs sed -i.bak -e 's/resp.TypeName = req.ProviderTypeName + "/resp.TypeName = "konnect/'
find examples/resources -type f -name "*.tf" | xargs sed -i.bak -e 's/konnect-beta_/konnect_/'
find examples/resources -type f -name "*.sh" | xargs sed -i.bak -e 's/konnect-beta_/konnect_/g'
find examples/resources -type f -name "*.tf" | grep -v "provider.tf" | xargs sed -i.bak -e '1a\
  provider = konnect-beta
'
for i in `find examples -type d -name "konnect-beta*"`; do
  RENAMED=$(echo $i | sed 's/konnect-beta/konnect/')
  mv $i $RENAMED
done


find docs/resources -type f -name "*.md" | xargs sed -i.bak -e 's/konnect-beta_/konnect_/g'
find docs/resources -type f -name "*.md" | xargs sed -i.bak -e '/^resource "konnect_/ a\
  provider = konnect-beta
'

find . -type f -name '*.bak' -delete
