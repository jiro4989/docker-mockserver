# docker-mockserver
単一のJSONを返すモックサーバをdocker上で起動するテスト

## 前提条件
下記のツールが必要です。

- docker-compose
- dep

## 使い方
下記のコマンドを実行してリポジトリを整える。

```bash
git clone https://github.com/jiro4989/docker-mockserver.git
cd docker-mockserver
dep ensure
ls -lah
```

下記と同様のファイル構造であることを確認する。

```bash
合計 40K
drwxr-xr-x  5 jiro jiro 4.0K  4月 14 11:46 .
drwxr-xr-x 14 jiro jiro 4.0K  4月 14 08:01 ..
-rw-r--r--  1 jiro jiro  589  4月 14 08:01 Gopkg.lock
-rw-r--r--  1 jiro jiro  812  4月 14 08:01 Gopkg.toml
-rw-r--r--  1 jiro jiro  125  4月 14 11:46 README.md
drwxr-xr-x  2 jiro jiro 4.0K  4月 14 08:02 cmd
-rw-r--r--  1 jiro jiro  217  4月 14 11:43 config.toml
-rw-r--r--  1 jiro jiro  218  4月 14 08:06 docker-compose.yml
drwxr-xr-x  4 jiro jiro 4.0K  4月 14 01:06 resp
drwxr-xr-x  3 jiro jiro 4.0K  4月 14 08:01 vendor
```

下記のコマンドを実行してコンテナを起動する。

```bash
docker-compose up -d
```

下記のコマンドを実行してコンテナにログインする

```bash
docker exec -it mock bash
```

下記のコマンドを実行してコンテナ内でプログラムのあるディレクトリに移動する。

```bash
cd /go/src/mock
```

下記のコマンドを実行してサーバを起動する。

```bash
go run cmd/main.go
```

別で端末を起動して、下記のコマンドを実行する。

```bash
curl 172.17.0.1:8088/get/user.json
```

jsonが返ってきたら成功です。

