# VV
映画検索アプリ

## 概要
映画を検索するWebアプリです。

[Webフロントエンド（SPA）の実装はこちら](https://github.com/shotanue/vv_web)


個人的に触ってみたかった技術を色々入れている実験的プロジェクトです。

例えば、  
バックエンド周りではマイクロサービス化、grpcを使った通信、GraphQLに興味があったので取り込んでいます。

BFF <-> マイクロサービス <-> マイクロサービス の通信はgrpcを使っています。

フロント <-> サーバー(BFF) の通信はgraphqlを採用しています。

また、マイクロサービス化にあたって、バックエンドの実装は今の所全てgolangを使っています。  
マイクロサービス内でもクリーンアーキテクチャを意識した構成にしており、DIにgoogleのwireを使っています。

(リソースとなるAPIに対してプロキシ・キャッシュしているだけなので、あまりテストは書けていませn...。  
せっかくinfra層の切り替えとか楽なのに...)

これらサーバーの実装に対してフロントは、Angularを使ってSPAにしています。
  
Angularについては何かこだわりがあるわけではないですが、
Angularのお作法に従って色々よしなにしてもらえれば、単純にサーバーの実装に集中できるかなと思った次第です。

(なお、vvについて意味は特にないです。  
...が、頭につけるnamespaceとして使いやすいというのはあります。)

## 構成

<img width="923" alt="vv_architecture" src="https://user-images.githubusercontent.com/22065594/61609386-ee03b800-ac90-11e9-832d-b69cac59a97f.png">



## 主に使っているもの
- フロントエンド
    - Angular
    - graphql(Apollo)
    - stylus
    
- BFF
    - golang
        - 実装言語
    - graphql
        - フロントとの通信
    - grpc
        - マイクロサービスとの通信
    
- 各マイクロサービス
    - golang
    - grpc
    - wire
    
    


##　動かし方
```bash
cd vv_compose
docker-compose up -d
```

##　実際に動いているところ(WIP)
https://mitaii.com

