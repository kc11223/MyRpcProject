package pool

import (
	"net"
	"time"
)

// const
const (
	DialTimeout = 5 * time.Second

	BackoffMaxDelay = 3 * time.Second

	KeepAliveTime = time.Duration(10) * time.Second

	KeepAliveTimeout = time.Duration(3) * time.Second

	InitialWindowSize = 1 << 30

	InitialConnWindowSize = 1 << 30

	MaxSendMsgSize = 4 << 30

	MaxRecvMsgSize = 4 << 30
)

// PoolOptions are params for creating grpc connect pool.
type PoolOptions struct {
	adrr        string
	newConnFunc func(string) (net.Conn, error)

	minConns  int // 最小连接数
	maxConns  int // 最大连接数
	maxIdle   int // 最大空闲连接数
	poolNum   int // connections大小
	shrinknum int // 缩容数目

	indexFreq         time.Duration // 获取index超时时间
	connectionTimeout time.Duration // 连接超时时间
	idleTimeout       time.Duration // 空闲连接超时时间
	keepAliveInterval time.Duration // 保活检查时间

	reuse bool
}

// DefaultPoolOptions sets a list of recommended options for good performance.
var DefaultPoolOptions = PoolOptions{
	minConns:          5,
	maxConns:          50,
	maxIdle:           20,
	poolNum:           5,
	shrinknum:         20,
	connectionTimeout: 30 * time.Second, // 连接超时时间
	idleTimeout:       30 * time.Second, // 空闲连接超时时间
	keepAliveInterval: 30 * time.Second, // 保活检查时间
	reuse:             true,
}