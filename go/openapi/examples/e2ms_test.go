package examples

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	"github.com/coralogix/coralogix-management-sdk/go/openapi/gen/events2metrics_service"
)

func TestE2MsLogsQuery(t *testing.T) {
	t.Skip("The API does not return the discriminating field for aggregations, and therefore we cannot deserialize the payload into any of the variants.")
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().WithHTTPLogging().Build()
	client := cxsdk.NewEvents2MetricsClient(cfg)

	description := "E2M Logs Query created by E2M logs query test"
	luceneQuery := "remote_addr_enriched:/.*/"
	geoLocationMetricName1 := "geo_point"
	// geoLocationMetricName2 := "geo_point2"
	// geoLocationMetricName3 := "geo_point3"
	// geoLocationMetricName4 := "geo_point4"
	// geoLocationMetricName5 := "geo_point5"
	enable := true

	createReq := events2metrics_service.Events2MetricServiceCreateE2MRequest{
		E2MCreateParamsLogsQuery: &events2metrics_service.E2MCreateParamsLogsQuery{
			Name:        "Test E2M Logs Query",
			Description: &description,
			LogsQuery: &events2metrics_service.V2LogsQuery{
				Lucene:                 &luceneQuery,
				ApplicationnameFilters: []string{"filter:startsWith:nginx"},
				SeverityFilters:        []events2metrics_service.Logs2metricsV2Severity{events2metrics_service.LOGS2METRICSV2SEVERITY_SEVERITY_DEBUG},
			},
			MetricFields: []events2metrics_service.V2MetricField{
				{
					SourceField:          "method",
					TargetBaseMetricName: "method",
					Aggregations: []events2metrics_service.V2Aggregation{
						{
							V2AggregationSamples: &events2metrics_service.V2AggregationSamples{
								AggType: events2metrics_service.AGGTYPE_AGG_TYPE_COUNT.Ptr(),
								Enabled: &enable,
								Samples: &events2metrics_service.E2MAggSamples{
									SampleType: events2metrics_service.SAMPLETYPE_SAMPLE_TYPE_MAX.Ptr().Ptr(),
								},
								TargetMetricName: &geoLocationMetricName1,
							},
						},
					},
				},
				{
					SourceField:          "location_geopoint",
					TargetBaseMetricName: "geo_point",
				},
			},
			MetricLabels: []events2metrics_service.MetricLabel{
				{
					SourceField: "Status",
					TargetLabel: "status",
				},
				{
					SourceField: "Path",
					TargetLabel: "http_referer",
				},
			},
			Type: events2metrics_service.E2MTYPE_E2_M_TYPE_LOGS2_METRICS.Ptr(),
		},
	}

	created, httpResp, err := client.
		Events2MetricServiceCreateE2M(context.Background()).
		Events2MetricServiceCreateE2MRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, created.E2m.E2MLogsQuery.Id)

}
