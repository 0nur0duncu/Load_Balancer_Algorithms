package main

import (
	"sync"
	"testing"
)

type Server struct {
	Name   string
	Weight int
}

type LoadBalancer struct {
	servers      []string   // Genişletilmiş sunucu listesi
	currentIndex int        // Mevcut indeks
	mu           sync.Mutex // Thread safety için mutex
}

func NewLoadBalancer(servers []Server) *LoadBalancer {
	expanded := make([]string, 0)

	// Ağırlıklara göre sunucu listesini genişlet
	for _, s := range servers {
		for i := 0; i < s.Weight; i++ {
			expanded = append(expanded, s.Name)
		}
	}

	return &LoadBalancer{
		servers:      expanded,
		currentIndex: 0,
	}
}

func (lb *LoadBalancer) GetNextServer() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if len(lb.servers) == 0 {
		return "" // Veya hata döndürülebilir
	}

	// Round-robin mantığıyla sıradaki sunucuyu al
	server := lb.servers[lb.currentIndex]
	lb.currentIndex = (lb.currentIndex + 1) % len(lb.servers)
	return server
}

func TestWeightedRoundRobin(t *testing.T) {
	testCases := []struct {
		name     string
		servers  []Server
		requests int
		expected map[string]int
	}{
		{
			name: "Basic weights 3:2:1",
			servers: []Server{
				{Name: "Server1", Weight: 3},
				{Name: "Server2", Weight: 2},
				{Name: "Server3", Weight: 1},
			},
			requests: 6,
			expected: map[string]int{
				"Server1": 3,
				"Server2": 2,
				"Server3": 1,
			},
		},
		{
			name: "Zero weight handling",
			servers: []Server{
				{Name: "ServerA", Weight: 2},
				{Name: "ServerB", Weight: 0},
				{Name: "ServerC", Weight: 1},
			},
			requests: 3,
			expected: map[string]int{
				"ServerA": 2,
				"ServerC": 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lb := NewLoadBalancer(tc.servers)
			counts := make(map[string]int)

			for i := 0; i < tc.requests; i++ {
				server := lb.GetNextServer()
				counts[server]++
			}

			for name, expectedCount := range tc.expected {
				if counts[name] != expectedCount {
					t.Errorf("%s için beklenen: %d, alınan: %d",
						name, expectedCount, counts[name])
				}
			}
		})
	}
}

func ExampleLoadBalancer() {
	servers := []Server{
		{Name: "Web-1", Weight: 5},
		{Name: "Web-2", Weight: 3},
		{Name: "Web-3", Weight: 2},
	}

	lb := NewLoadBalancer(servers)

	for i := 0; i < 10; i++ {
		lb.GetNextServer()
	}
}
