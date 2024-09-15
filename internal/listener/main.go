package listener

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quinn71/todolist/internal/config"
)

type Service struct {
	addr string
}

func (s *Service) Run(ctx context.Context) error {
	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pending": []string{
				"Buy groceries",
				"Complete the report",
				"Clean the living room and vacuum the carpets",
				"Daily running",
			},
		})
	})

	fmt.Printf("Started listener on %s", s.addr)
	if err := r.Run(s.addr); err != nil {
		return err
	}

	return nil
}

func New(cfg *config.Config) *Service {
	return &Service{addr: cfg.Listener.Addr}
}
