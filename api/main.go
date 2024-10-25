package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 書籍情報の構造体
type Book struct {
	BookID                 int    `json:"book_id"`                  // 書籍ID
	ISBN                   string `json:"isbn"`                     // ISBN番号
	StorageLocation        string `json:"storage_location"`         // 保管場所
	LabelType              string `json:"label_type"`               // テプラの種類
	LabelSpec              string `json:"label_spec"`               // テプラの仕様
	AttachmentLocationSpec string `json:"attachment_location_spec"` // 貼付場所の仕様
	Notes                  string `json:"notes"`                    // 特記事項
}

// 書籍のデータ（仮データ）
var books = []Book{
	{BookID: 1, ISBN: "978-4-12-003398-8", StorageLocation: "裏表紙左下部", LabelType: "brother", LabelSpec: "12mm、ゴールド色、フォント：ImpactURW、サイズ：20", AttachmentLocationSpec: "オフィス", Notes: "〇〇のスタンプは上部に押印する"},
	{BookID: 2, ISBN: "978-4-00-008613-4", StorageLocation: "裏表紙左下部", LabelType: "brother", LabelSpec: "12mm、ゴールド色、フォント：ImpactURW、サイズ：20", AttachmentLocationSpec: "オフィス", Notes: "〇〇のスタンプは上部に押印する"},
}

// Vercelが呼び出すハンドラ関数
func Handler(w http.ResponseWriter, r *http.Request) {
	// Ginをリリースモードに設定
	gin.SetMode(gin.ReleaseMode)

	// Ginの新しいルーターを作成
	router := gin.New()

	// 書籍リストを取得するルート
	router.GET("/api/v1/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books) // 書籍リストをJSON形式で返す
	})

	// リクエストに対するレスポンスを処理
	router.ServeHTTP(w, r)
}
