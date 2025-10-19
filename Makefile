# =====================================================
# 🧪 HLabGen Experiment Automation Makefile (JSON Edition - Final)
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

# =====================================================
# 🧩 Primary Targets
# =====================================================

.DEFAULT_GOAL := help

# 0️⃣ Help menu
help:
	@echo "$(COLOR_BLUE)🧪 HLabGen Experiment Automation (JSON Edition)$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_GREEN)Available targets:$(COLOR_RESET)"
	@echo "  make generate APP=<name>     - Generate a single app"
	@echo "  make experiment APP=<name>   - Run full pipeline for one app"
	@echo "  make all-experiments         - Run ALL experiments"
	@echo "  make report                  - Generate Markdown report from JSON metrics"
	@echo "  make list                    - List available experiments"
	@echo "  make clean                   - Clean all outputs and logs"
	@echo "  make quick-test              - Run a quick 3-app smoke test"
	@echo ""
	@echo "$(COLOR_YELLOW)Available apps:$(COLOR_RESET)"
	@for file in $(INPUT_FILES); do echo "  - $$(basename $$file .json)"; done

# 1️⃣ Generate one app
generate:
	@if [ -z "$(APP)" ]; then \
		echo "$(COLOR_RED)❌ Please specify APP=<AppName>$(COLOR_RESET)"; \
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
	@echo "$(COLOR_BLUE)🧬 Running all experiments in $(INPUT_DIR)...$(COLOR_RESET)"
	@mkdir -p $(LOG_DIR)
	@total=$$(echo "$(INPUT_FILES)" | wc -w); \
	current=0; failed=0; \
	for file in $(INPUT_FILES); do \
		app=$$(basename $$file .json); \
		current=$$((current + 1)); \
		echo ""; \
		echo "$(COLOR_BLUE)=========================================="; \
		echo "🚀 $$app ($$current of $$total)"; \
		echo "==========================================$(COLOR_RESET)"; \
		if go run ./cmd/hlabgen -input $$file -mode $(MODE) -out $(OUT_DIR)/$$app 2>&1 | tee $(LOG_DIR)/$$app.log; then \
			echo "$(COLOR_GREEN)✅ $$app completed successfully$(COLOR_RESET)"; \
		else \
			failed=$$((failed + 1)); \
			echo "$(COLOR_RED)❌ $$app failed$(COLOR_RESET)"; \
			echo "$$app" >> $(LOG_DIR)/failed_experiments.txt; \
		fi; \
	done; \
	echo ""; \
	echo "$(COLOR_BLUE)📊 Generating summary report...$(COLOR_RESET)"; \
	$(MAKE) report; \
	echo ""; \
	echo "$(COLOR_GREEN)✅ Completed $$((total - failed))/$$total experiments$(COLOR_RESET)"

# 5️⃣ Generate Markdown report from JSON metrics
report:
	@echo "$(COLOR_BLUE)📊 Generating Markdown report from JSON metrics...$(COLOR_RESET)"
	@go run ./cmd/report
	@echo "$(COLOR_GREEN)✅ Markdown report ready: $(RESULTS_MD)$(COLOR_RESET)"

# 6️⃣ Clean everything
clean:
	@echo "$(COLOR_YELLOW)🧹 Cleaning outputs and logs...$(COLOR_RESET)"
	@rm -rf $(OUT_DIR)/*
	@rm -rf $(LOG_DIR)/*
	@echo "$(COLOR_GREEN)✅ Cleaned$(COLOR_RESET)"

# 7️⃣ Quick smoke test
quick-test:
	@echo "$(COLOR_BLUE)🧪 Running quick test (3 apps)...$(COLOR_RESET)"
	@$(MAKE) experiment APP=LibraryAPI
	@$(MAKE) experiment APP=BlogAPI
	@$(MAKE) experiment APP=TaskManagerAPI
	@echo "$(COLOR_GREEN)✅ Quick test complete$(COLOR_RESET)"

# 8️⃣ Utility - List available experiments
list:
	@echo "$(COLOR_BLUE)📂 Available experiment configurations:$(COLOR_RESET)"
	@for file in $(INPUT_FILES); do \
		app=$$(basename $$file .json); \
		diff=$$(grep -o '"difficulty"[[:space:]]*:[[:space:]]*"[^"]*"' $$file | cut -d'"' -f4); \
		printf "  $(COLOR_GREEN)%-15s$(COLOR_RESET) [%s]\n" $$app $$diff; \
	done

.PHONY: help generate validate experiment all-experiments report clean list quick-test
