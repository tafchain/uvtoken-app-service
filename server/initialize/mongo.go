package initialize

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"os"
	"service/server/global"
)

func Mongo() *mongo.Client {
	var err error
	uri := fmt.Sprintf("mongodb://%s:%s@%s", global.GVA_CONFIG.Mongo.Username, global.GVA_CONFIG.Mongo.Password, global.GVA_CONFIG.Mongo.Host)
	clientOptions := options.Client().ApplyURI(uri)

	mgoCli, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		global.GVA_LOG.Error("", zap.Any("err", err))
		os.Exit(0)
	}

	return mgoCli
}

func InitMongoIndex() {

}
