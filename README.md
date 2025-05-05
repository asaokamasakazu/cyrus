# 🌐 Go Hello Word Web Server

これはGo言語で実装したシンプルなWebサーバーです。
いくつかの基本的なエンドポイントとロギング機能を備えています。

## ✅ 機能

- ルーティング機能
- HTTPリクエストのロギング（ミドルウェア実装）
- 3種類のエンドポイント：
    - `/` - テキストレスポンス
    - `/form` - HTMLフォームとPOSTリクエスト処理
    - `/json` - JSONレスポンス

## 🚀 実行方法
```bash
go run main.go
```

サーバーは http://localhost:8080 で起動します。

## 📝 使用例

### テキストレスポンス
```bash
curl http://localhost:8080
```

### GETリクエストでフォームを表示
```bash
curl http://localhost:8080/form
```

### POSTリクエストでフォームデータを送信
```bash
curl -X POST -d "name=foo bar" http://localhost:8080/form
```

### JSONレスポンス
```bash
curl http://localhost:8080/json
```
