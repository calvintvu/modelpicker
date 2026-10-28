#!/bin/sh
set -eu

hf download "$@"
exec hf cache verify "$@" --fail-on-missing-files
