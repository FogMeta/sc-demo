# sc-demo

## 1. 使用脚本启动lotus的轻节点
```
#!/bin/bash
export FULLNODE_API_INFO=wss://api.chain.love
export LOTUS_PATH=/data/.lotus
lotus daemon --lite  >> daemon.log 2>&1 &
```

## 2. 创建钱包
```
lotus wallet new
```

## 3. 转账
```
lotus send --from <address> <new_wallet_address> amount
```

## 4.拉取swan-client的二进制文件：
```
wget --no-check-certificate https://github.com/filswan/go-swan-client/releases/download/v2.1.0-rc1/swan-client-2.1.0-rc1-linux-amd64
wget --no-check-certificate https://github.com/filswan/go-swan-client/releases/download/v2.1.0-rc1/config.toml.example
wget --no-check-certificate https://github.com/filswan/go-swan-client/releases/download/v2.1.0-rc1/chain-rpc.json
```

## 5. 修改swan-client的配置文件(config.toml):
```

```


6. 编译：
```
go build
```
 