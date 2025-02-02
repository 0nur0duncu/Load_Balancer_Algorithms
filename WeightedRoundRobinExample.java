import java.util.ArrayList;
import java.util.List;

public class WeightedRoundRobinExample {
    public static void main(String[] args) {
        // Sample list of servers with weights
        List<WeightedRoundRobinBalancer.Server> serverList = new ArrayList<>();
        serverList.add(new WeightedRoundRobinBalancer.Server("Server1", 3));
        serverList.add(new WeightedRoundRobinBalancer.Server("Server2", 2));
        serverList.add(new WeightedRoundRobinBalancer.Server("Server3", 1));

        // Create a weighted round-robin load balancer with the server list
        WeightedRoundRobinBalancer balancer = new WeightedRoundRobinBalancer(serverList);

        // Simulate requests to the load balancer
        for (int i = 0; i < 10; i++) {
            WeightedRoundRobinBalancer.Server nextServer = balancer.getNextServer();
            System.out.println("Request " + (i + 1) + ": Routed to " + nextServer.getName());
        }
    }
}