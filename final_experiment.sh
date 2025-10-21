# Script for 5 runs
#!/bin/bash
for mode in rules ml hybrid; do
  make multi-run MODE=$mode RUNS=5
done
make reports-all