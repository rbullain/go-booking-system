package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-booking-system/internal/domain/booking/entity"
	"testing"
)

func TestDatabaseBookingRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)

	repo := DatabaseBookingRepository{db: db}

	user := &entity.UserEntity{
		Name:    "foo",
		Balance: 100,
	}

	mock.ExpectExec(createUserQuery).
		WithArgs(user.Name, user.Balance).
		WillReturnResult(sqlmock.NewResult(1, 1))

	createdUser, err := repo.CreateUser(user)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)

	assert.EqualValues(t, createdUser.ID, 1)
	assert.EqualValues(t, createdUser.Name, user.Name)
	assert.EqualValues(t, createdUser.Balance, user.Balance)
}

func TestDatabaseBookingRepository_GetUserByID_NoExist(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)

	repo := DatabaseBookingRepository{db: db}

	expectedUserID := int64(42)

	rows := sqlmock.NewRows([]string{"id", "name", "balance"})
	mock.ExpectQuery(getUserByIDQuery).
		WithArgs(expectedUserID).
		WillReturnRows(rows)

	user, err := repo.GetUserByID(expectedUserID)

	assert.Equal(t, sql.ErrNoRows, err)
	assert.Nil(t, user)
}

func TestDatabaseBookingRepository_GetUserByID_Exist(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)

	repo := DatabaseBookingRepository{db: db}

	expectedUser := entity.UserEntity{
		ID:      42,
		Name:    "foo",
		Balance: 100,
	}

	// Query result
	rows := sqlmock.NewRows([]string{"id", "name", "balance"}).
		AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Balance)
	mock.ExpectQuery(getUserByIDQuery).
		WithArgs(expectedUser.ID).
		WillReturnRows(rows)

	user, err := repo.GetUserByID(expectedUser.ID)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, user.ID, expectedUser.ID)
	assert.Equal(t, user.Name, expectedUser.Name)
	assert.EqualValues(t, user.Balance, expectedUser.Balance)
}

func TestDatabaseBookingRepository_CreateRoom(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)

	repo := DatabaseBookingRepository{db: db}

	room := &entity.RoomEntity{
		Name:  "foo",
		Price: 40,
	}

	mock.ExpectExec(createRoomQuery).
		WithArgs(room.Name, room.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	createdRoom, err := repo.CreateRoom(room)

	assert.NoError(t, err)
	assert.NotNil(t, createdRoom)

	assert.EqualValues(t, createdRoom.ID, 1)
	assert.EqualValues(t, createdRoom.Name, room.Name)
	assert.EqualValues(t, createdRoom.Price, room.Price)
}

func TestDatabaseBookingRepository_CreateReservation(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)

	repo := DatabaseBookingRepository{db: db}

	reservation := &entity.ReservationEntity{
		UserID: 1,
		RoomID: 1,
	}

	mock.ExpectExec(createReservationQuery).
		WithArgs(reservation.UserID, reservation.RoomID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	createdReservation, err := repo.CreateReservation(reservation)

	assert.NoError(t, err)
	assert.NotNil(t, createdReservation)

	assert.EqualValues(t, createdReservation.ID, 1)
	assert.EqualValues(t, createdReservation.UserID, reservation.UserID)
	assert.EqualValues(t, createdReservation.RoomID, reservation.RoomID)
}
