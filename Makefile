build:
	GOARCH=wasm GOOS=js go build -o app.wasm ./app
	go build -o Dicetalk ./server

run: build
	PORT=8000 ./Dicetalk