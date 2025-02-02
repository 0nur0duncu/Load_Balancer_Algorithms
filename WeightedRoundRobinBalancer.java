import java.util.ArrayList;
import java.util.List;
import java.util.Random;

class WeightedRoundRobinBalancer {
    private List<Server> servers;
    private int[] cumulativeWeights;
    private int totalWeight;
    private int currentIndex;
    private Random random;

    public WeightedRoundRobinBalancer(List<Server> servers) {
        this.servers = new ArrayList<>(servers);
        this.totalWeight = calculateTotalWeight(servers);
        this.cumulativeWeights = calculateCumulativeWeights(servers);
        this.currentIndex = 0;
        this.random = new Random();
    }

    private int calculateTotalWeight(List<Server> servers) {
        int totalWeight = 0;
        for (Server server : servers) {
            totalWeight += server.getWeight();
        }
        return totalWeight;
    }

    private int[] calculateCumulativeWeights(List<Server> servers) {
        int[] cumulativeWeights = new int[servers.size()];
        cumulativeWeights[0] = servers.get(0).getWeight();
        for (int i = 1; i < servers.size(); i++) {
            cumulativeWeights[i] = cumulativeWeights[i - 1] + servers.get(i).getWeight();
        }
        return cumulativeWeights;
    }

    public Server getNextServer() {
        int randomValue = random.nextInt(totalWeight);
        for (int i = 0; i < cumulativeWeights.length; i++) {
            if (randomValue < cumulativeWeights[i]) {
                currentIndex = i;
                break;
            }
        }
        return servers.get(currentIndex);
    }

    // Inner class representing a server with a weight
    static class Server {
        private String name;
        private int weight;

        public Server(String name, int weight) {
            this.name = name;
            this.weight = weight;
        }

        public String getName() {
            return name;
        }

        public int getWeight() {
            return weight;
        }
    }
}