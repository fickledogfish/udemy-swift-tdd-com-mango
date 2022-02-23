#!/bin/sh

for scheme in $(
	xcodebuild -"$PROJECT_TYPE" "$PROJECT_FILE" -list -json |
	jq '.project.targets[] | select(endswith("Tests"))' |
	tr -d '"'
); do
	printf '\n\n=== Running %s ===\n\n' "$scheme"
	xcodebuild test -"$PROJECT_TYPE" "$PROJECT_FILE" -scheme "$scheme"

	retcode=$?
	if [ $retcode -ne 0 ]; then
		exit $retcode
	fi
done
