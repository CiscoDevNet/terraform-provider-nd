#!/usr/bin/env bash
# Run all unit tests (excluding acceptance tests prefixed with TestAcc).
#
# Usage:
#   ./run_unit_tests.sh              # run all unit tests
#   ./run_unit_tests.sh -v           # verbose output
#   ./run_unit_tests.sh ./internal/manage/...  # specific package(s)

set -uo pipefail
cd "$(git rev-parse --show-toplevel)"

# Separate flags/packages from positional args
GOTEST_FLAGS=()
PACKAGES=()
for arg in "$@"; do
    case "$arg" in
        -*)  GOTEST_FLAGS+=("$arg") ;;
        *)   PACKAGES+=("$arg") ;;
    esac
done

# Default to all packages
if [ ${#PACKAGES[@]} -eq 0 ]; then
    PACKAGES=("./...")
fi

# Skip any test function starting with TestAcc
echo "==> Running unit tests (excluding TestAcc*)"
go test "${GOTEST_FLAGS[@]+"${GOTEST_FLAGS[@]}"}" -count=1 -run '^Test($|[^A]|A[^c]|Ac[^c])' "${PACKAGES[@]}"
rc=$?

if [ $rc -eq 0 ]; then
    echo "==> All unit tests passed"
else
    echo "==> Unit tests finished with failures (exit $rc)"
fi
exit $rc
