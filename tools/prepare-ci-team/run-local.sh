#!/usr/bin/env bash
set +x
set -euo pipefail

if (( $# != 0 )); then
  printf 'Usage: run-local.sh (no arguments)\n' >&2
  exit 2
fi

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)
cd "$script_dir/../.."

docker build -t prepare-ci-team:local tools/prepare-ci-team

need_terminal=false
for name in CORALOGIX_API_KEY CORALOGIX_SEND_YOUR_DATA_API_KEY PAGERDUTY_REST_API_KEY PAGERDUTY_USER_EMAIL CI_PREP_AI_APPLICATION_PREFIX; do
  if [[ -z ${!name-} ]]; then
    need_terminal=true
  fi
done
if [[ -z ${ARCHIVE_LOGS_BUCKET+x} ]]; then
  need_terminal=true
fi
if $need_terminal; then
  if ! exec 3</dev/tty; then
    printf 'An interactive terminal is required for missing inputs.\n' >&2
    exit 1
  fi
fi

for name in CORALOGIX_API_KEY CORALOGIX_SEND_YOUR_DATA_API_KEY PAGERDUTY_REST_API_KEY PAGERDUTY_USER_EMAIL CI_PREP_AI_APPLICATION_PREFIX; do
  case "$name" in
    CORALOGIX_API_KEY) label='Coralogix management API key' ;;
    CORALOGIX_SEND_YOUR_DATA_API_KEY) label='Coralogix Send Your Data API key' ;;
    PAGERDUTY_REST_API_KEY) label='PagerDuty REST API key' ;;
    PAGERDUTY_USER_EMAIL) label='PagerDuty user email' ;;
    CI_PREP_AI_APPLICATION_PREFIX) label='AI application prefix' ;;
  esac
  if [[ -n ${!name-} ]]; then
    printf '%s: using exported value\n' "$label" >&2
    continue
  fi
  printf '%s: ' "$label" >&2
  # The loop value is the name of the environment variable to set.
  # shellcheck disable=SC2229
  if ! IFS= read -r -s "$name" <&3; then
    printf '\nInput ended before %s was set.\n' "$name" >&2
    exit 1
  fi
  printf '\n' >&2
  if [[ -z ${!name} ]]; then
    printf '%s is required.\n' "$name" >&2
    exit 1
  fi
  # shellcheck disable=SC2163
  export "$name"
done

if [[ -z ${ARCHIVE_LOGS_BUCKET+x} ]]; then
  printf 'ARCHIVE_LOGS_BUCKET (optional bucket name): ' >&2
  if ! IFS= read -r -s ARCHIVE_LOGS_BUCKET <&3; then
    printf '\nInput ended before ARCHIVE_LOGS_BUCKET was set.\n' >&2
    exit 1
  fi
  printf '\n' >&2
  export ARCHIVE_LOGS_BUCKET
else
  printf 'ARCHIVE_LOGS_BUCKET: using exported value\n' >&2
fi

docker run --rm --read-only --tmpfs /tmp:rw,nosuid,nodev,mode=1777,size=32m \
  --cap-drop=ALL --security-opt no-new-privileges \
  -e CORALOGIX_API_KEY -e CORALOGIX_SEND_YOUR_DATA_API_KEY \
  -e PAGERDUTY_REST_API_KEY -e PAGERDUTY_USER_EMAIL \
  -e CI_PREP_AI_APPLICATION_PREFIX \
  -e ARCHIVE_LOGS_BUCKET prepare-ci-team:local
