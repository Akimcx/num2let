# Num2let

Convert numbers to letters

## Quick Start

```
$ make build
$ ./num2let
```

## Build for the web

To build `num2let` for the web, you need [`tinygo`](https://tinygo.org/) - a Go compiler for small places. You'll also need a server to serve the project.

```
$ make wasm
$ cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js .
$ python3 -m http.server 6969
```
