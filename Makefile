# =====================================================
# 🧪 HLabGen Experiment Automation Makefile (Enhanced)
# =====================================================

# --- Configuration Defaults ---
MODE        ?= hybrid
INPUT_DIR   ?= experiments/input
OUT_DIR     ?= experiments/out
LOG_DIR     ?= experiments/logs
RESULTS_MD  ?= $(LOG_DIR)/results.md
SUMMARY_CSV ?= $(LOG_DIR)/summary.csv

# Automatically detect all input files (JSONs)
INPUT_FILES := $(wildcard $(INPUT_DIR)/*.json)
APP_NAMES   := $(basename $(notdir $(INPUT_FILES)))

# Colors for output
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
	@echo "$(COLOR_BLUE)🧪 HLabGen Experiment Automation$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_GREEN)Available targets:$(COLOR_RESET)"
	@echo "  make generate APP=<name>     - Generate a single app"
	@echo "  make experiment APP=<name>   - Run full pipeline for one app"
	@echo "  make all-experiments         - Run ALL experiments"
	@echo "  make analyze                 - Aggregate and analyze results"
	@echo "  make list                    - List available experiments"
	@echo "  make clean                   - Clean all outputs"
	@echo "  make compare                 - Compare experiment results"
	@echo ""
	@echo "$(COLOR_YELLOW)Available apps:$(COLOR_RESET)"
	@for file in $(INPUT_FILES); do echo "  - $$(basename $$file .json)"; done

# 1️⃣ Generate one app (use: make generate APP=LibraryAPI)
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

# 2️⃣ Run validator only for one project
validate:
	@if [ -z "$(APP)" ]; then \
		echo "$(COLOR_RED)❌ Please specify APP=<AppName>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_BLUE)🔍 Validating $(APP)...$(COLOR_RESET)"
	@go run ./cmd/hlabgen -validate -out $(OUT_DIR)/$(APP)
	@echo "$(COLOR_GREEN)✅ Validation done for $(APP)$(COLOR_RESET)"

# 3️⃣ Run full pipeline for one app (generate + analyze)
experiment:
	@if [ -z "$(APP)" ]; then \
		echo "$(COLOR_RED)❌ Please specify APP=<AppName>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@$(MAKE) generate APP=$(APP)
	@$(MAKE) analyze

# 4️⃣ Run all experiments (loop through all JSONs in /input)
all-experiments:
	@echo "$(COLOR_BLUE)🧬 Running all experiments in $(INPUT_DIR)...$(COLOR_RESET)"
	@mkdir -p $(LOG_DIR)
	@total_count=$$(echo "$(INPUT_FILES)" | wc -w); \
	current=0; \
	failed=0; \
	echo "" > $(LOG_DIR)/failed_experiments.txt; \
	echo "| App | Build | Tests | Coverage (%) | Duration (s) | Repairs |" > $(RESULTS_MD); \
	echo "|-----|--------|--------|---------------|---------------|----------|" >> $(RESULTS_MD); \
	for file in $(INPUT_FILES); do \
		app=$$(basename $$file .json); \
		current=$$((current + 1)); \
		echo ""; \
		echo "$(COLOR_BLUE)=========================================="; \
		echo "🚀 Experiment $$current/$$total_count: $$app"; \
		echo "==========================================$(COLOR_RESET)"; \
		if go run ./cmd/hlabgen -input $$file -mode $(MODE) -out $(OUT_DIR)/$$app 2>&1 | tee $(LOG_DIR)/$$app.log; then \
			cov=$$(grep "Coverage" $(LOG_DIR)/$$app.log | tail -1 | awk '{print $$4}' | tr -d '%'); \
			dur=$$(grep "ML Duration" $(LOG_DIR)/$$app.log | awk '{print $$5}'); \
			build=$$(grep "BuildSuccess" $(LOG_DIR)/$$app.log | awk '{print $$4}'); \
			tests=$$(grep "TestsPass" $(LOG_DIR)/$$app.log | awk '{print $$4}'); \
			repairs=$$(grep "repair" $(LOG_DIR)/$$app.log | awk '{print $$6}' | tr -d ')'); \
			[ -z "$$cov" ] && cov="0.0"; \
			[ -z "$$dur" ] && dur="0.0"; \
			[ -z "$$build" ] && build="false"; \
			[ -z "$$tests" ] && tests="false"; \
			[ -z "$$repairs" ] && repairs="0"; \
			echo "| $$app | $$build | $$tests | $$cov | $$dur | $$repairs |" >> $(RESULTS_MD); \
			echo "$(COLOR_GREEN)✅ $$app completed successfully$(COLOR_RESET)"; \
		else \
			failed=$$((failed + 1)); \
			echo "$$app" >> $(LOG_DIR)/failed_experiments.txt; \
			echo "| $$app | ❌ | ❌ | 0.0 | 0.0 | 0 |" >> $(RESULTS_MD); \
			echo "$(COLOR_RED)❌ $$app failed$(COLOR_RESET)"; \
		fi; \
	done; \
	echo ""; \
	echo "$(COLOR_BLUE)=========================================="; \
	echo "📊 Computing Summary Statistics..."; \
	echo "==========================================$(COLOR_RESET)"; \
	total=$$(grep -v "App" $(RESULTS_MD) | grep -v "^$$" | awk -F'|' 'BEGIN{cov=0;dur=0;count=0;durcount=0} {gsub(/^[ \t]+|[ \t]+$$/, "", $$4); gsub(/^[ \t]+|[ \t]+$$/, "", $$5); if($$4+0>0){cov+=$$4;count++} if($$5+0>0){dur+=$$5;durcount++}} END {if(count>0) printf "Mean Coverage: %.1f%%\n", cov/count; else printf "Mean Coverage: 0.0%%\n"; if(durcount>0) printf "Mean Duration: %.2fs\n", dur/durcount; else printf "Mean Duration: 0.00s\n"}'); \
	passrate=$$(grep -c "| true | true |" $(RESULTS_MD) || echo "0"); \
	buildrate=$$(grep -c "| true |" $(RESULTS_MD) || echo "0"); \
	totalapps=$$(grep -c "^|" $(RESULTS_MD)); \
	totalapps=$$((totalapps - 2)); \
	echo "" >> $(RESULTS_MD); \
	echo "---" >> $(RESULTS_MD); \
	echo "" >> $(RESULTS_MD); \
	echo "## 📊 Summary Statistics" >> $(RESULTS_MD); \
	echo "" >> $(RESULTS_MD); \
	echo "$$total" >> $(RESULTS_MD); \
	echo "Build Success Rate: $$buildrate / $$totalapps projects ($$((buildrate * 100 / totalapps))%)" >> $(RESULTS_MD); \
	echo "Test Pass Rate: $$passrate / $$totalapps projects ($$((passrate * 100 / totalapps))%)" >> $(RESULTS_MD); \
	if [ $$failed -gt 0 ]; then \
		echo "" >> $(RESULTS_MD); \
		echo "## ⚠️ Failed Experiments" >> $(RESULTS_MD); \
		echo "" >> $(RESULTS_MD); \
		cat $(LOG_DIR)/failed_experiments.txt | while read line; do echo "- $$line" >> $(RESULTS_MD); done; \
	fi; \
	echo ""; \
	echo "$(COLOR_GREEN)✅ Completed $$((total_count - failed))/$$total_count experiments$(COLOR_RESET)"; \
	if [ $$failed -gt 0 ]; then \
		echo "$(COLOR_RED)❌ Failed: $$failed experiments$(COLOR_RESET)"; \
	fi
	@$(MAKE) analyze
	@echo ""
	@echo "$(COLOR_GREEN)✅ Results written to:$(COLOR_RESET)"
	@echo "  📄 Markdown: $(RESULTS_MD)"
	@echo "  📊 CSV: $(SUMMARY_CSV)"

# 5️⃣ Aggregate coverage + ML metrics
analyze:
	@echo "$(COLOR_BLUE)📊 Aggregating results...$(COLOR_RESET)"
	@go run ./cmd/analyze
	@echo "$(COLOR_GREEN)✅ Analysis complete$(COLOR_RESET)"

# 6️⃣ Compare results across experiments
compare:
	@echo "$(COLOR_BLUE)📊 Comparing Experiments$(COLOR_RESET)"
	@echo ""
	@if [ ! -f "$(SUMMARY_CSV)" ]; then \
		echo "$(COLOR_RED)❌ No summary.csv found. Run 'make all-experiments' first.$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_GREEN)By Difficulty:$(COLOR_RESET)"
	@echo "Beginner Projects:"
	@grep -E "Blog|Recipe" $(SUMMARY_CSV) 2>/dev/null || echo "  No beginner projects found"
	@echo ""
	@echo "Intermediate Projects:"
	@grep -E "Library|Task|Hotel|Inventory|Fitness|Event" $(SUMMARY_CSV) 2>/dev/null || echo "  No intermediate projects found"
	@echo ""
	@echo "Advanced Projects:"
	@grep -E "Shop|Social|Music" $(SUMMARY_CSV) 2>/dev/null || echo "  No advanced projects found"

# 7️⃣ Clean everything
clean:
	@echo "$(COLOR_YELLOW)🧹 Cleaning all outputs and logs...$(COLOR_RESET)"
	@rm -rf $(OUT_DIR)/*
	@rm -rf $(LOG_DIR)/*
	@echo "$(COLOR_GREEN)✅ Cleaned$(COLOR_RESET)"

# 8️⃣ Quick test - run 3 sample experiments
quick-test:
	@echo "$(COLOR_BLUE)🧪 Running quick test (3 experiments)...$(COLOR_RESET)"
	@$(MAKE) experiment APP=LibraryAPI
	@$(MAKE) experiment APP=BlogAPI
	@$(MAKE) experiment APP=TaskAPI
	@$(MAKE) analyze
	@echo "$(COLOR_GREEN)✅ Quick test complete$(COLOR_RESET)"

# =====================================================
# Utility Shortcuts
# =====================================================

# Quickly list available input experiments
list:
	@echo "$(COLOR_BLUE)📂 Available experiment configurations:$(COLOR_RESET)"
	@for file in $(INPUT_FILES); do \
		app=$$(basename $$file .json); \
		difficulty=$$(grep -o '"difficulty"[[:space:]]*:[[:space:]]*"[^"]*"' $$file | cut -d'"' -f4); \
		entities=$$(grep -o '"entities"[[:space:]]*:[[:space:]]*\[[^]]*\]' $$file | grep -o '"[^"]*"' | wc -l); \
		printf "  $(COLOR_GREEN)%-15s$(COLOR_RESET) [%-12s] (%d entities)\n" $$app $$difficulty $$entities; \
	done

# Re-analyze without regeneration
recalc:
	@$(MAKE) analyze

# Show current statistics
stats:
	@if [ -f "$(SUMMARY_CSV)" ]; then \
		echo "$(COLOR_BLUE)📊 Current Statistics:$(COLOR_RESET)"; \
		tail -n +2 $(SUMMARY_CSV) | awk -F',' '{ \
			total++; \
			if($$5=="true") builds++; \
			if($$6=="true") tests++; \
			cov+=$$7; \
		} END { \
			print "Total Projects:", total; \
			print "Build Success:", builds"/"total, "("int(builds*100/total)"%)"; \
			print "Test Pass:", tests"/"total, "("int(tests*100/total)"%)"; \
			print "Avg Coverage:", sprintf("%.1f%%", cov/total); \
		}'; \
	else \
		echo "$(COLOR_RED)❌ No statistics available. Run experiments first.$(COLOR_RESET)"; \
	fi

# Watch mode - re-run specific experiment on file change (requires entr)
watch:
	@if [ -z "$(APP)" ]; then \
		echo "$(COLOR_RED)❌ Please specify APP=<AppName>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_BLUE)👁️  Watching $(INPUT_DIR)/$(APP).json for changes...$(COLOR_RESET)"
	@echo "$(APP).json" | entr -c make experiment APP=$(APP)

.PHONY: help generate validate experiment all-experiments analyze compare clean list recalc stats quick-test watch