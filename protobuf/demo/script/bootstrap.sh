#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/demoproto"
exec "$CURDIR/bin/demoproto"
