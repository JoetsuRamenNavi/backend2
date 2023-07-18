# backend2
go_restapi

### run
- go run *.go
- ./restapi
- sudo systemctl start goapi

### build
- go build -o restapi

### url

#### curl http://localhost:8080/users
- ユーザー全件取得

#### curl -X POST -H "Content-Type: application/json" -d '{"nickname":"John Doe", "password":"mypassword", "mail":"john.doe@example.com"}' http://localhost:8080/users
- ユーザー追加
- POSTでnickname,password,mailを

#### curl -X PUT -H "Content-Type: application/json" -d '{"nickname":"newNickname","password":"newPassword","mail":"new@example.com"}' http://localhost:8080/users/101
- ユーザー更新
- PUTメソッドを使っているので注意

#### curl http://localhost:8080/users/1
- ユーザー１件取得

#### curl http://localhost:8080/stores
- ストア全件取得

#### curl http://localhost:8080/storesvotes/ターム１
- 指定タームに参加していた店と投票数全件取得

#### curl -X POST http://localhost:8080/incrementvote/ターム１/1
- /ターム名/store_id でvoteが１増える
- ない場合は作成
