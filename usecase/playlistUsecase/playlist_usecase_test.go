package playlistUsecase

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/repository/collab"
	"music-api-go/repository/playlist_song_repository"
	"music-api-go/repository/playlists_repository"
	"music-api-go/repository/song_repository"
	"music-api-go/repository/users"
	"testing"
)

func TestAddPlaylist(t *testing.T) {
	t.Run("success add playlist", func(t *testing.T) {
		id := uuid.NewString()

		MockPlaylist := model.Playlists{
			ID:      id,
			Name:    "Lilac",
			User_id: uuid.NewString(),
		}

		mockPlaylistRepository := new(playlists_repository.MockPlaylistsRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		mockUserRepository := new(users.MockUserRepository)

		playlist := NewPlaylistUsecase(mockPlaylistRepository, mockPlaylistSongRepository, mockCollabRepository, mockSongRepository, mockUserRepository)

		mockPlaylistRepository.On("AddPlaylist", MockPlaylist).Return(nil)

		err := playlist.AddPlaylist(MockPlaylist)
		assert.NoError(t, err)
		mockSongRepository.AssertExpectations(t)
	})
}

func TestDeletePlaylist(t *testing.T) {
	t.Run("success delete playlist", func(t *testing.T) {
		id := uuid.NewString()

		mockPlaylistRepository := new(playlists_repository.MockPlaylistsRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		mockUserRepository := new(users.MockUserRepository)

		playlist := NewPlaylistUsecase(mockPlaylistRepository, mockPlaylistSongRepository, mockCollabRepository, mockSongRepository, mockUserRepository)

		mockPlaylistRepository.On("DeletePlaylist", id).Return(nil)

		err := playlist.DeletePlaylist(id)
		assert.NoError(t, err)
		mockSongRepository.AssertExpectations(t)
	})
}

func TestGetAllPlaylist(t *testing.T) {
	t.Run("success get all playlist", func(t *testing.T) {
		id := uuid.NewString()
		u_id := uuid.NewString()
		var mockdtoplArr []dto.Playlist

		mockPlaylistArr := []model.Playlists{
			{
				ID:      id,
				Name:    "Lilac",
				User_id: u_id,
			},
		}

		mockPlaylist := model.Playlists{
			ID:      id,
			Name:    "Lilac",
			User_id: u_id,
		}

		mockDTOplaylist := *mockPlaylist.ToDTOPlaylists()

		mockPlaylistRepository := new(playlists_repository.MockPlaylistsRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		mockUserRepository := new(users.MockUserRepository)

		playlist := NewPlaylistUsecase(mockPlaylistRepository, mockPlaylistSongRepository, mockCollabRepository, mockSongRepository, mockUserRepository)

		mockPlaylistRepository.On("GetAllPlaylists").Return(mockPlaylistArr, nil)
		mockPlaylistSongRepository.On("GetTotalSongs", id).Return(mockDTOplaylist.TotalSong, nil)
		mockPlaylistSongRepository.On("GetDurationPlaylist", id).Return(mockDTOplaylist.TotalDuration, nil)
		mockCollabRepository.On("GetAllUserID", id).Return([]string{id}, mockDTOplaylist.TotalUserSharing, nil)

		mockdtoplArr = append(mockdtoplArr, mockDTOplaylist)

		res, err := playlist.GetAllPlaylists()
		assert.NoError(t, err)
		assert.Equal(t, mockdtoplArr, res)
		mockSongRepository.AssertExpectations(t)
	})
}

func TestGetPlaylistByID(t *testing.T) {
	t.Run("success get playlist by id", func(t *testing.T) {
		id := uuid.NewString()
		u_id := uuid.NewString()

		mockPlaylist := model.Playlists{
			ID:      id,
			Name:    "Lilac",
			User_id: u_id,
		}

		mockDTOplaylist := *mockPlaylist.ToDTOPlaylists()

		mockPlaylistRepository := new(playlists_repository.MockPlaylistsRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		mockUserRepository := new(users.MockUserRepository)

		playlist := NewPlaylistUsecase(mockPlaylistRepository, mockPlaylistSongRepository, mockCollabRepository, mockSongRepository, mockUserRepository)

		mockPlaylistRepository.On("GetPlaylist", id).Return(mockPlaylist, nil)
		mockPlaylistSongRepository.On("GetTotalSongs", id).Return(mockDTOplaylist.TotalSong, nil)
		mockPlaylistSongRepository.On("GetDurationPlaylist", id).Return(mockDTOplaylist.TotalDuration, nil)
		mockCollabRepository.On("GetAllUserID", id).Return([]string{id}, mockDTOplaylist.TotalUserSharing, nil)

		res, err := playlist.GetPlaylistByID(id)
		assert.NoError(t, err)
		assert.Equal(t, mockDTOplaylist, res)
		mockSongRepository.AssertExpectations(t)
	})
}

func TestGetPlaylistDetail(t *testing.T) {
	t.Run("success get playlist detail", func(t *testing.T) {
		id := uuid.NewString()
		u_id := uuid.NewString()

		mockUser := model.Users{
			ID:       u_id,
			Username: "fathazhar",
			Email:    "fathazhar@gmail.com",
			Fullname: "Fath Azhar",
		}

		mockDTOuser := *mockUser.ToDTOUser()

		mockPlaylist := model.Playlists{
			ID:      id,
			Name:    "Lilac",
			User_id: u_id,
		}

		mockDTOplaylist := *mockPlaylist.ToDTOPlaylistDetails()
		songArr := make([]string, 0)
		userArr := make([]string, 0)

		mockPlaylistRepository := new(playlists_repository.MockPlaylistsRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		mockUserRepository := new(users.MockUserRepository)

		playlist := NewPlaylistUsecase(mockPlaylistRepository, mockPlaylistSongRepository, mockCollabRepository, mockSongRepository, mockUserRepository)

		mockPlaylistRepository.On("GetPlaylist", id).Return(mockPlaylist, nil)
		mockUserRepository.On("GetUserById", u_id).Return(mockUser, nil)
		mockPlaylistSongRepository.On("GetAllSongID", id).Return(songArr, nil)
		mockPlaylistSongRepository.On("GetTotalSongs", id).Return(mockDTOplaylist.TotalSong, nil)
		mockPlaylistSongRepository.On("GetDurationPlaylist", id).Return(mockDTOplaylist.TotalDuration, nil)
		mockCollabRepository.On("GetAllUserID", id).Return(userArr, mockDTOplaylist.TotalUserSharing, nil)

		mockDTOuserArr := []dto.User{}
		mockDTOuserArr = append(mockDTOuserArr, mockDTOuser)

		mockDTOplaylist.User = mockDTOuserArr

		res, err := playlist.GetPlaylistDetail(id)
		assert.NoError(t, err)
		assert.Equal(t, mockDTOplaylist, res)
		mockSongRepository.AssertExpectations(t)
	})
}

func TestGetPlaylistByUser(t *testing.T) {
	t.Run("success get playlist by user", func(t *testing.T) {
		id := uuid.NewString()
		u_id := uuid.NewString()
		mockDtoPlArr := make([]dto.Playlist, 0)

		mockPlaylistArr := []model.Playlists{
			{
				ID:      id,
				Name:    "Lilac",
				User_id: u_id,
			},
		}

		mockPlaylist := model.Playlists{
			ID:      id,
			Name:    "Lilac",
			User_id: u_id,
		}

		mockDTOplaylist := *mockPlaylist.ToDTOPlaylists()

		mockPlaylistRepository := new(playlists_repository.MockPlaylistsRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		mockSongRepository := new(song_repository.MockSongRepository)
		mockPlaylistSongRepository := new(playlist_song_repository.MockPlaylistSongRepository)
		mockUserRepository := new(users.MockUserRepository)

		playlist := NewPlaylistUsecase(mockPlaylistRepository, mockPlaylistSongRepository, mockCollabRepository, mockSongRepository, mockUserRepository)

		mockPlaylistRepository.On("GetPlaylistByUser", u_id).Return(mockPlaylistArr, nil)

		mockDtoPlArr = append(mockDtoPlArr, mockDTOplaylist)

		res, err := playlist.GetPlaylistByUser(u_id)
		assert.NoError(t, err)
		assert.Equal(t, mockDtoPlArr, res)
		mockSongRepository.AssertExpectations(t)
	})
}
