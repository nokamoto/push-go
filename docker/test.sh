#!/bin/bash

set -ex

names=`echo push-go-{app,endpoint,log,notification,subscription}`

for name in $names
do
    docker run -d --name $name nokamotohub/$name
done

# app
docker exec -it push-go-app pushcli app get | grep '{}'

docker exec -it push-go-app pushcli app set default

docker exec -it push-go-app pushcli app get | grep '{"apiKey":"default"}'

# endpoint
[[ -z $(docker exec -it push-go-endpoint pushcli endpoint get x) ]]

docker exec -it push-go-endpoint pushcli endpoint set x y

docker exec -it push-go-endpoint pushcli endpoint get x | grep '{"token":"y"}'

docker exec -it push-go-endpoint pushcli endpoint delete y

[[ -z $(docker exec -it push-go-endpoint pushcli endpoint get x) ]]

docker exec -it push-go-endpoint pushcli endpoint set x y

docker exec -it push-go-endpoint pushcli endpoint update y z

docker exec -it push-go-endpoint pushcli endpoint get x | grep '{"token":"z"}'

docker exec -it push-go-endpoint pushcli endpoint delete z

[[ -z $(docker exec -it push-go-endpoint pushcli endpoint get x) ]]

# log
[[ -z $(docker exec -it push-go-log pushcli log tail) ]]

docker exec -it push-go-log pushcli log info x

docker exec -it push-go-log pushcli log tail | grep '"text":"x"'

docker exec -it push-go-log pushcli log info y

docker exec -it push-go-log pushcli log tail | head -n 1 | grep '"text":"x"'

docker exec -it push-go-log pushcli log tail | tail -n 1 | grep '"text":"y"'

# subscription
[[ -z $(docker exec -it push-go-subscription pushcli subscription get x) ]]

docker exec -it push-go-subscription pushcli subscription subscribe x y

docker exec -it push-go-subscription pushcli subscription get x | grep '"id":"y"'

docker exec -it push-go-subscription pushcli subscription unsubscribe x y

[[ -z $(docker exec -it push-go-subscription pushcli subscription get x) ]]

for name in $names
do
    docker stop $name
    docker rm $name
done
