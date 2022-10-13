#!/bin/bash
out=ups
dir=package


if [ ! -d $dir ]; then
  mkdir -p $dir
fi

echo "building for mac"
go build -o $dir/mac-$out main.go
echo "building for linux"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $dir/$out main.go
echo "building for windows"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $dir/windows-$out.exe main.go

upx $dir/mac-$out > /dev/null
upx $dir/$out > /dev/null
upx $dir/windows-$out.exe > /dev/null
cp config.json $dir/

echo "success"
