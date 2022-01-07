// パッケージ mongorepo は MongoDB への接続を管理する。
package mongorepo

import (
	"context"
	"os"
	"time"

	"github.com/kkntzw/bookmark/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

// DBハンドラ。
// パッケージ bookmark/infrastructure/mongorepo の内部でのみアクセスを許可する。
var db *mongo.Database

// 環境変数に設定された URI へ接続し、DBハンドラを初期化する。
// URI の形式が不正な場合、あるいは認証情報が不正な場合は異常終了する。
// 環境変数 MONGO_URI には MongoDB の URI を設定する。
// 環境変数 MONGO_DB_NAME にはハンドリングを行う MongoDB のデータベース名を設定する。
func init() {
	config.Logger.Debug("Initializing the client of MongoDB...")

	// コンテキストを生成する。
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// クライアントを初期化する。
	uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		config.Logger.Fatal("Failed to initialize the client of MongoDB", zap.Error(err))
	}

	// MongoDB への疎通を確認する。
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		config.Logger.Fatal("Failed to ping MongoDB.", zap.Error(err))
	}

	// DBハンドラを設定する。
	name := os.Getenv("MONGO_DB_NAME")
	db = client.Database(name)

	config.Logger.Debug("Initialized the client of MongoDB")
}
