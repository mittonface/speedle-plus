// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Principal
	Subject
	ContextRequest
	IsAllowedResponse
	AndPrincipals
	RolePolicy
	Policy
	EvaluatedCondition
	EvaluatedRolePolicy
	EvaluatedPolicy
	EvaluationDebugResponse
	AllRoleResponse
	AllPermissionResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Principal struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Idd  string `protobuf:"bytes,3,opt,name=idd" json:"idd,omitempty"`
}

func (m *Principal) Reset()                    { *m = Principal{} }
func (m *Principal) String() string            { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()               {}
func (*Principal) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Principal) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Principal) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Principal) GetIdd() string {
	if m != nil {
		return m.Idd
	}
	return ""
}

type Subject struct {
	Principals []*Principal `protobuf:"bytes,1,rep,name=principals" json:"principals,omitempty"`
	TokenType  string       `protobuf:"bytes,2,opt,name=tokenType" json:"tokenType,omitempty"`
	Token      string       `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *Subject) Reset()                    { *m = Subject{} }
func (m *Subject) String() string            { return proto.CompactTextString(m) }
func (*Subject) ProtoMessage()               {}
func (*Subject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Subject) GetPrincipals() []*Principal {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *Subject) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *Subject) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ContextRequest struct {
	Subject     *Subject          `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	ServiceName string            `protobuf:"bytes,2,opt,name=serviceName" json:"serviceName,omitempty"`
	Resource    string            `protobuf:"bytes,3,opt,name=resource" json:"resource,omitempty"`
	Action      string            `protobuf:"bytes,4,opt,name=action" json:"action,omitempty"`
	Attributes  map[string]string `protobuf:"bytes,5,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ContextRequest) Reset()                    { *m = ContextRequest{} }
func (m *ContextRequest) String() string            { return proto.CompactTextString(m) }
func (*ContextRequest) ProtoMessage()               {}
func (*ContextRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ContextRequest) GetSubject() *Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *ContextRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ContextRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *ContextRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ContextRequest) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type IsAllowedResponse struct {
	Allowed bool   `protobuf:"varint,1,opt,name=allowed" json:"allowed,omitempty"`
	Reason  int32  `protobuf:"varint,2,opt,name=reason" json:"reason,omitempty"`
	ErrMsg  string `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *IsAllowedResponse) Reset()                    { *m = IsAllowedResponse{} }
func (m *IsAllowedResponse) String() string            { return proto.CompactTextString(m) }
func (*IsAllowedResponse) ProtoMessage()               {}
func (*IsAllowedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IsAllowedResponse) GetAllowed() bool {
	if m != nil {
		return m.Allowed
	}
	return false
}

func (m *IsAllowedResponse) GetReason() int32 {
	if m != nil {
		return m.Reason
	}
	return 0
}

func (m *IsAllowedResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type AndPrincipals struct {
	Principals []string `protobuf:"bytes,1,rep,name=principals" json:"principals,omitempty"`
}

func (m *AndPrincipals) Reset()                    { *m = AndPrincipals{} }
func (m *AndPrincipals) String() string            { return proto.CompactTextString(m) }
func (*AndPrincipals) ProtoMessage()               {}
func (*AndPrincipals) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AndPrincipals) GetPrincipals() []string {
	if m != nil {
		return m.Principals
	}
	return nil
}

type RolePolicy struct {
	ID                  string   `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name                string   `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Effect              string   `protobuf:"bytes,3,opt,name=Effect" json:"Effect,omitempty"`
	Roles               []string `protobuf:"bytes,4,rep,name=Roles" json:"Roles,omitempty"`
	Principals          []string `protobuf:"bytes,5,rep,name=Principals" json:"Principals,omitempty"`
	Resources           []string `protobuf:"bytes,6,rep,name=Resources" json:"Resources,omitempty"`
	ResourceExpressions []string `protobuf:"bytes,7,rep,name=ResourceExpressions" json:"ResourceExpressions,omitempty"`
	Condition           string   `protobuf:"bytes,8,opt,name=Condition" json:"Condition,omitempty"`
}

func (m *RolePolicy) Reset()                    { *m = RolePolicy{} }
func (m *RolePolicy) String() string            { return proto.CompactTextString(m) }
func (*RolePolicy) ProtoMessage()               {}
func (*RolePolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RolePolicy) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *RolePolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RolePolicy) GetEffect() string {
	if m != nil {
		return m.Effect
	}
	return ""
}

func (m *RolePolicy) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *RolePolicy) GetPrincipals() []string {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *RolePolicy) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *RolePolicy) GetResourceExpressions() []string {
	if m != nil {
		return m.ResourceExpressions
	}
	return nil
}

func (m *RolePolicy) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

type Policy struct {
	ID          string               `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Effect      string               `protobuf:"bytes,3,opt,name=Effect" json:"Effect,omitempty"`
	Permissions []*Policy_Permission `protobuf:"bytes,4,rep,name=permissions" json:"permissions,omitempty"`
	Principals  []*AndPrincipals     `protobuf:"bytes,5,rep,name=Principals" json:"Principals,omitempty"`
	Condition   string               `protobuf:"bytes,6,opt,name=Condition" json:"Condition,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Policy) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetEffect() string {
	if m != nil {
		return m.Effect
	}
	return ""
}

func (m *Policy) GetPermissions() []*Policy_Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Policy) GetPrincipals() []*AndPrincipals {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *Policy) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

type Policy_Permission struct {
	Resource           string   `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	ResourceExpression string   `protobuf:"bytes,2,opt,name=resourceExpression" json:"resourceExpression,omitempty"`
	Actions            []string `protobuf:"bytes,3,rep,name=actions" json:"actions,omitempty"`
}

func (m *Policy_Permission) Reset()                    { *m = Policy_Permission{} }
func (m *Policy_Permission) String() string            { return proto.CompactTextString(m) }
func (*Policy_Permission) ProtoMessage()               {}
func (*Policy_Permission) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *Policy_Permission) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *Policy_Permission) GetResourceExpression() string {
	if m != nil {
		return m.ResourceExpression
	}
	return ""
}

func (m *Policy_Permission) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

type EvaluatedCondition struct {
	ConditionExpression string `protobuf:"bytes,1,opt,name=ConditionExpression" json:"ConditionExpression,omitempty"`
	EvaluationResult    string `protobuf:"bytes,2,opt,name=EvaluationResult" json:"EvaluationResult,omitempty"`
}

func (m *EvaluatedCondition) Reset()                    { *m = EvaluatedCondition{} }
func (m *EvaluatedCondition) String() string            { return proto.CompactTextString(m) }
func (*EvaluatedCondition) ProtoMessage()               {}
func (*EvaluatedCondition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *EvaluatedCondition) GetConditionExpression() string {
	if m != nil {
		return m.ConditionExpression
	}
	return ""
}

func (m *EvaluatedCondition) GetEvaluationResult() string {
	if m != nil {
		return m.EvaluationResult
	}
	return ""
}

type EvaluatedRolePolicy struct {
	Status              string              `protobuf:"bytes,1,opt,name=Status" json:"Status,omitempty"`
	ID                  string              `protobuf:"bytes,2,opt,name=ID" json:"ID,omitempty"`
	Name                string              `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Effect              string              `protobuf:"bytes,4,opt,name=Effect" json:"Effect,omitempty"`
	Roles               []string            `protobuf:"bytes,5,rep,name=Roles" json:"Roles,omitempty"`
	Principals          []string            `protobuf:"bytes,6,rep,name=Principals" json:"Principals,omitempty"`
	Resources           []string            `protobuf:"bytes,7,rep,name=Resources" json:"Resources,omitempty"`
	ResourceExpressions []string            `protobuf:"bytes,8,rep,name=ResourceExpressions" json:"ResourceExpressions,omitempty"`
	Condition           *EvaluatedCondition `protobuf:"bytes,9,opt,name=Condition" json:"Condition,omitempty"`
}

func (m *EvaluatedRolePolicy) Reset()                    { *m = EvaluatedRolePolicy{} }
func (m *EvaluatedRolePolicy) String() string            { return proto.CompactTextString(m) }
func (*EvaluatedRolePolicy) ProtoMessage()               {}
func (*EvaluatedRolePolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *EvaluatedRolePolicy) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *EvaluatedRolePolicy) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *EvaluatedRolePolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EvaluatedRolePolicy) GetEffect() string {
	if m != nil {
		return m.Effect
	}
	return ""
}

func (m *EvaluatedRolePolicy) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *EvaluatedRolePolicy) GetPrincipals() []string {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *EvaluatedRolePolicy) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *EvaluatedRolePolicy) GetResourceExpressions() []string {
	if m != nil {
		return m.ResourceExpressions
	}
	return nil
}

func (m *EvaluatedRolePolicy) GetCondition() *EvaluatedCondition {
	if m != nil {
		return m.Condition
	}
	return nil
}

type EvaluatedPolicy struct {
	Status      string                        `protobuf:"bytes,1,opt,name=Status" json:"Status,omitempty"`
	ID          string                        `protobuf:"bytes,2,opt,name=ID" json:"ID,omitempty"`
	Name        string                        `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Effect      string                        `protobuf:"bytes,4,opt,name=Effect" json:"Effect,omitempty"`
	Permissions []*EvaluatedPolicy_Permission `protobuf:"bytes,5,rep,name=permissions" json:"permissions,omitempty"`
	Principals  []string                      `protobuf:"bytes,6,rep,name=Principals" json:"Principals,omitempty"`
	Condition   *EvaluatedCondition           `protobuf:"bytes,7,opt,name=Condition" json:"Condition,omitempty"`
}

func (m *EvaluatedPolicy) Reset()                    { *m = EvaluatedPolicy{} }
func (m *EvaluatedPolicy) String() string            { return proto.CompactTextString(m) }
func (*EvaluatedPolicy) ProtoMessage()               {}
func (*EvaluatedPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *EvaluatedPolicy) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *EvaluatedPolicy) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *EvaluatedPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EvaluatedPolicy) GetEffect() string {
	if m != nil {
		return m.Effect
	}
	return ""
}

func (m *EvaluatedPolicy) GetPermissions() []*EvaluatedPolicy_Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *EvaluatedPolicy) GetPrincipals() []string {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *EvaluatedPolicy) GetCondition() *EvaluatedCondition {
	if m != nil {
		return m.Condition
	}
	return nil
}

type EvaluatedPolicy_Permission struct {
	Resource           string   `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	ResourceExpression string   `protobuf:"bytes,2,opt,name=resourceExpression" json:"resourceExpression,omitempty"`
	Actions            []string `protobuf:"bytes,3,rep,name=actions" json:"actions,omitempty"`
}

func (m *EvaluatedPolicy_Permission) Reset()                    { *m = EvaluatedPolicy_Permission{} }
func (m *EvaluatedPolicy_Permission) String() string            { return proto.CompactTextString(m) }
func (*EvaluatedPolicy_Permission) ProtoMessage()               {}
func (*EvaluatedPolicy_Permission) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

func (m *EvaluatedPolicy_Permission) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *EvaluatedPolicy_Permission) GetResourceExpression() string {
	if m != nil {
		return m.ResourceExpression
	}
	return ""
}

func (m *EvaluatedPolicy_Permission) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

type EvaluationDebugResponse struct {
	Allowed        bool                   `protobuf:"varint,1,opt,name=allowed" json:"allowed,omitempty"`
	Reason         string                 `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	RequestContext *ContextRequest        `protobuf:"bytes,3,opt,name=requestContext" json:"requestContext,omitempty"`
	GrantedRoles   []string               `protobuf:"bytes,4,rep,name=grantedRoles" json:"grantedRoles,omitempty"`
	RolePolicies   []*EvaluatedRolePolicy `protobuf:"bytes,5,rep,name=rolePolicies" json:"rolePolicies,omitempty"`
	Policies       []*EvaluatedPolicy     `protobuf:"bytes,6,rep,name=policies" json:"policies,omitempty"`
}

func (m *EvaluationDebugResponse) Reset()                    { *m = EvaluationDebugResponse{} }
func (m *EvaluationDebugResponse) String() string            { return proto.CompactTextString(m) }
func (*EvaluationDebugResponse) ProtoMessage()               {}
func (*EvaluationDebugResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *EvaluationDebugResponse) GetAllowed() bool {
	if m != nil {
		return m.Allowed
	}
	return false
}

func (m *EvaluationDebugResponse) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *EvaluationDebugResponse) GetRequestContext() *ContextRequest {
	if m != nil {
		return m.RequestContext
	}
	return nil
}

func (m *EvaluationDebugResponse) GetGrantedRoles() []string {
	if m != nil {
		return m.GrantedRoles
	}
	return nil
}

func (m *EvaluationDebugResponse) GetRolePolicies() []*EvaluatedRolePolicy {
	if m != nil {
		return m.RolePolicies
	}
	return nil
}

func (m *EvaluationDebugResponse) GetPolicies() []*EvaluatedPolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type AllRoleResponse struct {
	Roles []string `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
}

func (m *AllRoleResponse) Reset()                    { *m = AllRoleResponse{} }
func (m *AllRoleResponse) String() string            { return proto.CompactTextString(m) }
func (*AllRoleResponse) ProtoMessage()               {}
func (*AllRoleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AllRoleResponse) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type AllPermissionResponse struct {
	Permissions []*AllPermissionResponse_Permission `protobuf:"bytes,1,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *AllPermissionResponse) Reset()                    { *m = AllPermissionResponse{} }
func (m *AllPermissionResponse) String() string            { return proto.CompactTextString(m) }
func (*AllPermissionResponse) ProtoMessage()               {}
func (*AllPermissionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *AllPermissionResponse) GetPermissions() []*AllPermissionResponse_Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type AllPermissionResponse_Permission struct {
	Resource string   `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	Actions  []string `protobuf:"bytes,2,rep,name=actions" json:"actions,omitempty"`
}

func (m *AllPermissionResponse_Permission) Reset()         { *m = AllPermissionResponse_Permission{} }
func (m *AllPermissionResponse_Permission) String() string { return proto.CompactTextString(m) }
func (*AllPermissionResponse_Permission) ProtoMessage()    {}
func (*AllPermissionResponse_Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{12, 0}
}

func (m *AllPermissionResponse_Permission) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *AllPermissionResponse_Permission) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*Principal)(nil), "pb.Principal")
	proto.RegisterType((*Subject)(nil), "pb.Subject")
	proto.RegisterType((*ContextRequest)(nil), "pb.ContextRequest")
	proto.RegisterType((*IsAllowedResponse)(nil), "pb.IsAllowedResponse")
	proto.RegisterType((*AndPrincipals)(nil), "pb.AndPrincipals")
	proto.RegisterType((*RolePolicy)(nil), "pb.RolePolicy")
	proto.RegisterType((*Policy)(nil), "pb.Policy")
	proto.RegisterType((*Policy_Permission)(nil), "pb.Policy.Permission")
	proto.RegisterType((*EvaluatedCondition)(nil), "pb.EvaluatedCondition")
	proto.RegisterType((*EvaluatedRolePolicy)(nil), "pb.EvaluatedRolePolicy")
	proto.RegisterType((*EvaluatedPolicy)(nil), "pb.EvaluatedPolicy")
	proto.RegisterType((*EvaluatedPolicy_Permission)(nil), "pb.EvaluatedPolicy.Permission")
	proto.RegisterType((*EvaluationDebugResponse)(nil), "pb.EvaluationDebugResponse")
	proto.RegisterType((*AllRoleResponse)(nil), "pb.AllRoleResponse")
	proto.RegisterType((*AllPermissionResponse)(nil), "pb.AllPermissionResponse")
	proto.RegisterType((*AllPermissionResponse_Permission)(nil), "pb.AllPermissionResponse.Permission")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Evaluator service

type EvaluatorClient interface {
	IsAllowed(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*IsAllowedResponse, error)
	GetAllGrantedRoles(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*AllRoleResponse, error)
	GetAllPermissions(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*AllPermissionResponse, error)
	Discover(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*IsAllowedResponse, error)
	Diagnose(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*EvaluationDebugResponse, error)
}

type evaluatorClient struct {
	cc *grpc.ClientConn
}

func NewEvaluatorClient(cc *grpc.ClientConn) EvaluatorClient {
	return &evaluatorClient{cc}
}

func (c *evaluatorClient) IsAllowed(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*IsAllowedResponse, error) {
	out := new(IsAllowedResponse)
	err := grpc.Invoke(ctx, "/pb.Evaluator/IsAllowed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatorClient) GetAllGrantedRoles(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*AllRoleResponse, error) {
	out := new(AllRoleResponse)
	err := grpc.Invoke(ctx, "/pb.Evaluator/GetAllGrantedRoles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatorClient) GetAllPermissions(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*AllPermissionResponse, error) {
	out := new(AllPermissionResponse)
	err := grpc.Invoke(ctx, "/pb.Evaluator/GetAllPermissions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatorClient) Discover(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*IsAllowedResponse, error) {
	out := new(IsAllowedResponse)
	err := grpc.Invoke(ctx, "/pb.Evaluator/Discover", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatorClient) Diagnose(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*EvaluationDebugResponse, error) {
	out := new(EvaluationDebugResponse)
	err := grpc.Invoke(ctx, "/pb.Evaluator/Diagnose", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Evaluator service

type EvaluatorServer interface {
	IsAllowed(context.Context, *ContextRequest) (*IsAllowedResponse, error)
	GetAllGrantedRoles(context.Context, *ContextRequest) (*AllRoleResponse, error)
	GetAllPermissions(context.Context, *ContextRequest) (*AllPermissionResponse, error)
	Discover(context.Context, *ContextRequest) (*IsAllowedResponse, error)
	Diagnose(context.Context, *ContextRequest) (*EvaluationDebugResponse, error)
}

func RegisterEvaluatorServer(s *grpc.Server, srv EvaluatorServer) {
	s.RegisterService(&_Evaluator_serviceDesc, srv)
}

func _Evaluator_IsAllowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatorServer).IsAllowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Evaluator/IsAllowed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatorServer).IsAllowed(ctx, req.(*ContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evaluator_GetAllGrantedRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatorServer).GetAllGrantedRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Evaluator/GetAllGrantedRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatorServer).GetAllGrantedRoles(ctx, req.(*ContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evaluator_GetAllPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatorServer).GetAllPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Evaluator/GetAllPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatorServer).GetAllPermissions(ctx, req.(*ContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evaluator_Discover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatorServer).Discover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Evaluator/Discover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatorServer).Discover(ctx, req.(*ContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evaluator_Diagnose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatorServer).Diagnose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Evaluator/Diagnose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatorServer).Diagnose(ctx, req.(*ContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Evaluator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Evaluator",
	HandlerType: (*EvaluatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAllowed",
			Handler:    _Evaluator_IsAllowed_Handler,
		},
		{
			MethodName: "GetAllGrantedRoles",
			Handler:    _Evaluator_GetAllGrantedRoles_Handler,
		},
		{
			MethodName: "GetAllPermissions",
			Handler:    _Evaluator_GetAllPermissions_Handler,
		},
		{
			MethodName: "Discover",
			Handler:    _Evaluator_Discover_Handler,
		},
		{
			MethodName: "Diagnose",
			Handler:    _Evaluator_Diagnose_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 904 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x4e, 0x9c, 0xc4, 0x89, 0x4f, 0xba, 0xbb, 0xdd, 0x49, 0xbb, 0x35, 0x01, 0x55, 0xab, 0x11,
	0x88, 0x0a, 0x89, 0x14, 0x02, 0x52, 0xab, 0xa2, 0x0a, 0xd2, 0x26, 0x54, 0x7b, 0x01, 0x8a, 0x5c,
	0x6e, 0xb9, 0x70, 0x92, 0x69, 0x64, 0xea, 0x7a, 0xcc, 0xcc, 0x78, 0x69, 0xde, 0x82, 0x5b, 0xae,
	0x11, 0x2f, 0x82, 0x78, 0x1c, 0x6e, 0x78, 0x03, 0x34, 0x3f, 0xb6, 0xc7, 0xb1, 0xc3, 0xee, 0x4a,
	0x20, 0xee, 0x7c, 0xce, 0x9c, 0x39, 0x7f, 0xdf, 0x77, 0xe6, 0x18, 0x8e, 0x38, 0x61, 0x97, 0xd1,
	0x9a, 0x4c, 0x52, 0x46, 0x05, 0x45, 0x4e, 0xba, 0xc2, 0x0b, 0xf0, 0x96, 0x2c, 0x4a, 0xd6, 0x51,
	0x1a, 0xc6, 0x08, 0x41, 0x57, 0xec, 0x52, 0xe2, 0xb7, 0xcf, 0xdb, 0x0f, 0xbc, 0x40, 0x7d, 0x4b,
	0x5d, 0x12, 0xbe, 0x21, 0xbe, 0xa3, 0x75, 0xf2, 0x1b, 0xdd, 0x86, 0x4e, 0xb4, 0xd9, 0xf8, 0x1d,
	0xa5, 0x92, 0x9f, 0x38, 0x86, 0xfe, 0xcb, 0x6c, 0xf5, 0x03, 0x59, 0x0b, 0xf4, 0x31, 0x40, 0x9a,
	0x7b, 0xe4, 0x7e, 0xfb, 0xbc, 0xf3, 0x60, 0x38, 0x3d, 0x9a, 0xa4, 0xab, 0x49, 0x11, 0x27, 0xb0,
	0x0c, 0xd0, 0x7b, 0xe0, 0x09, 0xfa, 0x9a, 0x24, 0xdf, 0xc9, 0xc0, 0x3a, 0x48, 0xa9, 0x40, 0x77,
	0xa0, 0xa7, 0x04, 0x13, 0x4b, 0x0b, 0xf8, 0x67, 0x07, 0x8e, 0x9f, 0xd3, 0x44, 0x90, 0xb7, 0x22,
	0x20, 0x3f, 0x66, 0x84, 0x0b, 0xf4, 0x01, 0xf4, 0xb9, 0x4e, 0x40, 0x65, 0x3f, 0x9c, 0x0e, 0x65,
	0x48, 0x93, 0x53, 0x90, 0x9f, 0xa1, 0x73, 0x18, 0x9a, 0x1e, 0x7c, 0x5b, 0x16, 0x65, 0xab, 0xd0,
	0x18, 0x06, 0x8c, 0x70, 0x9a, 0xb1, 0x35, 0x31, 0x41, 0x0b, 0x19, 0x9d, 0x81, 0x1b, 0xae, 0x45,
	0x44, 0x13, 0xbf, 0xab, 0x4e, 0x8c, 0x84, 0x9e, 0x01, 0x84, 0x42, 0xb0, 0x68, 0x95, 0x09, 0xc2,
	0xfd, 0x9e, 0x2a, 0x19, 0xcb, 0xf8, 0xd5, 0x24, 0x27, 0xb3, 0xc2, 0x68, 0x91, 0x08, 0xb6, 0x0b,
	0xac, 0x5b, 0xe3, 0xa7, 0x70, 0xb2, 0x77, 0x2c, 0xdb, 0xfc, 0x9a, 0xec, 0x0c, 0x1a, 0xf2, 0x53,
	0xb6, 0xe3, 0x32, 0x8c, 0xb3, 0x3c, 0x71, 0x2d, 0x3c, 0x71, 0x1e, 0xb7, 0xf1, 0xf7, 0x70, 0x7a,
	0xc1, 0x67, 0x71, 0x4c, 0x7f, 0x22, 0x9b, 0x80, 0xf0, 0x94, 0x26, 0x9c, 0x20, 0x1f, 0xfa, 0xa1,
	0x56, 0x29, 0x27, 0x83, 0x20, 0x17, 0x65, 0x25, 0x8c, 0x84, 0x9c, 0x26, 0xca, 0x53, 0x2f, 0x30,
	0x92, 0xd4, 0x13, 0xc6, 0xbe, 0xe1, 0x5b, 0x53, 0xbb, 0x91, 0xf0, 0x43, 0x38, 0x9a, 0x25, 0x9b,
	0x65, 0x09, 0xdb, 0xfd, 0x1a, 0xca, 0x9e, 0x0d, 0x2b, 0xfe, 0xb3, 0x0d, 0x10, 0xd0, 0x98, 0x2c,
	0x69, 0x1c, 0xad, 0x77, 0xe8, 0x18, 0x9c, 0x8b, 0xb9, 0xa9, 0xc4, 0xb9, 0x98, 0x4b, 0x56, 0x59,
	0x00, 0xa8, 0x6f, 0x19, 0x7b, 0xf1, 0xea, 0x95, 0x44, 0xd0, 0xc4, 0xd6, 0x92, 0x2c, 0x5a, 0x7a,
	0xe2, 0x7e, 0x57, 0x45, 0xd1, 0x82, 0x4c, 0xa0, 0x4c, 0x47, 0xf5, 0xdc, 0x0b, 0x2c, 0x8d, 0xe4,
	0x55, 0x60, 0x70, 0xe3, 0xbe, 0xab, 0x8e, 0x4b, 0x05, 0xfa, 0x04, 0x46, 0xb9, 0xb0, 0x78, 0x9b,
	0x32, 0xc2, 0x79, 0x44, 0x13, 0xee, 0xf7, 0x95, 0x5d, 0xd3, 0x91, 0xf4, 0xf7, 0x9c, 0x26, 0x9b,
	0x48, 0xc1, 0x3f, 0xd0, 0x3c, 0x2d, 0x14, 0xf8, 0x77, 0x07, 0xdc, 0x7f, 0xa1, 0xd4, 0x47, 0x30,
	0x4c, 0x09, 0x7b, 0x13, 0x99, 0x74, 0xba, 0x8a, 0x49, 0x77, 0xd5, 0xf0, 0x28, 0xe7, 0x93, 0x65,
	0x71, 0x1a, 0xd8, 0x96, 0xe8, 0xd3, 0x5a, 0x37, 0x86, 0xd3, 0x53, 0x79, 0xaf, 0x82, 0xda, 0x7e,
	0x83, 0xca, 0x82, 0xdc, 0xbd, 0x82, 0xc6, 0x0c, 0xa0, 0x8c, 0x55, 0x19, 0x8a, 0xf6, 0xde, 0x50,
	0x4c, 0x00, 0xb1, 0x5a, 0xbf, 0x4c, 0xb5, 0x0d, 0x27, 0x8a, 0x94, 0x6a, 0x6c, 0xb8, 0xdf, 0x51,
	0xed, 0xce, 0x45, 0xcc, 0x00, 0x2d, 0x24, 0xa3, 0x43, 0x41, 0x36, 0x45, 0x26, 0x12, 0xaa, 0x42,
	0xb0, 0x02, 0xe8, 0x34, 0x9a, 0x8e, 0xd0, 0x47, 0x70, 0xdb, 0xf8, 0x91, 0x7d, 0x22, 0x3c, 0x8b,
	0x85, 0xc9, 0xa7, 0xa6, 0xc7, 0xbf, 0x39, 0x30, 0x2a, 0x82, 0x5a, 0x84, 0x3d, 0x03, 0xf7, 0xa5,
	0x08, 0x45, 0xc6, 0x4d, 0x20, 0x23, 0x19, 0x74, 0x9d, 0x1a, 0xba, 0x9d, 0x46, 0x74, 0xbb, 0xcd,
	0x44, 0xee, 0x1d, 0x26, 0xb2, 0xfb, 0xcf, 0x44, 0xee, 0x5f, 0x93, 0xc8, 0x83, 0xc3, 0x44, 0xfe,
	0xdc, 0xc6, 0xdd, 0x53, 0x6f, 0xe5, 0x99, 0x64, 0x4a, 0xbd, 0xf5, 0x36, 0xc1, 0xff, 0x72, 0xe0,
	0xa4, 0xb0, 0xf8, 0x0f, 0x7b, 0xf4, 0x55, 0x75, 0x02, 0x34, 0x93, 0xef, 0x57, 0xf2, 0xbb, 0x62,
	0x14, 0xae, 0xea, 0x67, 0xa5, 0xfe, 0xfe, 0x35, 0xeb, 0xff, 0x5f, 0xe6, 0xe1, 0x17, 0x07, 0xee,
	0x95, 0x84, 0x9d, 0x93, 0x55, 0xb6, 0xbd, 0xf1, 0xd3, 0xee, 0x15, 0x4f, 0xfb, 0x13, 0x38, 0x66,
	0x7a, 0x0f, 0x99, 0xad, 0xa4, 0xf0, 0x18, 0x4e, 0x51, 0x7d, 0x51, 0x05, 0x7b, 0x96, 0x08, 0xc3,
	0xad, 0x2d, 0x0b, 0x13, 0x33, 0x22, 0xf9, 0x4b, 0x5c, 0xd1, 0xa1, 0x2f, 0xe0, 0x16, 0xcb, 0xe7,
	0x27, 0x2a, 0xd6, 0xe0, 0xbd, 0x4a, 0x6b, 0xcb, 0x01, 0x0b, 0x2a, 0xc6, 0xe8, 0x21, 0x0c, 0xd2,
	0xfc, 0xa2, 0xab, 0x2e, 0x8e, 0x1a, 0x30, 0x0f, 0x0a, 0x23, 0xfc, 0x21, 0x9c, 0xcc, 0xe2, 0x58,
	0xfa, 0x2b, 0x5a, 0x72, 0x07, 0x7a, 0x4c, 0x65, 0xa7, 0xb7, 0x91, 0x16, 0xf0, 0xaf, 0x6d, 0xb8,
	0x3b, 0x8b, 0x63, 0x8b, 0x2d, 0xb9, 0xfd, 0xd7, 0x55, 0xaa, 0xe9, 0x3f, 0x95, 0xf7, 0xd5, 0xa3,
	0xd9, 0x64, 0x7f, 0x88, 0x70, 0xe3, 0x67, 0xd7, 0xa6, 0x86, 0x05, 0xb5, 0x53, 0x81, 0x7a, 0xfa,
	0x87, 0x03, 0x9e, 0x29, 0x96, 0x32, 0xf4, 0x18, 0xbc, 0x62, 0x99, 0xa3, 0x06, 0x7c, 0xc6, 0x6a,
	0x25, 0xd4, 0xf6, 0x3d, 0x6e, 0xa1, 0x2f, 0x01, 0xbd, 0x20, 0x62, 0x16, 0xc7, 0x2f, 0x6c, 0x68,
	0x9a, 0x5c, 0x8c, 0x4c, 0xa1, 0x76, 0x0b, 0x71, 0x0b, 0xcd, 0xe1, 0x54, 0x3b, 0x58, 0x5a, 0x23,
	0xd5, 0x74, 0xff, 0x9d, 0x83, 0x8d, 0xc2, 0x2d, 0xf4, 0x08, 0x06, 0xf3, 0x88, 0xaf, 0xe9, 0x25,
	0x61, 0x37, 0xcb, 0xff, 0xa9, 0xbc, 0x18, 0x6e, 0x13, 0xca, 0x49, 0xe3, 0xc5, 0x77, 0x2d, 0x56,
	0xec, 0xcf, 0x04, 0x6e, 0xad, 0x5c, 0xf5, 0x63, 0xfb, 0xd9, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x50, 0x10, 0x21, 0x84, 0xe9, 0x0a, 0x00, 0x00,
}
