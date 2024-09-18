package repository

import (
	"movielibrary/internal/domain"
	ws "movielibrary/internal/domain/enums/watchstatus"
	"movielibrary/internal/domain/models"
	"movielibrary/internal/domain/types"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MovieLibraryRepository struct {
	db *gorm.DB
}

func NewMovieLibraryRepository(client *gorm.DB) domain.MovieLibraryRepo {
	return MovieLibraryRepository{
		db: client,
	}
}

func (r MovieLibraryRepository) UpsertMovie(movie types.Movie) types.Result[models.Movie] {
	m := models.Movie{
		Title:  movie.Title,
		Status: int(movie.Status),
		Year:   movie.Year,
		ImdbID: movie.ImdbID.Get(),
		Liked:  &movie.Liked,
	}

	result := r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		UpdateAll: false,
		DoUpdates: clause.AssignmentColumns([]string{"watched"}),
	}).Create(&m)

	return *types.NewResult(m).SetErr(result.Error)
}

func (r MovieLibraryRepository) GetMovie(movieID int) types.Result[models.Movie] {
	m := models.Movie{ID: movieID}
	query := r.db.First(&m)
	return *types.NewResult(m).SetErr(query.Error)
}

func (r MovieLibraryRepository) GetMovies(filter ws.WatchStatus) types.Result[[]models.Movie] {
	ms := make([]models.Movie, 0)
	query := r.db.Find(&ms, "status = ?", int(filter))
	return *types.NewResult(ms).SetErr(query.Error)
}

func (r MovieLibraryRepository) GetAllMovies() types.Result[[]models.Movie] {
	ms := make([]models.Movie, 0)
	query := r.db.Find(&ms)
	return *types.NewResult(ms).SetErr(query.Error)
}

func (r MovieLibraryRepository) RemoveMovie(movieID int) error {
	m := models.Movie{ID: movieID}
	query := r.db.Delete(&m)
	return query.Error
}
