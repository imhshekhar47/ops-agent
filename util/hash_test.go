package util

import (
	"testing"
	"time"

	"github.com/imhshekhar47/ops-agent/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetHash(t *testing.T) {
	agent := &pb.Agent{
		Meta: &pb.Metadata{
			Timestamp: timestamppb.New(time.Unix(0, 0)),
			Version:   "0.0.0",
		},
	}

	code := GetHash(agent)
	copyCode := GetHash(agent)

	if copyCode != code {
		t.Errorf("inconsistent hash values")
		t.Fail()
	}

	agent.Address = "localhost:5702"

	copyCode = GetHash(agent)

	if copyCode == code {
		t.Errorf("inconsistent hash values")
		t.Fail()
	}

}
