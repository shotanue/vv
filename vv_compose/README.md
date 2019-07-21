# VV docker-compose

## 概要
開発用サーバーを立ち上げる。  
なるべく設定ファイルなど書かずにコマンド一発で立ち上がるように。

## How to
 ```bash
# 1. docker-compose用にAPIのトークンを環境変数を.envファイルで設定するのでテンプレをコピー
cp .env.template .env

# TMDB_TOKENを　https://www.themoviedb.org/settings/api　に行き、tokenを取得する
# tokenの取得には登録が必要
# 取得したtokenを.envに書き込む
vim .env

# 必要なtokenなどが用意できたらdocker-composeを立ち上げる
docker-compose up -d
 ``` 
