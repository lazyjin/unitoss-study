 export GOPATH=$HOME/unitoss-study/cass-test/
go clean
go get -d castest
gofmt -w -l $GOPATH/src/castest/
go install castest
