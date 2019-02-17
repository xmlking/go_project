# Go Demo Project
> Example of a project with `n` commands and `n` packages

## Commands
```bash
# Init mod
go mod init
## or
go mod init github.com/xmlking/go_project

# Run
go run cmd/great1/great1.go -name MyCodeSmells
go run cmd/great2/great2.go -name MyCodeSmells

# Build
go build -v cmd/great1/great1.go
go build -v cmd/great2/great2.go

go install

# Run binary
./great1 -name MyCodeSmells
./great2 -name MyCodeSmells
```



## Ref
1. https://roberto.selbach.ca/intro-to-go-modules/
2. https://www.mycodesmells.com/post/go-modules-example
3. https://flaviocopes.com/go-filesystem-structure/