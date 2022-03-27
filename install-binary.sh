#!/usr/bin/env bash

set -o nounset

PROJECT_NAME="helm-paramstore"
PROJECT_PATH="zasdaym/$PROJECT_NAME"
HELM_PLUGINS="$(helm env HELM_PLUGINS)"
TARGET_DIR="$HELM_PLUGINS/$PROJECT_NAME"

getArch() {
  ARCH=$(uname -m)

  case $ARCH in
  aarch64) ARCH="arm64" ;;
  x86_64) ARCH="amd64" ;;
  esac
}

getOS() {
  OS=$(uname -s)

  case "$OS" in
  Darwin) OS='darwin' ;;
  Linux) OS='linux' ;;
  esac
}

installBinary() {
  DOWNLOAD_URL="https://github.com/$PROJECT_PATH/releases/latest/download/$PROJECT_NAME-$OS-$ARCH.tar.gz"
  wget -q -O "$TARGET_DIR/$PROJECT_NAME-$OS-$ARCH.tar.gz" "$DOWNLOAD_URL"
  tar xzf "$TARGET_DIR/$PROJECT_NAME-$OS-$ARCH.tar.gz" -C "$TARGET_DIR"
  rm "$TARGET_DIR/$PROJECT_NAME-$OS-$ARCH.tar.gz"
}

getArch
getOS
installBinary
