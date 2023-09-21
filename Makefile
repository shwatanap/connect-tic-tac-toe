STANDARD-CLIENT=./bin/standard-client # 標準のgRPCによるクライアント
STANDARD-SERVER=./bin/standard-server # 標準のgRPCによるサーバー
CONNECT-CLIENT=./bin/connect-client   # connect-goによるクライアント
CONNECT-SERVER=./bin/connect-server   # connect-goによるサーバー

# コード生成
gen:
	cd proto && buf generate

# 実行
run-standard-client: ${STANDARD-CLIENT}
	$?

run-standard-server: ${STANDARD-SERVER}
	$?

run-connect-client: ${CONNECT-CLIENT}
	$?

run-connect-server: ${CONNECT-SERVER}
	$?

# ビルド
build: ${STANDARD-CLIENT} ${STANDARD-SERVER} ${CONNECT-CLIENT} ${CONNECT-SERVER}

# 実行ファイルの削除
clean:
	rm -rf ./bin

# リビルド
re: clean build

${STANDARD-CLIENT}:
	go build -o $@ ./src/standard/cmd/client/main.go

${STANDARD-SERVER}:
	go build -o $@ ./src/standard/cmd/server/main.go

${CONNECT-CLIENT}:
	go build -o $@ ./src/connect/cmd/client/main.go

${CONNECT-SERVER}:
	go build -o $@ ./src/connect/cmd/server/main.go
