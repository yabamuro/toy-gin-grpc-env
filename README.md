# toy-gin-grpc-env
## This is a toy enviroment for testing the gRPC connection with Gin via Nginx.
### A gin application in this repositry is a just function returning the same value received via gRPC request with the received time.
#### You can play it as bellows.
```
docker build -t gin-grpc ./
docker run -dit --privileged --name gin-grpc gin-grpc /sbin/init
docker exec -it gin-grpc bash
curl localhost/gin?test
Rcv(gRPC) test on 2020-03-11 19:01:13.1925168 +0000 UTC m=+150.262292001
```
