# Prepare one EU2 CI team

Run this command before live acceptance jobs. It sends synthetic GenAI traces
for 20 application names derived from one required prefix. The names stay the
same on every run when the prefix stays the same.
It sends a separate server span through a short-lived OpenTelemetry Collector
Contrib process. The Collector generates full and compact Span Metrics. The
command then polls the AI application catalog and the APM v2 service list
every 10 seconds for up to 10 minutes. It checks that all 20 AI applications
are visible and that enough evaluation targets are free for parallel tests.
The local Collector has a separate 15-second startup limit. It stays running
during catalog polling so that pending metric exports can retry. The command
stops it after polling.
During catalog polling, stderr shows the AI application, unused evaluation
target, and APM service (V2) states after each check. These lines contain
no IDs or credential values.
On success, stdout contains fixture names. Suppress this output in public CI.
Public CI logs must not contain application names, team integration IDs, or
generated connector values.

This command passed once on an existing EU2 team. A fresh-team run still needs
to confirm the AI entry, the APM v2 service list entry, and the APM v2 Span
Metrics path before this command is used in the shared acceptance workflow.

## Inputs

Set these environment variables for the same EU2 team:

- `CORALOGIX_API_KEY`: management API key.
- `CORALOGIX_SEND_YOUR_DATA_API_KEY`: telemetry key.
- `PAGERDUTY_REST_API_KEY`: PagerDuty REST API key.
- `PAGERDUTY_USER_EMAIL`: PagerDuty actor email.
- `CI_PREP_AI_APPLICATION_PREFIX`: a stable, private name for the AI application
  family. Use 3 to 40 lowercase letters, digits, or hyphens. Start with a
  letter and end with a letter or digit. The command adds `-02` through
  `-20` for the other applications.

`ARCHIVE_LOGS_BUCKET` is an optional, non-credential input. It names an existing
logs-and-traces S3 bucket in EU2. The command uses it only if the team has no
active archive target. The bucket must already allow Coralogix archive access.
There is no separate archive-name field in the target API.

The command reads credential values only from these four variables. It does not
accept credential arguments. It does not print credentials. Use a shell or a
private environment manager to set them. Do not put values in repository files
or command arguments.

## Run in a local container

Only Docker and Bash are needed on the host. The image builds the Go command.
It copies a pinned OpenTelemetry Collector Contrib
binary from its official image. The Collector runs inside the same container
as the command. It listens on container loopback, so no host port is needed.

Run `tools/prepare-ci-team/run-local.sh` by its absolute path from any working
directory. It changes to the SDK root and builds the current image before it
asks for missing inputs. It reuses nonempty exported credential variables and
an exported `ARCHIVE_LOGS_BUCKET`, including an empty value. It prompts for
missing values through the terminal with echo off. Values entered at a prompt
exist only in the launcher process; export them in your terminal to reuse them
on later runs. The Docker `-e` flags pass variable names, not values. Do not
put values in command arguments or an env file.

The local Docker daemon can inspect the container environment while it runs.
Use a private local Docker daemon. `--rm` removes the stopped container. The
command creates its temporary Collector configuration in container memory.

## Native alternative

Install a pinned `otelcol-contrib` binary on `PATH`. The command uses the EU2
APM v2 HTTP entity list for its service check. It starts `otelcol-contrib` with a temporary configuration that
contains an environment reference, not a key value. It removes that file and
stops the process after catalog polling. Run from the SDK root:

```sh
go run ./tools/prepare-ci-team
```

On success, stdout contains one JSON object with `application`, `subsystem`,
and `apmService`. Keep this output out of public CI logs.

For a native run that needs integration IDs, set `PREPARE_CI_TEAM_OUTPUT_FILE`
to a new private file path. The command creates that file with mode `0600` and
refuses to replace an existing file. It writes `pagerDutyIntegrationId`,
`pagerDutyConnectorKey`, and `slackIntegrationId` there with the fixture names.
The connector key is a new UUID-shaped test value for one-way PagerDuty
connector CRUD. It is not a PagerDuty service routing key. Keep this file on
the runner. Do not print it or copy its values into public job outputs. The
local Docker launcher does not mount a private output file; it prints only
fixture names to the local terminal.

## Read-only lookup in acceptance jobs

After the preparation job succeeds, each acceptance job can run
`go run ./tools/prepare-ci-team lookup` from its pinned SDK checkout. This
mode needs only `CORALOGIX_API_KEY` for the same team and a new
`PREPARE_CI_TEAM_OUTPUT_FILE` path in that job's private runner directory. It
sends only GET requests. It resolves the authenticated team's ID with the
Identity API and requires active Slack and PagerDuty registrations. It writes
the team and integration IDs and a new one-way connector test value to the
private file. It does not start the Collector or create anything.

The acceptance job can read the file into its own job environment. Remove
the file after use. Do not print it or pass the IDs through public job
outputs. Existing Terraform tests use one PagerDuty variable for two different
connector types. Split those test inputs before using the generated one-way
connector value.

## Team prerequisites

Use keys for one EU2 team. In Coralogix, open **Data Flow > Integrations >
Slack**. Create the integration if needed, then select **Authorize** and
approve it in Slack. The command checks for one active Slack registration and
stops before writing if it is missing or ambiguous. This OAuth step is needed
once per new team.

AI Center reads GenAI traces from the team's S3 archive. This is the existing
logs-and-traces archive for the team, not a new bucket for this command. Check
**Data Flow > Setup Archive** for an active logs-and-traces (CX data) target.
Check **Data Flow > TCO Optimizer > Traces** to make sure the fixture trace can
reach archive storage. A trace kept only in Frequent Search does not appear in
AI Center. The command reuses an active S3 target. If none exists, it can set
`ARCHIVE_LOGS_BUCKET` as the target through the logs archive API. It refuses
to replace a different inactive target. It cannot create the AWS bucket, grant
bucket or KMS access, or change trace routing policies.

## Current limits

The command checks Slack first. Slack needs one-time OAuth authorization for
each team. If a connected Slack integration is missing, setup fails before it
creates other fixtures.

The command reuses an active account-level PagerDuty integration. If none
exists, it reads the EU2 `pagerduty-incidents` definition and checks its
version and five required fields. It sends the requested values to the
validate-only API before creation. It sends `DataProcessingConsent=true` and
`EnableBidirectional=true`. It then creates the integration and waits for an
active registration. It does not retry a create request. If the response is
uncertain, it reads registrations before failing. A repeated run reuses the
existing registration.

The backend definition confirmed the fields and version. An EU2 validate-only
probe confirmed that `ApiKey` must use `stringValue`; `apiKey.value` returned
HTTP 400. An EU2 lifecycle probe then created the integration and observed
`ACTIVE`. The full preparation command passed on an existing EU2 team. A
fresh-team run is still required before use in the shared acceptance workflow.

The selected team must archive GenAI traces to S3 for AI Center discovery.
The Collector must export compact Span Metrics for APM v2 discovery. The
command checks the APM v2 entity list. This check does not test whether the
SLO API accepts that service. The first fresh-team run must verify SLO
creation separately. The command does not inspect the S3 object or exported
metric body.
Management reads use the EU2 `/mgmt/openapi/5` base path.

## Local checks

```sh
go test ./tools/prepare-ci-team
go vet ./tools/prepare-ci-team
```

The tests use local HTTP fakes. They do not contact a team.

## References

- [AI Center setup](https://coralogix.com/docs/user-guides/ai/getting_started/)
- [Compact Span Metrics](https://coralogix.com/docs/user-guides/apm-v2/getting-started/using-compact-span-metric/)
- [Coralogix Collector configuration](https://github.com/coralogix/telemetry-shippers/blob/master/otel-ecs-fargate/config.yaml)
- [APM v2 entity list route in the Coralogix CLI](https://github.com/coralogix/cx-cli/blob/master/src/commands/service_catalog/api.rs)
- [PagerDuty integration](https://coralogix.com/docs/user-guides/cases/pagerduty-integration/)

- [AI Center archive requirement](https://coralogix.com/docs/user-guides/ai/getting_started/)
- [Connect an S3 archive](https://coralogix.com/docs/user-guides/data-flow/s3-archive/connect-s3-archive/)
- [Connect a Slack workspace](https://coralogix.com/docs/integrations/pull/slack/)
- [Get logs archive target](https://coralogix.com/docs/api-reference/v5/target-service/get-target/)
- [Set logs archive target](https://coralogix.com/docs/api-reference/v5/target-service/set-target/)
