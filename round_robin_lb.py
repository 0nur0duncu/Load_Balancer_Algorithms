class LoadBalancer:
    def __init__(self, servers):
        """
        Create a new LoadBalancer from the given list of servers.

        The input list is copied, so it can be modified independently of the
        LoadBalancer.
        """
        self.servers = list(servers)
        self.current_index = 0

    def get_next_server(self):
        """
        Returns the next server in the round-robin sequence.

        This method updates the current index to maintain the order of 
        server selection in a thread-safe manner.
        """

        next_server = self.servers[self.current_index]
        self.current_index = (self.current_index + 1) % len(self.servers)
        return next_server

if __name__ == "__main__":
    server_list = ["Server1", "Server2", "Server3"]
    load_balancer = LoadBalancer(server_list)

    for i in range(10):
        next_server = load_balancer.get_next_server()
        print(f"Request {i+1}: Routed to {next_server}")