package com.nure.lazure.partola.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.ToString;

import java.math.BigDecimal;

/**
 * @author Ivan Partola
 */
@ToString
@Data
@AllArgsConstructor
public class SimpleUserHeaderDTO {
    private String username;
    private double balance;
    private String currency;
}
