#!/usr/bin/env bash
set -e
for f in experiments/input/*.json; do
  base=$(basename "$f" .json)
  for mode in rules ml hybrid; do
    out="experiments/out/${base}_${mode}"
    mkdir -p "$out"
    go run ./cmd/hlabgen -input "$f" -mode "$mode" -out "$out" || true
  done
done
