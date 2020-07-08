scripts=(
    # "scripts/linter.sh"
)

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
RESET='\033[0m'

echo "Running go tests..."
go test -cover -v -timeout 60s `go list ./... | grep -v vendor` | sed ''/PASS/s//$(printf "${GREEN}PASS${RESET}")/'' | sed ''/FAIL/s//$(printf "${RED}FAIL${RESET}")/''
GO_TEST_EXIT_CODE=${PIPESTATUS[0]}
if [[ $GO_TEST_EXIT_CODE -ne 0 ]]; then
    exit $GO_TEST_EXIT_CODE
fi

echo "Running validation scripts..."
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
fail=0
for s in "${scripts[@]}"; do
    SCRIPT=${DIR}/${s}
    echo "RUN ${SCRIPT}"
    set +e
    ${SCRIPT}
    result=$?
    set -e
    if [[ $result  -eq 1 ]]; then
        echo -e "${RED}FAILED${RESET} ${s}"
        fail=1
    else
        echo -e "${GREEN}PASSED${RESET} ${s}"
    fi
done
exit $fail