ROOT_DIR=$(shell pwd)
OUTPUT_DIR=$(ROOT_DIR)/_output

server:
	@go build -o $(OUTPUT_DIR)/server main.go
	@echo "编译产物位于 $(OUTPUT_DIR)"

.PHONY: clean
clean:
	@echo "清空 $(OUTPUT_DIR)"
	@rm -rf $(OUTPUT_DIR)
