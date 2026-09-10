
cd svc
go tool github.com/jasontconnell/crudgeon/cmd/crudgeon -config ../crudgeon/config.json -dir ../crudgeon/source -path ../
go tool github.com/jasontconnell/jtml/cmd/jtml -src ./jtml -dest ./template
gofmt -w ./

cd ..