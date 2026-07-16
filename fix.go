package memberlist

import "sync"

// FixedMemberlist wraps Memberlist with thread-safe operations
type FixedMemberlist struct {
    mu       sync.RWMutex
    nodeMap  map[string]*Node
}

// GetNode safely retrieves a node by name
func (m *FixedMemberlist) GetNode(name string) *Node {
    m.mu.RLock()
    defer m.mu.RUnlock()
    return m.nodeMap[name]
}

// SetNode safely adds/updates a node
func (m *FixedMemberlist) SetNode(name string, node *Node) {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.nodeMap[name] = node
}

// RemoveNode safely removes a node
func (m *FixedMemberlist) RemoveNode(name string) {
    m.mu.Lock()
    defer m.mu.Unlock()
    delete(m.nodeMap, name)
}

// IsAlive checks if a node is alive with proper locking
func (m *FixedMemberlist) IsAlive(name string) bool {
    m.mu.RLock()
    defer m.mu.RUnlock()
    node, ok := m.nodeMap[name]
    return ok && node != nil && node.State == StateAlive
}
