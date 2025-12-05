package log

import (
	"fmt"
	"log/slog"
	"os"
)

// Log 構造体
// *slog.Logger を埋め込むことで、Info, Error, InfoContext などの全メソッドをそのまま使えるようにする.
type Log struct {
	*slog.Logger
}

func NewLogger() *Log {
	// JSONハンドラーの設定
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,

		// ログの属性（キーバリュー）を書き換える関数
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			// キーが "error" で、中身が error型 の場合
			if attr.Key == "error" {
				if err, ok := attr.Value.Any().(error); ok {
					// slog.Group を使って、JSONの中でオブジェクト(入れ子)にする
					return slog.Group("error",
						// 1. 短いメッセージ (一覧表示用)
						slog.String("message", err.Error()),

						// 2. 詳細なスタックトレース (詳細確認用)
						// ここには改行コード(\n)が含まれるが、JSONとしては正しい
						slog.String("stack_trace", fmt.Sprintf("%+v", err)),
					)
				}
			}

			return attr
		},
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return &Log{Logger: logger}

	// // グローバルロガーにも設定しておく（サードパーティライブラリなどが使う場合のため）

	// return &Log{
	// 	Logger: logger,
	// }
}
