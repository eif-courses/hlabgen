# =====================================================
# 🧪 HLabGen Experiment Automation Makefile (Complete Edition)
# =====================================================

# --- Configuration Defaults ---
MODE        ?= hybrid
INPUT_DIR   ?= experiments/input
OUT_DIR     ?= experiments/out
LOG_DIR     ?= experiments/logs
RESULTS_MD  ?= $(LOG_DIR)/results.md

# Automatically detect all input files (JSONs)
INPUT_FILES := $(wildcard $(INPUT_DIR)/*.json)
APP_NAMES   := $(basename $(notdir $(INPUT_FILES)))

# --- Colors ---
COLOR_RESET  := \033[0m
COLOR_BLUE   := \033[34m
COLOR_GREEN  := \033[32m
COLOR_YELLOW := \033[33m
COLOR_RED    := \033[31m
COLOR_CYAN   := \033[36m
COLOR_PURPLE := \033[35m

# =====================================================
# 🧩 Primary Targets
# =====================================================

.DEFAULT_GOAL := help

# 0️⃣ Help menu
help:
	@echo "$(COLOR_BLUE)╔════════════════════════════════════════════════════════════╗$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)║  🧪 HLabGen Experiment Automation (Complete Edition)      ║$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)╚════════════════════════════════════════════════════════════╝$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_GREEN)📦 Experiment Commands:$(COLOR_RESET)"
	@echo "  make generate APP=<name>     - Generate a single app"
	@echo "  make validate APP=<name>     - Validate an existing app"
	@echo "  make experiment APP=<name>   - Run full pipeline for one app"
	@echo "  make all-experiments         - Run ALL experiments"
	@echo "  make quick-test              - Run a quick 3-app smoke test"
	@echo ""
	@echo "$(COLOR_CYAN)📊 Report Generation Commands:$(COLOR_RESET)"
	@echo "  make report                  - Generate standard Markdown report"
	@echo "  make reports-all             - Generate ALL report types"
	@echo "  make report-comparative      - Comparative analysis (modes)"
	@echo "  make report-statistics       - Statistical analysis + correlations"
	@echo "  make report-failures         - Detailed failure analysis"
	@echo "  make report-latex            - LaTeX tables for papers"
	@echo "  make academic-package        - Complete academic report package"
	@echo ""
	@echo "$(COLOR_PURPLE)🔧 Utility Commands:$(COLOR_RESET)"
	@echo "  make list                    - List available experiments"
	@echo "  make stats                   - Quick statistics summary"
	@echo "  make clean                   - Clean all outputs and logs"
	@echo "  make clean-logs              - Clean only log files"
	@echo "  make watch APP=<name>        - Watch and auto-rerun experiment"
	@echo ""
	@echo "$(COLOR_YELLOW)📋 Available apps:$(COLOR_RESET)"
	@for file in $(INPUT_FILES); do \
		app=$$(basename $$file .json); \
		diff=$$(grep -o '"difficulty"[[:space:]]*:[[:space:]]*"[^"]*"' $$file 2>/dev/null | cut -d'"' -f4 || echo "unknown"); \
		printf "  $(COLOR_GREEN)%-20s$(COLOR_RESET) [%s]\n" $$app $$diff; \
	done
	@echo ""
	@echo "$(COLOR_CYAN)💡 Examples:$(COLOR_RESET)"
	@echo "  make experiment APP=LibraryAPI"
	@echo "  make all-experiments"
	@echo "  make academic-package"

# =====================================================
# 🧬 Experiment Execution
# =====================================================

# 1️⃣ Generate one app
generate:
	@if [ -z "$(APP)" ]; then \
		echo "$(COLOR_RED)❌ Please specify APP=<AppName>$(COLOR_RESET)"; \
		echo "$(COLOR_YELLOW)Available apps:$(COLOR_RESET)"; \
		for file in $(INPUT_FILES); do echo "  - $$(basename $$file .json)"; done; \
		exit 1; \
	fi
	@if [ ! -f "$(INPUT_DIR)/$(APP).json" ]; then \
		echo "$(COLOR_RED)❌ File $(INPUT_DIR)/$(APP).json not found$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_BLUE)🚀 Generating $(APP)...$(COLOR_RESET)"
	@go run ./cmd/hlabgen -input $(INPUT_DIR)/$(APP).json -mode $(MODE) -out $(OUT_DIR)/$(APP)
	@echo "$(COLOR_GREEN)✅ Finished generating $(APP)$(COLOR_RESET)"

# 2️⃣ Validate only
validate:
	@if [ -z "$(APP)" ]; then \
		echo "$(COLOR_RED)❌ Please specify APP=<AppName>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_BLUE)🔍 Validating $(APP)...$(COLOR_RESET)"
	@go run ./cmd/hlabgen -validate -out $(OUT_DIR)/$(APP)
	@echo "$(COLOR_GREEN)✅ Validation done for $(APP)$(COLOR_RESET)"

# 3️⃣ Full pipeline for one app (generate + report)
experiment:
	@if [ -z "$(APP)" ]; then \
		echo "$(COLOR_RED)❌ Please specify APP=<AppName>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@$(MAKE) generate APP=$(APP)
	@$(MAKE) report

# 4️⃣ Run all experiments
all-experiments:
	@echo "$(COLOR_BLUE)╔════════════════════════════════════════════════════════════╗$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)║  🧬 Running all experiments in $(INPUT_DIR)...            ║$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)╚════════════════════════════════════════════════════════════╝$(COLOR_RESET)"
	@mkdir -p $(LOG_DIR)
	@rm -f $(LOG_DIR)/failed_experiments.txt
	@total=$$(echo "$(INPUT_FILES)" | wc -w | tr -d ' '); \
	current=0; failed=0; \
	start_time=$$(date +%s); \
	for file in $(INPUT_FILES); do \
		app=$$(basename $$file .json); \
		current=$$((current + 1)); \
		echo ""; \
		echo "$(COLOR_BLUE)══════════════════════════════════════════════════════════════"; \
		echo "🚀 [$$current/$$total] $$app"; \
		echo "══════════════════════════════════════════════════════════════$(COLOR_RESET)"; \
		if go run ./cmd/hlabgen -input $$file -mode $(MODE) -out $(OUT_DIR)/$$app 2>&1 | tee $(LOG_DIR)/$$app.log; then \
			echo "$(COLOR_GREEN)✅ $$app completed successfully$(COLOR_RESET)"; \
		else \
			failed=$$((failed + 1)); \
			echo "$(COLOR_RED)❌ $$app failed$(COLOR_RESET)"; \
			echo "$$app" >> $(LOG_DIR)/failed_experiments.txt; \
		fi; \
	done; \
	end_time=$$(date +%s); \
	duration=$$((end_time - start_time)); \
	echo ""; \
	echo "$(COLOR_BLUE)╔════════════════════════════════════════════════════════════╗$(COLOR_RESET)"; \
	echo "$(COLOR_BLUE)║  📊 Generating comprehensive reports...                   ║$(COLOR_RESET)"; \
	echo "$(COLOR_BLUE)╚════════════════════════════════════════════════════════════╝$(COLOR_RESET)"; \
	$(MAKE) reports-all; \
	echo ""; \
	echo "$(COLOR_BLUE)╔════════════════════════════════════════════════════════════╗$(COLOR_RESET)"; \
	echo "$(COLOR_BLUE)║  ✅ EXPERIMENT RUN COMPLETE                                ║$(COLOR_RESET)"; \
	echo "$(COLOR_BLUE)╚════════════════════════════════════════════════════════════╝$(COLOR_RESET)"; \
	echo "$(COLOR_GREEN)✅ Completed $$((total - failed))/$$total experiments ($$((100 * (total - failed) / total))%)$(COLOR_RESET)"; \
	if [ $$failed -gt 0 ]; then \
		echo "$(COLOR_RED)❌ Failed: $$failed experiments$(COLOR_RESET)"; \
		echo "$(COLOR_YELLOW)📋 See $(LOG_DIR)/failed_experiments.txt for details$(COLOR_RESET)"; \
	fi; \
	echo "$(COLOR_CYAN)⏱️  Total duration: $$duration seconds$(COLOR_RESET)"; \
	echo "$(COLOR_CYAN)📁 Reports available in: $(LOG_DIR)/$(COLOR_RESET)"

# 5️⃣ Quick smoke test
quick-test:
	@echo "$(COLOR_BLUE)🧪 Running quick test (3 apps)...$(COLOR_RESET)"
	@$(MAKE) experiment APP=LibraryAPI
	@$(MAKE) experiment APP=BlogAPI
	@$(MAKE) experiment APP=TaskManagerAPI
	@echo "$(COLOR_GREEN)✅ Quick test complete$(COLOR_RESET)"

# =====================================================
# 📊 Report Generation
# =====================================================

# Standard report (existing functionality)
report:
	@echo "$(COLOR_BLUE)📊 Generating standard Markdown report...$(COLOR_RESET)"
	@go run ./cmd/report
	@echo "$(COLOR_GREEN)✅ Report ready: $(RESULTS_MD)$(COLOR_RESET)"

# Generate ALL report types
reports-all:
	@echo "$(COLOR_BLUE)╔════════════════════════════════════════════════════════════╗$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)║  📊 Generating ALL report types...                        ║$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)╚════════════════════════════════════════════════════════════╝$(COLOR_RESET)"
	@go run ./cmd/report -mode all
	@echo ""
	@echo "$(COLOR_GREEN)✅ All reports generated successfully!$(COLOR_RESET)"
	@echo "$(COLOR_CYAN)📁 Reports available:$(COLOR_RESET)"
	@echo "  • $(LOG_DIR)/results.md          - Standard results"
	@echo "  • $(LOG_DIR)/comparative.md      - Mode comparison"
	@echo "  • $(LOG_DIR)/statistics.md       - Statistical analysis"
	@echo "  • $(LOG_DIR)/failures.md         - Failure analysis"
	@echo "  • $(LOG_DIR)/tables.tex          - LaTeX tables"

# Individual report types
report-comparative:
	@echo "$(COLOR_BLUE)📊 Generating comparative analysis...$(COLOR_RESET)"
	@go run ./cmd/report -mode comparative
	@echo "$(COLOR_GREEN)✅ Comparative analysis: $(LOG_DIR)/comparative.md$(COLOR_RESET)"

report-statistics:
	@echo "$(COLOR_BLUE)📊 Generating statistical analysis...$(COLOR_RESET)"
	@go run ./cmd/report -mode statistics
	@echo "$(COLOR_GREEN)✅ Statistics report: $(LOG_DIR)/statistics.md$(COLOR_RESET)"

report-failures:
	@echo "$(COLOR_BLUE)📊 Generating failure analysis...$(COLOR_RESET)"
	@go run ./cmd/report -mode failures
	@echo "$(COLOR_GREEN)✅ Failures analysis: $(LOG_DIR)/failures.md$(COLOR_RESET)"

report-latex:
	@echo "$(COLOR_BLUE)📊 Generating LaTeX tables...$(COLOR_RESET)"
	@go run ./cmd/report -mode latex
	@echo "$(COLOR_GREEN)✅ LaTeX tables: $(LOG_DIR)/tables.tex$(COLOR_RESET)"

# Complete academic package (experiments + all reports)
academic-package: all-experiments reports-all
	@echo ""
	@echo "$(COLOR_BLUE)╔════════════════════════════════════════════════════════════╗$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)║  ✅ COMPLETE ACADEMIC PACKAGE READY                        ║$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)╚════════════════════════════════════════════════════════════╝$(COLOR_RESET)"
	@echo "$(COLOR_GREEN)📦 All experiments completed$(COLOR_RESET)"
	@echo "$(COLOR_GREEN)📊 All reports generated$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_CYAN)📁 Package contents (experiments/logs/):$(COLOR_RESET)"
	@ls -lh $(LOG_DIR)/*.md $(LOG_DIR)/*.tex 2>/dev/null || echo "  (no reports found)"
	@echo ""
	@echo "$(COLOR_YELLOW)💡 Ready for paper submission!$(COLOR_RESET)"

# =====================================================
# 🔧 Utility Commands
# =====================================================

# List available experiments with details
list:
	@echo "$(COLOR_BLUE)📂 Available experiment configurations:$(COLOR_RESET)"
	@echo ""
	@printf "$(COLOR_CYAN)%-25s %-15s %-10s$(COLOR_RESET)\n" "NAME" "DIFFICULTY" "ENTITIES"
	@echo "────────────────────────────────────────────────────────"
	@for file in $(INPUT_FILES); do \
		app=$$(basename $$file .json); \
		diff=$$(grep -o '"difficulty"[[:space:]]*:[[:space:]]*"[^"]*"' $$file 2>/dev/null | cut -d'"' -f4 || echo "unknown"); \
		entities=$$(grep -o '"entities"[[:space:]]*:[[:space:]]*\[[^]]*\]' $$file 2>/dev/null | grep -o '"[^"]*"' | wc -l | tr -d ' '); \
		printf "$(COLOR_GREEN)%-25s$(COLOR_RESET) %-15s %-10s\n" $$app $$diff $$entities; \
	done
	@echo ""
	@echo "$(COLOR_YELLOW)Total: $$(echo $(INPUT_FILES) | wc -w | tr -d ' ') experiments$(COLOR_RESET)"

# Quick statistics summary
stats:
	@echo "$(COLOR_BLUE)📊 Quick Statistics Summary$(COLOR_RESET)"
	@echo "────────────────────────────────────────────────────────"
	@if [ -f "$(LOG_DIR)/summary.csv" ]; then \
		echo "$(COLOR_GREEN)CSV Summary:$(COLOR_RESET)"; \
		total=$$(tail -n +2 $(LOG_DIR)/summary.csv | wc -l | tr -d ' '); \
		success=$$(tail -n +2 $(LOG_DIR)/summary.csv | cut -d',' -f5 | grep -c "true" || echo 0); \
		echo "  • Total experiments: $$total"; \
		echo "  • Successful builds: $$success"; \
		if [ $$total -gt 0 ]; then \
			rate=$$((100 * success / total)); \
			echo "  • Success rate: $$rate%"; \
		fi; \
	else \
		echo "$(COLOR_YELLOW)No summary.csv found. Run experiments first.$(COLOR_RESET)"; \
	fi
	@echo ""
	@if [ -f "$(LOG_DIR)/results.md" ]; then \
		echo "$(COLOR_GREEN)Markdown Report:$(COLOR_RESET)"; \
		echo "  • $(LOG_DIR)/results.md"; \
	fi
	@if [ -f "$(LOG_DIR)/statistics.md" ]; then \
		echo "$(COLOR_GREEN)Statistics Report:$(COLOR_RESET)"; \
		echo "  • $(LOG_DIR)/statistics.md"; \
	fi

# Clean everything
clean:
	@echo "$(COLOR_YELLOW)🧹 Cleaning all outputs and logs...$(COLOR_RESET)"
	@rm -rf $(OUT_DIR)/*
	@rm -rf $(LOG_DIR)/*
	@echo "$(COLOR_GREEN)✅ Cleaned$(COLOR_RESET)"

# Clean only logs (keep generated code)
clean-logs:
	@echo "$(COLOR_YELLOW)🧹 Cleaning log files...$(COLOR_RESET)"
	@rm -rf $(LOG_DIR)/*
	@echo "$(COLOR_GREEN)✅ Logs cleaned$(COLOR_RESET)"

# Watch and auto-rerun (requires entr)
watch:
	@if [ -z "$(APP)" ]; then \
		echo "$(COLOR_RED)❌ Please specify APP=<AppName>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@if ! command -v entr > /dev/null; then \
		echo "$(COLOR_RED)❌ entr not found. Install: brew install entr$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_BLUE)👀 Watching $(INPUT_DIR)/$(APP).json...$(COLOR_RESET)"
	@echo "$(APP).json" | entr -c make experiment APP=$(APP)

# Show file sizes and disk usage
disk-usage:
	@echo "$(COLOR_BLUE)💾 Disk Usage Summary$(COLOR_RESET)"
	@echo "────────────────────────────────────────────────────────"
	@echo "$(COLOR_CYAN)Output directory:$(COLOR_RESET)"
	@du -sh $(OUT_DIR) 2>/dev/null || echo "  (empty)"
	@echo ""
	@echo "$(COLOR_CYAN)Logs directory:$(COLOR_RESET)"
	@du -sh $(LOG_DIR) 2>/dev/null || echo "  (empty)"
	@echo ""
	@echo "$(COLOR_CYAN)Generated apps:$(COLOR_RESET)"
	@du -sh $(OUT_DIR)/* 2>/dev/null | head -10 || echo "  (none)"

# Verify environment and dependencies
verify-env:
	@echo "$(COLOR_BLUE)🔍 Verifying environment...$(COLOR_RESET)"
	@echo "────────────────────────────────────────────────────────"
	@echo -n "Go version: "
	@go version || echo "$(COLOR_RED)❌ Go not found$(COLOR_RESET)"
	@echo -n "OpenAI API Key: "
	@if [ -z "$$OPENAI_API_KEY" ]; then \
		echo "$(COLOR_RED)❌ Not set$(COLOR_RESET)"; \
	else \
		echo "$(COLOR_GREEN)✅ Set (length: $${#OPENAI_API_KEY})$(COLOR_RESET)"; \
	fi
	@echo -n "golangci-lint: "
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint version | head -1; \
	else \
		echo "$(COLOR_YELLOW)⚠️  Not installed (optional)$(COLOR_RESET)"; \
	fi
	@echo -n "gocyclo: "
	@if command -v gocyclo > /dev/null; then \
		echo "$(COLOR_GREEN)✅ Installed$(COLOR_RESET)"; \
	else \
		echo "$(COLOR_YELLOW)⚠️  Not installed (optional)$(COLOR_RESET)"; \
	fi
	@echo -n "entr: "
	@if command -v entr > /dev/null; then \
		echo "$(COLOR_GREEN)✅ Installed$(COLOR_RESET)"; \
	else \
		echo "$(COLOR_YELLOW)⚠️  Not installed (optional for watch)$(COLOR_RESET)"; \
	fi

# Show experiment status
status:
	@echo "$(COLOR_BLUE)📊 Experiment Status$(COLOR_RESET)"
	@echo "────────────────────────────────────────────────────────"
	@echo "$(COLOR_CYAN)Input configurations:$(COLOR_RESET) $$(echo $(INPUT_FILES) | wc -w | tr -d ' ')"
	@echo "$(COLOR_CYAN)Generated apps:$(COLOR_RESET) $$(ls -d $(OUT_DIR)/*/ 2>/dev/null | wc -l | tr -d ' ')"
	@echo "$(COLOR_CYAN)Log files:$(COLOR_RESET) $$(ls $(LOG_DIR)/*.log 2>/dev/null | wc -l | tr -d ' ')"
	@if [ -f "$(LOG_DIR)/failed_experiments.txt" ]; then \
		echo "$(COLOR_RED)Failed experiments:$(COLOR_RESET) $$(wc -l < $(LOG_DIR)/failed_experiments.txt | tr -d ' ')"; \
	fi
	@echo ""
	@if [ -f "$(LOG_DIR)/results.md" ]; then \
		echo "$(COLOR_GREEN)✅ Reports generated$(COLOR_RESET)"; \
	else \
		echo "$(COLOR_YELLOW)⚠️  No reports yet (run: make report)$(COLOR_RESET)"; \
	fi

# Archive results for backup
archive:
	@timestamp=$$(date +%Y%m%d_%H%M%S); \
	archive_name="hlabgen_results_$$timestamp.tar.gz"; \
	echo "$(COLOR_BLUE)📦 Creating archive: $$archive_name$(COLOR_RESET)"; \
	tar -czf $$archive_name $(OUT_DIR) $(LOG_DIR); \
	echo "$(COLOR_GREEN)✅ Archive created: $$archive_name$(COLOR_RESET)"

# =====================================================
# 🎯 Phony Targets
# =====================================================

.PHONY: help generate validate experiment all-experiments quick-test \
        report reports-all report-comparative report-statistics report-failures report-latex \
        academic-package list stats clean clean-logs watch disk-usage verify-env status archive