# study-go

[A Tour of Go](https://go.dev/tour/list "A Tour of Go") をやっていき。

## 環境作成

```
$ make build
```
or
```
$ docker-compose up -d --build
```

## 環境に入る

```
$ make in
```
or
```
$ docker-compose exec go bash
```

## 実行

```
# 環境に入ってから
go run hello.go
```