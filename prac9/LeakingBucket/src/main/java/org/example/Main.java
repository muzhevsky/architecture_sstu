package org.example;

import com.sun.net.httpserver.HttpServer;

import java.net.InetSocketAddress;

public class Main {
    public static void main(String[] args) throws Exception {
        HttpServer server = HttpServer.create(new InetSocketAddress(8000), 0);
        var trafficManager = new LeakingBucket(3, 5);

        var intervalHandler = new MyHttpHandler(trafficManager);
        trafficManager.start();
        server.createContext("/", intervalHandler);

        server.setExecutor(null);
        server.start();
    }
}
