package personal.jeremyxu.archrevolutition.demo6.aggregationservice.service;

import personal.jeremyxu.archrevolutition.demo6.aggregationservice.dto.Blog;

import java.util.List;

public interface BlogProxyService {
    void deleteBlogsByUserId(Integer userId);

    List<Blog> getBlogs();

    Blog getBlog(Integer blogId);
}
