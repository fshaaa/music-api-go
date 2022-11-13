package songUsecase

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/repository/playlist_song_repository"
	"music-api-go/repository/song_repository"
	"testing"
)

func TestAddSong(t *testing.T) {
	t.Run("success add song", func(t *testing.T) {
		MockSong := model.Songs{
			ID:        uuid.NewString(),
			Title:     "Tomboy",
			Year:      2022,
			Performer: "(G)I-DLe",
			Genre:     "K-Pop",
			Duration:  180,
			Album_id:  uuid.NewString(),
		}

		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		song := NewSongUsecase(mockSongRepository, mockPlaylistSongRepository)

		mockSongRepository.On("AddSong", MockSong).Return(nil)

		err := song.AddSong(MockSong)
		assert.NoError(t, err)
		mockSongRepository.AssertExpectations(t)
	})
}

func TestGetAllSongs(t *testing.T) {
	t.Run("success get all songs", func(t *testing.T) {
		id := uuid.NewString()

		mockSongArr := []model.Songs{
			{
				ID:        id,
				Title:     "Tomboy",
				Year:      2022,
				Performer: "(G)I-DLe",
				Genre:     "K-Pop",
				Duration:  180,
				Album_id:  uuid.NewString(),
			},
		}

		mockSongRes := []dto.Song{
			{
				ID:        id,
				Title:     "Tomboy",
				Year:      2022,
				Performer: "(G)I-DLe",
				Genre:     "K-Pop",
				Duration:  180,
			},
		}

		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		song := NewSongUsecase(mockSongRepository, mockPlaylistSongRepository)

		mockSongRepository.On("GetAllSongs").Return(mockSongArr, nil)

		res, err := song.GetAllSongs()
		assert.NoError(t, err)
		assert.Equal(t, mockSongRes, res)
		mockSongRepository.AssertExpectations(t)
	})
}

func TestGetSongByID(t *testing.T) {
	t.Run("success get song by id", func(t *testing.T) {
		mockSong := model.Songs{
			ID:        uuid.NewString(),
			Title:     "Tomboy",
			Year:      2022,
			Performer: "(G)I-DLe",
			Genre:     "K-Pop",
			Duration:  180,
			Album_id:  uuid.NewString(),
		}

		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		song := NewSongUsecase(mockSongRepository, mockPlaylistSongRepository)

		mockSongRepository.On("GetSongById", mockSong.ID).Return(mockSong, nil)

		songs, err := song.GetSongById(mockSong.ID)
		assert.NoError(t, err)
		assert.Equal(t, songs, *mockSong.ToDTOSong())
		mockSongRepository.AssertExpectations(t)
	})
}

func TestUpdateSong(t *testing.T) {
	t.Run("success update song", func(t *testing.T) {
		id := uuid.NewString()
		album_id := uuid.NewString()

		mockSong := model.Songs{
			ID:        id,
			Title:     "Tomboy",
			Year:      2022,
			Performer: "(G)I-DLe",
			Genre:     "K-Pop",
			Duration:  180,
			Album_id:  album_id,
		}

		mockRes := map[string]interface{}{
			"title":     "Tomboy",
			"year":      2022,
			"performer": "(G)I-DLe",
			"genre":     "K-Pop",
			"duration":  180,
			"album_id":  album_id,
		}

		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		song := NewSongUsecase(mockSongRepository, mockPlaylistSongRepository)

		mockSongRepository.On("UpdateSong", id, mock.Anything, mock.Anything).Return(nil)

		songs, err := song.UpdateSong(id, mockSong)
		assert.NoError(t, err)
		assert.Equal(t, songs, mockRes)
		mockSongRepository.AssertExpectations(t)
	})
}

func TestDeleteSong(t *testing.T) {
	t.Run("success delete song", func(t *testing.T) {
		id := uuid.NewString()

		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		song := NewSongUsecase(mockSongRepository, mockPlaylistSongRepository)

		mockSongRepository.On("DeleteSong", id).Return(nil)

		err := song.DeleteSong(id)
		assert.NoError(t, err)
		mockSongRepository.AssertExpectations(t)
	})
}

func TestSearchSong(t *testing.T) {
	t.Run("success search song", func(t *testing.T) {
		id := uuid.NewString()
		album_id := uuid.NewString()

		mockSong := []model.Songs{
			{
				ID:        id,
				Title:     "Tomboy",
				Year:      2022,
				Performer: "(G)I-DLe",
				Genre:     "K-Pop",
				Duration:  180,
				Album_id:  album_id,
			},
		}

		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		song := NewSongUsecase(mockSongRepository, mockPlaylistSongRepository)

		mockSongRepository.On("SearchSong", mock.Anything).Return(mockSong, nil)

		_, err := song.SearchSong(mock.Anything)
		assert.NoError(t, err)
		mockSongRepository.AssertExpectations(t)
	})
}

func TestGetAllSongsInPlaylist(t *testing.T) {
	t.Run("success get all song in playlist", func(t *testing.T) {
		id := uuid.NewString()
		album_id := uuid.NewString()
		song_id := []string{id}

		mockSongArr := []dto.Song{
			{
				ID:        id,
				Title:     "Tomboy",
				Year:      2022,
				Performer: "(G)I-DLe",
				Genre:     "K-Pop",
				Duration:  180,
			},
		}

		mockSong := model.Songs{
			ID:        id,
			Title:     "Tomboy",
			Year:      2022,
			Performer: "(G)I-DLe",
			Genre:     "K-Pop",
			Duration:  180,
			Album_id:  album_id,
		}

		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		song := NewSongUsecase(mockSongRepository, mockPlaylistSongRepository)

		mockPlaylistSongRepository.On("GetAllSongID", mock.Anything).Return(song_id, nil)
		mockSongRepository.On("GetSongById", mock.Anything).Return(mockSong, nil)

		res, _, _, err := song.GetAllSongsInPlaylist(mock.Anything)
		assert.NoError(t, err)
		assert.Equal(t, mockSongArr, res)
		mockSongRepository.AssertExpectations(t)
	})
}
