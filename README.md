## Proxy Download

![Bazel](https://github.com/fzhyzamt/proxy-download/workflows/Bazel/badge.svg?branch=master)

### Usage

构建执行文件
```shell
git clone git@github.com:fzhyzamt/proxy-download.git
cd proxy-download
go build -o target/build proxy-download/src
```

如果直接访问端口速度很快，但通过nginx后非常慢，尝试关闭nginx的buffer。
```
proxy_buffering off;
```

限制并发量
```
limit_conn_zone $server_name zone=perserver:10m;
limit_conn perserver 2;
```

### example
```
limit_conn_zone $server_name zone=perserver:10m;

...

server {
        listen 80;
        listen 443 ssl;
        server_name d.example.com;

        ssl_certificate /etc/letsencrypt/live/example.com/fullchain.pem;
        ssl_certificate_key /etc/letsencrypt/live/example.com/privkey.pem;

        proxy_buffering off;

        limit_conn perserver 2;

        location / {
                proxy_redirect off;
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_pass http://127.0.0.1:8081;
        }
}
```

### 编译
```
# 使用bazel编译所有目标平台
bazel build //:all

# 使用go编译单个平台
go build -o target/build proxy-download/src
```
