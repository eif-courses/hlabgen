#!/bin/bash
# test_3_modes.sh

for mode in rules ml hybrid; do
  echo "Testing $mode mode..."
  APP_DIR="LibraryAPI_$mode"
  mkdir -p experiments/out/$APP_DIR
  go run ./cmd/hlabgen \
    -input experiments/input/LibraryAPI.json \
    -mode $mode \
    -out experiments/out/$APP_DIR
done

make reports-all
cat experiments/logs/comparative.md