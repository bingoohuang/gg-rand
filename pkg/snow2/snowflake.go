// Package snow2 provides a very simple Twitter snowflake generator and parser.
package snow2

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/bingoohuang/gg/pkg/randx"
)

// Node12 每秒可以生成 256 个，可以有 4 个节点，可以使用 27.8 年，计算过程，见 cid_test.go
// 97656251 = 100000000000 >> 10 + 1
var Node12, _ = NewNode(WithEpochAdd(97656251), WithNodeBits(2), WithStepBits(8), WithTimeRound(time.Second))

// NodeUint32 creates and returns a unique snowflake ID for uint32
// CAUTION: only for low frequency usages，低频使用场景.
// 每秒可以生成 2^3 = 8 个，可以有 1 个节点，可以使用 17 年
// timestamp(29) + node(0) + step(3), 可用 2^29/60/60/24/365 ≈ 17 年
// max uint32 is 4294967295 (0b11111111_11111111_11111111_11111111)
var NodeUint32, _ = NewNode(WithNodeBits(0), WithStepBits(3), WithTimeRound(time.Second))

const (
	// NodeBits holds the number of bits to use for Node
	// Remember, you have a total 22 bits to share between Node/Step
	NodeBits uint8 = 10

	// StepBits holds the number of bits to use for Step
	// Remember, you have a total 22 bits to share between Node/Step
	StepBits uint8 = 12
)

type Config struct {
	Epoch     time.Time
	EpochAdd  int64
	TimeRound time.Duration
	NodeBits  uint8
	StepBits  uint8
	NodeID    int64
}

type ConfigFn func(*Config)

func WithTimeRound(n time.Duration) ConfigFn { return func(c *Config) { c.TimeRound = n } }
func WithNodeID(n int64) ConfigFn            { return func(c *Config) { c.NodeID = n } }
func WithEpochAdd(n int64) ConfigFn          { return func(c *Config) { c.EpochAdd = n } }
func WithEpoch(n time.Time) ConfigFn         { return func(c *Config) { c.Epoch = n } }
func WithStepBits(n uint8) ConfigFn          { return func(c *Config) { c.StepBits = n } }
func WithNodeBits(n uint8) ConfigFn          { return func(c *Config) { c.NodeBits = n } }

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	Config

	mu      sync.Mutex
	elapsed time.Duration
	node    int64
	step    int64

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8
	start     time.Time
}

// NewNode returns a new snowflake node that can be used to generate snowflake IDs
func NewNode(fns ...ConfigFn) (*Node, error) {
	epoch, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-02-10 00:11:25", time.Local)
	c := Config{
		Epoch:     epoch,
		TimeRound: time.Millisecond,
		NodeBits:  NodeBits,
		StepBits:  StepBits,
		NodeID:    0,
	}
	for _, fn := range fns {
		fn(&c)
	}

	n := Node{Config: c}
	n.node = c.NodeID
	n.nodeMax = -1 ^ (-1 << c.NodeBits)
	if n.node == 0 && n.nodeMax > 0 {
		n.node = int64(randx.IntN(int(n.nodeMax)))
	}

	n.nodeMask = n.nodeMax << c.StepBits
	n.stepMask = -1 ^ (-1 << c.StepBits)
	n.timeShift = c.NodeBits + c.StepBits
	n.nodeShift = c.StepBits

	if n.node < 0 || n.node > n.nodeMax {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(n.nodeMax, 10))
	}

	n.start = c.Epoch
	// n.step = int64(randx.IntN(int(n.stepMask)))

	return &n, nil
}

var Default, _ = NewNode()

func Next() int64 {
	return Default.Next()
}

// Next creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (n *Node) Next() int64 {
	elapsed := time.Since(n.start) / n.TimeRound
	if elapsed < n.elapsed { // 处理时间回拨（每次时间回拨时，不应回拨过多，例如1次回拨1秒）
		time.Sleep((n.elapsed - elapsed) * n.TimeRound)
		elapsed = time.Since(n.start) / n.TimeRound
	}

	n.mu.Lock()

	if elapsed == n.elapsed {
		if n.step = (n.step + 1) & n.stepMask; n.step == 0 {
			for elapsed <= n.elapsed {
				elapsed = time.Since(n.start) / n.TimeRound
			}
		}
	} else {
		n.step = 0
	}

	n.elapsed = elapsed
	n.mu.Unlock()

	tt := (int64(elapsed) + n.EpochAdd) << n.timeShift
	nn := n.node << n.nodeShift
	return tt | nn | n.step
}
