import random
from collections import defaultdict

class Server:
    def __init__(self, name, weight):
        self.name = name
        self.weight = weight

class WeightedRoundRobinBalancer:
    def __init__(self, servers):
        self.servers = servers.copy()
        self.total_weight = sum(server.weight for server in servers)
        self.cumulative_weights = self._calculate_cumulative_weights()
        self.current_index = 0

    def _calculate_cumulative_weights(self):
        cumulative_weights = []
        current = 0
        for server in self.servers:
            current += server.weight
            cumulative_weights.append(current)
        return cumulative_weights

    def get_next_server(self):
        if self.total_weight == 0:
            raise ValueError("Total weight of servers must be greater than 0")
        
        random_value = random.randrange(self.total_weight)
        for i, cum_weight in enumerate(self.cumulative_weights):
            if random_value < cum_weight:
                self.current_index = i
                return self.servers[i]
        return self.servers[0]  # Fallback, should never reach here

def test_balancer_distribution():
    servers = [
        Server("Server1", 3),
        Server("Server2", 2),
        Server("Server3", 1)
    ]
    
    balancer = WeightedRoundRobinBalancer(servers)
    counts = defaultdict(int)
    
    total_requests = 100000
    for _ in range(total_requests):
        server = balancer.get_next_server()
        counts[server.name] += 1
    
    total_weight = sum(s.weight for s in servers)
    expected_ratios = {s.name: s.weight/total_weight for s in servers}
    
    for name, count in counts.items():
        actual_ratio = count / total_requests
        expected = expected_ratios[name]
        assert abs(actual_ratio - expected) < 0.01, (
            f"{name} has ratio {actual_ratio:.3f} (expected {expected:.3f})"
        )
    
    print("Distribution test passed!")

def sample_usage():
    servers = [
        Server("Web Server 1", 5),
        Server("Web Server 2", 3),
        Server("Web Server 3", 2)
    ]
    
    balancer = WeightedRoundRobinBalancer(servers)
    
    print("10 sample requests:")
    for i in range(10):
        server = balancer.get_next_server()
        print(f"Request {i+1} → {server.name}")

if __name__ == "__main__":
    sample_usage()
    print("\nRunning distribution test...")
    test_balancer_distribution()