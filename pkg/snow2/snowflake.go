// Package snow2 provides a very simple Twitter snowflake generator and parser.
package snow2

import (
	"errors"
	"github.com/bingoohuang/gg/pkg/randx"
	"strconv"
	"sync"
	"time"
)

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

	mu   sync.Mutex
	time time.Duration
	node int64
	step int64

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8
	start     time.Time
}

// NewNode returns a new snowflake node that can be used to generate snowflake IDs
func NewNode(fns ...ConfigFn) (*Node, error) {
	epoch, _ := time.Parse("2006-01-02", "2022-02-09")
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
	if n.node == 0 {
		n.node = int64(randx.IntN(int(n.nodeMax)))
	}

	n.nodeMask = n.nodeMax << c.StepBits
	n.stepMask = -1 ^ (-1 << c.StepBits)
	n.timeShift = c.NodeBits + c.StepBits
	n.nodeShift = c.StepBits

	if n.node < 0 || n.node > n.nodeMax {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(n.nodeMax, 10))
	}

	n.start = time.Now()
	n.step = int64(randx.IntN(int(n.stepMask)))

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
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Since(n.start) / n.TimeRound
	if now == n.time {
		if n.step = (n.step + 1) & n.stepMask; n.step == 0 {
			for now <= n.time {
				now = time.Since(n.start) / n.TimeRound
			}
		}
	} else {
		n.step = 0
	}

	n.time = now
	r := (int64(now)+n.EpochAdd)<<n.timeShift | (n.node << n.nodeShift) | (n.step)
	return r
}
