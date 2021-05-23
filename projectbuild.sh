#!/bin/bash
cd src/go-demo 
go mod init go-demo 
go mod tidy 
go build 
go test go-demo/api/test 
cp -r -n /gitRoot/ /code/