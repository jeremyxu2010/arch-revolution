package personal.jeremyxu.archrevolutition.demo2.blogservice.service;


import personal.jeremyxu.archrevolutition.demo2.blogservice.model.Blog;

import java.util.List;

public interface BlogService {
    List<Blog> findAll();

    void save(Blog blog);

    Blog findById(int blogId);

    void deleteById(int blogId);

    void deleteByUserId(Integer userId);
}
