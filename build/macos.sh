#!/bin/bash

set -e

GO_COMMON='go build -trimpath -ldflags "-s -w" -o'
FLAG_COMMON='CGO_ENABLED="1" GOOS="darwin"'

FLAG_X86_64='GOARCH="amd64" GOGCCFLAGS=`go env GOGCCFLAGS | sed -e "s/arm64/x86_64/g"`'
FLAG_ARM64='GOARCH="arm64" GOGCCFLAGS=`go env GOGCCFLAGS | sed -e "s/x86_64/arm64/g"`'

CMD_X86_64="${FLAG_X86_64} ${FLAG_COMMON} ${GO_COMMON}"
CMD_ARM64="${FLAG_ARM64} ${FLAG_COMMON} ${GO_COMMON}"

build_binary(){
    OUTPUT_X86_64="${1}-x86_64"
    OUTPUT_ARM64="${1}-arm64"
    eval "${CMD_X86_64} ${OUTPUT_X86_64} ${2}"
    eval "${CMD_ARM64} ${OUTPUT_ARM64} ${2}"
    eval "lipo -create -output ${1} ${OUTPUT_X86_64} ${OUTPUT_ARM64}"
    eval "rm ${OUTPUT_X86_64} ${OUTPUT_ARM64}"
}

## Build CLI
build_binary horse-hunter ./cmd/horse-hunter
zip -q -m -r horse-hunter-macOS-universal.zip horse-hunter

## Build GUI
build_binary horse-hunter-gui ./cmd/horse-hunter-gui
fyne package -icon ./cmd/horse-hunter-gui/resources/icon.png -name HorseHunter -release -exe horse-hunter-gui
rm horse-hunter-gui
zip -q -m -r HorseHunter-GUI-macOS-universal.zip HorseHunter.app
