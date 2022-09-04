// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/rtaa/streamer/twitter/v1/twitterwatch.proto

package twitterv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/cvcio/rtaa-72/internal/proto/rtaa/streamer/twitter/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TwitterWatchServiceName is the fully-qualified name of the TwitterWatchService service.
	TwitterWatchServiceName = "proto.rtaa.streamer.twitter.v1.TwitterWatchService"
)

// TwitterWatchServiceClient is a client for the proto.rtaa.streamer.twitter.v1.TwitterWatchService
// service.
type TwitterWatchServiceClient interface {
	// Autorize implements twitter's oauth2 authorization
	// read more: https://developer.twitter.com/en/docs/authentication/oauth-2-0
	Authorize(context.Context, *connect_go.Request[v1.AuthorizeRequest]) (*connect_go.Response[v1.AuthResponse], error)
	// Obb implements twitter's pin-based oath 1.0a authorization
	// read more: https://developer.twitter.com/en/docs/authentication/oauth-1-0a/pin-based-oauth
	Obb(context.Context, *connect_go.Request[v1.Empty]) (*connect_go.Response[v1.ObbResponse], error)
	// Verify twitter's authorization
	Verify(context.Context, *connect_go.Request[v1.VerifyRequest]) (*connect_go.Response[v1.AuthResponse], error)
	// Addrules adds a new rule to twitter streaming api
	AddRules(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error)
	// GetRules returns all rules stored on twitter streaming api
	GetRules(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error)
	// GetRule by id, specific to user id
	GetRule(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error)
	// DeleteRules deletes all rules
	DeleteRules(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error)
	// DeleteRule deletes a specific rule
	DeleteRule(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error)
	// Connect establish a connection with the service
	Connect(context.Context, *connect_go.Request[v1.ConnectRequest]) (*connect_go.Response[v1.ConnectResponse], error)
	// Disconnect establish a connection with the service
	Disconnect(context.Context, *connect_go.Request[v1.ConnectRequest]) (*connect_go.Response[v1.ConnectResponse], error)
	// Stream data from twitter api
	Stream(context.Context, *connect_go.Request[v1.StreamRequest]) (*connect_go.ServerStreamForClient[v1.StreamResponse], error)
}

// NewTwitterWatchServiceClient constructs a client for the
// proto.rtaa.streamer.twitter.v1.TwitterWatchService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTwitterWatchServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TwitterWatchServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &twitterWatchServiceClient{
		authorize: connect_go.NewClient[v1.AuthorizeRequest, v1.AuthResponse](
			httpClient,
			baseURL+"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Authorize",
			opts...,
		),
		obb: connect_go.NewClient[v1.Empty, v1.ObbResponse](
			httpClient,
			baseURL+"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Obb",
			opts...,
		),
		verify: connect_go.NewClient[v1.VerifyRequest, v1.AuthResponse](
			httpClient,
			baseURL+"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Verify",
			opts...,
		),
		addRules: connect_go.NewClient[v1.RulesRequest, v1.RulesResponse](
			httpClient,
			baseURL+"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/AddRules",
			opts...,
		),
		getRules: connect_go.NewClient[v1.RulesRequest, v1.RulesResponse](
			httpClient,
			baseURL+"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/GetRules",
			opts...,
		),
		getRule: connect_go.NewClient[v1.RulesRequest, v1.RulesResponse](
			httpClient,
			baseURL+"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/GetRule",
			opts...,
		),
		deleteRules: connect_go.NewClient[v1.RulesRequest, v1.RulesResponse](
			httpClient,
			baseURL+"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/DeleteRules",
			opts...,
		),
		deleteRule: connect_go.NewClient[v1.RulesRequest, v1.RulesResponse](
			httpClient,
			baseURL+"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/DeleteRule",
			opts...,
		),
		connect: connect_go.NewClient[v1.ConnectRequest, v1.ConnectResponse](
			httpClient,
			baseURL+"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Connect",
			opts...,
		),
		disconnect: connect_go.NewClient[v1.ConnectRequest, v1.ConnectResponse](
			httpClient,
			baseURL+"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Disconnect",
			opts...,
		),
		stream: connect_go.NewClient[v1.StreamRequest, v1.StreamResponse](
			httpClient,
			baseURL+"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Stream",
			opts...,
		),
	}
}

// twitterWatchServiceClient implements TwitterWatchServiceClient.
type twitterWatchServiceClient struct {
	authorize   *connect_go.Client[v1.AuthorizeRequest, v1.AuthResponse]
	obb         *connect_go.Client[v1.Empty, v1.ObbResponse]
	verify      *connect_go.Client[v1.VerifyRequest, v1.AuthResponse]
	addRules    *connect_go.Client[v1.RulesRequest, v1.RulesResponse]
	getRules    *connect_go.Client[v1.RulesRequest, v1.RulesResponse]
	getRule     *connect_go.Client[v1.RulesRequest, v1.RulesResponse]
	deleteRules *connect_go.Client[v1.RulesRequest, v1.RulesResponse]
	deleteRule  *connect_go.Client[v1.RulesRequest, v1.RulesResponse]
	connect     *connect_go.Client[v1.ConnectRequest, v1.ConnectResponse]
	disconnect  *connect_go.Client[v1.ConnectRequest, v1.ConnectResponse]
	stream      *connect_go.Client[v1.StreamRequest, v1.StreamResponse]
}

// Authorize calls proto.rtaa.streamer.twitter.v1.TwitterWatchService.Authorize.
func (c *twitterWatchServiceClient) Authorize(ctx context.Context, req *connect_go.Request[v1.AuthorizeRequest]) (*connect_go.Response[v1.AuthResponse], error) {
	return c.authorize.CallUnary(ctx, req)
}

// Obb calls proto.rtaa.streamer.twitter.v1.TwitterWatchService.Obb.
func (c *twitterWatchServiceClient) Obb(ctx context.Context, req *connect_go.Request[v1.Empty]) (*connect_go.Response[v1.ObbResponse], error) {
	return c.obb.CallUnary(ctx, req)
}

// Verify calls proto.rtaa.streamer.twitter.v1.TwitterWatchService.Verify.
func (c *twitterWatchServiceClient) Verify(ctx context.Context, req *connect_go.Request[v1.VerifyRequest]) (*connect_go.Response[v1.AuthResponse], error) {
	return c.verify.CallUnary(ctx, req)
}

// AddRules calls proto.rtaa.streamer.twitter.v1.TwitterWatchService.AddRules.
func (c *twitterWatchServiceClient) AddRules(ctx context.Context, req *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error) {
	return c.addRules.CallUnary(ctx, req)
}

// GetRules calls proto.rtaa.streamer.twitter.v1.TwitterWatchService.GetRules.
func (c *twitterWatchServiceClient) GetRules(ctx context.Context, req *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error) {
	return c.getRules.CallUnary(ctx, req)
}

// GetRule calls proto.rtaa.streamer.twitter.v1.TwitterWatchService.GetRule.
func (c *twitterWatchServiceClient) GetRule(ctx context.Context, req *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error) {
	return c.getRule.CallUnary(ctx, req)
}

// DeleteRules calls proto.rtaa.streamer.twitter.v1.TwitterWatchService.DeleteRules.
func (c *twitterWatchServiceClient) DeleteRules(ctx context.Context, req *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error) {
	return c.deleteRules.CallUnary(ctx, req)
}

// DeleteRule calls proto.rtaa.streamer.twitter.v1.TwitterWatchService.DeleteRule.
func (c *twitterWatchServiceClient) DeleteRule(ctx context.Context, req *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error) {
	return c.deleteRule.CallUnary(ctx, req)
}

// Connect calls proto.rtaa.streamer.twitter.v1.TwitterWatchService.Connect.
func (c *twitterWatchServiceClient) Connect(ctx context.Context, req *connect_go.Request[v1.ConnectRequest]) (*connect_go.Response[v1.ConnectResponse], error) {
	return c.connect.CallUnary(ctx, req)
}

// Disconnect calls proto.rtaa.streamer.twitter.v1.TwitterWatchService.Disconnect.
func (c *twitterWatchServiceClient) Disconnect(ctx context.Context, req *connect_go.Request[v1.ConnectRequest]) (*connect_go.Response[v1.ConnectResponse], error) {
	return c.disconnect.CallUnary(ctx, req)
}

// Stream calls proto.rtaa.streamer.twitter.v1.TwitterWatchService.Stream.
func (c *twitterWatchServiceClient) Stream(ctx context.Context, req *connect_go.Request[v1.StreamRequest]) (*connect_go.ServerStreamForClient[v1.StreamResponse], error) {
	return c.stream.CallServerStream(ctx, req)
}

// TwitterWatchServiceHandler is an implementation of the
// proto.rtaa.streamer.twitter.v1.TwitterWatchService service.
type TwitterWatchServiceHandler interface {
	// Autorize implements twitter's oauth2 authorization
	// read more: https://developer.twitter.com/en/docs/authentication/oauth-2-0
	Authorize(context.Context, *connect_go.Request[v1.AuthorizeRequest]) (*connect_go.Response[v1.AuthResponse], error)
	// Obb implements twitter's pin-based oath 1.0a authorization
	// read more: https://developer.twitter.com/en/docs/authentication/oauth-1-0a/pin-based-oauth
	Obb(context.Context, *connect_go.Request[v1.Empty]) (*connect_go.Response[v1.ObbResponse], error)
	// Verify twitter's authorization
	Verify(context.Context, *connect_go.Request[v1.VerifyRequest]) (*connect_go.Response[v1.AuthResponse], error)
	// Addrules adds a new rule to twitter streaming api
	AddRules(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error)
	// GetRules returns all rules stored on twitter streaming api
	GetRules(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error)
	// GetRule by id, specific to user id
	GetRule(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error)
	// DeleteRules deletes all rules
	DeleteRules(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error)
	// DeleteRule deletes a specific rule
	DeleteRule(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error)
	// Connect establish a connection with the service
	Connect(context.Context, *connect_go.Request[v1.ConnectRequest]) (*connect_go.Response[v1.ConnectResponse], error)
	// Disconnect establish a connection with the service
	Disconnect(context.Context, *connect_go.Request[v1.ConnectRequest]) (*connect_go.Response[v1.ConnectResponse], error)
	// Stream data from twitter api
	Stream(context.Context, *connect_go.Request[v1.StreamRequest], *connect_go.ServerStream[v1.StreamResponse]) error
}

// NewTwitterWatchServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTwitterWatchServiceHandler(svc TwitterWatchServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Authorize", connect_go.NewUnaryHandler(
		"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Authorize",
		svc.Authorize,
		opts...,
	))
	mux.Handle("/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Obb", connect_go.NewUnaryHandler(
		"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Obb",
		svc.Obb,
		opts...,
	))
	mux.Handle("/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Verify", connect_go.NewUnaryHandler(
		"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Verify",
		svc.Verify,
		opts...,
	))
	mux.Handle("/proto.rtaa.streamer.twitter.v1.TwitterWatchService/AddRules", connect_go.NewUnaryHandler(
		"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/AddRules",
		svc.AddRules,
		opts...,
	))
	mux.Handle("/proto.rtaa.streamer.twitter.v1.TwitterWatchService/GetRules", connect_go.NewUnaryHandler(
		"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/GetRules",
		svc.GetRules,
		opts...,
	))
	mux.Handle("/proto.rtaa.streamer.twitter.v1.TwitterWatchService/GetRule", connect_go.NewUnaryHandler(
		"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/GetRule",
		svc.GetRule,
		opts...,
	))
	mux.Handle("/proto.rtaa.streamer.twitter.v1.TwitterWatchService/DeleteRules", connect_go.NewUnaryHandler(
		"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/DeleteRules",
		svc.DeleteRules,
		opts...,
	))
	mux.Handle("/proto.rtaa.streamer.twitter.v1.TwitterWatchService/DeleteRule", connect_go.NewUnaryHandler(
		"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/DeleteRule",
		svc.DeleteRule,
		opts...,
	))
	mux.Handle("/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Connect", connect_go.NewUnaryHandler(
		"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Connect",
		svc.Connect,
		opts...,
	))
	mux.Handle("/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Disconnect", connect_go.NewUnaryHandler(
		"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Disconnect",
		svc.Disconnect,
		opts...,
	))
	mux.Handle("/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Stream", connect_go.NewServerStreamHandler(
		"/proto.rtaa.streamer.twitter.v1.TwitterWatchService/Stream",
		svc.Stream,
		opts...,
	))
	return "/proto.rtaa.streamer.twitter.v1.TwitterWatchService/", mux
}

// UnimplementedTwitterWatchServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTwitterWatchServiceHandler struct{}

func (UnimplementedTwitterWatchServiceHandler) Authorize(context.Context, *connect_go.Request[v1.AuthorizeRequest]) (*connect_go.Response[v1.AuthResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.rtaa.streamer.twitter.v1.TwitterWatchService.Authorize is not implemented"))
}

func (UnimplementedTwitterWatchServiceHandler) Obb(context.Context, *connect_go.Request[v1.Empty]) (*connect_go.Response[v1.ObbResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.rtaa.streamer.twitter.v1.TwitterWatchService.Obb is not implemented"))
}

func (UnimplementedTwitterWatchServiceHandler) Verify(context.Context, *connect_go.Request[v1.VerifyRequest]) (*connect_go.Response[v1.AuthResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.rtaa.streamer.twitter.v1.TwitterWatchService.Verify is not implemented"))
}

func (UnimplementedTwitterWatchServiceHandler) AddRules(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.rtaa.streamer.twitter.v1.TwitterWatchService.AddRules is not implemented"))
}

func (UnimplementedTwitterWatchServiceHandler) GetRules(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.rtaa.streamer.twitter.v1.TwitterWatchService.GetRules is not implemented"))
}

func (UnimplementedTwitterWatchServiceHandler) GetRule(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.rtaa.streamer.twitter.v1.TwitterWatchService.GetRule is not implemented"))
}

func (UnimplementedTwitterWatchServiceHandler) DeleteRules(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.rtaa.streamer.twitter.v1.TwitterWatchService.DeleteRules is not implemented"))
}

func (UnimplementedTwitterWatchServiceHandler) DeleteRule(context.Context, *connect_go.Request[v1.RulesRequest]) (*connect_go.Response[v1.RulesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.rtaa.streamer.twitter.v1.TwitterWatchService.DeleteRule is not implemented"))
}

func (UnimplementedTwitterWatchServiceHandler) Connect(context.Context, *connect_go.Request[v1.ConnectRequest]) (*connect_go.Response[v1.ConnectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.rtaa.streamer.twitter.v1.TwitterWatchService.Connect is not implemented"))
}

func (UnimplementedTwitterWatchServiceHandler) Disconnect(context.Context, *connect_go.Request[v1.ConnectRequest]) (*connect_go.Response[v1.ConnectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.rtaa.streamer.twitter.v1.TwitterWatchService.Disconnect is not implemented"))
}

func (UnimplementedTwitterWatchServiceHandler) Stream(context.Context, *connect_go.Request[v1.StreamRequest], *connect_go.ServerStream[v1.StreamResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.rtaa.streamer.twitter.v1.TwitterWatchService.Stream is not implemented"))
}