// Code generated by Boiler; YOU MUST CHANGE THIS.

package dummy_service

import (
	"context"
	desc "github.com/ushakovn/ushakovn-org/loghouse-test/internal/pb/loghouse_test"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetDummy implementation stub. Change this.
func (s *DummyService) GetDummy(ctx context.Context, req *desc.GetDummyRequest) (*desc.GetDummyResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetDummy not implemented")
}