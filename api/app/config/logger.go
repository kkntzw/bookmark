package config

import (
	"io/ioutil"
	"log"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// アプリケーションに共通するロガー。
// ロギングは log.Print[f|ln] ではなく config.Logger.[Debug|Info|Warn|Fatal|Error] で処理する。
var Logger *zap.Logger

// 設定ファイル ./resource/logging.yml に記載された内容を基に、ロガーを構築する。
// 設定ファイルが存在しない場合、あるいは内容が不正な場合は異常終了する。
func init() {
	// 設定ファイルを開く。
	file, err := ioutil.ReadFile("./resource/logging.yml")
	if err != nil {
		log.Fatalf("Failed to open the file \"logging.yml\": %v", err)
	}

	// 設定ファイルの内容を zap.Config にアサインする。
	var config zap.Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatalf("Failed to unmarshal the file \"logging.yml\": %v", err)
	}

	// 設定を反映したロガーを構築する。
	if Logger, err = config.Build(); err != nil {
		log.Fatalf("Failed to build the Logger: %v", err)
	}
	Logger.Debug("Initialized the Logger")
}
