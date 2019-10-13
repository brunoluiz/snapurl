package service

import (
	"context"
	"time"

	"github.com/brunoluiz/snapurl"
	"github.com/brunoluiz/snapurl/api"
	"google.golang.org/genproto/googleapis/api/httpbody"
)

type Service struct{}

func (h *Service) Snapshot(ctx context.Context, req *api.SnapshotRequest) (*httpbody.HttpBody, error) {
	duration := int32(5) // In seconds
	if req.WaitPeriod != 0 {
		duration = req.WaitPeriod
	}

	buf, err := snapurl.Snap(context.Background(), req.Url, snapurl.Params{
		WaitPeriod: time.Duration(duration) * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return &httpbody.HttpBody{
		ContentType: "image/png",
		Data:        buf,
	}, nil
}
