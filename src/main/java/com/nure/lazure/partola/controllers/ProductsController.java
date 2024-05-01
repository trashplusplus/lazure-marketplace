package com.nure.lazure.partola.controllers;

import com.nure.lazure.partola.models.Product;
import jakarta.servlet.http.HttpSession;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

/**
 * @author Ivan Partola
 */
@RequiredArgsConstructor
@RequestMapping("/api/listings")
@RestController
public class ProductsController {
    private final RestTemplate restTemplate;

    @Autowired
    public ProductsController(RestTemplateBuilder restTemplateBuilder) {
        this.restTemplate = restTemplateBuilder.build();
    }
    @PostMapping("/new")
    public ResponseEntity<?> add(@ModelAttribute Product product, HttpSession session) {
        try {
            String jwtToken = session.getAttribute("jwtToken").toString();
            HttpHeaders headers = new HttpHeaders();
            headers.set(HttpHeaders.AUTHORIZATION, "Bearer " + jwtToken);
            HttpEntity<Product> request = new HttpEntity<>(product, headers);

            restTemplate.exchange(
                    "https://productsapi-954ed826b909.herokuapp.com/product",
                    HttpMethod.POST,
                    request,
                    String.class
            );

            return ResponseEntity.ok("Product added successfully.");
        } catch (Exception e) {
            return ResponseEntity.badRequest().body("Error while adding a new product.");
        }
    }

}
