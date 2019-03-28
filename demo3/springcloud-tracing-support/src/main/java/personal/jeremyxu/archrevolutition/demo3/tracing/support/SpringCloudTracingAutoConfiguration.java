package personal.jeremyxu.archrevolutition.demo3.tracing.support;

import feign.RequestInterceptor;
import org.springframework.boot.autoconfigure.condition.ConditionalOnBean;
import org.springframework.boot.autoconfigure.condition.ConditionalOnClass;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class SpringCloudTracingAutoConfiguration {

  @Bean
  @ConditionalOnClass(RequestInterceptor.class)
  public RequestInterceptor springCloudTracingFeignClientRequestInterceptor(){
    return new SpringCloudTracingFeignClientRequestInterceptor();
  }

  @Bean
  @ConditionalOnBean(RestTemplate.class)
  public SpringCloudTracingRestTemplateInterceptor springCloudTracingRestTemplateInterceptor(){
    return new SpringCloudTracingRestTemplateInterceptor();
  }
}
