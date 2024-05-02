package com.nure.lazure.partola.controllers;

import com.nure.lazure.partola.models.User;
import jakarta.servlet.http.HttpSession;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.util.Objects;

/**
 * @author Ivan Partola
 */
@RestController
@RequestMapping("/api/users")
public class UserLoginController {
    private final RestTemplate restTemplate;

    @Autowired
    public UserLoginController(RestTemplateBuilder restTemplateBuilder) {
        this.restTemplate = restTemplateBuilder.build();
    }

    @PostMapping("/login")
    public ResponseEntity<String> login(@RequestBody User user, HttpSession session) {
        HttpHeaders headers = new HttpHeaders();
        headers.set(HttpHeaders.AUTHORIZATION, "Bearer " + System.getenv("PASSWORD"));
        HttpEntity<User> request = new HttpEntity<>(user, headers);

        ResponseEntity<String> response = restTemplate.exchange(
                "https://accountsapi-3a5f92f4b3d5.herokuapp.com/users/login",
                HttpMethod.POST,
                request,
                String.class
        );

        HttpHeaders responseHeaders = response.getHeaders();
        String jwtToken = Objects.requireNonNull(responseHeaders.getFirst(HttpHeaders.AUTHORIZATION)).replace("Bearer ", "");
        session.setAttribute("jwtToken", jwtToken);

        headers = new HttpHeaders();
        headers.set(HttpHeaders.AUTHORIZATION, jwtToken);
        return ResponseEntity.ok().headers(headers).body("Wallet was successfully connected!");
    }

    @PostMapping("/logout")
    public ResponseEntity<?> logout(HttpSession session) {
        session.removeAttribute("jwtToken");
        return ResponseEntity.ok("Wallet was successfully disconnected!");
    }
}
