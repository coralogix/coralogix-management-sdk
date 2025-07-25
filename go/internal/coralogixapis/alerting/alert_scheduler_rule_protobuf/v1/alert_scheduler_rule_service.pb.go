// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_rule_service.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_rawDesc = "" +
	"\n" +
	"^com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_rule_service.proto\x12;com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1\x1aYcom/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_request.proto\x1aZcom/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_response.proto\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto2\xe5\x17\n" +
	"\x19AlertSchedulerRuleService\x12\xa0\x03\n" +
	"\x15GetAlertSchedulerRule\x12Y.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleRequest\x1aZ.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleResponse\"\xcf\x01\x92A\x90\x01\n" +
	"\x1cAlert Scheduler Rule service\x12\x1bGet an alert scheduler ruleJ\x14\n" +
	"\x03400\x12\r\n" +
	"\vBad RequestJ\x1d\n" +
	"\x03401\x12\x16\n" +
	"\x14Unauthorized requestJ\x1e\n" +
	"\x03500\x12\x17\n" +
	"\x15Internal server error\x82\xd3\xe4\x93\x025\x123/v1/alert-scheduler-rules/{alert_scheduler_rule_id}\x12\x95\x03\n" +
	"\x18CreateAlertSchedulerRule\x12\\.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleRequest\x1a].com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleResponse\"\xbb\x01\x92A\x93\x01\n" +
	"\x1cAlert Scheduler Rule service\x12\x1eCreate an alert scheduler ruleJ\x14\n" +
	"\x03400\x12\r\n" +
	"\vBad RequestJ\x1d\n" +
	"\x03401\x12\x16\n" +
	"\x14Unauthorized requestJ\x1e\n" +
	"\x03500\x12\x17\n" +
	"\x15Internal server error\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/v1/alert-scheduler-rules\x12\x95\x03\n" +
	"\x18UpdateAlertSchedulerRule\x12\\.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleRequest\x1a].com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleResponse\"\xbb\x01\x92A\x93\x01\n" +
	"\x1cAlert Scheduler Rule service\x12\x1eUpdate an alert scheduler ruleJ\x14\n" +
	"\x03400\x12\r\n" +
	"\vBad RequestJ\x1d\n" +
	"\x03401\x12\x16\n" +
	"\x14Unauthorized requestJ\x1e\n" +
	"\x03500\x12\x17\n" +
	"\x15Internal server error\x82\xd3\xe4\x93\x02\x1e:\x01*\x1a\x19/v1/alert-scheduler-rules\x12\xac\x03\n" +
	"\x18DeleteAlertSchedulerRule\x12\\.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleRequest\x1a].com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleResponse\"\xd2\x01\x92A\x93\x01\n" +
	"\x1cAlert Scheduler Rule service\x12\x1eDelete an alert scheduler ruleJ\x14\n" +
	"\x03400\x12\r\n" +
	"\vBad RequestJ\x1d\n" +
	"\x03401\x12\x16\n" +
	"\x14Unauthorized requestJ\x1e\n" +
	"\x03500\x12\x17\n" +
	"\x15Internal server error\x82\xd3\xe4\x93\x025*3/v1/alert-scheduler-rules/{alert_scheduler_rule_id}\x12\x9e\x03\n" +
	"\x19GetBulkAlertSchedulerRule\x12].com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleRequest\x1a^.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleResponse\"\xc1\x01\x92A\x97\x01\n" +
	"\x1cAlert Scheduler Rule service\x12\"Get multiple alert scheduler rulesJ\x14\n" +
	"\x03400\x12\r\n" +
	"\vBad RequestJ\x1d\n" +
	"\x03401\x12\x16\n" +
	"\x14Unauthorized requestJ\x1e\n" +
	"\x03500\x12\x17\n" +
	"\x15Internal server error\x82\xd3\xe4\x93\x02 \x12\x1e/v1/alert-scheduler-rules/bulk\x12\xad\x03\n" +
	"\x1cCreateBulkAlertSchedulerRule\x12`.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleRequest\x1aa.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleResponse\"\xc7\x01\x92A\x9a\x01\n" +
	"\x1cAlert Scheduler Rule service\x12%Create multiple alert scheduler rulesJ\x14\n" +
	"\x03400\x12\r\n" +
	"\vBad RequestJ\x1d\n" +
	"\x03401\x12\x16\n" +
	"\x14Unauthorized requestJ\x1e\n" +
	"\x03500\x12\x17\n" +
	"\x15Internal server error\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/v1/alert-scheduler-rules/bulk\x12\xad\x03\n" +
	"\x1cUpdateBulkAlertSchedulerRule\x12`.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleRequest\x1aa.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleResponse\"\xc7\x01\x92A\x9a\x01\n" +
	"\x1cAlert Scheduler Rule service\x12%Update multiple alert scheduler rulesJ\x14\n" +
	"\x03400\x12\r\n" +
	"\vBad RequestJ\x1d\n" +
	"\x03401\x12\x16\n" +
	"\x14Unauthorized requestJ\x1e\n" +
	"\x03500\x12\x17\n" +
	"\x15Internal server error\x82\xd3\xe4\x93\x02#:\x01*\x1a\x1e/v1/alert-scheduler-rules/bulk\x1aE\x92AB\n" +
	"\x1cAlert Scheduler Rule service\x12\"Manage your alert scheduler rules.b\x06proto3"

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_goTypes = []any{
	(*GetAlertSchedulerRuleRequest)(nil),         // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleRequest
	(*CreateAlertSchedulerRuleRequest)(nil),      // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleRequest
	(*UpdateAlertSchedulerRuleRequest)(nil),      // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleRequest
	(*DeleteAlertSchedulerRuleRequest)(nil),      // 3: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleRequest
	(*GetBulkAlertSchedulerRuleRequest)(nil),     // 4: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleRequest
	(*CreateBulkAlertSchedulerRuleRequest)(nil),  // 5: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleRequest
	(*UpdateBulkAlertSchedulerRuleRequest)(nil),  // 6: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleRequest
	(*GetAlertSchedulerRuleResponse)(nil),        // 7: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleResponse
	(*CreateAlertSchedulerRuleResponse)(nil),     // 8: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleResponse
	(*UpdateAlertSchedulerRuleResponse)(nil),     // 9: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleResponse
	(*DeleteAlertSchedulerRuleResponse)(nil),     // 10: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleResponse
	(*GetBulkAlertSchedulerRuleResponse)(nil),    // 11: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleResponse
	(*CreateBulkAlertSchedulerRuleResponse)(nil), // 12: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleResponse
	(*UpdateBulkAlertSchedulerRuleResponse)(nil), // 13: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleResponse
}
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_depIdxs = []int32{
	0,  // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.GetAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleRequest
	1,  // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.CreateAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleRequest
	2,  // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.UpdateAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleRequest
	3,  // 3: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.DeleteAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleRequest
	4,  // 4: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.GetBulkAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleRequest
	5,  // 5: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.CreateBulkAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleRequest
	6,  // 6: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.UpdateBulkAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleRequest
	7,  // 7: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.GetAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleResponse
	8,  // 8: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.CreateAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleResponse
	9,  // 9: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.UpdateAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleResponse
	10, // 10: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.DeleteAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleResponse
	11, // 11: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.GetBulkAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleResponse
	12, // 12: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.CreateBulkAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleResponse
	13, // 13: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.UpdateBulkAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_init()
}
func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_init() {
	if File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto != nil {
		return
	}
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_init()
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_rawDesc), len(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_depIdxs,
	}.Build()
	File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto = out.File
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_goTypes = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_depIdxs = nil
}
