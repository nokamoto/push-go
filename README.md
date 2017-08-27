# push-go

[![CircleCI](https://circleci.com/gh/nokamoto/push-go/tree/master.svg?style=svg)](https://circleci.com/gh/nokamoto/push-go/tree/master)

## Google Cloud Platform

[Getting Started with gRPC on Container Engine](https://cloud.google.com/endpoints/docs/quickstart-grpc-container-engine) に従いサービスをデプロイします

```
./gcp/setup-api_config.sh [PROJECT_ID]
```

```
./gcp/setup-k8s.sh [SERVICE] [SERVICE_NAME] [SERVICE_CONFIG_ID]
```

認証情報を作成し、以下でデプロイされたサービスを確認します

```
pushcli -host [EXTERNAL_IP] -port 80 -apikey [API_KEY] ...
```
