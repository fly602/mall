FROM golang:1.17

LABEL maintainer="fuleyi <fuleyi@outlook.com>"

# 创建一个普通用户
RUN useradd -ms /bin/bash normal

USER root
RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security-cdn.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \  
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    curl \
    zip \
    unzip \
    git \
    openssh-server \
    bash-completion \
    vim \
    tzdata \
    iputils-ping

ENV GOPROXY=https://goproxy.cn/,direct \
    GOPATH="/go:/usr" \
    PATH=$GOPATH/bin:$PATH

USER normal
# 安装vscode的go插件/protoc-gen-go/protoc-gen-go-grpc/goctl
RUN export GOPROXY=https://goproxy.cn/,direct && \ 
    export GO111MODULE="on" && \
    go get -v golang.org/x/tools/gopls@v0.11.0 && \
    go get -v github.com/go-delve/delve/cmd/dlv@v1.7.0 && \
    go get -v github.com/mdempsky/gocode && \
    go get -v golang.org/x/tools/cmd/goimports && \
    go get -v github.com/rogpeppe/godef && \
    go get -v github.com/cweill/gotests/gotests@v1.6.0 && \
    go get -v github.com/fatih/gomodifytags@v1.16.0 && \
    go get -v github.com/josharian/impl@v1.1.0 && \
    go get -v github.com/haya14busa/goplay/cmd/goplay@v1.0.0 && \
    go get -v honnef.co/go/tools/cmd/staticcheck@v0.3.3 && \
    go get -v google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0 &&\
    go get -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 &&\
    go get -v github.com/zeromicro/go-zero/tools/goctl@v1.4.5

USER root
# 安装 protoc
COPY protoc-21.12-linux-x86_64.zip /tmp/protoc.zip
RUN unzip -d /tmp/protoc /tmp/protoc.zip && \
    mv /tmp/protoc/bin/protoc /usr/bin

# 命令行开启自动补全
RUN echo 'if ! shopt -oq posix; then \n\
  if [ -f /usr/share/bash-completion/bash_completion ]; then \n\
    . /usr/share/bash-completion/bash_completion \n\
  elif [ -f /etc/bash_completion ]; then \n\
    . /etc/bash_completion \n\
  fi \n\
fi'\
>> /etc/bash.bashrc
# goctl自动补全
RUN goctl completion bash > /etc/bash_completion.d/goctl

# 清理垃圾
USER root
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    rm /var/log/lastlog /var/log/faillog

RUN mkdir -p /opt/mall && \
    chown -R normal /opt/mall/

# 设置工作目录
WORKDIR /usr/src/