package com.nure.lazure.partola.models;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * @author Ivan Partola
 */
@AllArgsConstructor
@NoArgsConstructor
@Data
public class Product {
    private String name;
    private String description;
    private Double price;
    private String resourceLink;
    private String wallet;
}
