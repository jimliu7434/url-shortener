# url-shortener

## 需求

* 將長網址縮短為 n 碼短網址
* 需能設置 ExpireAt 
* 在未達 ExpireAt 時，呼叫短網址將轉址到原長網址
* 超過 ExpireAt 後，回傳 404 notfound

## 設計思維

* 以 Restful API 為基礎
* 縮網址服務通常 **讀多寫少** ，以能快速搜索到 Original URL 為主要考量，**Key-Value DB** 是首選
* 有 ExpireAt 需求，需要可以在 DB 的每一筆資料上設置過期時間， **Redis** 有此功能，過期的資料會自動消失，不需另外實作清除
* 日後也許會增加 metrics 記錄使用次數等等，可以使用 Redis 做基本的 Counting

## 技術選型

目前需求較為單一，選擇單一 DB 即可完成需求，待需求越漸複雜，再漸次導入其他適合的工具或資料庫  

其他 lib / module 則是我比較常用的  

* DB: **Redis**
* Web: **Gin**
* Logger: **logrus** + **lumberjack** (logrotate)

## API

[文件](./api.md)

## Run / Test / Build

need: **Golang 1.16 up**

Run  

```bash
go run main.go
```

Test  

```bash
go test ./_test/main_test.go
```

Build (linux)  

```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags="jsoniter" -o service .
```

## TODO & Thought (僅有想法，尚未實作)

* [TODO] 是否需要記錄 Metrics？ 例如 locale , total access times , access times each client ip
* [TODO] 防堵大量呼叫「不存在」的短網址，可以在每一個短網址 uid 末尾加上 1~2 碼 CheckSum。只要 CheckSum 與前面 uid 算出的內容不符合，則直接回傳 **404NotFound**，避免查詢 DB。等於使用 Application 層的 CPU 消耗來換取減少無謂存取 DB (前提：Application 層的擴展比較方便)
* [TODO] 檢查 OrigURL 是否是可用的 URL，避免 Redirect 失敗
* [TODO] 檢查 OrigURL 是否指向本服務，避免造成無限 Redirect 循環
* [Thought] 同 OrigURL + 同 ExpireAt 是否需要產生相同結果？ 此時 uid 是否適合用亂數，或需改用 URL hash？
* [Thought] 超過 ExpireAt 的 URL 是否需在 DB 留存記錄？
