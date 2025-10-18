# =====================================================
# 🧪 HLabGen Experiment Automation Makefile (with Summary Stats)
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

# =====================================================
# 🧩 Primary Targets
# =====================================================

# 1️⃣ Generate one app (use: make generate APP=LibraryAPI)
generate:
	@if [ -z "$(APP)" ]; then echo "❌ Please specify APP=<AppName>"; exit 1; fi
	@echo "🚀 Generating $(APP)..."
	go run ./cmd/hlabgen -input $(INPUT_DIR)/$(APP).json -mode $(MODE) -out $(OUT_DIR)/$(APP)
	@echo "✅ Finished generating $(APP)"

# 2️⃣ Run validator only for one project
validate:
	@if [ -z "$(APP)" ]; then echo "❌ Please specify APP=<AppName>"; exit 1; fi
	@echo "🔍 Validating $(APP)..."
	go run ./cmd/hlabgen -validate -out $(OUT_DIR)/$(APP)
	@echo "✅ Validation done for $(APP)"

# 3️⃣ Run full pipeline for one app (generate + analyze)
experiment:
	@if [ -z "$(APP)" ]; then echo "❌ Please specify APP=<AppName>"; exit 1; fi
	@$(MAKE) generate APP=$(APP)
	@$(MAKE) analyze

# 4️⃣ Run all experiments (loop through all JSONs in /input)
all-experiments:
	@echo "🧬 Running all experiments in $(INPUT_DIR)..."
	@mkdir -p $(LOG_DIR)
	@echo "| App | Build | Tests | Coverage (%) | Duration (s) | Repairs |" > $(RESULTS_MD)
	@echo "|-----|--------|--------|---------------|---------------|----------|" >> $(RESULTS_MD)
	@for file in $(INPUT_FILES); do \
		app=$$(basename $$file .json); \
		echo "🚀 Running experiment for $$app..."; \
		go run ./cmd/hlabgen -input $$file -mode $(MODE) -out $(OUT_DIR)/$$app | tee $(LOG_DIR)/$$app.log; \
		cov=$$(grep "Coverage" $(LOG_DIR)/$$app.log | tail -1 | awk '{print $$3}' | tr -d '%'); \
		dur=$$(grep "ML Duration" $(LOG_DIR)/$$app.log | awk '{print $$5}'); \
		build=$$(grep "BuildSuccess" $(LOG_DIR)/$$app.log | awk '{print $$3}'); \
		tests=$$(grep "TestsPass" $(LOG_DIR)/$$app.log | awk '{print $$3}'); \
		repairs=$$(grep "repair" $(LOG_DIR)/$$app.log | awk '{print $$6}'); \
		echo "| $$app | $$build | $$tests | $$cov | $$dur | $$repairs |" >> $(RESULTS_MD); \
	done; \
	total=$$(grep -v "App" $(RESULTS_MD) | awk -F'|' '{if($$4+0>0){cov+=$$4;count++} if($$5+0>0){dur+=$$5}} END {if(count>0) printf "Mean Coverage: %.1f%%\nMean Duration: %.2fs\n", cov/count, dur/count}'); \
	passrate=$$(grep -c "| true | true |" $(RESULTS_MD)); \
	totalapps=$$(grep -c "^|" $(RESULTS_MD)); \
	echo "" >> $(RESULTS_MD); \
	echo "**Summary Statistics:**" >> $(RESULTS_MD); \
	echo "$$total" >> $(RESULTS_MD); \
	echo "Test Pass Rate: $$passrate / $$totalapps projects passed tests" >> $(RESULTS_MD)
	@$(MAKE) analyze
	@echo "\n✅ Markdown summary with averages written to $(RESULTS_MD)"

# 5️⃣ Aggregate coverage + ML metrics
analyze:
	@echo "📊 Aggregating results..."
	go run ./cmd/analyze

# 6️⃣ Clean everything
clean:
	rm -rf $(OUT_DIR)/*
	rm -rf $(LOG_DIR)/*
	@echo "🧹 Cleaned all outputs and logs"

# =====================================================
# Utility Shortcuts
# =====================================================

# Quickly list available input experiments
list:
	@echo "📂 Available input files:"
	@for file in $(INPUT_FILES); do echo " - $$(basename $$file)"; done

# Re-analyze without regeneration
recalc:
	@$(MAKE) analyze

.PHONY: generate validate experiment all-experiments analyze clean list recalc
