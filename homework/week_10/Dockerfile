FROM ubuntu
ENV MY_SERVICE_PORT=80
# ENV MY_SERVICE_PORT1=80
# ENV MY_SERVICE_PORT2=80
# ENV MY_SERVICE_PORT3=80
# LABEL multi.label1="value1" multi.label2="value2" other="value3"

LABEL label="httpserver"
ADD httpserver /httpserver

# RUN apt-get update && apt-get install -y golang
# RUN go env -w GO111MODULE=on
# RUN go env -w GOSUMDB=off
# RUN go env -w GOPROXY=https://goproxy.cn,direct
# RUN chmod 777 httpserver
RUN chmod +x /httpserver

EXPOSE 80
ENTRYPOINT ["./httpserver/httpserver"]