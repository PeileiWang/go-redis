package redis

import (
	"context"
	"github.com/go-redis/redis/v8/internal/pool"
)

type ProxyOptions struct {
	Addrs []string
	*Options
}

type ProxyClient struct {
	*Client
	multiServerPool *MultiServerPool
}

func (c *ProxyClient) getConn(ctx context.Context) (*pool.Conn, error) {
	return c.multiServerPool.GetConn(ctx)
}

func (c *ProxyClient) releaseConn(ctx context.Context, cn *pool.Conn, err error) {
	c.multiServerPool.ReleaseConn(ctx, cn, err)
}

func NewProxyClient(opt *ProxyOptions) *ProxyClient {
	opt.init()

	serverList := opt.Addrs
	if len(serverList) == 0 {
		//todo
	}

	serverCh := make(chan []string, 1)
	serverCh <- serverList

	c := &ProxyClient{
		Client:          NewClient(opt.Options),
		multiServerPool: NewMultiServerPool(serverList, serverCh, opt),
	}

	c.GetConnAddition(c.getConn)
	c.ReleaseConnAddition(c.releaseConn)

	return c
}
