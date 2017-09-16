# 透明的修改 HTTP 请求
实现一个简单的数据校验功能，对请求body做md5计算，然后把结果转为十六进制添加到请求header中，头部名称为X-Md5。也就是在HTTP 请求header中添加一个X-Md5: <hex md5 of body> 键值对，如果body为空那么就不填。
