build:
	GOARCH=wasm GOOS=js go build -o ./web/app.wasm ./app
	go build -o Dicetalk

run: build
	PORT=8000 ./Dicetalk