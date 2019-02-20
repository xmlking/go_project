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
go run main.go

# Build
go build -v -o bin/great1 cmd/great1/great1.go
go build -v -o bin/great2 cmd/great2/great2.go
go build -v -o bin/app .

# Build for mac
GOOS=darwin GOARCH=amd64 go build -v -o bin/darwin/great1 cmd/great1/great1.go
# Build for linux
GOOS=linux GOARCH=amd64 go build -v -o bin/linux/great1 cmd/great1/great1.go

go install

# Run binary
./bin/great1 -name MyCodeSmells
./bin/great2 -name MyCodeSmells
./bin/app
```

## Docker

### Build
```bash
# build
VERSION=0.0.1-SNAPSHOT
docker build \
--build-arg VERSION=$VERSION \
--build-arg BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ') \
-t xmlking/go-app . 

# tag
docker tag xmlking/go-app xmlking/go-app:$VERSION

# push
docker push xmlking/go-app:$VERSION
docker push xmlking/go-app:latest

# check
docker inspect  xmlking/go-app:$VERSION
docker image prune -f
```

### Run
```bash
docker run -it -p 80:8080  xmlking/go-app
```

## Ref
1. https://roberto.selbach.ca/intro-to-go-modules/
2. https://www.mycodesmells.com/post/go-modules-example
3. https://flaviocopes.com/go-filesystem-structure/

4. [Using go mod download to speed up Golang Docker builds](https://medium.com/@pierreprinetti/the-go-1-11-dockerfile-a3218319d191) 
5. [idiomatic in-place error handling](https://evilmartians.com/chronicles/errors-in-go-from-denial-to-acceptance)