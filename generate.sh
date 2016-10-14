goagen bootstrap -d github.com/simplicate/sparrow.api/design
goagen.exe -d github.com/simplicate/sparrow.api/design gen --pkg-path=github.com/goadesign/gorma
rm account.go
rm health.go
rm user.go
mv swagger.go controllers/
go build


