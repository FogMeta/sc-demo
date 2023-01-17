# sc-demo
Before the `swan-client` have an SDK, this demo can help developers store their data to the IPFS and Filecoin network using [swan-client](https://github.com/filswan/go-swan-client). If you are not familiar with the Filecoin network and do not have a Filecoin full node, the following instructions can help you. 

## 1. Start a lite node of Filecoin
For the client, a lite node is enough to send deals to the storage providers.
```
export FULLNODE_API_INFO=wss://api.chain.love
export LOTUS_PATH=/data/.lotus
lotus daemon --lite  >> daemon.log 2>&1 &
```
> **Note**: `FULLNODE_API_INFO ` is maintained by the Protocol Labs

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
## 4. Install swan-client
```
mkdir swan-client
cd swan-client
wget --no-check-certificate https://github.com/filswan/go-swan-client/releases/download/v2.1.0-rc1/install.sh
chmod +x install.sh
./install.sh

```

## 5. Config the swan-client:

The default config file is in the `~/.swan/client/config.toml`
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

## 6. Build `sc-demo`：
```
go build
```
## 7. Run `sc-demo`：
The demo shows you a complete process, including the following functions:
 - Generate the CAR file from the source file, and get the path of the metadata file
 - Upload the CAR file to the IPFS gateway, and update metadata files (JSON and CSV), return the download link from IPFS gateway.
 - Send the task(deals) to the storage providers (including private tasks and auto-bid tasks) using the metadata file. After 24 hours, the files will be stored to the Filecoin network
> Note: the test files are in the `./sc-demo/test/sources/`
```
git clone https://github.com/FogMeta/sc-demo.git
cd sc-demo/
./sc-demo 
```
Output:
```

CreateCar-->  metaDataPath: /home/xxx/xxx/sc-demo/test/car/car.json 

UploadCar-->  downloadUrl: http://ipfs1.storefrontiers.cn/ipfs/QmWfmpGWx2bXyfgnz8AMgH3HUkEKLoKtJ8bQhUjKYgPXQ7 

=======private task =======
time="2023-01-17 01:31:19.949" level=info msg="Your metadata file is: /home/xxx/xxx/sc-demo/test/car/car.json" func=func9 file="main.go:368"
time="2023-01-17 01:31:19.949" level=error msg="stat /home/xxx/xxx/sc-demo/test/car/deal: no such file or directory" func=GetPathType file="file.go:65"
time="2023-01-17 01:31:19.949" level=info msg="output directory: /home/xxx/xxx/sc-demo/test/car/deal created" func=CreateDirIfNotExists file="file.go:320"
time="2023-01-17 01:31:19.949" level=info msg="Your output dir: /home/xxx/xxx/sc-demo/test/car/deal" func=CreateTask file="task.go:154"
time="2023-01-17 01:31:20.249" level=info msg="File:/home/xxx/xxx/sc-demo/test/car/QmYqkd9ZKyUhP57wbLJTh4GHTmsXj64bedApEXLAABuBMj.car,current epoch:2520422, start epoch:2532062" func=sendDeals2Miner file="deal.go:234"
time="2023-01-17 01:31:26.908" level=info msg="miner: f01955030, price: 0" func=CheckDealConfig file="client.go:595"
time="2023-01-17 01:31:28.185" level=info msg="deal sent successfully, task name:swan-task-21qnsu, car file:/home/xxx/xxx/sc-demo/test/car/QmYqkd9ZKyUhP57wbLJTh4GHTmsXj64bedApEXLAABuBMj.car, dealCID|dealUuid:bafyreic367xlczyzckatlp2adcci653edckwfm3mbpujjnxrw5qi6hhjvu, start epoch:2532062, miner:f01955030" func=sendDeals2Miner file="deal.go:294"
You are using the MARKET(version=1.1 built-in Lotus) send deals, but it is deprecated, will remove soon. Please set [main.market_version=“1.2”]
time="2023-01-17 01:31:28.686" level=info msg="1 deal(s) has(ve) been sent for task: swan-task-21qnsu, minerID: [f01955030]" func=sendDeals2Miner file="deal.go:307"
time="2023-01-17 01:31:28.687" level=info msg="Metadata json file generated: /home/xxx/xxx/sc-demo/test/car/deal/swan-task-21qnsu-deals.json" func=WriteFileDescsToJsonFile file="common.go:103"
time="2023-01-17 01:31:28.687" level=info msg="Metadata csv generated: /home/xxx/xxx/sc-demo/test/car/deal/swan-task-21qnsu-deals.csv" func=WriteCarFilesToCsvFile file="common.go:214"
time="2023-01-17 01:31:28.687" level=info msg="Metadata json file generated: /home/xxx/xxx/sc-demo/test/car/deal/swan-task-21qnsu-metadata.json" func=WriteFileDescsToJsonFile file="common.go:103"
time="2023-01-17 01:31:28.687" level=info msg="Metadata csv generated: /home/xxx/xxx/sc-demo/test/car/deal/swan-task-21qnsu-metadata.csv" func=WriteCarFilesToCsvFile file="common.go:214"
time="2023-01-17 01:31:28.687" level=info msg="Working in Online Mode. A swan task will be created on the filwan.com after process done. " func=sendTask2Swan file="task.go:292"
time="2023-01-17 01:31:45.688" level=info msg="status:success, message:task created successfully" func=sendTask2Swan file="task.go:311"
time="2023-01-17 01:31:45.688" level=info msg="task information is in:/home/xxx/xxx/sc-demo/test/car/deal/swan-task-21qnsu-metadata.json" func=CreateTaskByConfig file="task.go:105"
SendDeal-->  minerIdAndDealCids: [{MinerId:f01955030 DealCid:bafyreic367xlczyzckatlp2adcci653edckwfm3mbpujjnxrw5qi6hhjvu}]

=======auto deal =======
time="2023-01-17 01:31:45.759" level=info msg="Your metadata file is: /home/xxx/xxx/sc-demo/test/car/car.json" func=func9 file="main.go:368"
time="2023-01-17 01:31:45.759" level=warning msg="cmdDeal is unnecessary for auto-bid or manual-bid task" func=CreateTask file="task.go:125"
time="2023-01-17 01:31:45.759" level=info msg="Your output dir: /home/xxx/xxx/sc-demo/test/car/deal" func=CreateTask file="task.go:154"
time="2023-01-17 01:31:45.912" level=info msg="Metadata json file generated: /home/xxx/xxx/sc-demo/test/car/deal/swan-task-gd01n1-metadata.json" func=WriteFileDescsToJsonFile file="common.go:103"
time="2023-01-17 01:31:45.912" level=info msg="Metadata csv generated: /home/xxx/xxx/sc-demo/test/car/deal/swan-task-gd01n1-metadata.csv" func=WriteCarFilesToCsvFile file="common.go:214"
time="2023-01-17 01:31:45.912" level=info msg="Working in Online Mode. A swan task will be created on the filwan.com after process done. " func=sendTask2Swan file="task.go:292"
time="2023-01-17 01:32:01.049" level=info msg="status:success, message:task created successfully" func=sendTask2Swan file="task.go:311"
time="2023-01-17 01:32:01.049" level=info msg="task information is in:/home/xxx/xxx/sc-demo/test/car/deal/swan-task-gd01n1-metadata.json" func=CreateTaskByConfig file="task.go:105"
time="2023-01-17 01:32:12.556" level=info msg="no offline deals to be sent" func=sendAutoBidDeals4Task file="auto.go:196"

time="2023-01-17 01:33:14.074" level=info msg="send deal for task:90358,e7bdd2aa-c557-403e-810d-549da629c353, deal:636306" func=sendAutobidDeal file="auto.go:316"
time="2023-01-17 01:33:14.976" level=info msg="miner: f0717969, price: 0" func=CheckDealConfig file="client.go:595"
time="2023-01-17 01:33:16.185" level=info msg="deal sent successfully, task:90358, uuid:e7bdd2aa-c557-403e-810d-549da629c353, deal:636306, task name:0xc000e77220, deal CID:bafyreieahpkremc7glrufpzijrhwgklam3hpmozk4mcaz2iayp7mdfjnva, start epoch:2531943, miner:f0717969" func=sendAutobidDeal file="auto.go:377"
time="2023-01-17 01:33:16.813" level=info msg="Metadata json file generated: /home/xxx/xxx/sc-demo/test/car/deal/swan-task-gd01n1-auto-deals.json" func=WriteFileDescsToJsonFile file="common.go:103"
time="2023-01-17 01:33:16.813" level=info msg="Metadata csv generated: /home/xxx/xxx/sc-demo/test/car/deal/swan-task-gd01n1-auto-deals.csv" func=WriteCarFilesToCsvFile file="common.go:214"
time="2023-01-17 01:33:16.814" level=info msg="auto send deal end,dealTotalCount: 1,successed: 1,failed: 0" func=SendAutoBidDealsBySwanClientSourceId file="auto.go:471"
SendDeal-->  minerIdAndDealCids: [{MinerId:f0717969 DealCid:bafyreieahpkremc7glrufpzijrhwgklam3hpmozk4mcaz2iayp7mdfjnva}]## 8. Retrieval and restore

```
The `swan-task-gd01n1-auto-deals.json` can be seen as follows:
```
[
 {
  "Uuid": "e7bdd2aa-c557-403e-810d-549da629c353",
  "SourceFileName": "sources",
  "SourceFilePath": "/home/XXX/XXX/sc-demo/test/sources",
  "SourceFileMd5": "",
  "SourceFileSize": 224228,
  "CarFileName": "QmYqkd9ZKyUhP57wbLJTh4GHTmsXj64bedApEXLAABuBMj.car",
  "CarFilePath": "/home/XXX/XXX/test/car/QmYqkd9ZKyUhP57wbLJTh4GHTmsXj64bedApEXLAABuBMj.car",
  "CarFileMd5": "",
  "CarFileUrl": "http://ipfs1.storefrontiers.cn/ipfs/QmWfmpGWx2bXyfgnz8AMgH3HUkEKLoKtJ8bQhUjKYgPXQ7",
  "CarFileSize": 260096,
  "PayloadCid": "QmYqkd9ZKyUhP57wbLJTh4GHTmsXj64bedApEXLAABuBMj",
  "PieceCid": "baga6ea4seaqp5d6djsucgfv3gcgajckd2mgeslbasljrzoujy5pkvk7g57ax4by",
  "StartEpoch": null,
  "SourceId": 2,
  "Deals": [
   {
    "DealCid": "bafyreieahpkremc7glrufpzijrhwgklam3hpmozk4mcaz2iayp7mdfjnva",
    "MinerFid": "f0717969",
    "StartEpoch": 2531943,
    "Cost": ""
   }
  ]
 }
]
```


## 8. Retrieval and Restore
 - If you want to retrieve the data from the IPFS, you can download the CAR file from the above link and restore it.

Step1:
```
wget -c "http://ipfs1.storefrontiers.cn/ipfs/QmWfmpGWx2bXyfgnz8AMgH3HUkEKLoKtJ8bQhUjKYgPXQ7" -O QmWfmpGWx2bXyfgnz8AMgH3HUkEKLoKtJ8bQhUjKYgPXQ7.car

```

Step2:
```
swan-client generate-car graphsplit restore --input-dir=<CAR_file_path> --output-dir=<output_path>

```

 - If you want to retrieve them from the Filecoin network, Lotus can help you:
```
lotus client retrieve --provider=<miner_ID> <PayloadCid> ~/output-file
```

Swan-client is developing a series of API, include `generate-car`, `upload to IPFS`, `send deals to storage providers`, `Retrieval and Restore`, and so on. please keep your eyes on the [here](https://github.com/filswan/go-swan-client)
