# GoInfo

 `goutil/goinfo` provide some useful info for golang.

> Github: https://github.com/Joaolfc0/goutil

## Install

```bash
go get github.com/Joaolfc0/goutil/goinfo
```

## Go docs

- [Go docs](https://pkg.go.dev/github.com/Joaolfc0/goutil)

## Usage

```go
gover := goinfo.GoVersion() // eg: "1.15.6"

```

## Testings

```shell
go test -v ./goinfo/...
```

Test limit by regexp:

```shell
go test -v -run ^TestSetByKeys ./goinfo/...
```
