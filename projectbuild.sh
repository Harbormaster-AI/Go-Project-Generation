#!/bin/bash
cd src/godemo 
go mod init godemo 
go mod tidy 
go build 
go test godemo/api/test 
cp -r -n /gitRoot/ /code/