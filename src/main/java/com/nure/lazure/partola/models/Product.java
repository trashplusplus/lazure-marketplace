package com.nure.lazure.partola.models;

import com.fasterxml.jackson.annotation.JsonProperty;
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

    @JsonProperty("resource_link")
    private String resourceLink;

    @JsonProperty("category_id")
    private Integer categoryId;
}
