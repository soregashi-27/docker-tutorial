# ベストプラクティスアプリケーション

## graceful shutdownの確認

コンテナの起動

```
make run
```

定期的なリクエストの発行

```
watch curl -s localhost:8080
```

時間のかかるリクエストの発行

```
curl -s localhost:8080/heavy
```

コンテナの終了

```
make stop
```

## healthcheckの確認

Statusにhealthyと表示される

```
$ docker ps
CONTAINER ID        IMAGE                  COMMAND                  CREATED              STATUS                        PORTS                       NAMES
9b86ccfa1f13        bestpractice:v1        "/server"                About a minute ago   Up About a minute (healthy)   0.0.0.0:8080->8080/tcp      bestpractice
```
