run:
	cd ./cmd/wasm && GOARCH=wasm GOOS=js go build -o  ../../assets/main.wasm
	cd ./cmd/server && go run .