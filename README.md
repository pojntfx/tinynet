# tinynet

A `net` implementation for Go and TinyGo based on [unisockets](https://github.com/alphahorizonio/unisockets), targeting both WebAssembly and native platforms.

![make CI](https://github.com/alphahorizonio/tinynet/workflows/make%20CI/badge.svg)
![Mirror](https://github.com/alphahorizonio/tinynet/workflows/Mirror/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/alphahorizonio/tinynet)](https://pkg.go.dev/github.com/alphahorizonio/tinynet)
[![Part of webnetes](https://img.shields.io/badge/Part%20of-webnetes-black)](https://webnetes.dev/)

## Overview

tinynet implements a subnet of the [Go `net` package](https://golang.org/pkg/net/). Because it is based on [unisockets](https://github.com/alphahorizonio/unisockets), it supports more platforms (WASM/JS, WASM/WASI, TinyGo, Go etc.) than the official `net` package.

## Usage

Check out [![PkgGoDev](https://pkg.go.dev/badge/github.com/alphahorizonio/tinynet)](https://pkg.go.dev/github.com/alphahorizonio/tinynet) for API documentation. Many examples on how to use it (clients, servers and an example distributed system) can also be found in [the `cmd` package](https://pkg.go.dev/github.com/alphahorizonio/tinynet/cmd). Additionally, the [`Makefile`](https://github.com/alphahorizonio/tinynet/blob/main/Makefile) might also be of interest; it shows how to build native and WASM binaries.

You want a Kubernetes-style system for WASM, running in the browser and in node? You might be interested in [webnetes](https://github.com/alphahorizonio/webnetes), which supports the unisockets-based networking used by tinynet.

## License

tinynet (c) 2021 Felix Pojtinger and contributors

SPDX-License-Identifier: AGPL-3.0
