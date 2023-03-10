FROM golang:1.18

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
    go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest &&\
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest &&\
    go install github.com/zeromicro/go-zero/tools/goctl@latest

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