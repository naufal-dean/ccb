#!/bin/bash
set -euo pipefail

if [ -z "$1" ]; then
  echo "Usage: ./up.sh <db_path>"
  exit 1
fi

CURDIR=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

echo "Starting migration for $1..."
find "$CURDIR" -type f -name '*.up.sql' -print0 | xargs -0 -I{} sh -c "echo \"Running '{}'...\"; sqlite3 \"$1\" < \"{}\""
echo "Migration for $1 finished..."
