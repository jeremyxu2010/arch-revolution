package personal.jeremyxu.archrevolutition.demo2.aggregationservice.service;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import personal.jeremyxu.archrevolutition.demo2.aggregationservice.dto.UserDTO;

import java.util.List;

@FeignClient("user-service")
@RequestMapping(value = "/api/v1")
public interface UserFeignService {
    @DeleteMapping("/users/{id}")
    Integer deleteUser(@PathVariable("id") Integer userId);

    @GetMapping("/users/{id}")
    UserDTO getUser(@PathVariable("id") Integer userId);
}
