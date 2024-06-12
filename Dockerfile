FROM golang:1.22.4-bullseye as build
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

RUN  echo "deb http://mirrors.ustc.edu.cn/debian bullseye main contrib non-free" > /etc/apt/sources.list && \
     echo "deb http://mirrors.ustc.edu.cn/debian bullseye-updates main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb http://mirrors.ustc.edu.cn/debian bullseye-backports main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb http://mirrors.ustc.edu.cn/debian-security/ bullseye-security/updates main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb-src http://mirrors.ustc.edu.cn/debian bullseye main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb-src http://mirrors.ustc.edu.cn/debian bullseye-updates main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb-src http://mirrors.ustc.edu.cn/debian bullseye-backports main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb-src http://mirrors.ustc.edu.cn/debian-security/ bullseye-security/updates main contrib non-free" >> /etc/apt/sources.list && \
     echo cat /etc/apt/sources.list && \
    apt-get update

COPY . /go/src/notion-blog
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

WORKDIR /go/src/notion-blog
RUN go build -o /go/bin/runner

FROM debian:bullseye-slim
RUN  echo "deb http://mirrors.ustc.edu.cn/debian bullseye main contrib non-free" > /etc/apt/sources.list && \
     echo "deb http://mirrors.ustc.edu.cn/debian bullseye-updates main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb http://mirrors.ustc.edu.cn/debian bullseye-backports main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb http://mirrors.ustc.edu.cn/debian-security/ bullseye-security/updates main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb-src http://mirrors.ustc.edu.cn/debian bullseye main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb-src http://mirrors.ustc.edu.cn/debian bullseye-updates main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb-src http://mirrors.ustc.edu.cn/debian bullseye-backports main contrib non-free" >> /etc/apt/sources.list && \
     echo "deb-src http://mirrors.ustc.edu.cn/debian-security/ bullseye-security/updates main contrib non-free" >> /etc/apt/sources.list && \
     echo cat /etc/apt/sources.list && \
    apt-get update

# 64 system run aapt
# RUN apt-get install -y libc6-amd64-i386-cross lib32stdc++6 lib32gcc1 lib32z1 libncurses5-dev cron
RUN apt-get install -y cron
RUN apt-get clean

RUN mkdir /worker
WORKDIR /worker
COPY --from=build /go/bin/runner ./runner
COPY --from=build /go/src/notion-blog/.env ./.env
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
COPY --from=build /go/src/notion-blog/crontab /etc/cron.d/notion-cron

RUN chmod 0644 /etc/cron.d/notion-cron
RUN chmod 0744 /worker/runner
RUN chmod 0744 /worker/.env

# 设置时区
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/timezone
RUN apt-get install -y ca-certificates openssl && update-ca-certificates

RUN crontab /etc/cron.d/notion-cron

#ENTRYPOINT ["/runner -l ./logs -c ./env"]
CMD ["cron", "-f", "-l", "2"]
