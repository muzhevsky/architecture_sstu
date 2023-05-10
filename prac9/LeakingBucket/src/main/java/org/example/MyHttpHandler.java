package org.example;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;

import java.io.IOException;
import java.io.OutputStream;

public class MyHttpHandler implements HttpHandler {
    private TrafficManager manager;

    public MyHttpHandler(TrafficManager manager){
        this.manager = manager;
    }

    public void handle(HttpExchange exchange) throws IOException {
        String response = "";
        if (exchange.getRequestURI().toString().equals("/favicon.ico")) return;
        var succeeded = manager.addRequest(exchange);
        if (succeeded)
            response = "query added";
        else{
            response = "query rejected";
        }
        exchange.sendResponseHeaders(200, response.length());
        OutputStream os = exchange.getResponseBody();
        os.write(response.getBytes());
        os.close();
    }
}
