package sse

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
)

type clientConn struct {
	projectID string
	ch        chan []byte
}

type EventServer struct {
	context       context.Context
	cancel        context.CancelFunc
	ConnectClient chan clientConn
	CloseClient   chan clientConn
	clients       map[string]map[chan []byte]struct{} // projectID -> set of client channels

	SrvWg  sync.WaitGroup
	reqWg  sync.WaitGroup
	logger *slog.Logger
	mu     sync.Mutex
}

func NewSSEServer(parentCtx context.Context, logger *slog.Logger) *EventServer {
	ctx, cancel := context.WithCancel(parentCtx)
	return &EventServer{
		context:       ctx,
		cancel:        cancel,
		ConnectClient: make(chan clientConn),
		CloseClient:   make(chan clientConn),
		clients:       make(map[string]map[chan []byte]struct{}),
		logger:        logger,
	}
}

func (sseServer *EventServer) Run() {
	sseServer.logger.Info("Starting SSE server")
	defer sseServer.SrvWg.Done()

	for {
		select {
		case <-sseServer.context.Done():
			sseServer.logger.Info("SSE server shutdown", "err", sseServer.context.Err())
			sseServer.reqWg.Wait()
			return

		case conn := <-sseServer.ConnectClient:
			sseServer.mu.Lock()
			if sseServer.clients[conn.projectID] == nil {
				sseServer.clients[conn.projectID] = make(map[chan []byte]struct{})
			}
			sseServer.clients[conn.projectID][conn.ch] = struct{}{}
			sseServer.logger.Debug("client connected", "projectID", conn.projectID)
			sseServer.mu.Unlock()

		case conn := <-sseServer.CloseClient:
			sseServer.mu.Lock()
			if projectClients, ok := sseServer.clients[conn.projectID]; ok {
				delete(projectClients, conn.ch)
				if len(projectClients) == 0 {
					delete(sseServer.clients, conn.projectID)
				}
			}
			sseServer.logger.Debug("client disconnected", "projectID", conn.projectID)
			sseServer.mu.Unlock()
		}
	}
}

func (sseServer *EventServer) Stop() {
	sseServer.logger.Info("Waiting for SSE server routines to finish")
	for projectID, projectClients := range sseServer.clients {
		for ch := range projectClients {
			close(ch)
		}
		delete(sseServer.clients, projectID)
	}
	sseServer.SrvWg.Wait()
	sseServer.cancel()
}

func (sseServer *EventServer) HandleConnection(c fiber.Ctx) error {
	projectID := c.Params("projectId")
	if projectID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "projectId is required"})
	}

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("X-Accel-Buffering", "no")

	message := make(chan []byte)
	conn := clientConn{projectID: projectID, ch: message}
	sseServer.ConnectClient <- conn

	keepAliveTicker := time.NewTicker(15 * time.Second)
	keepAliveMsg := []byte(":keepalive\n\n")

	sseServer.logger.Debug("SSE client connected", "projectID", projectID)

	c.SendStreamWriter(func(w *bufio.Writer) {
		sseServer.reqWg.Add(1)
		defer sseServer.reqWg.Done()
		defer func() { sseServer.CloseClient <- conn }()
		defer keepAliveTicker.Stop()

		for {
			select {
			case msg, ok := <-message:
				if !ok {
					return
				}
				_, err := fmt.Fprintf(w, "data: %s\n\n", msg)
				if err != nil {
					log.Printf("error writing message: %v", err)
					return
				}
				w.Flush()
			case <-keepAliveTicker.C:
				_, err := w.Write(keepAliveMsg)
				if err != nil {
					sseServer.logger.Debug("error writing keepalive", "err", err)
					return
				}
				w.Flush()
			case <-sseServer.context.Done():
				return
			}
		}
	})
	return nil
}

func (s *EventServer) BroadcastToProject(projectID string, msg []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	projectClients, ok := s.clients[projectID]
	if !ok {
		return
	}
	for client := range projectClients {
		select {
		case client <- msg:
		default:
		}
	}
}
