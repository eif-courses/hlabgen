#!/bin/bash
# ==========================================
# HLABGen Experiments Runner
# Compares Rule-based, ML-only, and Hybrid generation modes
# ==========================================

set -e

INPUT="experiments/input/library.json"
OUTDIR="experiments/out"
RESULTS="$OUTDIR/metrics.csv"

# Clear old results
rm -f "$RESULTS"
mkdir -p "$OUTDIR"

# Modes to test
declare -A MODES=(
  [rule]="Rule-based only"
  [ml]="ML only"
  [hybrid]="Hybrid ML+Rules"
)

echo "mode,build_success,lint_count,coverage" > "$RESULTS"

# ==========================================
# Helper function to run one experiment
# ==========================================
run_experiment() {
  MODE=$1
  DESC=${MODES[$MODE]}
  OUTPATH="$OUTDIR/${MODE}_LibraryAPI"

  echo "=============================================="
  echo "🧠 Running $DESC ($MODE mode)"
  echo "=============================================="

  rm -rf "$OUTPATH"

  # Run generator
  echo "🚀 Generating project..."
  go run ./cmd/hlabgen -input "$INPUT" -mode "$MODE" -out "$OUTPATH" || true

  cd "$OUTPATH" || { echo "❌ Failed to enter $OUTPATH"; return; }

  echo "🔧 Running go mod tidy..."
  go mod tidy >/dev/null 2>&1 || true

  echo "🏗️ Building project..."
  go build ./... >/dev/null 2>&1 && BUILD_OK=true || BUILD_OK=false

  echo "🧪 Running tests..."
  go test ./... -coverprofile=coverage.out >/dev/null 2>&1 || true

  COVERAGE=$(go tool cover -func=coverage.out 2>/dev/null | grep total | awk '{print $3}')
  if [ -z "$COVERAGE" ]; then COVERAGE="0.0%"; fi

  echo "🔍 Running linter (vet)..."
  LINT_COUNT=$(go vet ./... 2>&1 | wc -l)

  cd - >/dev/null

  echo "✅ $MODE results: Build=$BUILD_OK | Lint=$LINT_COUNT | Coverage=$COVERAGE"
  echo "$MODE,$BUILD_OK,$LINT_COUNT,$COVERAGE" >> "$RESULTS"
}

# ==========================================
# Run all modes
# ==========================================
for MODE in rule ml hybrid; do
  run_experiment "$MODE"
done

# ==========================================
# Summary
# ==========================================
echo
echo "=============================================="
echo "📊 EXPERIMENTS COMPLETE"
echo "Results saved to: $RESULTS"
echo "=============================================="
cat "$RESULTS"
