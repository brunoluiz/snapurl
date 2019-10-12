package service

import (
	"context"

	"github.com/brunoluiz/snapurl"
)

type Service struct{}

func (h *Service) Snapshot(ctx context.Context, req *snapurl.SnapshotRequest) (*snapurl.SnapshotResponse, error) {
	return &snapurl.SnapshotResponse{
		SnapshotUrl: "test",
	}, nil
}
