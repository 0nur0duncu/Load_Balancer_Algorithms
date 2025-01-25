package main

import (
	"sync"
	"testing"
)

type LoadBalancer struct {
	servers      []string
	currentIndex int
	mu           sync.Mutex
}

// NewLoadBalancer creates a new LoadBalancer instance from the given list of servers.
//
// The input slice is copied, so it can be modified independently of the LoadBalancer.
func NewLoadBalancer(servers []string) *LoadBalancer {
	copiedServers := make([]string, len(servers))
	copy(copiedServers, servers)
	return &LoadBalancer{
		servers:      copiedServers,
		currentIndex: 0,
	}
}

// GetNextServer returns the next server in the round-robin sequence.
// It locks the LoadBalancer to ensure thread safety and updates the current index
// to maintain the order of server selection.

func (lb *LoadBalancer) GetNextServer() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	server := lb.servers[lb.currentIndex]
	lb.currentIndex = (lb.currentIndex + 1) % len(lb.servers)
	return server
}

// TestRoundRobin verifies that the LoadBalancer correctly rotates through the servers
// in a round-robin sequence.
func TestRoundRobin(t *testing.T) {
	serverList := []string{"Server1", "Server2", "Server3"}
	lb := NewLoadBalancer(serverList)

	expected := []string{"Server1", "Server2", "Server3", "Server1", "Server2"}
	for i := 0; i < 5; i++ {
		server := lb.GetNextServer()
		if server != expected[i] {
			t.Errorf("Expected %s, got %s at iteration %d", expected[i], server, i)
		}
	}
}
