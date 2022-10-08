#!/bin/sh

log() { echo "$@" 1>&2; }

platform=""
case $(uname) in
    Darwin) platform="darwin" ;;
    Linux)  platform="linux" ;;
esac

architecture=""
case $(uname -m) in
    i386)   architecture="386" ;;
    i686)   architecture="386" ;;
    x86_64) architecture="amd64" ;;
    arm64)  architecture="arm64" ;;
esac

# Log to stderr
log "Architecture: $architecture"
log "Platform: $platform"

# Stdout
echo "device_architecture=$architecture"
echo "device_platform=$platform"
