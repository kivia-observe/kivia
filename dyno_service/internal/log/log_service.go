package log

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	logger "log"

	amqp "github.com/rabbitmq/amqp091-go"
	apikey "github.com/winnerx0/dyno/internal/api_key"
	"github.com/winnerx0/dyno/internal/rabbitmq"
)

type Logservice struct {
	repo           *Repository
	apiKeyRepo     *apikey.Repository
	rabbitMQClient *rabbitmq.RabbitMQClient
}

func NewLogService(repo *Repository, apiKeyRepo *apikey.Repository, rabbitMQClient *rabbitmq.RabbitMQClient) *Logservice {
	return &Logservice{
		repo:           repo,
		apiKeyRepo: apiKeyRepo,
		rabbitMQClient: rabbitMQClient,
	}
}

func (s Logservice) CreateLog(createLogRequest createLogRequest, apiKey string) error {
	
	queue, err := s.rabbitMQClient.Channel.QueueDeclare("log_queue", true, false, false, false, nil)

	if err != nil {
		return err
	}

	apiKeyHash := sha256.Sum256([]byte(apiKey))

	apiKeyHex := hex.EncodeToString(apiKeyHash[:])

	projectId, err := s.apiKeyRepo.FindProjectIdByKey(apiKeyHex)

	if err != nil {
		return err
	}

	latency := fmt.Sprintf("%d ms", createLogRequest.Latency)

	fmt.Println("latency", latency)

	log := Log{
		Id:        createLogRequest.Id,
		Path:      createLogRequest.Path,
		IPAddress: createLogRequest.IPAddress,
		Status:    createLogRequest.Status,
		Timestamp: createLogRequest.Timestamp,
		Latency:   latency,
		ProjectId: projectId,
	}

	logBytes, err := json.Marshal(log)

	if err != nil {
		return err
	}

	err = s.rabbitMQClient.Channel.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         logBytes,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s Logservice) GetLogsByProjectId(projectId string, startDate *string, endDate *string) ([]Log, error) {

	return s.repo.GetLogsByProjectId(projectId, startDate, endDate)
}

func (s Logservice) LogConsumer(ctx context.Context) error {

	msgs, err := s.rabbitMQClient.ConsumeRabbitMQQueue("log_queue")
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-msgs:

				log.Println("received log")
				if !ok {
					return
				}

				var log Log
				if err := json.Unmarshal(d.Body, &log); err != nil {
					logger.Println("Invalid log format:", err)
					continue
				}

				if err := s.repo.Save(log); err != nil {
					logger.Println("Failed to save log:", err)
				}
			}
		}
	}()

	return nil

}
