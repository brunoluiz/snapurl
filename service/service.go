package service

import (
	"context"

	"github.com/brunoluiz/snapurl"
	"github.com/brunoluiz/snapurl/snapshot"
	"google.golang.org/genproto/googleapis/api/httpbody"
)

type Service struct{}

func (h *Service) Snapshot(ctx context.Context, req *snapurl.SnapshotRequest) (*httpbody.HttpBody, error) {
	params := snapshot.Params{
		WaitPeriod: req.WaitPeriod,
	}

	buf, err := snapshot.Snap(context.Background(), req.Url, params)
	if err != nil {
		return nil, err
	}

	return &httpbody.HttpBody{
		ContentType: "image/png",
		Data:        buf,
	}, nil
}
