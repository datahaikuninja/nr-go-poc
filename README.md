## nrugVol09 demo app

## setup
1. build docker image
```shell
docker build --tag nrug-vol-09:v0.0.x --build-arg GO_VERSION="$(go version | sed 's/^.*go\([0-9][^ ]*\) .*$/\1/g')" --build-arg NR_APM_APP_NAME=${NR_APM_APP_NAME} --build-arg NR_LICENSE_KEY=${NR_LICENSE_KEY} .
```

2. run app
```shell
docker run --publish 8080:8080 nrug-vol-09:v0.0.x
```
