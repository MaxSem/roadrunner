package jobs

import (
	"net"
	"net/rpc"
	"testing"

	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
	jobsv1beta "github.com/spiral/roadrunner/v2/proto/jobs/v1beta"
	"github.com/stretchr/testify/assert"
)

func resumePipes(pipes ...string) func(t *testing.T) {
	return func(t *testing.T) {
		conn, err := net.Dial("tcp", "127.0.0.1:6001")
		assert.NoError(t, err)
		client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

		pipe := &jobsv1beta.Pipelines{Pipelines: make([]string, len(pipes))}

		for i := 0; i < len(pipes); i++ {
			pipe.GetPipelines()[i] = pipes[i]
		}

		er := &jobsv1beta.Empty{}
		err = client.Call("jobs.Resume", pipe, er)
		assert.NoError(t, err)
	}
}

func pushToDisabledPipe(pipeline string) func(t *testing.T) {
	return func(t *testing.T) {
		conn, err := net.Dial("tcp", "127.0.0.1:6001")
		assert.NoError(t, err)
		client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

		req := &jobsv1beta.PushRequest{Job: &jobsv1beta.Job{
			Job:     "some/php/namespace",
			Id:      "1",
			Payload: `{"hello":"world"}`,
			Headers: nil,
			Options: &jobsv1beta.Options{
				Priority: 1,
				Pipeline: pipeline,
			},
		}}

		er := &jobsv1beta.Empty{}
		err = client.Call("jobs.Push", req, er)
		assert.Error(t, err)
	}
}

func pushToPipe(pipeline string) func(t *testing.T) {
	return func(t *testing.T) {
		conn, err := net.Dial("tcp", "127.0.0.1:6001")
		assert.NoError(t, err)
		client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

		req := &jobsv1beta.PushRequest{Job: &jobsv1beta.Job{
			Job:     "some/php/namespace",
			Id:      "1",
			Payload: `{"hello":"world"}`,
			Headers: map[string]*jobsv1beta.HeaderValue{"test": {Value: []string{"test2"}}},
			Options: &jobsv1beta.Options{
				Priority:   1,
				Pipeline:   pipeline,
				Delay:      0,
				Attempts:   0,
				RetryDelay: 0,
				Timeout:    0,
			},
		}}

		er := &jobsv1beta.Empty{}
		err = client.Call("jobs.Push", req, er)
		assert.NoError(t, err)
	}
}

func pausePipelines(pipes ...string) func(t *testing.T) {
	return func(t *testing.T) {
		conn, err := net.Dial("tcp", "127.0.0.1:6001")
		assert.NoError(t, err)
		client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

		pipe := &jobsv1beta.Pipelines{Pipelines: make([]string, len(pipes))}

		for i := 0; i < len(pipes); i++ {
			pipe.GetPipelines()[i] = pipes[i]
		}

		er := &jobsv1beta.Empty{}
		err = client.Call("jobs.Pause", pipe, er)
		assert.NoError(t, err)
	}
}

func destroyPipelines(pipes ...string) func(t *testing.T) {
	return func(t *testing.T) {
		conn, err := net.Dial("tcp", "127.0.0.1:6001")
		assert.NoError(t, err)
		client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

		pipe := &jobsv1beta.Pipelines{Pipelines: make([]string, len(pipes))}

		for i := 0; i < len(pipes); i++ {
			pipe.GetPipelines()[i] = pipes[i]
		}

		er := &jobsv1beta.Empty{}
		err = client.Call("jobs.Destroy", pipe, er)
		assert.NoError(t, err)
	}
}
