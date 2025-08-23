# This was generated partly with Claude (due to my lack of Make experience)

# Usage: make 001_dna (compiles and runs with input file if cpp file is in form
# 001_dna.cpp and input file is in the format 001_dna_data.txt)
	# NOTE: it's useful to add the flag -B to force a recompile

# WARN: NOT TESTED:
# Usage: make compile-001_dna (compile only)
# Usage: make run-001_dna (run only, assumes already compiled)

CXX = g++
CXXFLAGS = -std=c++20
INPUT_SUFFIX = _data.txt

# Pattern rule: compile and run
%: %.cpp %_data.txt
	@$(CXX) $(CXXFLAGS) $< -o $@
	@cat $@$(INPUT_SUFFIX) | ./$@

# Pattern rule: compile only
compile-%: %.cpp
	@echo "Compiling $*..."
	$(CXX) $(CXXFLAGS) $< -o $*

# Pattern rule: run only (assumes executable exists)
run-%: % %_data.txt
	@echo "Running $* with input file $*$(INPUT_SUFFIX)..."
	@echo $*$(INPUT_SUFFIX) | ./$*

# Clean up executables
.PHONY: clean
clean:
	@echo "Cleaning up executables..."
	@find . -maxdepth 1 -type f -executable -name '[0-9][0-9][0-9]_*' -delete
	@echo "Done."

# List available problems
.PHONY: list
list:
	@echo "Available problems:"
	@ls -1 [0-9][0-9][0-9]_*.cpp 2>/dev/null | sed 's/\.cpp$$//' | sort || echo "No problems found"

# Prevent make from deleting intermediate files
.PRECIOUS: %

# Help target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make <problem_name>        - Compile and run (e.g., make 001_dna)"
	@echo "  make compile-<problem_name> - Compile only"
	@echo "  make run-<problem_name>     - Run only (must be compiled first)"
	@echo "  make clean                 - Remove all executables"
	@echo "  make list                  - List all available problems"
