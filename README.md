# 0x3F 抖声后端

## 整体设计
TODO

## API接口
见[API文档](https://www.apifox.cn/apidoc/shared-1d70e29b-3744-4165-ad83-07ab4d156409)

## 文件组织
```
├── build 构建脚本
├── configs 配置文件
├── deploy 部署脚本与文档
├── internal 私有程序和库代码
│   ├── dao 数据访问层
│   ├── middleware 中间件
│   ├── model 模型
│   ├── pkg 公共代码
│   ├── routers 路由与控制器
│   └── service 业务逻辑
├── pkg 公共代码
└── third_party 外部工具和脚本
```

## 服务依赖
- Minio（或任意兼容 Amazon S3 接口的对象存储）
- MySQL
- Jaeger
- FFmpeg

## 部署
见[部署指南](deploy/README.md)