package redis

import (
	"context"
	"github.com/go-redis/redis/v8/internal/pool"
	"sync"
	"time"
)

type MultiServerPool struct {
	serverList  []string
	servPoolMap map[string]*ServPool
	ch          chan []string // update server list
	cursor      uint32
	sync.RWMutex
}

type ServPool struct {
	connPool pool.Pooler
	servOpt  *ServOpt
	isDead   bool
}

type ServOpt struct {
	network     string
	dialTimeout time.Duration
	serv        string
}

func NewMultiServerPool(serverList []string, ch chan []string, opt *ProxyOptions) *MultiServerPool {
	// todo
	return nil
}

func (m *MultiServerPool) GetConn(ctx context.Context) (*pool.Conn, error) {
	return nil, nil
}

func (m *MultiServerPool) ReleaseConn(ctx context.Context, cn *pool.Conn, err error) {

}
