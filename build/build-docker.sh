CMD_PATH=$(dirname $0)
cd $CMD_PATH
cd ..
docker build -t douyin -f build/Dockerfile .