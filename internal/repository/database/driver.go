package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/shanto-323/rely/model"
	"github.com/shanto-323/rely/model/dto"
	"github.com/shanto-323/rely/model/entity"
)

// Driver is an interface for database.
// It contains all methods that database should implement.
type Driver interface {
	// Database specific methods
	Ping(ctx context.Context) error
	IsInitialized(ctx context.Context) bool
	Close() error

	// Other methods related to database operation
	GetStudentByStudentID(ctx context.Context, studentId int) (*entity.Student, error)

	StudentAttendanceOverview(ctx context.Context, id uuid.UUID) (*dto.StudentAttendanceOverview, error)
	StudentsAttendanceOverview(ctx context.Context, paginate *dto.PaginationDto) (*model.PaginatedResponse[dto.StudentsOverview], error)
}
