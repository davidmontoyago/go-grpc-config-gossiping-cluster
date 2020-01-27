package cluster

import (
	"bytes"
	"encoding/gob"
	"log"
	"sync"
)

// ConfigStore is the Delegate invoked by memberlist during gossiping to sync config across members
type ConfigStore struct {
	mu sync.Mutex

	// useful to share node details with other nodes
	metadata map[string]string

	// node internal state - this is the actual config being gossiped
	config map[string]string
}

func newConfigStore(md map[string]string) *ConfigStore {
	return &ConfigStore{
		metadata: md,
		config:   make(map[string]string),
	}
}

// NodeMeta is used to retrieve meta-data about the current node
// when broadcasting an alive message. It's length is limited to
// the given byte size. This metadata is available in the Node structure.
func (c *ConfigStore) NodeMeta(limit int) []byte {
	c.mu.Lock()
	defer c.mu.Unlock()

	var network bytes.Buffer
	encoder := gob.NewEncoder(&network)
	err := encoder.Encode(c.metadata)
	if err != nil {
		log.Fatal("failed to encode metadata", err)
	}
	return network.Bytes()
}

// NotifyMsg is called when a user-data message is received.
// Care should be taken that this method does not block, since doing
// so would block the entire UDP packet receive loop. Additionally, the byte
// slice may be modified after the call returns, so it should be copied if needed
func (c *ConfigStore) NotifyMsg(b []byte) {
	// not expecting messages - push/pull sync should suffice
}

// GetBroadcasts is called when user data messages can be broadcast.
// It can return a list of buffers to send. Each buffer should assume an
// overhead as provided with a limit on the total byte size allowed.
// The total byte size of the resulting data to send must not exceed
// the limit. Care should be taken that this method does not block,
// since doing so would block the entire UDP packet receive loop.
func (c *ConfigStore) GetBroadcasts(overhead, limit int) [][]byte {
	// nothing to broadcast
	return nil
}

// LocalState is used for a TCP Push/Pull. This is sent to
// the remote side in addition to the membership information. Any
// data can be sent here. See MergeRemoteState as well. The `join`
// boolean indicates this is for a join instead of a push/pull.
func (c *ConfigStore) LocalState(join bool) []byte {
	c.mu.Lock()
	defer c.mu.Unlock()

	var network bytes.Buffer
	encoder := gob.NewEncoder(&network)
	err := encoder.Encode(c.config)
	if err != nil {
		log.Fatal("failed to encode local state", err)
	}
	return network.Bytes()
}

// MergeRemoteState is invoked after a TCP Push/Pull. This is the
// state received from the remote side and is the result of the
// remote side's LocalState call. The 'join'
// boolean indicates this is for a join instead of a push/pull.
func (c *ConfigStore) MergeRemoteState(buf []byte, join bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	network := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(network)
	config := make(map[string]string)
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal("failed to decode remote state", err)
	}

	for key, value := range config {
		if c.config[key] != value {
			log.Printf("updating config %s=%v", key, value)
			c.config[key] = value
		}
	}
	log.Println("successfully merged remote state.")
}

// Put adds config property to config store
func (c *ConfigStore) Put(key string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.config[key] = value
}

// Get returns a property value
func (c *ConfigStore) Get(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.config[key]
}
