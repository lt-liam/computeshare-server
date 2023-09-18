// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.2
// source: api/compute/v1/compute_power.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationComputePowerCancelExecPythonPackage = "/api.compute.v1.ComputePower/CancelExecPythonPackage"
const OperationComputePowerDownloadScriptExecuteResult = "/api.compute.v1.ComputePower/DownloadScriptExecuteResult"
const OperationComputePowerGetScriptInfo = "/api.compute.v1.ComputePower/GetScriptInfo"
const OperationComputePowerGetScriptList = "/api.compute.v1.ComputePower/GetScriptList"
const OperationComputePowerRunPythonPackage = "/api.compute.v1.ComputePower/RunPythonPackage"
const OperationComputePowerUploadScriptFile = "/api.compute.v1.ComputePower/UploadScriptFile"

type ComputePowerHTTPServer interface {
	// CancelExecPythonPackage取消执行
	CancelExecPythonPackage(context.Context, *CancelExecPythonPackageRequest) (*CancelExecPythonPackageReply, error)
	// DownloadScriptExecuteResult下载执行结果（http接口另外写）
	DownloadScriptExecuteResult(context.Context, *DownloadScriptExecuteResultRequest) (*DownloadScriptExecuteResultReply, error)
	// GetScriptInfo通过id
	GetScriptInfo(context.Context, *GetScriptInfoRequest) (*GetScriptInfoReply, error)
	// GetScriptList查询脚本列表
	GetScriptList(context.Context, *GetScriptListRequest) (*GetScriptListReply, error)
	// RunPythonPackage执行脚本
	RunPythonPackage(context.Context, *RunPythonPackageServerRequest) (*RunPythonPackageServerReply, error)
	// UploadScriptFile上传脚本（http接口另外写）
	UploadScriptFile(context.Context, *UploadScriptFileRequest) (*UploadScriptFileReply, error)
}

func RegisterComputePowerHTTPServer(s *http.Server, srv ComputePowerHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/compute-power", _ComputePower_UploadScriptFile0_HTTP_Handler(srv))
	r.GET("/v1/compute-power/script/list", _ComputePower_GetScriptList0_HTTP_Handler(srv))
	r.POST("/v1/compute-power/python", _ComputePower_RunPythonPackage0_HTTP_Handler(srv))
	r.POST("/v1/compute-power/python", _ComputePower_CancelExecPythonPackage0_HTTP_Handler(srv))
	r.GET("/v1/compute-power/script/info/{id}", _ComputePower_GetScriptInfo0_HTTP_Handler(srv))
	r.POST("/v1/compute-power/{id}", _ComputePower_DownloadScriptExecuteResult0_HTTP_Handler(srv))
}

func _ComputePower_UploadScriptFile0_HTTP_Handler(srv ComputePowerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadScriptFileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputePowerUploadScriptFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadScriptFile(ctx, req.(*UploadScriptFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadScriptFileReply)
		return ctx.Result(200, reply)
	}
}

func _ComputePower_GetScriptList0_HTTP_Handler(srv ComputePowerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetScriptListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputePowerGetScriptList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetScriptList(ctx, req.(*GetScriptListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetScriptListReply)
		return ctx.Result(200, reply)
	}
}

func _ComputePower_RunPythonPackage0_HTTP_Handler(srv ComputePowerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RunPythonPackageServerRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputePowerRunPythonPackage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RunPythonPackage(ctx, req.(*RunPythonPackageServerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RunPythonPackageServerReply)
		return ctx.Result(200, reply)
	}
}

func _ComputePower_CancelExecPythonPackage0_HTTP_Handler(srv ComputePowerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CancelExecPythonPackageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputePowerCancelExecPythonPackage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CancelExecPythonPackage(ctx, req.(*CancelExecPythonPackageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CancelExecPythonPackageReply)
		return ctx.Result(200, reply)
	}
}

func _ComputePower_GetScriptInfo0_HTTP_Handler(srv ComputePowerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetScriptInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputePowerGetScriptInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetScriptInfo(ctx, req.(*GetScriptInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetScriptInfoReply)
		return ctx.Result(200, reply)
	}
}

func _ComputePower_DownloadScriptExecuteResult0_HTTP_Handler(srv ComputePowerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DownloadScriptExecuteResultRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputePowerDownloadScriptExecuteResult)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DownloadScriptExecuteResult(ctx, req.(*DownloadScriptExecuteResultRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DownloadScriptExecuteResultReply)
		return ctx.Result(200, reply)
	}
}

type ComputePowerHTTPClient interface {
	CancelExecPythonPackage(ctx context.Context, req *CancelExecPythonPackageRequest, opts ...http.CallOption) (rsp *CancelExecPythonPackageReply, err error)
	DownloadScriptExecuteResult(ctx context.Context, req *DownloadScriptExecuteResultRequest, opts ...http.CallOption) (rsp *DownloadScriptExecuteResultReply, err error)
	GetScriptInfo(ctx context.Context, req *GetScriptInfoRequest, opts ...http.CallOption) (rsp *GetScriptInfoReply, err error)
	GetScriptList(ctx context.Context, req *GetScriptListRequest, opts ...http.CallOption) (rsp *GetScriptListReply, err error)
	RunPythonPackage(ctx context.Context, req *RunPythonPackageServerRequest, opts ...http.CallOption) (rsp *RunPythonPackageServerReply, err error)
	UploadScriptFile(ctx context.Context, req *UploadScriptFileRequest, opts ...http.CallOption) (rsp *UploadScriptFileReply, err error)
}

type ComputePowerHTTPClientImpl struct {
	cc *http.Client
}

func NewComputePowerHTTPClient(client *http.Client) ComputePowerHTTPClient {
	return &ComputePowerHTTPClientImpl{client}
}

func (c *ComputePowerHTTPClientImpl) CancelExecPythonPackage(ctx context.Context, in *CancelExecPythonPackageRequest, opts ...http.CallOption) (*CancelExecPythonPackageReply, error) {
	var out CancelExecPythonPackageReply
	pattern := "/v1/compute-power/python"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationComputePowerCancelExecPythonPackage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputePowerHTTPClientImpl) DownloadScriptExecuteResult(ctx context.Context, in *DownloadScriptExecuteResultRequest, opts ...http.CallOption) (*DownloadScriptExecuteResultReply, error) {
	var out DownloadScriptExecuteResultReply
	pattern := "/v1/compute-power/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationComputePowerDownloadScriptExecuteResult))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputePowerHTTPClientImpl) GetScriptInfo(ctx context.Context, in *GetScriptInfoRequest, opts ...http.CallOption) (*GetScriptInfoReply, error) {
	var out GetScriptInfoReply
	pattern := "/v1/compute-power/script/info/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationComputePowerGetScriptInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputePowerHTTPClientImpl) GetScriptList(ctx context.Context, in *GetScriptListRequest, opts ...http.CallOption) (*GetScriptListReply, error) {
	var out GetScriptListReply
	pattern := "/v1/compute-power/script/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationComputePowerGetScriptList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputePowerHTTPClientImpl) RunPythonPackage(ctx context.Context, in *RunPythonPackageServerRequest, opts ...http.CallOption) (*RunPythonPackageServerReply, error) {
	var out RunPythonPackageServerReply
	pattern := "/v1/compute-power/python"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationComputePowerRunPythonPackage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputePowerHTTPClientImpl) UploadScriptFile(ctx context.Context, in *UploadScriptFileRequest, opts ...http.CallOption) (*UploadScriptFileReply, error) {
	var out UploadScriptFileReply
	pattern := "/v1/compute-power"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationComputePowerUploadScriptFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
