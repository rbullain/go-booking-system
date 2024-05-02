package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"go-booking-system/internal/api/dto"
	"go-booking-system/internal/rabbitmq"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockBookingController struct {
}

func (m MockBookingController) GetUserByID(ctx *gin.Context) (*dto.UserRetrieveResponseDTO, error) {
	return nil, nil
}

func (m MockBookingController) CreateUser(ctx *gin.Context) (*dto.UserCreateResponseDTO, error) {
	return &dto.UserCreateResponseDTO{
		ID:      42,
		Name:    "foo",
		Balance: 100,
	}, nil
}

func (m MockBookingController) GetRoomByID(ctx *gin.Context) (*dto.RoomRetrieveResponseDTO, error) {
	return nil, nil
}

func (m MockBookingController) CreateRoom(ctx *gin.Context) (*dto.RoomCreateResponseDTO, error) {
	return nil, nil
}

func (m MockBookingController) GetReservationByID(ctx *gin.Context) (*dto.ReservationRetrieveResponseDTO, error) {
	return nil, nil
}

func (m MockBookingController) CreateReservation(ctx *gin.Context) (*dto.ReservationCreateResponseDTO, error) {
	return nil, nil
}

type MockBookingRabbitMQ struct {
}

func (m MockBookingRabbitMQ) PublishOnQueue(payload rabbitmq.IRabbitMQPayload, queueName string) error {
	return nil
}

func (m MockBookingRabbitMQ) Subscribe(consumerName string, handler func(amqp.Delivery)) chan error {
	return nil
}

func (m MockBookingRabbitMQ) Close() {
}

func TestApplication_CreateUser(t *testing.T) {
	app := Application{
		bookingController: MockBookingController{},
		rabbitmqClient:    MockBookingRabbitMQ{},
	}

	req, _ := http.NewRequest("POST", "/user", nil)
	resp := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(resp)
	ctx.Request = req

	app.CreateUser(ctx)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var userDTO dto.UserCreateResponseDTO
	err := json.Unmarshal(resp.Body.Bytes(), &userDTO)
	assert.NoError(t, err)

	assert.EqualValues(t, userDTO.ID, 42)
	assert.EqualValues(t, userDTO.Name, "foo")
	assert.EqualValues(t, userDTO.Balance, 100)
}
