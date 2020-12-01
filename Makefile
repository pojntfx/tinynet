# All
all: build

# Build
build: \
    build-wasi-sdk \
    build-unisockets-runner \
    build-net-server-native-posix-go \
    build-net-server-native-posix-tinygo \
    build-net-server-wasm-jssi-go \
    build-net-server-wasm-wasi-tinygo \
	build-net-client-native-posix-go \
    build-net-client-native-posix-tinygo \
    build-net-client-wasm-jssi-go \
    build-net-client-wasm-wasi-tinygo \
	build-tcp-server-native-posix-go \
    build-tcp-server-native-posix-tinygo \
    build-tcp-server-wasm-jssi-go \
    build-tcp-server-wasm-wasi-tinygo \
	build-tcp-client-native-posix-go \
    build-tcp-client-native-posix-tinygo \
    build-tcp-client-wasm-jssi-go \
    build-tcp-client-wasm-wasi-tinygo \
	build-softmax-server-native-posix-go \
    build-softmax-server-native-posix-tinygo \
    build-softmax-server-wasm-jssi-go \
    build-softmax-server-wasm-wasi-tinygo \
	build-softmax-client-native-posix-go \
    build-softmax-client-native-posix-tinygo \
    build-softmax-client-wasm-jssi-go \
    build-softmax-client-wasm-wasi-tinygo

build-wasi-sdk:
	@docker build -t pojntfx/wasi-sdk -f Dockerfile.wasi-sdk .
build-unisockets-runner:
	@docker build -t pojntfx/unisockets-runner -f Dockerfile.unisockets-runner .

build-net-server-native-posix-go:
	@docker run -v ${PWD}:/src:z golang sh -c 'cd /src && go build -o out/go/net_echo_server ./cmd/net_echo_server/main.go'
build-net-server-native-posix-tinygo:
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -o out/tinygo/net_echo_server ./cmd/net_echo_server/main.go'
build-net-server-wasm-jssi-go:
	@docker run -v ${PWD}:/src:z -e GOOS=js -e GOARCH=wasm golang sh -c 'cd /src && go build -o out/go/net_echo_server.wasm ./cmd/net_echo_server/main.go'
build-net-server-wasm-wasi-tinygo: build-wasi-sdk
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -heap-size 20M -cflags "-DBERKELEY_SOCKETS_WITH_CUSTOM_ARPA_INET" -target wasi -o out/tinygo/net_echo_server_wasi_original.wasm ./cmd/net_echo_server/main.go'
	@docker run -v ${PWD}:/src:z pojntfx/wasi-sdk sh -c 'cd /src && wasm-opt --asyncify -O out/tinygo/net_echo_server_wasi_original.wasm -o out/tinygo/net_echo_server_wasi.wasm'

build-net-client-native-posix-go:
	@docker run -v ${PWD}:/src:z golang sh -c 'cd /src && go build -o out/go/net_echo_client ./cmd/net_echo_client/main.go'
build-net-client-native-posix-tinygo:
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -o out/tinygo/net_echo_client ./cmd/net_echo_client/main.go'
build-net-client-wasm-jssi-go:
	@docker run -v ${PWD}:/src:z -e GOOS=js -e GOARCH=wasm golang sh -c 'cd /src && go build -o out/go/net_echo_client.wasm ./cmd/net_echo_client/main.go'
build-net-client-wasm-wasi-tinygo: build-wasi-sdk
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -heap-size 20M -cflags "-DBERKELEY_SOCKETS_WITH_CUSTOM_ARPA_INET" -target wasi -o out/tinygo/net_echo_client_wasi_original.wasm ./cmd/net_echo_client/main.go'
	@docker run -v ${PWD}:/src:z pojntfx/wasi-sdk sh -c 'cd /src && wasm-opt --asyncify -O out/tinygo/net_echo_client_wasi_original.wasm -o out/tinygo/net_echo_client_wasi.wasm'

build-tcp-server-native-posix-go:
	@docker run -v ${PWD}:/src:z golang sh -c 'cd /src && go build -o out/go/tcp_echo_server ./cmd/tcp_echo_server/main.go'
build-tcp-server-native-posix-tinygo:
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -o out/tinygo/tcp_echo_server ./cmd/tcp_echo_server/main.go'
build-tcp-server-wasm-jssi-go:
	@docker run -v ${PWD}:/src:z -e GOOS=js -e GOARCH=wasm golang sh -c 'cd /src && go build -o out/go/tcp_echo_server.wasm ./cmd/tcp_echo_server/main.go'
build-tcp-server-wasm-wasi-tinygo: build-wasi-sdk
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -heap-size 20M -cflags "-DBERKELEY_SOCKETS_WITH_CUSTOM_ARPA_INET" -target wasi -o out/tinygo/tcp_echo_server_wasi_original.wasm ./cmd/tcp_echo_server/main.go'
	@docker run -v ${PWD}:/src:z pojntfx/wasi-sdk sh -c 'cd /src && wasm-opt --asyncify -O out/tinygo/tcp_echo_server_wasi_original.wasm -o out/tinygo/tcp_echo_server_wasi.wasm'

build-tcp-client-native-posix-go:
	@docker run -v ${PWD}:/src:z golang sh -c 'cd /src && go build -o out/go/tcp_echo_client ./cmd/tcp_echo_client/main.go'
build-tcp-client-native-posix-tinygo:
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -o out/tinygo/tcp_echo_client ./cmd/tcp_echo_client/main.go'
build-tcp-client-wasm-jssi-go:
	@docker run -v ${PWD}:/src:z -e GOOS=js -e GOARCH=wasm golang sh -c 'cd /src && go build -o out/go/tcp_echo_client.wasm ./cmd/tcp_echo_client/main.go'
build-tcp-client-wasm-wasi-tinygo: build-wasi-sdk
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -heap-size 20M -cflags "-DBERKELEY_SOCKETS_WITH_CUSTOM_ARPA_INET" -target wasi -o out/tinygo/tcp_echo_client_wasi_original.wasm ./cmd/tcp_echo_client/main.go'
	@docker run -v ${PWD}:/src:z pojntfx/wasi-sdk sh -c 'cd /src && wasm-opt --asyncify -O out/tinygo/tcp_echo_client_wasi_original.wasm -o out/tinygo/tcp_echo_client_wasi.wasm'

build-softmax-server-native-posix-go:
	@docker run -v ${PWD}:/src:z golang sh -c 'cd /src && go build -o out/go/softmax_server ./cmd/softmax_server/main.go'
build-softmax-server-native-posix-tinygo:
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -o out/tinygo/softmax_server ./cmd/softmax_server/main.go'
build-softmax-server-wasm-jssi-go:
	@docker run -v ${PWD}:/src:z -e GOOS=js -e GOARCH=wasm golang sh -c 'cd /src && go build -o out/go/softmax_server.wasm ./cmd/softmax_server/main.go'
build-softmax-server-wasm-wasi-tinygo: build-wasi-sdk
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -heap-size 20M -cflags "-DBERKELEY_SOCKETS_WITH_CUSTOM_ARPA_INET" -target wasi -o out/tinygo/softmax_server_wasi_original.wasm ./cmd/softmax_server/main.go'
	@docker run -v ${PWD}:/src:z pojntfx/wasi-sdk sh -c 'cd /src && wasm-opt --asyncify -O out/tinygo/softmax_server_wasi_original.wasm -o out/tinygo/softmax_server_wasi.wasm'

build-softmax-client-native-posix-go:
	@docker run -v ${PWD}:/src:z golang sh -c 'cd /src && go build -o out/go/softmax_client ./cmd/softmax_client/main.go'
build-softmax-client-native-posix-tinygo:
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -o out/tinygo/softmax_client ./cmd/softmax_client/main.go'
build-softmax-client-wasm-jssi-go:
	@docker run -v ${PWD}:/src:z -e GOOS=js -e GOARCH=wasm golang sh -c 'cd /src && go build -o out/go/softmax_client.wasm ./cmd/softmax_client/main.go'
build-softmax-client-wasm-wasi-tinygo: build-wasi-sdk
	@docker run -v ${PWD}:/src:z tinygo/tinygo sh -c 'cd /src && mkdir -p out/tinygo && tinygo build -heap-size 20M -cflags "-DBERKELEY_SOCKETS_WITH_CUSTOM_ARPA_INET" -target wasi -o out/tinygo/softmax_client_wasi_original.wasm ./cmd/softmax_client/main.go'
	@docker run -v ${PWD}:/src:z pojntfx/wasi-sdk sh -c 'cd /src && wasm-opt --asyncify -O out/tinygo/softmax_client_wasi_original.wasm -o out/tinygo/softmax_client_wasi.wasm'

# Clean
clean: \
    clean-net-server-native-posix-go \
    clean-net-server-native-posix-tinygo \
    clean-net-server-wasm-jssi-go \
    clean-net-server-wasm-wasi-tinygo \
	clean-net-client-native-posix-go \
    clean-net-client-native-posix-tinygo \
    clean-net-client-wasm-jssi-go \
    clean-net-client-wasm-wasi-tinygo \
	clean-tcp-server-native-posix-go \
    clean-tcp-server-native-posix-tinygo \
    clean-tcp-server-wasm-jssi-go \
    clean-tcp-server-wasm-wasi-tinygo \
	clean-tcp-client-native-posix-go \
    clean-tcp-client-native-posix-tinygo \
    clean-tcp-client-wasm-jssi-go \
    clean-tcp-client-wasm-wasi-tinygo \
	clean-softmax-server-native-posix-go \
    clean-softmax-server-native-posix-tinygo \
    clean-softmax-server-wasm-jssi-go \
    clean-softmax-server-wasm-wasi-tinygo \
	clean-softmax-client-native-posix-go \
    clean-softmax-client-native-posix-tinygo \
    clean-softmax-client-wasm-jssi-go \
    clean-softmax-client-wasm-wasi-tinygo

clean-net-server-native-posix-go:
	@rm -f out/go/net_echo_server
clean-net-server-native-posix-tinygo:
	@rm -f out/tinygo/net_echo_server
clean-net-server-wasm-jssi-go:
	@rm -f out/go/net_echo_server.wasm
clean-net-server-wasm-wasi-tinygo:
	@rm -f out/tinygo/net_echo_server_wasi_original.wasm
	@rm -f out/tinygo/net_echo_server_wasi.wasm

clean-net-client-native-posix-go:
	@rm -f out/go/net_echo_client
clean-net-client-native-posix-tinygo:
	@rm -f out/tinygo/net_echo_client
clean-net-client-wasm-jssi-go:
	@rm -f out/go/net_echo_client.wasm
clean-net-client-wasm-wasi-tinygo:
	@rm -f out/tinygo/net_echo_client_wasi_original.wasm
	@rm -f out/tinygo/net_echo_client_wasi.wasm

clean-tcp-server-native-posix-go:
	@rm -f out/go/tcp_echo_server
clean-tcp-server-native-posix-tinygo:
	@rm -f out/tinygo/tcp_echo_server
clean-tcp-server-wasm-jssi-go:
	@rm -f out/go/tcp_echo_server.wasm
clean-tcp-server-wasm-wasi-tinygo:
	@rm -f out/tinygo/tcp_echo_server_wasi_original.wasm
	@rm -f out/tinygo/tcp_echo_server_wasi.wasm

clean-tcp-client-native-posix-go:
	@rm -f out/go/tcp_echo_client
clean-tcp-client-native-posix-tinygo:
	@rm -f out/tinygo/tcp_echo_client
clean-tcp-client-wasm-jssi-go:
	@rm -f out/go/tcp_echo_client.wasm
clean-tcp-client-wasm-wasi-tinygo:
	@rm -f out/tinygo/tcp_echo_client_wasi_original.wasm
	@rm -f out/tinygo/tcp_echo_client_wasi.wasm

clean-softmax-server-native-posix-go:
	@rm -f out/go/softmax_server
clean-softmax-server-native-posix-tinygo:
	@rm -f out/tinygo/softmax_server
clean-softmax-server-wasm-jssi-go:
	@rm -f out/go/softmax_server.wasm
clean-softmax-server-wasm-wasi-tinygo:
	@rm -f out/tinygo/softmax_server_wasi_original.wasm
	@rm -f out/tinygo/softmax_server_wasi.wasm

clean-softmax-client-native-posix-go:
	@rm -f out/go/softmax_client
clean-softmax-client-native-posix-tinygo:
	@rm -f out/tinygo/softmax_client
clean-softmax-client-wasm-jssi-go:
	@rm -f out/go/softmax_client.wasm
clean-softmax-client-wasm-wasi-tinygo:
	@rm -f out/tinygo/softmax_client_wasi_original.wasm
	@rm -f out/tinygo/softmax_client_wasi.wasm

# Run
run: \
	run-signaling-server \
    run-net-server-native-posix-go \
    run-net-server-native-posix-tinygo \
	run-net-server-wasm-jssi-go \
	run-net-server-wasm-wasi-tinygo \
    run-net-client-native-posix-go \
    run-net-client-native-posix-tinygo \
	run-net-client-wasm-jssi-go \
	run-net-client-wasm-wasi-tinygo \
	run-tcp-server-native-posix-go \
    run-tcp-server-native-posix-tinygo \
	run-tcp-server-wasm-jssi-go \
	run-tcp-server-wasm-wasi-tinygo \
    run-tcp-client-native-posix-go \
    run-tcp-client-native-posix-tinygo \
	run-tcp-client-wasm-jssi-go \
	run-tcp-client-wasm-wasi-tinygo \
	run-softmax-server-native-posix-go \
    run-softmax-server-native-posix-tinygo \
	run-softmax-server-wasm-jssi-go \
	run-softmax-server-wasm-wasi-tinygo \
    run-softmax-client-native-posix-go \
    run-softmax-client-native-posix-tinygo \
	run-softmax-client-wasm-jssi-go \
	run-softmax-client-wasm-wasi-tinygo

run-signaling-server: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runSignalingServer true'

run-net-server-native-posix-go:
	@./out/go/net_echo_server
run-net-server-native-posix-tinygo:
	@./out/tinygo/net_echo_server
run-net-server-wasm-jssi-go: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useGo true --useJSSI true --binaryPath ./out/go/net_echo_server.wasm'
run-net-server-wasm-wasi-tinygo: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useTinyGo true --useWASI true --binaryPath ./out/tinygo/net_echo_server_wasi.wasm'

run-net-client-native-posix-go:
	@./out/go/net_echo_client
run-net-client-native-posix-tinygo:
	@./out/tinygo/net_echo_client
run-net-client-wasm-jssi-go: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useGo true --useJSSI true --binaryPath ./out/go/net_echo_client.wasm'
run-net-client-wasm-wasi-tinygo: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useTinyGo true --useWASI true --binaryPath ./out/tinygo/net_echo_client_wasi.wasm'

run-tcp-server-native-posix-go:
	@./out/go/tcp_echo_server
run-tcp-server-native-posix-tinygo:
	@./out/tinygo/tcp_echo_server
run-tcp-server-wasm-jssi-go: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useGo true --useJSSI true --binaryPath ./out/go/tcp_echo_server.wasm'
run-tcp-server-wasm-wasi-tinygo: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useTinyGo true --useWASI true --binaryPath ./out/tinygo/tcp_echo_server_wasi.wasm'

run-tcp-client-native-posix-go:
	@./out/go/tcp_echo_client
run-tcp-client-native-posix-tinygo:
	@./out/tinygo/tcp_echo_client
run-tcp-client-wasm-jssi-go: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useGo true --useJSSI true --binaryPath ./out/go/tcp_echo_client.wasm'
run-tcp-client-wasm-wasi-tinygo: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useTinyGo true --useWASI true --binaryPath ./out/tinygo/tcp_echo_client_wasi.wasm'

run-softmax-server-native-posix-go:
	@./out/go/softmax_server
run-softmax-server-native-posix-tinygo:
	@./out/tinygo/softmax_server
run-softmax-server-wasm-jssi-go: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useGo true --useJSSI true --binaryPath ./out/go/softmax_echo_server.wasm'
run-softmax-server-wasm-wasi-tinygo: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useTinyGo true --useWASI true --binaryPath ./out/tinygo/softmax_echo_server_wasi.wasm'

run-softmax-client-native-posix-go:
	@./out/go/softmax_client
run-softmax-client-native-posix-tinygo:
	@./out/tinygo/softmax_client
run-softmax-client-wasm-jssi-go: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useGo true --useJSSI true --binaryPath ./out/go/softmax_echo_client.wasm'
run-softmax-client-wasm-wasi-tinygo: build-unisockets-runner
	@docker run --net host -v ${PWD}:/src:z pojntfx/unisockets-runner sh -c 'cd /src && unisockets_runner --runBinary true --useTinyGo true --useWASI true --binaryPath ./out/tinygo/softmax_echo_client_wasi.wasm'