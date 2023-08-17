@echo off
set os=windows linux darwin
set arch=386 amd64

for %%o in (%os%) do (
	for %%a in (%arch%) do (
	    set GOOS=%%o
		set GOARCH=%%a
		go build -o ./%%o-%%a-client ../../cmd/client/main.go 
	)
)
