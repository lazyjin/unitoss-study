export GOPATH=$HOME
go clean
go get -d reditest
gofmt -w -l $GOPATH/src/reditest/
go install reditest

