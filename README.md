# Go言語の勉強リポジトリ

Go言語の触りながら「クリーンアーキテクチャとは何ぞや？」を理解するための勉強リポジトリ。

## 使ってるもの

- フレームワーク:[Gin](https://github.com/gin-gonic/gin)
- ORM:[GORM](https://github.com/go-gorm/gorm)
- ORM:[Air](https://github.com/cosmtrek/air)

## 使用方法

```bash
# MySQLコンテナ起動
docker-compose up -d mysql

# Appコンテナ起動 TODO::コンテナの起動順を制御するスクリプトは今後対応
docker-compose up -d app

# 疎通確認
curl --request GET 'http://localhost:8080/ping'
> {"message":"pong"}
```

## 参考サイト

- http://psychedelicnekopunch.com/archives/1308
- https://qiita.com/ogady/items/34aae1b2af3080e0fec4
- https://github.com/bxcodec/go-clean-arch/tree/v2/author
- https://tech.mirrativ.stream/entry/2020/11/30/142354
- https://qiita.com/stranger1989/items/914bef8cc6ad77ef8350
- https://gorm.io/ja_JP/docs/connecting_to_the_database.html
- https://gorm.io/ja_JP/docs/query.html
- https://zenn.dev/nobonobo/articles/0b722c9c2b18d5
- https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode#go-get-%E3%81%AF%E3%82%AA%E3%83%AF%E3%82%B3%E3%83%B3%EF%BC%9F
- https://nrslib.com/clean-architecture/