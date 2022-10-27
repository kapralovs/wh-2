package local

import (
	"sync"

	"github.com/kapralovs/wh-2/internal/models"
)

type LocalRepo struct {
	mu    *sync.Mutex
	users map[int64]*models.User
}

func NewRepo() *LocalRepo {
	return &LocalRepo{
		mu:    new(sync.Mutex),
		users: make(map[int64]*models.User),
	}
}
