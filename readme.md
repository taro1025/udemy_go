go get "gopkg.in/go-ini/ini.v1"
go initはコンフィグだったり初期化に必要なファイルを読み込むのに必要

DBを使うときはドライバが必要
SQLite３の場合
go get github.com/mattn/go-sqlite3"

UUId
g unique id
go get "github.com/google/uuid"

# Railsでいうと

routes.rb >> server.goのStartMainServer()

model.rb >> model配下の

    type User struct {
    ID        int
    UUID      string
    Name      string
    Email     string
    PassWord  string
    CreatedAt time.Time
    Todos     []Todo

モデルメソッドはその下に。


controller >> controller配下

その他バックエンド的なことはserver.rb

ex. session, htmlのレイアウトをくっつける。

