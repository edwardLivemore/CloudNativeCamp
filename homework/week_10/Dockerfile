FROM ubuntu
ENV MY_SERVICE_PORT=80
# ENV MY_SERVICE_PORT1=80
# ENV MY_SERVICE_PORT2=80
# ENV MY_SERVICE_PORT3=80
# LABEL multi.label1="value1" multi.label2="value2" other="value3"

LABEL label="httpserver"
ADD httpserver /httpserver/

RUN apt-get update && apt-get install -y golang
RUN go env -w GO111MODULE=on
RUN go env -w GOSUMDB=off
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN cd /httpserver
RUN go mod tidy
# RUN go install github.com/golang/glog@v1.0.0
# RUN go install github.com/google/uuid@v1.3.0
# RUN go install github.com/prometheus/client_golang@v1.14.0
# RUN go install github.com/beorn7/perks@v1.0.1 // indirect
# RUN go install github.com/cespare/xxhash/v2@v2.1.2 // indirect
# RUN go install github.com/golang/protobuf@v1.5.2 // indirect
# RUN go install github.com/matttproud/golang_protobuf_extensions@v1.0.1 // indirect
# RUN go install github.com/prometheus/client_model@v0.3.0 // indirect
# RUN go install github.com/prometheus/common@v0.37.0 // indirect
# RUN go install github.com/prometheus/procfs@v0.8.0 // indirect
# RUN go install golang.org/x/sys@v0.0.0-20220520151302-bc2c85ada10a // indirect
# RUN go install google.golang.org/protobuf@v1.28.1 // indirect

EXPOSE 80
ENTRYPOINT ["nohup", "go", "run", "/httpserver/main.go", "&"]