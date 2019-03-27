package personal.jeremyxu.archrevolutition.demo2.aggregationservice.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import personal.jeremyxu.archrevolutition.demo2.aggregationservice.service.BlogProxyService;
import personal.jeremyxu.archrevolutition.demo2.aggregationservice.service.UserFeignService;

@RestController
@RequestMapping(value = "/api/v1/aggregation")
public class AggregationController {

    @Autowired
    UserFeignService userFeignService;

    @Autowired
    BlogProxyService blogProxyService;


    @DeleteMapping("/users/{id}")
    public ResponseEntity deleteUser(@PathVariable("id") Integer userId){
        // 先删除该用户的博文
        blogProxyService.deleteBlogsByUserId(userId);
        // 再删除该用户
        userFeignService.deleteUser(userId);
        return new ResponseEntity(userId, HttpStatus.OK);
    }


}
