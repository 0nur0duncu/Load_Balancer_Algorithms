import java.util.ArrayList;
import java.util.List;

public class RoundRobinExample {
    /**
     * This example demonstrates how to use the LoadBalancer class to
     * route requests to a list of servers in a round-robin sequence.
     * The LoadBalancer is created with a sample list of servers and
     * 10 requests are simulated to demonstrate how the load balancer
     * routes the requests to the servers in the list in a round-robin
     * sequence.
     */
    public static void main(String[] args) {

        List<String> serverList = new ArrayList<>();
        serverList.add("Server1");
        serverList.add("Server2");
        serverList.add("Server3");

        LoadBalancer loadBalancer = new LoadBalancer(serverList);

        for (int i = 0; i < 10; i++) {
            String nextServer = loadBalancer.getNextServer();
            System.out.println("Request " + (i + 1) + ": Routed to " + nextServer);
        }
    }
}