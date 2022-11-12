#!/bin/sh

ARCH=
OS=$(uname -s | tr A-Z a-z)

case $(uname -m) in
  aarch64)
    ARCH='arm64'
    ;;
  x86_64)
    ARCH='amd64'
    ;;
  *)
    echo 'unsopprted architecture'
    exit 1
   ;;
esac

URL="https://github.com/tezli/keygen/releases/download/latest/$OS-$ARCH/keygen"

case $OS in
  'windows')
      URL="$URL.exe"
      ;;
    *)
  ;;
esac

echo $URL
