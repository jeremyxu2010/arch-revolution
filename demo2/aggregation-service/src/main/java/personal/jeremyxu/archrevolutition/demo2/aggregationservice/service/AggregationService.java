package personal.jeremyxu.archrevolutition.demo2.aggregationservice.service;

import personal.jeremyxu.archrevolutition.demo2.aggregationservice.dto.BlogDTO;

import java.util.List;

public interface AggregationService {
    void deleteUser(Integer userId);

    List<BlogDTO> getBlogs();

    BlogDTO getBlog(Integer blogId);
}
