A simple micro service for slack messages to store in a local database.

go get github.com/mingderwang/ginger

go get ./...
make
go install
goslack migratedb
goslack serve
