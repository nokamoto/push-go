# push-go

[![CircleCI](https://circleci.com/gh/nokamoto/push-go/tree/master.svg?style=svg)](https://circleci.com/gh/nokamoto/push-go/tree/master)

## Google Cloud Platform

![2017-08-28 3 20 18](https://user-images.githubusercontent.com/4374383/29752810-dc2a9d10-8b9f-11e7-8e79-bbe4621fcf46.png)

### gRPC Clound Endpoints
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

### Cloud Functions
### functions/pub

Cloud Pub/Sub トピックを作成します

```
gcloud beta pubsub topics create resolve-endpoints
```

`config.json` を作成してください

```
{
  "SUBSCRIPTION_HOST": "[EXTERNAL_IP]",
  "SUBSCRIPTION_PORT": 80,
  "API_KEY": "[API_KEY]"
}
```

[HTTP Tutorial](https://cloud.google.com/functions/docs/tutorials/http) に従いソースコードをデプロイします

```
./functions/gcp-deploy.sh [PUB_BUCKET_NAME] [RESOLV_BUCKET_NAME]
```

### functions/resolv

[Cloud Pub/Sub Tutorial](https://cloud.google.com/functions/docs/tutorials/pubsub) に従いソースコードをデプロイします

```
./functions/gcp-deploy.sh [PUB_BUCKET_NAME] [RESOLV_BUCKET_NAME]
```
