# recaptcha-google-sample

## 手順

- http://www.google.com/recaptcha/admin にアクセスしてAPIキーペアを取得する
- client/index.htmlにAPIキーを設定する
- `make serve`してhttp://localhost:8000/にアクセスしてconsole.logで出力されるtokenを取得する
- api/main.goにAPIキーとtokenを設定する
- `make check`してreCAPTCHAの結果を確認する
