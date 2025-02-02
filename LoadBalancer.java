import java.util.ArrayList;
import java.util.List;

class LoadBalancer {
    private List<String> servers;
    private int currentIndex;

    public LoadBalancer(List<String> servers) {
        if (servers == null) {
            throw new NullPointerException("Server list cannot be null");
        }
        if (servers.isEmpty()) {
            throw new IllegalArgumentException("Server list cannot be empty");
        }
        this.servers = new ArrayList<>(servers);
        this.currentIndex = 0;
    }

    /**
     * Returns the next server in the round-robin sequence.
     *
     * This method updates the current index to maintain the order of
     * server selection.
     *
     * @return the next server as a String
     */

    public String getNextServer() {
        String nextServer = servers.get(currentIndex);
        currentIndex = (currentIndex + 1) % servers.size();
        return nextServer;
    }
}