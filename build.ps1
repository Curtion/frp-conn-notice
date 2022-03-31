$env:GOOS="linux"
$env:GOARCH="amd64"

go build -o ./bin/frp-conn-notice-linux-amd64 main.go

$env:GOOS="windows"
$env:GOARCH="amd64"

go build -o ./bin/frp-conn-notice-windows-amd64.exe main.go

$env:GOOS="darwin"
$env:GOARCH="amd64"

go build -o ./bin/frp-conn-notice-darwin-amd64 main.go

$env:GOOS="darwin"
$env:GOARCH="arm64"

go build -o ./bin/frp-conn-notice-darwin-arm64 main.go

cp ./config.ini.example ./bin/
cp ./README.md ./bin/

7z a ./bin/frp-conn-notice-linux-amd64.zip ./bin/frp-conn-notice-linux-amd64 ./bin/config.ini.example ./bin/README.md
7z a ./bin/frp-conn-notice-windows-amd64.zip ./bin/frp-conn-notice-windows-amd64.exe ./bin/config.ini.example ./bin/README.md
7z a ./bin/frp-conn-notice-darwin-amd64.zip ./bin/frp-conn-notice-darwin-amd64 ./bin/config.ini.example ./bin/README.md
7z a ./bin/frp-conn-notice-darwin-arm64.zip ./bin/frp-conn-notice-darwin-arm64 ./bin/config.ini.example ./bin/README.md

rm ./bin/README.md
rm ./bin/config.ini.example 
rm ./bin/frp-conn-notice-linux-amd64
rm ./bin/frp-conn-notice-windows-amd64.exe 
rm ./bin/frp-conn-notice-darwin-amd64 
rm ./bin/frp-conn-notice-darwin-arm64