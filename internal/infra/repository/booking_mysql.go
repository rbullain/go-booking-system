package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-booking-system/internal/domain/booking/entity"
)

var _ entity.BookingRepository = DatabaseBookingRepository{}

const (
	getUserByIDQuery = "SELECT id, name, balance FROM user WHERE id = ?"
	createUserQuery  = "INSERT INTO user (name, balance) VALUES (?, ?)"

	getRoomByIDQuery = "SELECT id, name, price FROM room WHERE id = ?"
	createRoomQuery  = "INSERT INTO room (name, price) VALUES (?, ?)"

	getReservationByIDQuery = "SELECT (id, user_id, room_id) FROM reservation WHERE id = ?"
	createReservationQuery  = "INSERT INTO reservation (user_id, room_id) VALUES (?, ?)"
)

type DatabaseBookingRepository struct {
	db *sql.DB
}

func connect(username, password, host, port, database string) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewDatabaseBookingRepository(username, password, host, port, database string) *DatabaseBookingRepository {
	db, err := connect(username, password, host, port, database)
	if err != nil {
		panic(err)
	}

	return &DatabaseBookingRepository{
		db: db,
	}
}

func (repo DatabaseBookingRepository) CreateReservation(reservation *entity.ReservationEntity) (*entity.ReservationEntity, error) {
	result, err := repo.db.Exec(createReservationQuery, reservation.UserID, reservation.RoomID)
	if err != nil {
		return nil, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createdReservation := &entity.ReservationEntity{
		ID:     insertedID,
		UserID: reservation.UserID,
		RoomID: reservation.RoomID,
	}
	return createdReservation, nil
}

func (repo DatabaseBookingRepository) GetReservationByID(id int64) (*entity.ReservationEntity, error) {
	row := repo.db.QueryRow(getReservationByIDQuery, id)

	var reservation entity.ReservationEntity

	err := row.Scan(&reservation.ID, &reservation.UserID, &reservation.RoomID)
	if err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (repo DatabaseBookingRepository) CreateUser(user *entity.UserEntity) (*entity.UserEntity, error) {
	result, err := repo.db.Exec(createUserQuery, user.Name, user.Balance)
	if err != nil {
		return nil, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createdUser := &entity.UserEntity{
		ID:      insertedID,
		Name:    user.Name,
		Balance: user.Balance,
	}
	return createdUser, nil
}

func (repo DatabaseBookingRepository) GetUserByID(id int64) (*entity.UserEntity, error) {
	row := repo.db.QueryRow(getUserByIDQuery, id)

	var user entity.UserEntity

	err := row.Scan(&user.ID, &user.Name, &user.Balance)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo DatabaseBookingRepository) CreateRoom(room *entity.RoomEntity) (*entity.RoomEntity, error) {
	result, err := repo.db.Exec(createRoomQuery, room.Name, room.Price)
	if err != nil {
		return nil, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createdRoom := &entity.RoomEntity{
		ID:    insertedID,
		Name:  room.Name,
		Price: room.Price,
	}
	return createdRoom, nil
}

func (repo DatabaseBookingRepository) GetRoomByID(id int64) (*entity.RoomEntity, error) {
	row := repo.db.QueryRow(getRoomByIDQuery, id)

	var room entity.RoomEntity

	err := row.Scan(&room.ID, &room.Name, &room.Price)
	if err != nil {
		return nil, err
	}

	return &room, nil
}
