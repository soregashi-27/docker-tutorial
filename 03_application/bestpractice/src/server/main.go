package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/cybozu-go/log"
	"github.com/cybozu-go/well"
)

func main() {
	// ログのデフォルト設定を適用する。STDERRにログが出力される。
	well.LogConfig{}.Apply()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World\n")
	})

	// Graceful Shutdownの確認用に重たいリクエストを模擬
	mux.HandleFunc("/heavy", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		io.WriteString(w, "Hello World\n")
	})

	// ヘルスチェック用のエンドポイント
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "healthy\n")
	})

	// 環境変数でポート番号を設定可能
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	// well.HTTPServerを利用
	// - アクセスログをSTDERRに出力
	// - SIGTERMなどのシグナルを受けるとGraceful Shutdown
	s := &well.HTTPServer{
		Server: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}

	log.Info("Start listening", map[string]interface{}{
		log.FnHTTPHost: addr,
	})

	err := s.ListenAndServe()
	if err != nil {
		log.ErrorExit(err)
	}

	err = well.Wait()
	if err != nil && !well.IsSignaled(err) {
		log.ErrorExit(err)
	}
}
