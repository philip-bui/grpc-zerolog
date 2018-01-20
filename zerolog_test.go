package zerolog

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/grpclog"
)

func TestSetNewZeroLogger(t *testing.T) {
	GrpcLogSetNewZeroLogger()
}

func TestSetZeroLogger(t *testing.T) {
	GrpcLogSetZeroLogger(NewGrpcZeroLogger(zerolog.Logger{}))
}

type TestLogSuite struct {
	suite.Suite
	out     *bytes.Buffer
	format  string
	args    []interface{}
	logger  zerolog.Logger
	grpclog GrpcZeroLogger
}

func (s *TestLogSuite) SetupSuite() {
	s.format = "Philip%v%v"
	s.args = []interface{}{"Was", "Here"}
}

func (s *TestLogSuite) SetupTest() {
	s.out = &bytes.Buffer{}
	s.logger = zerolog.New(s.out)
	s.grpclog = NewGrpcZeroLogger(s.logger)
	GrpcLogSetZeroLogger(s.grpclog)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

func TestGrpcZeroLogger(t *testing.T) {
	suite.Run(t, new(TestLogSuite))
}

func (s *TestLogSuite) TestError() {
	grpclog.Error(s.args...)
	s.JSONEq(`{"level":"error","message":"WasHere"}`, s.out.String())
}

func (s *TestLogSuite) TestErrorf() {
	grpclog.Errorf(s.format, s.args...)
	s.JSONEq(`{"level":"error","message":"PhilipWasHere"}`, s.out.String())
}

func (s *TestLogSuite) TestErrorln() {
	grpclog.Errorln(s.args...)
	s.JSONEq(`{"level":"error","message":"WasHere"}`, s.out.String())
}

func (s *TestLogSuite) TestInfo() {
	grpclog.Info(s.args...)
	s.JSONEq(`{"level":"info","message":"WasHere"}`, s.out.String())
}

func (s *TestLogSuite) TestInfof() {
	grpclog.Infof(s.format, s.args...)
	s.JSONEq(`{"level":"info","message":"PhilipWasHere"}`, s.out.String())
}

func (s *TestLogSuite) TestInfoln() {
	grpclog.Infoln(s.args...)
	s.JSONEq(`{"level":"info","message":"WasHere"}`, s.out.String())
}

func (s *TestLogSuite) TestWarning() {
	grpclog.Warning(s.args...)
	s.JSONEq(`{"level":"warn","message":"WasHere"}`, s.out.String())
}

func (s *TestLogSuite) TestWarningf() {
	grpclog.Warningf(s.format, s.args...)
	s.JSONEq(`{"level":"warn","message":"PhilipWasHere"}`, s.out.String())
}

func (s *TestLogSuite) TestWarningln() {
	grpclog.Warningln(s.args...)
	s.JSONEq(`{"level":"warn","message":"WasHere"}`, s.out.String())
}

// Note: gRPC Logger v2 deprecated Print methods to instead use Info methods.
func (s *TestLogSuite) TestPrint() {
	s.grpclog.Print(s.args...)
	s.JSONEq(`{"level":"info","message":"WasHere"}`, s.out.String())
}

func (s *TestLogSuite) TestPrintf() {
	s.grpclog.Printf(s.format, s.args...)
	s.JSONEq(`{"level":"info","message":"PhilipWasHere"}`, s.out.String())
}

func (s *TestLogSuite) TestPrintln() {
	s.grpclog.Println(s.args...)
	s.JSONEq(`{"level":"info","message":"WasHere"}`, s.out.String())
}

func (s *TestLogSuite) TestV() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	s.CheckV(0) // Check gRPC Info Logs

	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	s.CheckV(1) // Check gRPC Warn Logs

	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	s.CheckV(2) // Check gRPC Error Logs

	zerolog.SetGlobalLevel(zerolog.FatalLevel)
	s.CheckV(3) // Check gRPC Fatal Logs

	s.PanicsWithValue("unhandled gRPC logger level", func() {
		s.grpclog.V(4)
	})
}

func (s *TestLogSuite) CheckV(l int) {
	for i := 0; i <= l; i++ {
		s.Truef(s.grpclog.V(i), "level %v not enabled for %v", i, zerolog.GlobalLevel().String())
	}

	for i := l + 1; i < 4; i++ {
		s.Falsef(s.grpclog.V(i), "level %v enabled for %v", i, zerolog.GlobalLevel().String())
	}
}
