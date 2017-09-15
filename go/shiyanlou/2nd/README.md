# 缓存请求执行结果
在分布式系统中，我们经常会处理一些请求，这些请求有一个特点:

    改动不频繁，平均每个条目改动时间达到秒，十几秒级别
    允许数据出现延迟
    请求非常频繁，请求的ops有十几甚至几百
    数据存储在数据库

理论上针对这些请求，我们可以在程序内做缓存，不用把所有的请求打到后端的数据库。

这里先看 https://godoc.org/github.com/golang/groupcache/singleflight ，利用这个库我们能在很多并发的请求到达的时候，只有一个请求到达后端。

但是我们可以做更多的事情，我们可以在这个之上加一层缓存，缓存请求的执行结果，并在缓存结果过期之前，都返回同样的结果，这样就可以避免直接访问请求直接访问后端系统。

## 需求
在 singleflight 的基础上，实现一个带过期时间缓存库。

文件中基本的库调用方式已经有了，需要在 cacheflight.go 里面完成剩余的代码。

cachefligt.go 中使用函数 NewGroup(cacheExpiration time.Duration) (group *Group 返回过缓存库，其中 cacheExpiration 指明了该缓存库所有缓存项的过期时间，缓存项生存 cacheExpiration 时长后就会过期。
