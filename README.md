
# kafka demo

### 目录结构
```
.
├── README.md
├── cmd             #可执行程序文件夹
│   ├── consumer    #消费者入口
│   │   ├── config.toml.example
│   │   ├── logs
│   │   └── main.go
│   └── producer    #生产者入口
│   │   ├── config.toml.example
│   │   ├── logs
│   │   └── main.go
├── internal        
│   ├── common      #公共包
│   ├── consumer    #消费者
│   └── producer    #生产者
├── pkg             #工具包文件夹
│   └── kafka       #kafka调用封装
└── vendor          #依赖包文件夹

```

### 执行步骤

#### 运行消费者程序
1.复制配置文件
```
cp  cmd/consumer/config.toml.example  cmd/consumer/config.toml
```
请在config.toml文件中，填写kafka的相关配置   

2.编译
```
cd cmd/consumer
go build .
./consumer
```

#### 运行生产者程序
1.复制配置文件
```
cp  cmd/producer/config.toml.example  cmd/producer/config.toml
```
请在config.toml文件中，填写kafka的相关配置   

2.编译
```
cd cmd/producer
go build .
./producer
```