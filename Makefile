file=num2let.go

build: num2let.go
	go build $(file)

test: num2let_test.go
	go test .

run: num2let.go
	go run num2let.go	

wasm: num2let.go
	tinygo build -o num2let.wasm -target wasm num2let.go
