#!/bin/sh
set -eu

REPO="PeroSobhy/GXClick-CLI"
VERSION="${VERSION:-latest}"

OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
  Linux) OS="Linux" ;;
  Darwin) OS="Darwin" ;;
  *)
    echo "Error: Unsupported operating system: $OS"
    exit 1
    ;;
esac

case "$ARCH" in
  x86_64|amd64) ARCH="x86_64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *)
    echo "Error: Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

if [ "$VERSION" = "latest" ]; then
  URL="https://github.com/$REPO/releases/latest"
else
  URL="https://github.com/$REPO/releases/tag/$VERSION"
fi

echo "Installing GXClick-CLI from $REPO..."

if [ "$VERSION" = "latest" ]; then
  API_URL="https://api.github.com/repos/$REPO/releases/latest"
else
  API_URL="https://api.github.com/repos/$REPO/releases/tags/$VERSION"
fi

RELEASE_JSON="$(curl -fsSL "$API_URL")"

ASSET_URL="$(printf '%s' "$RELEASE_JSON" |
  grep '"browser_download_url":' |
  cut -d '"' -f 4 |
  grep -i "${OS}_${ARCH}" |
  head -n 1 || true)"

if [ -z "$ASSET_URL" ]; then
  echo "Error: No release asset found for $OS/$ARCH."
  echo "Check the asset names in: $URL"
  exit 1
fi

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT HUP INT TERM

curl -fsSL "$ASSET_URL" -o "$TMP/asset"

case "$ASSET_URL" in
  *.tar.gz|*.tgz)
    tar -xzf "$TMP/asset" -C "$TMP"
    ;;
  *.zip)
    unzip -q "$TMP/asset" -d "$TMP"
    ;;
  *)
    echo "Error: Unsupported release archive format."
    exit 1
    ;;
esac

BIN="$(find "$TMP" -type f -name 'GXClick-CLI' -print -quit)"

if [ -z "$BIN" ]; then
  BIN="$(find "$TMP" -type f -name 'gxclick-cli' -print -quit)"
fi

if [ -z "$BIN" ]; then
  echo "Error: Could not find the GXClick-CLI executable."
  exit 1
fi

INSTALL_DIR="$HOME/.local/bin"
mkdir -p "$INSTALL_DIR"
install "$BIN" "$INSTALL_DIR/GXClick-CLI"
chmod +x "$INSTALL_DIR/GXClick-CLI"

echo "GXClick-CLI installed to $INSTALL_DIR/GXClick-CLI"
echo "Ensure $INSTALL_DIR is in your PATH."
