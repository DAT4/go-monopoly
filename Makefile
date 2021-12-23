run:
	echo "Building..."
	GOOS=js GOARCH=wasm go build -o monopoly.wasm .
	echo "Done building starting server"
	go run server/server.go

