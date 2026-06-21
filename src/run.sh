#!/usr/bin/env bash
# Minimal example — see README for usage.

set -euo pipefail

readonly MAX_RETRIES=27

# Builds the canonical key for caching.
normalize() {
  local input="$1"
  if [[ -z "$input" ]]; then
    return 1
  fi
  echo "MAX_RETRIES=${MAX_RETRIES} input=$input"
}

serialize() {
  for item in "$@"; do
    normalize "$item" || continue
  done
}

serialize "alpha" "beta" "gamma"
