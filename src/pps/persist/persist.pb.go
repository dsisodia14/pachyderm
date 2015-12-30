// Code generated by protoc-gen-go.
// source: pps/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	pps/persist/persist.proto

It has these top-level messages:
	JobInfo
	JobInfos
	JobOutput
	PipelineInfo
	PipelineInfos
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/google-protobuf"
import google_protobuf1 "go.pedge.io/google-protobuf"
import pfs "github.com/pachyderm/pachyderm/src/pfs"
import pachyderm_pps "github.com/pachyderm/pachyderm/src/pps"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type JobInfo struct {
	JobId        string                      `protobuf:"bytes,1,opt,name=job_id" json:"job_id,omitempty"`
	Transform    *pachyderm_pps.Transform    `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	PipelineName string                      `protobuf:"bytes,3,opt,name=pipeline_name" json:"pipeline_name,omitempty"`
	Shards       uint64                      `protobuf:"varint,4,opt,name=shards" json:"shards,omitempty"`
	InputCommit  []*pfs.Commit               `protobuf:"bytes,5,rep,name=input_commit" json:"input_commit,omitempty"`
	ParentJob    *pachyderm_pps.Job          `protobuf:"bytes,6,opt,name=parent_job" json:"parent_job,omitempty"`
	CreatedAt    *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=created_at" json:"created_at,omitempty"`
	CommitIndex  string                      `protobuf:"bytes,8,opt,name=commit_index" json:"commit_index,omitempty"`
}

func (m *JobInfo) Reset()         { *m = JobInfo{} }
func (m *JobInfo) String() string { return proto.CompactTextString(m) }
func (*JobInfo) ProtoMessage()    {}

func (m *JobInfo) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *JobInfo) GetInputCommit() []*pfs.Commit {
	if m != nil {
		return m.InputCommit
	}
	return nil
}

func (m *JobInfo) GetParentJob() *pachyderm_pps.Job {
	if m != nil {
		return m.ParentJob
	}
	return nil
}

func (m *JobInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type JobInfos struct {
	JobInfo []*JobInfo `protobuf:"bytes,1,rep,name=job_info" json:"job_info,omitempty"`
}

func (m *JobInfos) Reset()         { *m = JobInfos{} }
func (m *JobInfos) String() string { return proto.CompactTextString(m) }
func (*JobInfos) ProtoMessage()    {}

func (m *JobInfos) GetJobInfo() []*JobInfo {
	if m != nil {
		return m.JobInfo
	}
	return nil
}

type JobOutput struct {
	JobId        string      `protobuf:"bytes,1,opt,name=job_id" json:"job_id,omitempty"`
	OutputCommit *pfs.Commit `protobuf:"bytes,2,opt,name=output_commit" json:"output_commit,omitempty"`
}

func (m *JobOutput) Reset()         { *m = JobOutput{} }
func (m *JobOutput) String() string { return proto.CompactTextString(m) }
func (*JobOutput) ProtoMessage()    {}

func (m *JobOutput) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

type PipelineInfo struct {
	PipelineName string                      `protobuf:"bytes,1,opt,name=pipeline_name" json:"pipeline_name,omitempty"`
	Transform    *pachyderm_pps.Transform    `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Shards       uint64                      `protobuf:"varint,3,opt,name=shards" json:"shards,omitempty"`
	InputRepo    []*pfs.Repo                 `protobuf:"bytes,4,rep,name=input_repo" json:"input_repo,omitempty"`
	CreatedAt    *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *PipelineInfo) Reset()         { *m = PipelineInfo{} }
func (m *PipelineInfo) String() string { return proto.CompactTextString(m) }
func (*PipelineInfo) ProtoMessage()    {}

func (m *PipelineInfo) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *PipelineInfo) GetInputRepo() []*pfs.Repo {
	if m != nil {
		return m.InputRepo
	}
	return nil
}

func (m *PipelineInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type PipelineInfos struct {
	PipelineInfo []*PipelineInfo `protobuf:"bytes,1,rep,name=pipeline_info" json:"pipeline_info,omitempty"`
}

func (m *PipelineInfos) Reset()         { *m = PipelineInfos{} }
func (m *PipelineInfos) String() string { return proto.CompactTextString(m) }
func (*PipelineInfos) ProtoMessage()    {}

func (m *PipelineInfos) GetPipelineInfo() []*PipelineInfo {
	if m != nil {
		return m.PipelineInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*JobInfo)(nil), "pachyderm.pps.persist.JobInfo")
	proto.RegisterType((*JobInfos)(nil), "pachyderm.pps.persist.JobInfos")
	proto.RegisterType((*JobOutput)(nil), "pachyderm.pps.persist.JobOutput")
	proto.RegisterType((*PipelineInfo)(nil), "pachyderm.pps.persist.PipelineInfo")
	proto.RegisterType((*PipelineInfos)(nil), "pachyderm.pps.persist.PipelineInfos")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for API service

type APIClient interface {
	// job_id cannot be set
	// timestamp cannot be set
	CreateJobInfo(ctx context.Context, in *JobInfo, opts ...grpc.CallOption) (*JobInfo, error)
	GetJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error)
	// ordered by time, latest to earliest
	ListJobInfos(ctx context.Context, in *pachyderm_pps.ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error)
	// should only be called when rolling back if a Job does not start!
	DeleteJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	CreateJobOutput(ctx context.Context, in *JobOutput, opts ...grpc.CallOption) (*JobOutput, error)
	GetJobOutput(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobOutput, error)
	CreatePipelineInfo(ctx context.Context, in *PipelineInfo, opts ...grpc.CallOption) (*PipelineInfo, error)
	GetPipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*PipelineInfo, error)
	// ordered by time, latest to earliest
	ListPipelineInfos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PipelineInfos, error)
	DeletePipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateJobInfo(ctx context.Context, in *JobInfo, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/GetJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListJobInfos(ctx context.Context, in *pachyderm_pps.ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error) {
	out := new(JobInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/ListJobInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/DeleteJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateJobOutput(ctx context.Context, in *JobOutput, opts ...grpc.CallOption) (*JobOutput, error) {
	out := new(JobOutput)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobOutput", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetJobOutput(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobOutput, error) {
	out := new(JobOutput)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/GetJobOutput", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreatePipelineInfo(ctx context.Context, in *PipelineInfo, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreatePipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetPipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/GetPipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListPipelineInfos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PipelineInfos, error) {
	out := new(PipelineInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/ListPipelineInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeletePipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/DeletePipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	// job_id cannot be set
	// timestamp cannot be set
	CreateJobInfo(context.Context, *JobInfo) (*JobInfo, error)
	GetJobInfo(context.Context, *pachyderm_pps.Job) (*JobInfo, error)
	// ordered by time, latest to earliest
	ListJobInfos(context.Context, *pachyderm_pps.ListJobRequest) (*JobInfos, error)
	// should only be called when rolling back if a Job does not start!
	DeleteJobInfo(context.Context, *pachyderm_pps.Job) (*google_protobuf.Empty, error)
	CreateJobOutput(context.Context, *JobOutput) (*JobOutput, error)
	GetJobOutput(context.Context, *pachyderm_pps.Job) (*JobOutput, error)
	CreatePipelineInfo(context.Context, *PipelineInfo) (*PipelineInfo, error)
	GetPipelineInfo(context.Context, *pachyderm_pps.Pipeline) (*PipelineInfo, error)
	// ordered by time, latest to earliest
	ListPipelineInfos(context.Context, *google_protobuf.Empty) (*PipelineInfos, error)
	DeletePipelineInfo(context.Context, *pachyderm_pps.Pipeline) (*google_protobuf.Empty, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetJobInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListJobInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListJobInfos(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeleteJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeleteJobInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreateJobOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobOutput)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobOutput(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetJobOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetJobOutput(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreatePipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PipelineInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreatePipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetPipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetPipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListPipelineInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListPipelineInfos(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeletePipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeletePipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.persist.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJobInfo",
			Handler:    _API_CreateJobInfo_Handler,
		},
		{
			MethodName: "GetJobInfo",
			Handler:    _API_GetJobInfo_Handler,
		},
		{
			MethodName: "ListJobInfos",
			Handler:    _API_ListJobInfos_Handler,
		},
		{
			MethodName: "DeleteJobInfo",
			Handler:    _API_DeleteJobInfo_Handler,
		},
		{
			MethodName: "CreateJobOutput",
			Handler:    _API_CreateJobOutput_Handler,
		},
		{
			MethodName: "GetJobOutput",
			Handler:    _API_GetJobOutput_Handler,
		},
		{
			MethodName: "CreatePipelineInfo",
			Handler:    _API_CreatePipelineInfo_Handler,
		},
		{
			MethodName: "GetPipelineInfo",
			Handler:    _API_GetPipelineInfo_Handler,
		},
		{
			MethodName: "ListPipelineInfos",
			Handler:    _API_ListPipelineInfos_Handler,
		},
		{
			MethodName: "DeletePipelineInfo",
			Handler:    _API_DeletePipelineInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
