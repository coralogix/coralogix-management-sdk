// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogix/quota/v1/policies_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
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

var File_com_coralogix_quota_v1_policies_service_proto protoreflect.FileDescriptor

const file_com_coralogix_quota_v1_policies_service_proto_rawDesc = "" +
	"\n" +
	"-com/coralogix/quota/v1/policies_service.proto\x12\x16com.coralogix.quota.v1\x1a'com/coralogix/common/v1/audit_log.proto\x1a/com/coralogix/quota/v1/get_policy_request.proto\x1a0com/coralogix/quota/v1/get_policy_response.proto\x1a2com/coralogix/quota/v1/create_policy_request.proto\x1a3com/coralogix/quota/v1/create_policy_response.proto\x1a2com/coralogix/quota/v1/update_policy_request.proto\x1a3com/coralogix/quota/v1/update_policy_response.proto\x1a9com/coralogix/quota/v1/get_company_policies_request.proto\x1a:com/coralogix/quota/v1/get_company_policies_response.proto\x1a2com/coralogix/quota/v1/delete_policy_request.proto\x1a3com/coralogix/quota/v1/delete_policy_response.proto\x1a5com/coralogix/quota/v1/reorder_policies_request.proto\x1a6com/coralogix/quota/v1/reorder_policies_response.proto\x1a;com/coralogix/quota/v1/bulk_test_log_policies_request.proto\x1a<com/coralogix/quota/v1/bulk_test_log_policies_response.proto\x1a2com/coralogix/quota/v1/toggle_policy_request.proto\x1a3com/coralogix/quota/v1/toggle_policy_response.proto\x1a7com/coralogix/quota/v1/bulk_create_policy_request.proto\x1a8com/coralogix/quota/v1/bulk_create_policy_response.proto\x1a?com/coralogix/quota/v1/atomic_overwrite_policies_response.proto\x1a>com/coralogix/quota/v1/atomic_overwrite_policies_request.proto\x1a\x1cgoogle/api/annotations.proto2\xb2\x10\n" +
	"\x0fPoliciesService\x12\x8b\x01\n" +
	"\tGetPolicy\x12(.com.coralogix.quota.v1.GetPolicyRequest\x1a).com.coralogix.quota.v1.GetPolicyResponse\")¸\x02\f\n" +
	"\n" +
	"Get Policy\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/policies/{id}\x12\x95\x01\n" +
	"\fCreatePolicy\x12+.com.coralogix.quota.v1.CreatePolicyRequest\x1a,.com.coralogix.quota.v1.CreatePolicyResponse\"*¸\x02\x0f\n" +
	"\rCreate Policy\x82\xd3\xe4\x93\x02\x11:\x01*\"\f/v1/policies\x12\x95\x01\n" +
	"\fUpdatePolicy\x12+.com.coralogix.quota.v1.UpdatePolicyRequest\x1a,.com.coralogix.quota.v1.UpdatePolicyResponse\"*¸\x02\x0f\n" +
	"\rUpdate Policy\x82\xd3\xe4\x93\x02\x11:\x01*\x1a\f/v1/policies\x12\xab\x01\n" +
	"\x12GetCompanyPolicies\x121.com.coralogix.quota.v1.GetCompanyPoliciesRequest\x1a2.com.coralogix.quota.v1.GetCompanyPoliciesResponse\".¸\x02\x16\n" +
	"\x14Get Company Policies\x82\xd3\xe4\x93\x02\x0e\x12\f/v1/policies\x12\x97\x01\n" +
	"\fDeletePolicy\x12+.com.coralogix.quota.v1.DeletePolicyRequest\x1a,.com.coralogix.quota.v1.DeletePolicyResponse\",¸\x02\x0f\n" +
	"\rDelete Policy\x82\xd3\xe4\x93\x02\x13*\x11/v1/policies/{id}\x12\xa9\x01\n" +
	"\x0fReorderPolicies\x12..com.coralogix.quota.v1.ReorderPoliciesRequest\x1a/.com.coralogix.quota.v1.ReorderPoliciesResponse\"5¸\x02\x12\n" +
	"\x10Reorder Policies\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/v1/policies:reorder\x12\xbf\x01\n" +
	"\x13BulkTestLogPolicies\x122.com.coralogix.quota.v1.BulkTestLogPoliciesRequest\x1a3.com.coralogix.quota.v1.BulkTestLogPoliciesResponse\"?¸\x02\x18\n" +
	"\x16Bulk Test Log Policies\x82\xd3\xe4\x93\x02\x1d:\x01*\"\x18/v1/policies:bulkTestLog\x12\x9c\x01\n" +
	"\fTogglePolicy\x12+.com.coralogix.quota.v1.TogglePolicyRequest\x1a,.com.coralogix.quota.v1.TogglePolicyResponse\"1¸\x02\x0f\n" +
	"\rToggle Policy\x82\xd3\xe4\x93\x02\x18:\x01*\"\x13/v1/policies:toggle\x12\xc6\x01\n" +
	"\x17AtomicBatchCreatePolicy\x126.com.coralogix.quota.v1.AtomicBatchCreatePolicyRequest\x1a7.com.coralogix.quota.v1.AtomicBatchCreatePolicyResponse\":¸\x02\x14\n" +
	"\x12Bulk Create Policy\x82\xd3\xe4\x93\x02\x1c:\x01*\"\x17/v1/policies:bulkCreate\x12\x9d\x02\n" +
	"\x1aAtomicOverwriteLogPolicies\x129.com.coralogix.quota.v1.AtomicOverwriteLogPoliciesRequest\x1a:.com.coralogix.quota.v1.AtomicOverwriteLogPoliciesResponse\"\x87\x01¸\x02Q\n" +
	"ODeletes all existing log policies and creates the newly provided one atomically\x82\xd3\xe4\x93\x02,:\x01*\"'/v1/policies:atomicOverwriteLogPolicies\x12\xa2\x02\n" +
	"\x1bAtomicOverwriteSpanPolicies\x12:.com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesRequest\x1a;.com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesResponse\"\x89\x01¸\x02R\n" +
	"PDeletes all existing span policies and creates the newly provided one atomically\x82\xd3\xe4\x93\x02-:\x01*\"(/v1/policies:atomicOverwriteSpanPoliciesb\x06proto3"

var file_com_coralogix_quota_v1_policies_service_proto_goTypes = []any{
	(*GetPolicyRequest)(nil),                    // 0: com.coralogix.quota.v1.GetPolicyRequest
	(*CreatePolicyRequest)(nil),                 // 1: com.coralogix.quota.v1.CreatePolicyRequest
	(*UpdatePolicyRequest)(nil),                 // 2: com.coralogix.quota.v1.UpdatePolicyRequest
	(*GetCompanyPoliciesRequest)(nil),           // 3: com.coralogix.quota.v1.GetCompanyPoliciesRequest
	(*DeletePolicyRequest)(nil),                 // 4: com.coralogix.quota.v1.DeletePolicyRequest
	(*ReorderPoliciesRequest)(nil),              // 5: com.coralogix.quota.v1.ReorderPoliciesRequest
	(*BulkTestLogPoliciesRequest)(nil),          // 6: com.coralogix.quota.v1.BulkTestLogPoliciesRequest
	(*TogglePolicyRequest)(nil),                 // 7: com.coralogix.quota.v1.TogglePolicyRequest
	(*AtomicBatchCreatePolicyRequest)(nil),      // 8: com.coralogix.quota.v1.AtomicBatchCreatePolicyRequest
	(*AtomicOverwriteLogPoliciesRequest)(nil),   // 9: com.coralogix.quota.v1.AtomicOverwriteLogPoliciesRequest
	(*AtomicOverwriteSpanPoliciesRequest)(nil),  // 10: com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesRequest
	(*GetPolicyResponse)(nil),                   // 11: com.coralogix.quota.v1.GetPolicyResponse
	(*CreatePolicyResponse)(nil),                // 12: com.coralogix.quota.v1.CreatePolicyResponse
	(*UpdatePolicyResponse)(nil),                // 13: com.coralogix.quota.v1.UpdatePolicyResponse
	(*GetCompanyPoliciesResponse)(nil),          // 14: com.coralogix.quota.v1.GetCompanyPoliciesResponse
	(*DeletePolicyResponse)(nil),                // 15: com.coralogix.quota.v1.DeletePolicyResponse
	(*ReorderPoliciesResponse)(nil),             // 16: com.coralogix.quota.v1.ReorderPoliciesResponse
	(*BulkTestLogPoliciesResponse)(nil),         // 17: com.coralogix.quota.v1.BulkTestLogPoliciesResponse
	(*TogglePolicyResponse)(nil),                // 18: com.coralogix.quota.v1.TogglePolicyResponse
	(*AtomicBatchCreatePolicyResponse)(nil),     // 19: com.coralogix.quota.v1.AtomicBatchCreatePolicyResponse
	(*AtomicOverwriteLogPoliciesResponse)(nil),  // 20: com.coralogix.quota.v1.AtomicOverwriteLogPoliciesResponse
	(*AtomicOverwriteSpanPoliciesResponse)(nil), // 21: com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesResponse
}
var file_com_coralogix_quota_v1_policies_service_proto_depIdxs = []int32{
	0,  // 0: com.coralogix.quota.v1.PoliciesService.GetPolicy:input_type -> com.coralogix.quota.v1.GetPolicyRequest
	1,  // 1: com.coralogix.quota.v1.PoliciesService.CreatePolicy:input_type -> com.coralogix.quota.v1.CreatePolicyRequest
	2,  // 2: com.coralogix.quota.v1.PoliciesService.UpdatePolicy:input_type -> com.coralogix.quota.v1.UpdatePolicyRequest
	3,  // 3: com.coralogix.quota.v1.PoliciesService.GetCompanyPolicies:input_type -> com.coralogix.quota.v1.GetCompanyPoliciesRequest
	4,  // 4: com.coralogix.quota.v1.PoliciesService.DeletePolicy:input_type -> com.coralogix.quota.v1.DeletePolicyRequest
	5,  // 5: com.coralogix.quota.v1.PoliciesService.ReorderPolicies:input_type -> com.coralogix.quota.v1.ReorderPoliciesRequest
	6,  // 6: com.coralogix.quota.v1.PoliciesService.BulkTestLogPolicies:input_type -> com.coralogix.quota.v1.BulkTestLogPoliciesRequest
	7,  // 7: com.coralogix.quota.v1.PoliciesService.TogglePolicy:input_type -> com.coralogix.quota.v1.TogglePolicyRequest
	8,  // 8: com.coralogix.quota.v1.PoliciesService.AtomicBatchCreatePolicy:input_type -> com.coralogix.quota.v1.AtomicBatchCreatePolicyRequest
	9,  // 9: com.coralogix.quota.v1.PoliciesService.AtomicOverwriteLogPolicies:input_type -> com.coralogix.quota.v1.AtomicOverwriteLogPoliciesRequest
	10, // 10: com.coralogix.quota.v1.PoliciesService.AtomicOverwriteSpanPolicies:input_type -> com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesRequest
	11, // 11: com.coralogix.quota.v1.PoliciesService.GetPolicy:output_type -> com.coralogix.quota.v1.GetPolicyResponse
	12, // 12: com.coralogix.quota.v1.PoliciesService.CreatePolicy:output_type -> com.coralogix.quota.v1.CreatePolicyResponse
	13, // 13: com.coralogix.quota.v1.PoliciesService.UpdatePolicy:output_type -> com.coralogix.quota.v1.UpdatePolicyResponse
	14, // 14: com.coralogix.quota.v1.PoliciesService.GetCompanyPolicies:output_type -> com.coralogix.quota.v1.GetCompanyPoliciesResponse
	15, // 15: com.coralogix.quota.v1.PoliciesService.DeletePolicy:output_type -> com.coralogix.quota.v1.DeletePolicyResponse
	16, // 16: com.coralogix.quota.v1.PoliciesService.ReorderPolicies:output_type -> com.coralogix.quota.v1.ReorderPoliciesResponse
	17, // 17: com.coralogix.quota.v1.PoliciesService.BulkTestLogPolicies:output_type -> com.coralogix.quota.v1.BulkTestLogPoliciesResponse
	18, // 18: com.coralogix.quota.v1.PoliciesService.TogglePolicy:output_type -> com.coralogix.quota.v1.TogglePolicyResponse
	19, // 19: com.coralogix.quota.v1.PoliciesService.AtomicBatchCreatePolicy:output_type -> com.coralogix.quota.v1.AtomicBatchCreatePolicyResponse
	20, // 20: com.coralogix.quota.v1.PoliciesService.AtomicOverwriteLogPolicies:output_type -> com.coralogix.quota.v1.AtomicOverwriteLogPoliciesResponse
	21, // 21: com.coralogix.quota.v1.PoliciesService.AtomicOverwriteSpanPolicies:output_type -> com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_policies_service_proto_init() }
func file_com_coralogix_quota_v1_policies_service_proto_init() {
	if File_com_coralogix_quota_v1_policies_service_proto != nil {
		return
	}
	file_com_coralogix_quota_v1_get_policy_request_proto_init()
	file_com_coralogix_quota_v1_get_policy_response_proto_init()
	file_com_coralogix_quota_v1_create_policy_request_proto_init()
	file_com_coralogix_quota_v1_create_policy_response_proto_init()
	file_com_coralogix_quota_v1_update_policy_request_proto_init()
	file_com_coralogix_quota_v1_update_policy_response_proto_init()
	file_com_coralogix_quota_v1_get_company_policies_request_proto_init()
	file_com_coralogix_quota_v1_get_company_policies_response_proto_init()
	file_com_coralogix_quota_v1_delete_policy_request_proto_init()
	file_com_coralogix_quota_v1_delete_policy_response_proto_init()
	file_com_coralogix_quota_v1_reorder_policies_request_proto_init()
	file_com_coralogix_quota_v1_reorder_policies_response_proto_init()
	file_com_coralogix_quota_v1_bulk_test_log_policies_request_proto_init()
	file_com_coralogix_quota_v1_bulk_test_log_policies_response_proto_init()
	file_com_coralogix_quota_v1_toggle_policy_request_proto_init()
	file_com_coralogix_quota_v1_toggle_policy_response_proto_init()
	file_com_coralogix_quota_v1_bulk_create_policy_request_proto_init()
	file_com_coralogix_quota_v1_bulk_create_policy_response_proto_init()
	file_com_coralogix_quota_v1_atomic_overwrite_policies_response_proto_init()
	file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_quota_v1_policies_service_proto_rawDesc), len(file_com_coralogix_quota_v1_policies_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_quota_v1_policies_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_policies_service_proto_depIdxs,
	}.Build()
	File_com_coralogix_quota_v1_policies_service_proto = out.File
	file_com_coralogix_quota_v1_policies_service_proto_goTypes = nil
	file_com_coralogix_quota_v1_policies_service_proto_depIdxs = nil
}
