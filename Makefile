ROOT_DIR=$(shell pwd)
OUTPUT_DIR=$(ROOT_DIR)/_output

## 编译服务端可执行文件
server:
	@go build -o $(OUTPUT_DIR)/server main.go
	@echo "编译产物位于 $(OUTPUT_DIR)"

## 构建Docker镜像
.PHONY: image
image:
	@$(shell sh ./build/build-docker.sh)
	@echo "Docker 镜像构建完成"

## 清空编译产物
.PHONY: clean
clean:
	@echo "清空 $(OUTPUT_DIR)"
	@rm -rf $(OUTPUT_DIR)
