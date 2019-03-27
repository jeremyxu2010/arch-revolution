package personal.jeremyxu.archrevolutition.demo2.aggregationservice.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;
import personal.jeremyxu.archrevolutition.demo2.aggregationservice.service.BlogProxyService;

import java.util.HashMap;
import java.util.Map;

@Service
public class BlogProxyServiceImpl implements BlogProxyService {

    @Autowired
    RestTemplate restTemplate;

    @Override
    public void deleteBlogsByUserId(Integer userId) {
        Map<String, Object> uriVariables = new HashMap();
        uriVariables.put("userId", userId);
        restTemplate.delete("http://blog-service/api/v1/blogs/by_userId/{userId}", uriVariables);
    }
}