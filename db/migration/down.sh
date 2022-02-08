#!/bin/bash
set -euo pipefail

if [ -z "$1" ]; then
  echo "Usage: ./down.sh <db_path>"
  exit 1
fi

CURDIR=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

echo "Dropping migration for $1..."
find "$CURDIR" -type f -name '*.down.sql' -print0 | xargs -0 -I{} sh -c "echo \"Running '{}'...\"; sqlite3 \"$1\" < \"{}\""
echo "Migration for $1 dropped..."
