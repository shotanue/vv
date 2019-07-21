# protocol buffer, grpcを生成するプロジェクト
## 概要
protobuf管理を行う

protobufを各マイクロサービス間であったり、  
graphql側から呼び出す際に必要なprotobufを用意したり、  
管理が厄介なことがわかったので、  
ひとまず専用のプロジェクトを設けることにする。

マイクロサービスにした場合のprotoファイル管理のベストプラクティスがわからない...)

## set up
```bash
# protobuf をコンパイルできるdockerイメージをローカルに用意して、コンパイルする
# (イメージ名が被る場合は調整が必要..)
docker image build -t local/proto .
```

## how to compile
```bash
# ローカル環境でmakeを叩いてコンパイルする

# 既にコンパイル済みの各言語ファイルを削除する
make clean

# 全言語のprotobufをビルドする
make 
make all

# js
make js
# golang
make golang
```

