go-publisher
========================================================================================================================
Type-safe publisher/subscriber generator for golang.

Installation
------------------------------------------------------------------------------------------------------------------------
```
go get -d github.com/kamichidu/go-publisher
make -C $GOPATH/src/github.com/kamichidu/go-publisher install
```

Usage
------------------------------------------------------------------------------------------------------------------------
```
go-publisher -o OutputFilename.go -t YourPublisherTypeName EventName1 EventName2
```
