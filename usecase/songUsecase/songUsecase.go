package songUsecase

import (
	"fmt"
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/repository/playlistSongsRepository"
	"music-api-go/repository/songRepository"
)

type SongUsecase interface {
	GetAllSongsInPlaylist(playlist_id string) ([]dto.Song, int, int, error)
	GetAllSongs() ([]dto.Song, error)
	GetSongByID(id string) (dto.Song, error)
	AddSong(song model.Songs) error
	UpdateSong(id string, song model.Songs) (map[string]interface{}, error)
	DeleteSong(id string) error
	SearchSong(title string) ([]dto.Song, error)
}

type songUsecase struct {
	song         songRepository.SongRepository
	playlistSong playlistSongsRepository.PlaylistSongsRepository
}

func NewSongUsecase(song songRepository.SongRepository, playlistSong playlistSongsRepository.PlaylistSongsRepository) *songUsecase {
	return &songUsecase{song, playlistSong}
}

func (s *songUsecase) GetAllSongsInPlaylist(playlist_id string) ([]dto.Song, int, int, error) {
	var songs []dto.Song
	var totalSong = 0
	var totalDuration = 0

	song_id, err := s.playlistSong.GetAllSongID(playlist_id)
	if err != nil {
		return nil, 0, 0, err
	}
	for _, id := range song_id {
		songModel, err := s.song.GetSongById(id)
		if err != nil {
			return nil, 0, 0, err
		}
		song := *songModel.ToDTOSong()
		songs = append(songs, song)
		totalSong++
		totalDuration += song.Duration
	}
	fmt.Println(songs)
	return songs, totalSong, totalDuration, nil
}

func (s *songUsecase) GetAllSongs() ([]dto.Song, error) {
	var songs []dto.Song
	songsModel, err := s.song.GetAllSongs()
	if err != nil {
		return nil, err
	}
	for _, songModel := range songsModel {
		song := *songModel.ToDTOSong()
		songs = append(songs, song)
	}
	return songs, nil
}

func (s *songUsecase) GetSongByID(id string) (dto.Song, error) {
	songModel, err := s.song.GetSongById(id)
	if err != nil {
		return dto.Song{}, err
	}
	return *songModel.ToDTOSong(), nil
}

func (s *songUsecase) AddSong(song model.Songs) error {
	return s.song.AddSong(song)
}

func (s *songUsecase) UpdateSong(id string, song model.Songs) (map[string]interface{}, error) {
	update := make(map[string]interface{})
	req := map[string]interface{}{
		"title":     song.Title,
		"year":      song.Year,
		"performer": song.Performer,
		"genre":     song.Genre,
		"duration":  song.Duration,
		"album_id":  song.Album_id,
	}
	var err error
	for key, value := range req {
		if value != nil {
			err = s.song.UpdateSong(id, key, value)
			update[key] = value
		}
	}
	return update, err
}

func (s *songUsecase) DeleteSong(id string) error {
	return s.song.DeleteSong(id)
}

func (s *songUsecase) SearchSong(title string) ([]dto.Song, error) {
	var songs []dto.Song
	songsModel, err := s.song.SearchSong(title)
	if err != nil {
		return nil, err
	}
	for _, songModel := range songsModel {
		song := *songModel.ToDTOSong()
		songs = append(songs, song)
	}
	return songs, nil
}
