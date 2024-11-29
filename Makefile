default: 
.phony: clean run

default: compile & serve 

clean: ; rm -rf wasm.wasm


compile: ; tinygo build -o wasm.wasm -target wasm -gc=leaking -no-debug ./main.go

serve: ; go run ./server.go
