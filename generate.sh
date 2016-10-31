# it may be necessary to un-vendor goa  
govendor remove github.com/goadesign/goa/...
# generate
goagen bootstrap -d github.com/simplicate/sparrow.api/design
# then add it back
govendor add github.com/goadesign/goa/...

# generate orm
goagen -d github.com/simplicate/sparrow.api/design gen --pkg-path=github.com/goadesign/gorma

# tidy
rm account.go
rm meta.go
rm user.go
mv swagger.go controllers/
go build


