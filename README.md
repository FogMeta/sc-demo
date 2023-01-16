# sc-demo

## 1. Start lotus lite-node
```
#!/bin/bash
export FULLNODE_API_INFO=wss://api.chain.love
export LOTUS_PATH=/data/.lotus
lotus daemon --lite  >> daemon.log 2>&1 &
```

## 2. Create new wallet
```
export LOTUS_PATH=/data/.lotus
lotus wallet new
```

## 3. Add funds to the wallets
```
export LOTUS_PATH=/data/.lotus
lotus send --from <address> <new_wallet_address> amount
```

## 4. Get the swan-client binaries and configuration files.
```
wget --no-check-certificate https://github.com/filswan/go-swan-client/releases/download/v2.1.0-rc1/swan-client-2.1.0-rc1-linux-amd64
wget --no-check-certificate https://github.com/filswan/go-swan-client/releases/download/v2.1.0-rc1/config.toml.example
wget --no-check-certificate https://github.com/filswan/go-swan-client/releases/download/v2.1.0-rc1/chain-rpc.json
```

## 5. Adjust the swan-client configuration file(config.toml):
```
  [lotus]
    client_api_url = "http://127.0.0.1:1234/rpc/v0"   # Url of lotus client web API, generally the [port] is 1234
    client_access_token = ""                       # Access token of lotus client web API, it should have admin access right

    [main]
    market_version = "1.1"                         # Send deal type, 1.1 or 1.2, config(market_version=1.1) is DEPRECATION, will REMOVE SOON (default: "1.1")
    api_url = "https://go-swan-server.filswan.com" # Swan API address. For Swan production, it is `https://go-swan-server.filswan.com`. It can be ignored if `[sender].offline_swan=true`
    api_key = ""                                   # Swan API key. Acquired from [Swan Platform](https://console.filswan.com/#/dashboard) -> "My Profile"->"Developer Settings". It can be ignored if `[sender].offline_swan=true`.
    access_token = ""                              # Swan API access token. Acquired from [Swan Platform](https://console.filswan.com/#/dashboard) -> "My Profile"->"Developer Settings". It can be ignored if `[sender].offline_swan=true`.

    [ipfs_server]
    download_url_prefix = "http://[ip]:[port]"     # IPFS server URL prefix. Store CAR files for downloading by the storage provider. The downloading URL will be `[download_url_prefix]/ipfs/[dataCID]`
    upload_url_prefix = "http://[ip]:[port]"       # IPFS server URL for uploading files

    [sender]
    offline_swan = false                           # Whether to create a task on [Swan Platform](https://console.filswan.com/#/dashboard), when set to true, only generate metadata for Storage Providers to import deals. 
    verified_deal = true                           # Whether deals in the task are going to be sent as verified or not
    fast_retrieval = true                          # Whether deals in the task are available for fast-retrieval or not
    skip_confirmation = false                      # Whether to skip manual confirmation of each deal before sending
    generate_md5 = false                           # Whether to generate md5 for each CAR file and source file(resource consuming)
    wallet = ""                                    # Wallet used for sending offline deals
  
```

## 6. Build `example`：
```
go build
```
## 7. Run `example`：
```
./example
```
 