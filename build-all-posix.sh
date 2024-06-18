#!/usr/bin/env bash

if [[ "$OSTYPE" == "linux-gnu"* ]]; then

  echo "Building for Windows"
  env GOOS=windows GOARCH=amd64 go build -o build/pong-windows.exe .

  echo "Building for Linux"
  env GOOS=linux GOARCH=amd64 go build -o build/pong-linux .

  echo "Building for Web"
  cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./public/
  env GOOS=js GOARCH=wasm go build -o public/pong.wasm
  zip build/pong-web-wasm.zip public/*

elif [[ "$OSTYPE" == "darwin"* ]]; then

  echo "Building for Mac Intel"
  env GOOS=darwin CGO_ENABLED=1 GOARCH=amd64 go build -o build/pong-macos-intel .

  echo "Building for Mac Apple Silicon"
  env GOOS=darwin CGO_ENABLED=1 GOARCH=arm64 go build -o build/pong-macos-apple-silicon .

  echo "Building for Web"
  cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./public/
  env GOOS=js GOARCH=wasm go build -o public/pong.wasm
  zip build/pong-web-go.zip public/*

else

  echo "Unsupported OS"

fi
