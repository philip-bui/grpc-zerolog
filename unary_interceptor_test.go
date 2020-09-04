package zerolog

import (
	"bytes"
	"testing"

	"github.com/philip-bui/grpc-zerolog/test"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestUnaryInterceptorSuite struct {
	suite.Suite
	out    *bytes.Buffer
	log    zerolog.Logger
	client *test.TestClient
}

func TestUnaryServerInterceptor(t *testing.T) {
	assert.NotNil(t, UnaryInterceptor())
}

func (s *TestUnaryInterceptorSuite) SetupSuite() {
	s.out = &bytes.Buffer{}
	s.log = zerolog.New(s.out)
	go func() {
		test.StartServer(NewUnaryServerInterceptorWithLogger(&s.log))
	}()
	s.client = test.GetClient()
}

func (s *TestUnaryInterceptorSuite) SetupTest() {
	s.out = &bytes.Buffer{}
	s.log = zerolog.New(s.out)
}

func TestUnaryInterceptor(t *testing.T) {
	suite.Run(t, new(TestUnaryInterceptorSuite))
}

func (s *TestUnaryInterceptorSuite) TestUnaryClientInterceptor() {
	resp, err := s.client.SendReq()
	s.NoError(err, "Expected no errors")
	s.Equal(s.client.ExampleReq.Test, resp.Test, "Expected request and response to be the same")
	s.NotEmpty(s.out.String())
}

func (s *TestUnaryInterceptorSuite) TestUnaryClientInterceptorError() {
	resp, err := s.client.SendErr()
	s.Error(err)
	s.Empty(resp)

	s.NotEmpty(s.out.String())
}
