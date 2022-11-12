package playlistActivityUsecase

import (
	"music-api-go/model"
	"music-api-go/repository/playlistActivitiesRepository"
)

type PlaylistActivityUsecase interface {
	AddPlaylistActivity(activity model.PlaylistActivities) error
	DeletePlaylistActivity(id string) error
}

type playlistActivityUsecase struct {
	playlistActivity playlistActivitiesRepository.PlaylistActivitiesRepository
}

func NewPlaylistActivityUsecase(pa playlistActivitiesRepository.PlaylistActivitiesRepository) *playlistActivityUsecase {
	return &playlistActivityUsecase{pa}
}

func (p *playlistActivityUsecase) AddPlaylistActivity(activity model.PlaylistActivities) error {
	return p.playlistActivity.AddPlaylistActivity(activity)
}

func (p *playlistActivityUsecase) DeletePlaylistActivity(id string) error {
	return p.playlistActivity.DeletePlaylistActivity(id)
}
