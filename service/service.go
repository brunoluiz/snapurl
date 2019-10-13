package service

import (
	"context"

	"github.com/brunoluiz/snapurl"
	"github.com/brunoluiz/snapurl/snapshot"
	"google.golang.org/genproto/googleapis/api/httpbody"
)

type Service struct{}

const chunkSize int = 64

func (h *Service) Snapshot(ctx context.Context, req *snapurl.SnapshotRequest) (*httpbody.HttpBody, error) {
	buf, err := snapshot.Snap(context.Background(), "https://google.co.uk", snapshot.Params{})
	if err != nil {
		return nil, err
	}

	return &httpbody.HttpBody{
		ContentType: "image/png",
		Data:        buf,
	}, nil
}
