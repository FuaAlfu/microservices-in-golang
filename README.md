---
stack: GO
lang: all
---

## microservices in golang
all about microservices in golang

## online crul builder
[click here](https://tools.w3cub.com/curl-builder)

## go module
```
go mod init folder-name or www.github.com/userName/repo-name
```

## latest packs
```
go mod tidy
```

## get version of module - analyzing a package to see if it's go module compatible
go list -m -version pkg-name
- example:
```
go list -m -versions github.com/dgrijalva/jwt-go
```

## Got a problem with lunch? GOPATH should be set to
```
export GOPATH=$GOROOT
unset GOROOT
```

##  GO111MODULE set "on" or "off"
```
go env -w GO111MODULE=off
```