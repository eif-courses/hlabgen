# =====================================================
# 🧪 HLabGen Experiment Automation Makefile (Complete Edition with Fixed Cleaning)
# =====================================================

# --- Configuration Defaults ---
MODE        ?= hybrid
INPUT_DIR   ?= experiments/input
OUT_DIR     ?= experiments/out
LOG_DIR     ?= experiments/logs
ARCHIVE_DIR ?= experiments/archives
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
	@echo "$(COLOR_PURPLE)🧹 Cleaning Commands:$(COLOR_RESET)"
	@echo "  make clean                   - Clean all outputs and logs"
	@echo "  make clean-safe              - Clean but preserve metrics"
	@echo "  make clean-code              - Clean only generated code"
	@echo "  make clean-logs              - Clean only log files"
	@echo "  make clean-archive           - Archive then clean (SAFEST)"
	@echo "  make clean-all               - Complete clean (with archive)"
	@echo "  make clean-dry-run           - Show what would be deleted"
	@echo "  make clean-force             - Nuclear option (removes everything)"
	@echo ""
	@echo "$(COLOR_CYAN)💾 Backup & Archive Commands:$(COLOR_RESET)"
	@echo "  make archive                 - Create compressed backup"
	@echo "  make archive-metrics         - Archive only metrics"
	@echo "  make backup                  - Full backup with timestamp"
	@echo "  make list-archives           - Show all archived data"
	@echo "  make restore-latest          - Restore from latest archive"
	@echo ""
	@echo "$(COLOR_PURPLE)🔧 Utility Commands:$(COLOR_RESET)"
	@echo "  make list                    - List available experiments"
	@echo "  make stats                   - Quick statistics summary"
	@echo "  make status                  - Check experiment status"
	@echo "  make verify-env              - Verify environment setup"
	@echo "  make watch APP=<name>        - Watch and auto-rerun experiment"
	@echo "  make disk-usage              - Check disk space usage"
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
	@echo "  make clean-archive           # Safe clean with backup"

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
# 🧹 Cleaning Commands (FIXED - Handles correct folder structure)
# =====================================================

# Standard clean - removes everything including empty dirs
clean:
	@echo "$(COLOR_YELLOW)🧹 Cleaning all outputs and logs...$(COLOR_RESET)"
	@echo "$(COLOR_RED)⚠️  This will delete all generated code and metrics!$(COLOR_RESET)"
	@echo "$(COLOR_YELLOW)💡 Use 'make clean-archive' to backup first$(COLOR_RESET)"
	@if [ -d "$(OUT_DIR)" ]; then \
		rm -rf $(OUT_DIR); \
		mkdir -p $(OUT_DIR); \
	fi
	@if [ -d "$(LOG_DIR)" ]; then \
		rm -rf $(LOG_DIR); \
		mkdir -p $(LOG_DIR); \
	fi
	@for dir in experiments/*/; do \
		if [ -d "$$dir" ] && [ "$$(basename $$dir)" != "input" ] && [ "$$(basename $$dir)" != "out" ] && [ "$$(basename $$dir)" != "logs" ] && [ "$$(basename $$dir)" != "archives" ]; then \
			rm -rf "$$dir"; \
		fi; \
	done
	@echo "$(COLOR_GREEN)✅ Cleaned$(COLOR_RESET)"

# Clean but preserve metrics files
clean-safe:
	@echo "$(COLOR_YELLOW)🧹 Cleaning outputs (preserving metrics)...$(COLOR_RESET)"
	@if [ -d "$(OUT_DIR)" ]; then \
		for dir in $(OUT_DIR)/*/; do \
			if [ -d "$$dir" ]; then \
				find "$$dir" -type f ! -name '*metrics*.json' ! -name 'coverage.json' -delete 2>/dev/null; \
			fi; \
		done; \
		find $(OUT_DIR) -type d -empty -delete 2>/dev/null; \
	fi
	@for dir in experiments/*/; do \
		dirname=$$(basename "$$dir"); \
		if [ "$$dirname" != "input" ] && [ "$$dirname" != "out" ] && [ "$$dirname" != "logs" ] && [ "$$dirname" != "archives" ] && [ -d "$$dir" ]; then \
			find "$$dir" -type f ! -name '*metrics*.json' ! -name 'coverage.json' -delete 2>/dev/null; \
			if [ -z "$$(ls -A $$dir 2>/dev/null | grep -v '.*metrics.*\.json\|coverage\.json')" ]; then \
				: ; \
			fi; \
		fi; \
	done
	@echo "$(COLOR_GREEN)✅ Cleaned (metrics preserved)$(COLOR_RESET)"

# Clean only generated code, keep all metrics
clean-code:
	@echo "$(COLOR_YELLOW)🧹 Cleaning generated code only...$(COLOR_RESET)"
	@if [ -d "$(OUT_DIR)" ]; then \
		for dir in $(OUT_DIR)/*/; do \
			if [ -d "$$dir" ]; then \
				rm -rf "$$dir/cmd" "$$dir/internal" "$$dir/go.mod" "$$dir/go.sum" "$$dir/tasks.md" 2>/dev/null || true; \
			fi; \
		done; \
		find $(OUT_DIR) -type d -empty -delete 2>/dev/null; \
	fi
	@echo "$(COLOR_GREEN)✅ Code cleaned (all metrics preserved)$(COLOR_RESET)"

# Clean only logs (keep generated code and metrics)
clean-logs:
	@echo "$(COLOR_YELLOW)🧹 Cleaning log files...$(COLOR_RESET)"
	@if [ -d "$(LOG_DIR)" ]; then \
		rm -rf $(LOG_DIR); \
		mkdir -p $(LOG_DIR); \
	fi
	@echo "$(COLOR_GREEN)✅ Logs cleaned$(COLOR_RESET)"

# Archive then clean (safest option)
clean-archive: archive-metrics
	@echo "$(COLOR_YELLOW)🧹 Performing clean after archiving...$(COLOR_RESET)"
	@if [ -d "$(OUT_DIR)" ]; then \
		rm -rf $(OUT_DIR); \
		mkdir -p $(OUT_DIR); \
	fi
	@if [ -d "$(LOG_DIR)" ]; then \
		rm -rf $(LOG_DIR); \
		mkdir -p $(LOG_DIR); \
	fi
	@for dir in experiments/*/; do \
		dirname=$$(basename "$$dir"); \
		if [ "$$dirname" != "input" ] && [ "$$dirname" != "out" ] && [ "$$dirname" != "logs" ] && [ "$$dirname" != "archives" ] && [ -d "$$dir" ]; then \
			rm -rf "$$dir"; \
		fi; \
	done
	@echo "$(COLOR_GREEN)✅ Cleaned after archiving$(COLOR_RESET)"

# Complete clean with automatic backup
clean-all: backup
	@echo "$(COLOR_YELLOW)🧹 Performing complete clean...$(COLOR_RESET)"
	@if [ -d "$(OUT_DIR)" ]; then \
		rm -rf $(OUT_DIR); \
		mkdir -p $(OUT_DIR); \
	fi
	@if [ -d "$(LOG_DIR)" ]; then \
		rm -rf $(LOG_DIR); \
		mkdir -p $(LOG_DIR); \
	fi
	@if [ -d "$(ARCHIVE_DIR)" ]; then \
		rm -rf $(ARCHIVE_DIR); \
	fi
	@for dir in experiments/*/; do \
		dirname=$$(basename "$$dir"); \
		if [ "$$dirname" != "input" ] && [ -d "$$dir" ]; then \
			rm -rf "$$dir"; \
		fi; \
	done
	@mkdir -p $(OUT_DIR) $(LOG_DIR)
	@echo "$(COLOR_GREEN)✅ Complete clean (backup saved)$(COLOR_RESET)"

# Show what would be deleted (dry run)
clean-dry-run:
	@echo "$(COLOR_BLUE)🔍 Files and directories that would be deleted by 'make clean':$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_YELLOW)📁 Directories in experiments/:$(COLOR_RESET)"
	@for dir in experiments/*/; do \
		dirname=$$(basename "$$dir"); \
		if [ "$$dirname" != "input" ] && [ "$$dirname" != "out" ] && [ "$$dirname" != "logs" ] && [ "$$dirname" != "archives" ] && [ -d "$$dir" ]; then \
			size=$$(du -sh "$$dir" 2>/dev/null | cut -f1); \
			files=$$(find "$$dir" -type f 2>/dev/null | wc -l); \
			echo "  ❌ $$dirname/ ($$size, $$files files)"; \
		fi; \
	done
	@echo ""
	@echo "$(COLOR_YELLOW)📁 experiments/out/:$(COLOR_RESET)"
	@if [ -d "$(OUT_DIR)" ]; then \
		size=$$(du -sh $(OUT_DIR) 2>/dev/null | cut -f1); \
		files=$$(find $(OUT_DIR) -type f 2>/dev/null | wc -l); \
		echo "  ❌ All contents ($$size, $$files files)"; \
	else \
		echo "  (empty)"; \
	fi
	@echo ""
	@echo "$(COLOR_YELLOW)📁 experiments/logs/:$(COLOR_RESET)"
	@if [ -d "$(LOG_DIR)" ]; then \
		size=$$(du -sh $(LOG_DIR) 2>/dev/null | cut -f1); \
		files=$$(find $(LOG_DIR) -type f 2>/dev/null | wc -l); \
		echo "  ❌ All contents ($$size, $$files files)"; \
	else \
		echo "  (empty)"; \
	fi
	@echo ""
	@echo "$(COLOR_RED)⚠️  Metrics files that would be LOST:$(COLOR_RESET)"
	@find experiments/ -name '*metrics*.json' -o -name 'coverage.json' 2>/dev/null | sed 's/^/  /' || echo "  (none)"
	@echo ""
	@echo "$(COLOR_CYAN)💡 Use 'make clean-archive' to backup before cleaning$(COLOR_RESET)"

# Force remove everything (nuclear option)
clean-force:
	@echo "$(COLOR_RED)💥 FORCE CLEAN - Removing all experiment data...$(COLOR_RESET)"
	@rm -rf $(OUT_DIR) $(LOG_DIR) $(ARCHIVE_DIR)
	@for dir in experiments/*/; do \
		dirname=$$(basename "$$dir"); \
		if [ "$$dirname" != "input" ] && [ -d "$$dir" ]; then \
			rm -rf "$$dir"; \
		fi; \
	done
	@mkdir -p $(OUT_DIR) $(LOG_DIR)
	@echo "$(COLOR_GREEN)✅ Force clean complete (directories recreated)$(COLOR_RESET)"

# =====================================================
# 💾 Backup & Archive Commands (FIXED - Correct paths)
# =====================================================

# Archive only metrics files with better naming
archive-metrics:
	@timestamp=$$(date +%Y%m%d_%H%M%S); \
	archive_dir="$(ARCHIVE_DIR)/metrics_$$timestamp"; \
	mkdir -p "$$archive_dir"; \
	echo "$(COLOR_BLUE)📦 Archiving metrics to $$archive_dir$(COLOR_RESET)"; \
	file_count=0; \
	for dir in experiments/*/; do \
		app_name=$$(basename "$$dir"); \
		if [ "$$app_name" != "input" ] && [ "$$app_name" != "out" ] && [ "$$app_name" != "logs" ] && [ "$$app_name" != "archives" ] && [ -d "$$dir" ]; then \
			for file in "$$dir"gen_metrics.json "$$dir"coverage.json "$$dir"*metrics*.json; do \
				if [ -f "$$file" ]; then \
					file_name=$$(basename "$$file"); \
					cp "$$file" "$$archive_dir/$${app_name}_$$file_name" 2>/dev/null && file_count=$$((file_count + 1)); \
				fi; \
			done; \
		fi; \
	done; \
	if [ -d "$(OUT_DIR)" ]; then \
		for file in $$(find $(OUT_DIR) -name '*metrics*.json' -o -name 'coverage.json' 2>/dev/null); do \
			app_name=$$(basename $$(dirname $$file)); \
			file_name=$$(basename $$file); \
			cp "$$file" "$$archive_dir/out_$${app_name}_$$file_name" 2>/dev/null && file_count=$$((file_count + 1)); \
		done; \
	fi; \
	if [ -d "$(LOG_DIR)" ]; then \
		for file in $(LOG_DIR)/*.csv $(LOG_DIR)/*.md $(LOG_DIR)/*.tex; do \
			if [ -f "$$file" ]; then \
				cp "$$file" "$$archive_dir/" 2>/dev/null && file_count=$$((file_count + 1)); \
			fi; \
		done; \
	fi; \
	if [ $$file_count -eq 0 ]; then \
		echo "$(COLOR_YELLOW)⚠️  No files to archive$(COLOR_RESET)"; \
		rmdir "$$archive_dir" 2>/dev/null || true; \
		rmdir "$(ARCHIVE_DIR)" 2>/dev/null || true; \
	else \
		echo "$(COLOR_GREEN)✅ Archived $$file_count files to $$archive_dir$(COLOR_RESET)"; \
	fi
# Create compressed archive of all results
archive:
	@timestamp=$$(date +%Y%m%d_%H%M%S); \
	archive_name="hlabgen_results_$$timestamp.tar.gz"; \
	echo "$(COLOR_BLUE)📦 Creating compressed archive: $$archive_name$(COLOR_RESET)"; \
	tar -czf $$archive_name $(OUT_DIR) $(LOG_DIR) 2>/dev/null; \
	size=$$(du -h $$archive_name | cut -f1); \
	echo "$(COLOR_GREEN)✅ Archive created: $$archive_name ($$size)$(COLOR_RESET)"

# Full backup with timestamp
backup:
	@timestamp=$$(date +%Y%m%d_%H%M%S); \
	backup_name="hlabgen_backup_$$timestamp.tar.gz"; \
	echo "$(COLOR_BLUE)📦 Creating full backup: $$backup_name$(COLOR_RESET)"; \
	tar -czf $$backup_name experiments/ 2>/dev/null; \
	size=$$(du -h $$backup_name | cut -f1); \
	echo "$(COLOR_GREEN)✅ Backup created: $$backup_name ($$size)$(COLOR_RESET)"; \
	echo "$(COLOR_CYAN)💡 To restore: tar -xzf $$backup_name$(COLOR_RESET)"

# List all archives
list-archives:
	@echo "$(COLOR_BLUE)📂 Available Archives:$(COLOR_RESET)"
	@echo ""
	@if [ -d "$(ARCHIVE_DIR)" ]; then \
		if [ -n "$$(ls -A $(ARCHIVE_DIR) 2>/dev/null)" ]; then \
			for dir in $(ARCHIVE_DIR)/*/; do \
				if [ -d "$$dir" ]; then \
					count=$$(find "$$dir" -type f | wc -l); \
					size=$$(du -sh "$$dir" | cut -f1); \
					echo "  📁 $$(basename $$dir) - $$count files ($$size)"; \
				fi; \
			done; \
		else \
			echo "  (no archives)"; \
		fi; \
	else \
		echo "  (no archives directory)"; \
	fi
	@echo ""
	@echo "$(COLOR_BLUE)📦 Backup Files:$(COLOR_RESET)"
	@ls -lht hlabgen_backup_*.tar.gz hlabgen_results_*.tar.gz 2>/dev/null || echo "  (no backups)"

# Restore from latest archive
restore-latest:
	@latest=$$(ls -td $(ARCHIVE_DIR)/metrics_*/ 2>/dev/null | head -1); \
	if [ -z "$$latest" ]; then \
		echo "$(COLOR_RED)❌ No archives found$(COLOR_RESET)"; \
		exit 1; \
	fi; \
	echo "$(COLOR_BLUE)📥 Restoring from: $$latest$(COLOR_RESET)"; \
	restored=0; \
	for file in "$$latest"/*; do \
		if [ -f "$$file" ]; then \
			filename=$$(basename "$$file"); \
			if echo "$$filename" | grep -q "_"; then \
				app=$$(echo "$$filename" | cut -d'_' -f1); \
				metric=$$(echo "$$filename" | cut -d'_' -f2-); \
				mkdir -p "$(OUT_DIR)/$$app"; \
				cp "$$file" "$(OUT_DIR)/$$app/$$metric" 2>/dev/null && restored=$$((restored + 1)); \
			else \
				cp "$$file" "$(LOG_DIR)/" 2>/dev/null && restored=$$((restored + 1)); \
			fi; \
		fi; \
	done; \
	echo "$(COLOR_GREEN)✅ Restored $$restored files from archive$(COLOR_RESET)"; \
	echo "$(COLOR_CYAN)💡 Run 'make reports-all' to regenerate reports$(COLOR_RESET)"

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
	@if [ -f "$(LOG_DIR)/coverage.csv" ]; then \
		echo "$(COLOR_GREEN)Coverage Data:$(COLOR_RESET)"; \
		avg=$$(tail -n +2 $(LOG_DIR)/coverage.csv | cut -d',' -f4 | awk '{s+=$$1; c++} END {if(c>0) printf "%.1f", s/c; else print "0.0"}'); \
		echo "  • Average coverage: $$avg%"; \
	fi
	@echo ""
	@echo "$(COLOR_CYAN)Reports:$(COLOR_RESET)"
	@ls -1 $(LOG_DIR)/*.md 2>/dev/null | xargs -I {} basename {} | sed 's/^/  • /' || echo "  (no reports)"

# Show experiment status
status:
	@echo "$(COLOR_BLUE)📊 Experiment Status$(COLOR_RESET)"
	@echo "────────────────────────────────────────────────────────"
	@echo "$(COLOR_CYAN)Input configurations:$(COLOR_RESET) $$(echo $(INPUT_FILES) | wc -w | tr -d ' ')"
	@echo "$(COLOR_CYAN)Generated apps (experiments/):$(COLOR_RESET) $$(find experiments/ -maxdepth 1 -mindepth 1 -type d ! -name 'input' ! -name 'out' ! -name 'logs' ! -name 'archives' 2>/dev/null | wc -l | tr -d ' ')"
	@echo "$(COLOR_CYAN)Generated apps (experiments/out/):$(COLOR_RESET) $$(find $(OUT_DIR) -mindepth 1 -maxdepth 1 -type d 2>/dev/null | wc -l | tr -d ' ')"
	@echo "$(COLOR_CYAN)Metrics files (all):$(COLOR_RESET) $$(find experiments/ -name '*metrics*.json' 2>/dev/null | wc -l | tr -d ' ')"
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
	@echo ""
	@if [ -d "$(ARCHIVE_DIR)" ] && [ -n "$$(ls -A $(ARCHIVE_DIR) 2>/dev/null)" ]; then \
		archive_count=$$(find $(ARCHIVE_DIR) -mindepth 1 -maxdepth 1 -type d 2>/dev/null | wc -l); \
		echo "$(COLOR_CYAN)Archives:$(COLOR_RESET) $$archive_count"; \
	fi

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
	@echo "$(COLOR_CYAN)Archives:$(COLOR_RESET)"
	@du -sh $(ARCHIVE_DIR) 2>/dev/null || echo "  (no archives)"
	@echo ""
	@echo "$(COLOR_CYAN)Largest apps (top 5):$(COLOR_RESET)"
	@du -sh $(OUT_DIR)/* 2>/dev/null | sort -hr | head -5 || echo "  (none)"
	@echo ""
	@echo "$(COLOR_CYAN)Total experiment data:$(COLOR_RESET)"
	@total_size=$$(du -sh experiments/ 2>/dev/null | cut -f1); \
	echo "  $$total_size"

# Count lines of code in generated projects
count-loc:
	@echo "$(COLOR_BLUE)📊 Lines of Code Summary$(COLOR_RESET)"
	@echo "────────────────────────────────────────────────────────"
	@total_loc=0; \
	for dir in $(OUT_DIR)/*/; do \
		if [ -d "$$dir" ]; then \
			app=$$(basename "$$dir"); \
			loc=$$(find "$$dir" -name '*.go' -type f -exec wc -l {} + 2>/dev/null | tail -1 | awk '{print $$1}'); \
			if [ -n "$$loc" ] && [ "$$loc" != "0" ]; then \
				printf "  $(COLOR_GREEN)%-25s$(COLOR_RESET) %s LOC\n" $$app $$loc; \
				total_loc=$$((total_loc + loc)); \
			fi; \
		fi; \
	done; \
	echo "────────────────────────────────────────────────────────"; \
	printf "  $(COLOR_CYAN)%-25s$(COLOR_RESET) %s LOC\n" "TOTAL" $$total_loc

# Show recent activity
activity:
	@echo "$(COLOR_BLUE)📈 Recent Activity$(COLOR_RESET)"
	@echo "────────────────────────────────────────────────────────"
	@echo "$(COLOR_CYAN)Recently modified experiments:$(COLOR_RESET)"
	@find $(OUT_DIR) -name 'gen_metrics.json' -type f -mtime -7 -exec dirname {} \; 2>/dev/null | \
		xargs -I {} basename {} | head -10 || echo "  (none in last 7 days)"
	@echo ""
	@echo "$(COLOR_CYAN)Recent reports:$(COLOR_RESET)"
	@ls -lht $(LOG_DIR)/*.md 2>/dev/null | head -5 || echo "  (no reports)"
	@echo ""
	@echo "$(COLOR_CYAN)Recent archives:$(COLOR_RESET)"
	@ls -lhdt $(ARCHIVE_DIR)/*/ 2>/dev/null | head -3 || echo "  (no archives)"

# Compare two experiments
compare:
	@if [ -z "$(APP1)" ] || [ -z "$(APP2)" ]; then \
		echo "$(COLOR_RED)❌ Please specify APP1=<name> APP2=<name>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_BLUE)🔍 Comparing $(APP1) vs $(APP2)$(COLOR_RESET)"
	@echo "────────────────────────────────────────────────────────"
	@metric1=$(OUT_DIR)/$(APP1)/gen_metrics.json; \
	metric2=$(OUT_DIR)/$(APP2)/gen_metrics.json; \
	if [ ! -f "$$metric1" ]; then \
		echo "$(COLOR_RED)❌ Metrics not found for $(APP1)$(COLOR_RESET)"; \
		exit 1; \
	fi; \
	if [ ! -f "$$metric2" ]; then \
		echo "$(COLOR_RED)❌ Metrics not found for $(APP2)$(COLOR_RESET)"; \
		exit 1; \
	fi; \
	echo "$(COLOR_CYAN)$(APP1):$(COLOR_RESET)"; \
	cat "$$metric1" | grep -E "duration|repair|success" | sed 's/^/  /'; \
	echo ""; \
	echo "$(COLOR_CYAN)$(APP2):$(COLOR_RESET)"; \
	cat "$$metric2" | grep -E "duration|repair|success" | sed 's/^/  /'

# Show configuration
config:
	@echo "$(COLOR_BLUE)⚙️  Current Configuration$(COLOR_RESET)"
	@echo "────────────────────────────────────────────────────────"
	@echo "$(COLOR_CYAN)Mode:$(COLOR_RESET)           $(MODE)"
	@echo "$(COLOR_CYAN)Input Dir:$(COLOR_RESET)      $(INPUT_DIR)"
	@echo "$(COLOR_CYAN)Output Dir:$(COLOR_RESET)     $(OUT_DIR)"
	@echo "$(COLOR_CYAN)Log Dir:$(COLOR_RESET)        $(LOG_DIR)"
	@echo "$(COLOR_CYAN)Archive Dir:$(COLOR_RESET)    $(ARCHIVE_DIR)"
	@echo ""
	@echo "$(COLOR_CYAN)Environment:$(COLOR_RESET)"
# =====================================================
# 🎯 Phony Targets
# =====================================================

.PHONY: help generate validate experiment all-experiments quick-test \
        report reports-all report-comparative report-statistics report-failures report-latex \
        academic-package clean clean-safe clean-code clean-logs clean-archive clean-all \
        clean-dry-run clean-force archive archive-metrics backup list-archives restore-latest \
        list stats status verify-env watch disk-usage count-loc activity compare config