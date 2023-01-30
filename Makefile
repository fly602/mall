SERVICE_DIRS=./service

target:
	@make -C ${SERVICE_DIRS}

build:
	@make build -C ${SERVICE_DIRS}

all: target run

clean:
	@+make clean -C ${SERVICE_DIRS}

run:
	@docker-compose build
	@docker-compose up -d

down:
	@docker-compose down

help:
	@echo "make 在docker环境中编译service服务"
	@echo "make build 在本地环境编译service服务"
	@echo "make clean 删除中间目标文件"
	@echo "make run 通过docker-compose构建服务并在容器中运行"
	@echo "make all 编译并在容器中运行微服务"
	@echo "make down 停止运行服务和容器"


.PHONY: all build clean run down

