package zerolog

import (
	"bytes"
	"context"
	"testing"

	pb "github.com/philip-bui/grpc-zerolog/protos"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type MockNetAddr struct {
	mock.Mock
}

func (n MockNetAddr) Network() string {
	return "tcp"
}

func (n MockNetAddr) String() string {
	return "127.0.0.1"
}

type TestUtilSuite struct {
	suite.Suite
	out    *bytes.Buffer
	log    *zerolog.Event
	msg    string
	msgDef string
	req    *pb.TestMessage
	resp   *pb.TestMessage
}

func (s *TestUtilSuite) SetupSuite() {
	s.msg = "PhilipB"
	s.msgDef = `{"level":"debug","message":"PhilipB"}`
	s.req = &pb.TestMessage{
		Test: "req",
	}
	s.resp = &pb.TestMessage{
		Test: "resp",
	}
}

func (s *TestUtilSuite) SetupTest() {
	s.out = &bytes.Buffer{}
	logger := zerolog.New(s.out)
	s.log = logger.Debug()

	TimestampLog = true
	ServiceField = "service"
	ServiceLog = true
	MethodField = "method"
	MethodLog = true
	DurationField = "dur"
	DurationLog = true
	IPField = "ip"
	IPLog = true
	MetadataField = "md"
	MetadataLog = true
	UserAgentField = "ua"
	UserAgentLog = true
	ReqField = "req"
	ReqLog = true
	RespField = "resp"
	RespLog = true
	MaxSize = 2048000
	CodeField = "code"
	MsgField = "msg"
	DetailsField = "details"
	UnaryMessageDefault = "unary"
}

func TestUtil(t *testing.T) {
	suite.Run(t, new(TestUtilSuite))
}

func (s *TestUtilSuite) TestLogIP() {
	LogIP(peer.NewContext(context.Background(), &peer.Peer{
		Addr: MockNetAddr{},
	}), s.log)
	s.log.Msg(s.msg)
	s.JSONEq(`{"level":"debug","ip":"127.0.0.1","message":"PhilipB"}`, s.out.String())
}

func (s *TestUtilSuite) TestLogIPInvalid() {
	LogIP(context.Background(), s.log)
	s.log.Msg(s.msg)
	s.JSONEq(s.msgDef, s.out.String())
}

func (s *TestUtilSuite) TestLogIPDisabled() {
	IPLog = false
	LogIP(context.Background(), s.log)
	s.log.Msg(s.msg)
	s.JSONEq(s.msgDef, s.out.String())
}

func (s *TestUtilSuite) TestLogRequest() {
	LogRequest(s.log, s.req)
	s.log.Msg(s.msg)
	s.JSONEq(`{"level":"debug","req":{"test":"req"},"message":"PhilipB"}`, s.out.String())
}

func (s *TestUtilSuite) TestLogRequestDisabled() {
	ReqLog = false
	LogRequest(s.log, s.req)
	s.log.Msg(s.msg)
	s.JSONEq(s.msgDef, s.out.String())
}

func (s *TestUtilSuite) TestLogRequestField() {
	ReqField = "philip"
	LogRequest(s.log, s.req)
	s.log.Msg(s.msg)
	s.JSONEq(`{"level":"debug","philip":{"test":"req"},"message":"PhilipB"}`, s.out.String())
}

func (s *TestUtilSuite) TestLogResponse() {
	LogResponse(s.log, s.req)
	s.log.Msg(s.msg)
	s.JSONEq(`{"level":"debug","resp":{"test":"req"},"message":"PhilipB"}`, s.out.String())
}

func (s *TestUtilSuite) TestLogResponseDisabled() {
	RespLog = false
	LogResponse(s.log, s.resp)
	s.log.Msg(s.msg)
	s.JSONEq(s.msgDef, s.out.String())
}

func (s *TestUtilSuite) TestLogResponseField() {
	RespField = "philip"
	LogResponse(s.log, s.resp)
	s.log.Msg(s.msg)
	s.JSONEq(`{"level":"debug","philip":{"test":"resp"},"message":"PhilipB"}`, s.out.String())
}

func (s *TestUtilSuite) TestGetRawJSONInvalid() {
	s.Nil(GetRawJSON(new(interface{})))
}

func (s *TestUtilSuite) TestLogIncomingMetadata() {
	LogIncomingMetadata(metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
		"philip": "WasHere",
	})), s.log)
	s.log.Msg(s.msg)
	s.JSONEq(`{"level":"debug","md":{"philip":"WasHere"},"message":"PhilipB"}`, s.out.String())
}

func (s *TestUtilSuite) TestLogIncomingMetadataField() {
	MetadataField = "metadata"
	LogIncomingMetadata(metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
		"philip": "WasHere",
	})), s.log)
	s.log.Msg(s.msg)
	s.JSONEq(`{"level":"debug","metadata":{"philip":"WasHere"},"message":"PhilipB"}`, s.out.String())
}

func (s *TestUtilSuite) TestLogIncomingMetadataInvalid() {
	LogIncomingMetadata(context.Background(), s.log)
	s.log.Msg(s.msg)
	s.JSONEq(s.msgDef, s.out.String())
}

func (s *TestUtilSuite) TestLogIncomingMetadataDisabledForIp() {
	MetadataLog = false
	LogIncomingMetadata(metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
		"user-agent": "test",
	})), s.log)
	s.log.Msg(s.msg)
	s.JSONEq(`{"message":"PhilipB", "level":"debug", "ua":"test"}`, s.out.String())
}

func (s *TestUtilSuite) TestLogIncomingMetadataDisabledForIpField() {
	MetadataLog = false
	UserAgentField = "user-agent"
	LogIncomingMetadata(metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
		"user-agent": "test",
	})), s.log)
	s.log.Msg(s.msg)
	s.JSONEq(`{"message":"PhilipB", "level":"debug", "user-agent":"test"}`, s.out.String())
}

func (s *TestUtilSuite) TestLogIncomingMetadataDisabledForIpInvalid() {
	MetadataLog = false
	LogIncomingMetadata(context.Background(), s.log)
	s.log.Msg(s.msg)
	s.JSONEq(s.msgDef, s.out.String())
}
