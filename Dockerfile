FROM grpc/go:latest
MAINTAINER Takayuki Yamamuro

RUN apt-get update
RUN apt-get -y install vim
RUN apt-get -y install sysv-rc-conf

# Make the Log directory
WORKDIR /var/log
RUN mkdir apl

# nginx settings
RUN apt-get -y install nginx
WORKDIR /etc/nginx
COPY nginx.conf .
RUN sysv-rc-conf nginx on

# gin settings
#RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-gonic/gin

#gRPC settings
RUN go get golang.org/x/text/secure/bidirule
RUN go get golang.org/x/net/context
WORKDIR /go/src
RUN mkdir ginApl
WORKDIR ginApl
COPY ginApl.pb.go .
COPY ginApl.proto .
WORKDIR /etc/init.d/
COPY grpcServerSvc .

#Apl settings
WORKDIR /usr/local
RUN mkdir ginApl
WORKDIR ginApl
COPY ginServer.go .
COPY grpcServer.go .
WORKDIR /etc/init.d
COPY gin .
RUN chmod +x gin

#Service start
ENTRYPOINT /etc/init.d/gin start && /etc/init.d/grpcServerSvc start && service nginx start && bash 
