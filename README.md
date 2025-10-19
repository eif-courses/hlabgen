# 🧪 HLabGen Experiment Automation Framework

HLabGen is a **hybrid AI-assisted Go code generator** combining **rule-based scaffolding** and **machine learning (GPT-based)** code synthesis.  
This framework automates **experiment execution**, **validation**, and **metric collection** for reproducible research in software automation and code generation.

---

## 📘 Table of Contents

1. [Overview](#-overview)  
2. [Repository Structure](#-repository-structure)  
3. [Setup & Requirements](#-setup--requirements)  
4. [Quick Start](#-quick-start)  
5. [Experiment Workflow](#-experiment-workflow)  
6. [Available Commands](#-available-commands)  
7. [Metrics Explained](#-metrics-explained)  
8. [For Researchers](#-for-researchers)  
9. [Reproducibility Notes](#-reproducibility-notes)  
10. [Example Experiment Input](#-example-experiment-input)  
11. [Credits](#-credits)

---

## 🧩 Overview

HLabGen enables fully automated testing of **AI-generated Go backends** from JSON specifications.  
Each experiment:
1. Generates Go code using GPT + rule-based templates  
2. Attempts self-repair for invalid JSON outputs  
3. Builds, vets, lints, and tests the generated code  
4. Records quantitative metrics  
5. Produces aggregated results for analysis and publication  

This framework is designed to support **academic experiments** and **automation research** on AI-assisted software engineering.

---

## 📂 Repository Structure

```
hlabgen/
├── cmd/
│   ├── hlabgen/          # Main CLI tool for generation + validation
│   ├── analyze/          # Legacy CSV-based summarizer
│   └── report/           # JSON-based Markdown report generator
├── internal/
│   ├── ml/               # GPT-based code generation logic
│   ├── rules/            # Rule-based code construction templates
│   ├── metrics/          # Shared experiment result structures
│   └── validate/         # Build/test/coverage measurement
├── experiments/
│   ├── input/            # JSON experiment definitions
│   ├── out/              # Generated Go projects
│   ├── logs/             # Logs + aggregated reports
│   └── LibraryAPI.json   # Example input specification
├── Makefile              # Automation of all experiments
└── README.md             # Documentation (you are here)
```

---

## ⚙️ Setup & Requirements

### 🧰 Prerequisites

| Requirement | Description |
|--------------|-------------|
| Go ≥ 1.23 | Required for building and validation |
| OpenAI API key | Needed for GPT-based code generation |
| (Optional) `golangci-lint`, `gocyclo` | For lint and complexity metrics |
| (Optional) `entr` | For live file watching |

### 🔑 Setup

```bash
git clone https://github.com/eif-courses/hlabgen.git
cd hlabgen

# Add your OpenAI key
export OPENAI_API_KEY="sk-..."

# Install validation tools (optional)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
brew install entr  # (optional, macOS/Linux)
```

---

## 🚀 Quick Start

Run all experiments in one command:

```bash
make all-experiments
```

This will:
- Generate and validate all JSON inputs in `experiments/input/`
- Collect build and ML metrics
- Save Markdown results in `experiments/logs/results.md`

---

## 🧬 Experiment Workflow

Each experiment follows these phases:

| Phase | Description |
|--------|--------------|
| **1️⃣ Input** | Load JSON description of entities and endpoints |
| **2️⃣ Generation** | GPT-4o-Mini generates Go code (via OpenAI API) |
| **3️⃣ Repair** | Invalid outputs are automatically repaired and re-parsed |
| **4️⃣ Validation** | `go build`, `go vet`, `golangci-lint`, and `go test -cover` |
| **5️⃣ Metrics** | Store ML + build metrics in JSON |
| **6️⃣ Reporting** | Summarize results in Markdown for analysis |

---

## 💻 Available Commands

| Command | Purpose |
|----------|----------|
| `make list` | List all available experiments in `experiments/input/` |
| `make generate APP=LibraryAPI` | Generate Go project for one app |
| `make experiment APP=LibraryAPI` | Full pipeline for one experiment |
| `make all-experiments` | Run all experiments automatically |
| `make report` | Generate Markdown summary from JSON metrics |
| `make analyze` | Legacy CSV summarizer (optional) |
| `make clean` | Remove all outputs and logs |
| `make quick-test` | Run 3 sample experiments |
| `make stats` | Show quick metrics summary |
| `make watch APP=LibraryAPI` | Watch a single file and rerun automatically |

---

## 🧾 Output Files

After running experiments, the following files are generated:

| File | Description |
|-------|--------------|
| `experiments/out/<App>/` | Generated Go project |
| `experiments/out/<App>/gen_metrics.json` | GPT generation + repair metrics |
| `experiments/out/<App>/coverage.json` | Per-package coverage report |
| `experiments/logs/results.md` | Final Markdown summary of all experiments |
| `experiments/logs/failed_experiments.txt` | List of failed cases |

---

## 📊 Metrics Explained

| Metric | Source | Description |
|---------|---------|-------------|
| **BuildSuccess** | Go compiler | Whether `go build ./...` succeeded |
| **TestsPass** | Go test | True if all tests passed |
| **CoveragePct** | Go test | Code coverage percentage |
| **VetWarnings** | `go vet` | Number of detected warnings |
| **LintWarnings** | `golangci-lint` | Linting issues found |
| **CyclomaticAvg** | `gocyclo` | Average complexity per function |
| **GenTimeSec** | Internal | Duration of ML generation |
| **PrimarySuccess** | ML | Whether the first GPT output was valid JSON |
| **RepairAttempts** | ML | Number of JSON repairs attempted |
| **FinalSuccess** | ML + Validation | Overall success after validation |
| **RuleFixes** | Rule-based | Count of automated code fixes applied |
| **ErrorMessage** | ML | Description of final failure (if any) |

---

## 🧠 For Researchers

This framework was built for **reproducible scientific experiments** on hybrid AI code generation.  
You can include its results directly in your **appendix**, **evaluation**, or **technical report**.

### 📈 Interpreting Results

| Category | Metrics | Meaning |
|-----------|----------|---------|
| **Model Performance** | `PrimarySuccess`, `RepairAttempts`, `FinalSuccess` | Indicates the model’s ability to produce syntactically valid code |
| **Code Quality** | `VetWarnings`, `LintWarnings`, `CyclomaticAvg` | Reflects structural and maintainability properties |
| **Functional Validity** | `BuildSuccess`, `TestsPass`, `CoveragePct` | Verifies executable and tested code |
| **Efficiency** | `GenTimeSec`, `Duration` | Measures time to generation and validation |
| **Stability** | `RuleFixes`, `RepairAttempts` | Shows resilience of rule/ML hybrid system |

### 📊 Data Aggregation

Use:
```bash
make report
```

This produces a Markdown table:
```markdown
| App | Primary Success | Repair Attempts | Rule Fixes | Final Success | Build Success | Tests Pass | Coverage (%) | Duration (s) |
|-----|-----------------|----------------|-------------|----------------|----------------|-------------|---------------|---------------|
| LibraryAPI | true | 0 | 2 | true | true | true | 89.3 | 8.42 |
| ShopAPI | false | 1 | 4 | false | false | false | 0.0 | 12.21 |
```

### 📚 Suggested Evaluation Metrics

In your paper or thesis, you can summarize:
- **Success Rate (%)** = #FinalSuccess / #Total
- **Average Coverage (%)**
- **Mean Generation Time (s)**
- **Repair Attempts per Experiment**
- **Correlation** between coverage and ML duration

---

## ♻️ Reproducibility Notes

To ensure full reproducibility:
1. Store all inputs (`experiments/input/*.json`) in version control.
2. Keep generated `gen_metrics.json` files for dataset reuse.
3. Use `go version` lock files (`go.mod`).
4. Document your OpenAI model version (currently GPT-4o-mini).
5. Run all experiments on a clean environment:
   ```bash
   make clean && make all-experiments
   ```
6. Export metrics as CSV or JSON for supplementary material.

---

## 🧩 Example Experiment Input

File: `experiments/input/LibraryAPI.json`
```json
{
  "title": "LibraryAPI",
  "description": "A REST API for managing books and authors",
  "difficulty": "intermediate",
  "entities": ["Book", "Author"],
  "endpoints": ["CreateBook", "GetBooks", "DeleteBook"]
}
```

---

## 🔍 Example Generated Output

After running:
```bash
make experiment APP=LibraryAPI
```

You’ll get:
```
experiments/out/LibraryAPI/
├── main.go
├── models/
│   └── book.go
├── handlers/
│   └── books.go
├── gen_metrics.json
└── coverage.json
```

---

## 🧮 Example Report Snippet

Generated at:
```
experiments/logs/results.md
```

Example:
```markdown
# Experimental Evaluation Results

| App | Primary Success | Repair Attempts | Rule Fixes | Final Success | Build Success | Tests Pass | Coverage (%) | Duration (s) |
|-----|-----------------|----------------|-------------|----------------|----------------|-------------|---------------|---------------|
| LibraryAPI | true | 0 | 2 | true | true | true | 89.3 | 8.42 |
| BlogAPI | true | 1 | 3 | true | true | true | 91.1 | 9.57 |
| TaskManagerAPI | true | 0 | 4 | true | true | true | 88.7 | 7.93 |

✅ 3/3 experiments succeeded (100%)
```

---

## 🧠 Suggested Research Questions

You can analyze this dataset to answer questions like:
- How often does GPT produce valid code on the first attempt?  
- Does hybrid (rule + ML) generation improve coverage and test success?  
- How do repairs affect complexity and maintainability?  
- Is generation time correlated with final build success?

---

## 🧾 Example Academic Citation

If you use this setup in your research, please cite:

> Gžegoževskis, M., Gžegoževskė, L., & Vasaitienė, J. (2025).  
> **Hybrid AI-Driven Code Generation in Educational and Organizational Systems.**  
> In *Proceedings of the International Conference on Applied Informatics and Automation*, pp. XX–XX.

---

## 🙌 Credits

Developed by  
**Marius Gžegoževskis**, ****, and ****  
as part of the *HLabGen Project* on AI-driven software automation in higher education.


# Prepare dependencies
go mod tidy

# Run one experiment
make experiment APP=BlogAPI

# Run all
make all-experiments

# Generate summary (if needed manually)
make report

# Check results
cat experiments/logs/results.md

# Clean up
make clean
