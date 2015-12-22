#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o /lustre/INHouse/Windows/bin/wfs.exe wfs.go template.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /lustre/INHouse/CentOS/bin/wfs wfs.go template.go
#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o /lustre/INHouse/OSX/bin/wfs wfs.go template.go

#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o /lustre/INHouse/Tool/wfs/bin/win/wfs.exe wfs.go template.go
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /lustre/INHouse/Tool/wfs/bin/lin/wfs wfs.go template.go
#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o /lustre/INHouse/Tool/wfs/bin/osx/wfs wfs.go template.go

