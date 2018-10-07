package io.github.flagbrew.servepkx;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.event.EventListener;

import java.awt.*;
import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;

/**
 * @author rproud
 * @since 2018-10-07
 */
@SpringBootApplication
public class Application {

    private static final Logger logger = LoggerFactory.getLogger(Application.class);

    @Value("${local.server.port:8080}")
    private Integer serverPort;

    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }

    @EventListener(ApplicationReadyEvent.class)
    public void applicationReadyEvent() {
        System.out.println("Application started ... launching browser now");
        Browse("http://localhost:" + serverPort);
    }

    public static void Browse(String url) {
        if(Desktop.isDesktopSupported()){
            final Desktop desktop = Desktop.getDesktop();
            try {
                desktop.browse(new URI(url));
            } catch (IOException | URISyntaxException e) {
                logger.error("Failed to open browser with Desktop support", e);
            }
        }else{
            final Runtime runtime = Runtime.getRuntime();
            try {
                runtime.exec("rundll32 url.dll,FileProtocolHandler " + url);
            } catch (IOException e) {
                logger.error("Failed to open url with FileProtocolHandler", e);
            }
        }
    }
}
