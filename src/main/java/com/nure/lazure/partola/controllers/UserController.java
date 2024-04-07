package com.nure.lazure.partola.controllers;

import com.nure.lazure.partola.dto.SimpleUserHeaderDTO;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

/**
 * @author Ivan Partola
 */
@Controller
@RequiredArgsConstructor
public class UserController {
    @GetMapping({"", "/"})
    public String mainPage(Model model) {
        model.addAttribute("headerUserInfo", new SimpleUserHeaderDTO("simord", 100, "USDT"));
        return "global/marketplace";
    }

    @GetMapping("/profile/{username}")
    public String profile(@PathVariable String username,
                          Model model) {
        model.addAttribute("headerUserInfo", new SimpleUserHeaderDTO("simord", 100, "USDT"));
        //model.addAttribute("user", simpleUserService.findUserByUsername(username).get());
        return "user/profile";
    }

    @GetMapping("/my-listings")
    public String listings(Model model) {
        model.addAttribute("headerUserInfo", new SimpleUserHeaderDTO("simord", 100, "USDT"));
        return "user/listings";
    }
}
