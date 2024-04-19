package com.nure.lazure.partola.controllers;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

/**
 * @author Ivan Partola
 */
@Controller
@RequiredArgsConstructor
public class UserController {
    @GetMapping({"", "/"})
    public String mainPage() {
        return "global/marketplace";
    }

    @GetMapping("/profile/{wallet}")
    public String profile(@PathVariable String wallet) {
        return "user/profile";
    }

    @GetMapping("/my-listings")
    public String listings() {
        return "user/listings";
    }
}
