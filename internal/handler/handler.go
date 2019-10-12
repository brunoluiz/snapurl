package handler

import (
	"context"

	"github.com/brunoluiz/snapurl"
)

type Handler struct{}

func (h *Handler) Snapshot(ctx context.Context, req *snapurl.SnapshotRequest) (*snapurl.SnapshotResponse, error) {
	return &snapurl.SnapshotResponse{
		SnapshotUrl: "test",
	}, nil
}
