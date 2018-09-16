# Task service

[API](https://github.com/sdaf47/go-knowledge-base/blob/master/small_programms/task/api.http)

## Docker run

```
cd ./small_programms/cmd/task
docker build -t task-manager .

docker run -it -P task-manager
```