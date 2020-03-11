# toy-gin-grpc-env
## This is a toy enviroment for testing the gRPC connection with Gin via Nginx.
### Application in this repositry is a just simple function implementing the four arithmetic operations.
#### You can play it as bellows.
```
docker build -t gin-grpc ./
docker run -dit --privileged --name gin-grpc gin-grpc /sbin/init
docker exec -it gin-grpc bash
curl localhost/calc?4+3
7
```
