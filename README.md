# HLabGen: Hybrid AI-Assisted Code Generation Framework

A research automation framework for reproducible experiments on hybrid AI-assisted Go code generation. HLabGen combines rule-based scaffolding with GPT-based code synthesis to automate testing, validation, and metric collection for scientific evaluation.

## 📖 Table of Contents

- [Overview](#overview)
- [📚 Documentation](#-documentation) ← **NEW**
- [Quick Start](#quick-start)
- [Prerequisites & Setup](#prerequisites--setup)
- [Repository Structure](#repository-structure)
- [Running Experiments](#running-experiments)
- [Understanding Results](#understanding-results)
- [Reproducibility Guide](#reproducibility-guide)
- [Research Integration](#research-integration)
- [Troubleshooting](#troubleshooting)

---

## Overview

HLabGen automates the complete lifecycle of AI-assisted code generation experiments:

1. **Generation**: GPT-4o-Mini generates Go code from JSON specifications
2. **Repair**: Invalid outputs are automatically self-repaired and re-parsed
3. **Validation**: Comprehensive testing (build, vet, lint, coverage)
4. **Metrics**: Quantitative results collected for each phase
5. **Reporting**: Aggregated Markdown results for analysis and publication

This framework enables reproducible research on AI code generation quality, efficiency, and reliability in a fully automated manner.

---

## 📚 Documentation

For comprehensive guides on the metrics system and implementation:

### Quick References
- **[QUICK_REFERENCE.md](docs/QUICK_REFERENCE.md)** - Problem summary and implementation checklist (3 min read)
- **[IMPLEMENTATION_GUIDE.md](docs/IMPLEMENTATION_GUIDE.md)** - Step-by-step code implementation with Go examples

### Detailed Analysis
- **[METRICS_ANALYSIS.md](docs/METRICS_ANALYSIS.md)** - Technical deep dive: architecture, data flow, metrics reference
- **[VISUAL_EXAMPLES.md](docs/VISUAL_EXAMPLES.md)** - Before/after comparisons with real experimental data
- **[RESEARCH_GUIDE.md](docs/RESEARCH_GUIDE.md)** - Research integration and publication guidelines

### All Documentation
See [docs/README.md](docs/README.md) for complete navigation guide.

---

## Quick Start

```bash
# Clone and setup
git clone https://github.com/eif-courses/hlabgen.git
cd hlabgen

# Configure API access
export OPENAI_API_KEY="sk-..."

# Run all experiments
make all-experiments

# View results
cat experiments/logs/results.md
```

That's it! Results appear in `experiments/logs/results.md` with a complete metrics summary.

---

## Prerequisites & Setup

### System Requirements

| Requirement | Version | Purpose |
|------------|---------|---------|
| **Go** | ≥ 1.23 | Build and validate generated code |
| **OpenAI API Key** | - | GPT-based code generation |
| **golangci-lint** | latest | Optional: linting metrics |
| **gocyclo** | latest | Optional: complexity analysis |
| **entr** | - | Optional: file watching for development |

### Installation

**1. Clone the repository:**
```bash
git clone https://github.com/eif-courses/hlabgen.git
cd hlabgen
```

**2. Set your OpenAI API key:**
```bash
export OPENAI_API_KEY="sk-..."
# To persist, add to ~/.bashrc or ~/.zshrc
```

**3. Install optional validation tools:**
```bash
# Linting
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Complexity analysis
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

# File watching (macOS/Linux only)
brew install entr
```

**4. Verify setup:**
```bash
go version
echo $OPENAI_API_KEY  # Should print your key
make list              # Should list available experiments
```

---

## Repository Structure

```
hlabgen/
├── README.md                 # This file
├── Makefile                  # Experiment automation
├── cmd/
│   ├── hlabgen/              # CLI: generation + validation orchestration
│   ├── analyze/              # CSV-based results summarizer
│   ├── gather-metrics/       # Metrics aggregation and reporting
│   └── report/               # JSON-to-Markdown report generator
├── internal/
│   ├── ml/                   # GPT-based code generation logic
│   ├── rules/                # Rule-based templates and construction
│   ├── metrics/              # Result structures (shared types)
│   └── validate/             # Build, vet, lint, coverage measurement
├── experiments/
│   ├── input/                # JSON experiment definitions (inputs)
│   ├── out/                  # Generated Go projects (outputs)
│   ├── logs/                 # Logs, aggregated reports, failures
│   └── LibraryAPI.json       # Example experiment specification
└── docs/                     # Documentation
    ├── README.md             # Documentation index
    ├── QUICK_REFERENCE.md
    ├── IMPLEMENTATION_GUIDE.md
    ├── METRICS_ANALYSIS.md
    ├── VISUAL_EXAMPLES.md
    └── RESEARCH_GUIDE.md
```

**Key directories for reproduction:**
- `experiments/input/` → Add your experiment JSON files here
- `experiments/out/` → Generated Go code (auto-created)
- `experiments/logs/` → Final results and reports
- `docs/` → Comprehensive documentation

---

## Running Experiments

### Available Commands

```bash
# List all available experiments
make list

# Generate Go project for one experiment
make generate APP=LibraryAPI

# Run full pipeline for one experiment
make experiment APP=LibraryAPI

# Run all experiments
make all-experiments

# Generate final Markdown report
make report

# Quick test (3 sample experiments)
make quick-test

# Show metrics summary
make stats

# Clean all outputs
make clean

# Watch a file and re-run automatically
make watch APP=LibraryAPI
```

### Typical Workflow

**For a single experiment:**
```bash
make experiment APP=LibraryAPI
# Output: experiments/out/LibraryAPI/ + experiments/logs/results.md
```

**For batch processing:**
```bash
make all-experiments
# Runs all .json files in experiments/input/
# Results: experiments/logs/results.md
```

**For reproducible runs:**
```bash
make clean
make all-experiments
# Clears previous results and reruns from scratch
```

---

## Understanding Results

### Output Structure

After running an experiment, you'll see:

```
experiments/out/LibraryAPI/
├── main.go                   # Generated entry point
├── models/
│   ├── book.go              # Generated domain models
│   └── author.go
├── handlers/
│   └── api.go               # Generated handlers
├── gen_metrics.json          # Generation metrics (JSON)
└── coverage.json             # Coverage report (JSON)

experiments/logs/
├── results.md               # Final Markdown report (ALL experiments)
├── failed_experiments.txt   # List of failures
└── hlabgen.log              # Detailed logs
```

### Metrics Reference

**Model Performance**
- `PrimarySuccess`: Was first GPT output valid JSON?
- `RepairAttempts`: How many self-repairs were needed?
- `FinalSuccess`: Did model produce valid code after all attempts?

**Code Quality**
- `VetWarnings`: Issues detected by `go vet`
- `LintWarnings`: Issues from `golangci-lint`
- `CyclomaticAvg`: Average cyclomatic complexity per function

**Functional Validity**
- `BuildSuccess`: Did `go build` succeed?
- `TestsPass`: Did all tests pass?
- `CoveragePct`: Code coverage percentage

**Efficiency & Stability**
- `GenTimeSec`: Seconds spent in GPT generation
- `RuleFixes`: Automated rule-based fixes applied
- `ErrorMessage`: Description of failure (if any)

### Reading the Report

The final report (`experiments/logs/results.md`) contains a summary table:

| App | Primary Success | Repair Attempts | Final Success | Build Success | Tests Pass | Coverage (%) |
|-----|-----------------|-----------------|----------------|----------------|-------------|---------------|
| LibraryAPI | true | 0 | true | true | true | 89.3 |
| BlogAPI | true | 1 | true | true | true | 91.1 |

**Key metrics to track:**
- **Success Rate** = (# Final Success) / (# Total) × 100%
- **Average Coverage** = Mean of all Coverage (%)
- **Mean Generation Time** = Average GenTimeSec
- **Repair Efficiency** = Mean RepairAttempts (lower is better)

---

## Reproducibility Guide

To ensure your experiments are fully reproducible:

### 1. Version Control
```bash
# Commit all input specifications
git add experiments/input/*.json
git commit -m "Add experiment definitions"

# Save metrics for reuse
git add experiments/out/*/gen_metrics.json
```

### 2. Environment Documentation
Document your exact setup:
```bash
go version
echo $OPENAI_API_KEY | cut -c1-10  # Print first 10 chars
golangci-lint version
gocyclo -version
```

### 3. Clean Reproducible Run
```bash
# Start fresh
make clean

# Run all experiments (captures all metrics)
make all-experiments

# Export results
cp experiments/logs/results.md results_${DATE}.md
cp -r experiments/out/ backup_${DATE}/
```

### 4. Sharing Results
Include in your appendix:
- `experiments/logs/results.md` (summary table)
- `experiments/out/*/gen_metrics.json` (detailed metrics per experiment)
- `experiments/logs/hlabgen.log` (execution logs)
- Your environment setup (Go version, API model, etc.)

### 5. Fixing Issues
If experiments fail:
```bash
# Check logs
tail -100 experiments/logs/hlabgen.log

# View failed cases
cat experiments/logs/failed_experiments.txt

# Re-run one experiment with verbose output
make experiment APP=FailedApp
```

---

## Research Integration

### For Academic Papers

Include this setup in your methodology:

> We conducted experiments using HLabGen, a hybrid AI-assisted code generation framework. All experiments were executed on [DATE] using GPT-4o-Mini with [N] JSON specifications. Code was validated using Go 1.23+ with linting (golangci-lint) and coverage analysis (go test -cover).

### Key Research Questions Enabled

- **Model Capability**: How often does GPT produce valid code on first attempt?
- **Hybrid Effectiveness**: Does rule+ML combination improve test coverage?
- **Resilience**: How many repairs are needed per experiment?
- **Correlation**: Is generation time correlated with final build success?
- **Quality**: How does generated code quality compare to hand-written benchmarks?

### Data Analysis

**Export metrics to CSV:**
```bash
# Using jq (if available)
jq -r '.[] | [.App, .PrimarySuccess, .CoveragePct, .GenTimeSec] | @csv' \
  experiments/out/*/gen_metrics.json > results.csv
```

**Suggested statistics to report:**
- Success rate with confidence intervals
- Mean and median generation time
- Coverage distribution (box plots)
- Correlation between repair attempts and final success

### Citation

If you use HLabGen in your research:

```bibtex
@inproceedings{hlabgen2025,
  title={Hybrid AI-Driven Code Generation in Educational and Organizational Systems},
  author={Gžegoževskė, L. and Gžegoževskis, M.},
  booktitle={Proceedings of the International Conference on Applied Informatics and Automation},
  year={2025}
}
```

---

## Troubleshooting

### API Errors

**Error**: `401 Unauthorized`
- **Cause**: Missing or invalid OpenAI API key
- **Fix**: `export OPENAI_API_KEY="sk-..."` and verify with `echo $OPENAI_API_KEY`

**Error**: `Rate limit exceeded`
- **Cause**: Too many concurrent API calls
- **Fix**: Wait 60 seconds or reduce experiment count

### Build Failures

**Error**: `go: package not found`
- **Cause**: Missing dependencies
- **Fix**: `go mod download` or `go mod tidy`

**Error**: `command not found: golangci-lint`
- **Cause**: Optional linting tool not installed
- **Fix**: Run `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` or skip with `make all-experiments SKIP_LINT=1`

### No Results Generated

**Issue**: `experiments/logs/results.md` not created
- **Check**: `make list` shows experiments available?
- **Check**: `echo $OPENAI_API_KEY` is set?
- **Check**: `tail -50 experiments/logs/hlabgen.log` for errors

### Partial Failures

**Issue**: Some experiments passed, some failed
- **Fix**: Check `experiments/logs/failed_experiments.txt` for details
- **Fix**: Review specific error: `cat experiments/out/FailedApp/gen_metrics.json | jq '.ErrorMessage'`

---

## Next Steps

1. **Read the docs**: Start with [docs/QUICK_REFERENCE.md](docs/QUICK_REFERENCE.md)
2. **Add your experiments**: Place JSON specs in `experiments/input/`
3. **Run**: `make all-experiments`
4. **Analyze**: Open `experiments/logs/results.md`
5. **Integrate**: Use results in your research or publication
6. **Share**: Commit inputs and metrics to version control

For questions or contributions, open an issue on GitHub.

**Happy experimenting!** 🚀
