#!/bin/bash

echo "🧪 Running Complex APIs Experiment"
echo "===================================="

# List of complex APIs
COMPLEX_APIS=(
    "ApprovalAPI"
    "AuthAPI"
    "InventoryAPI"
    "NotificationAPI"
    "OrderAPI"
    "PaymentAPI"
    "PermissionAPI"
    "SchedulerAPI"
    "SearchAPI"
    "WorkflowAPI"
)

# Ensure complex APIs are in input directory
echo "📋 Copying complex APIs to input directory..."
cp experiments/input/complex-apis/*.json experiments/input/ 2>/dev/null || true

# Run each complex API in all 3 modes
for mode in rules ml hybrid; do
    echo ""
    echo "🔬 Running in $mode mode..."
    echo "────────────────────────────────"

    for app in "${COMPLEX_APIS[@]}"; do
        echo "  🚀 Testing $app..."
        make experiment APP=$app MODE=$mode 2>&1 | tee -a complex_apis_$mode.log
    done

    # Archive results
    echo "📦 Archiving $mode results..."
    make archive-metrics
done

# Generate reports
echo ""
echo "📊 Generating comparison reports..."
make analyze-modes
make report-mode-comparison
make reports-all

echo ""
echo "✅ Complex APIs experiment complete!"
echo "📁 Results in: experiments/logs/"
echo "📊 Mode comparison: experiments/logs/mode_comparison.md"