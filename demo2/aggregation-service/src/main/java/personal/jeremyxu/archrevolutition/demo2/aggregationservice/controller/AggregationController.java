package personal.jeremyxu.archrevolutition.demo2.aggregationservice.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import personal.jeremyxu.archrevolutition.demo2.aggregationservice.dto.User;
import personal.jeremyxu.archrevolutition.demo2.aggregationservice.service.UserService;

import java.util.List;

@RestController
@RequestMapping(value = "/api/v1/aggregation")
public class AggregationController {

    @Autowired
    UserService userService;

    @GetMapping(value = "/users")
    public List<User> getUsers() {
        // 查询user表中所有记录
        return userService.getUsers();
    }
}
