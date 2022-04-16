# gmFS_backend
毕设后端

# 整体架构

```bash
gmFS_backend
├── build.sh # 编译脚本
├── config  # 配置信息相关
├── dal  # Data Access Layer
│   ├── cache # 缓存相关
│   ├── db # db相关
│   └── obj_store # 对象存储相关
├── go.mod
├── go.sum
├── grpc_server # server的grpc实现
├── http_server # server的http api实现（推荐）
├── logs # server日志配置
├── main.go
├── method # 业务逻辑代码
├── model # 实体定义
│   ├── ddl.sql
├── pb_gen # 根据IDL生成的文件
├── README.md
├── test
└── util # 工具包
    └── tool.go
```