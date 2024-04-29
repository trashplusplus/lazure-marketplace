package com.nure.lazure.partola.controllers;

import com.nure.lazure.partola.models.Product;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

/**
 * @author Ivan Partola
 */
@RequiredArgsConstructor
@RequestMapping("/api/listings")
@RestController
public class ListingsController {
    @PostMapping("/new")
    public ResponseEntity<?> add(@ModelAttribute Product product) {
        try {
            System.out.println(product);
            return ResponseEntity.ok("Product added successfully.");
        } catch (Exception e) {
            return ResponseEntity.badRequest().body("Error while adding a new product.");
        }
    }

}
