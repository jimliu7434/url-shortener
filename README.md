# url-shortener

## 需求

* 將長網址縮短為 n 碼短網址代碼
* 需能設置 ExpireAt 
* 在未達 ExpireAt 時，呼叫短網址將轉址到原長網址
* 超過 ExpireAt 後，回傳 404 notfound

## 設計思維

* 以 Restful API 為基礎
* 縮網址服務通常讀多寫少，以能快速搜索到 Original URL 為主要考量，Key-Value Pair DB 是首選
* 有 ExpireAt 需求，需要可以在 DB 的每一筆資料上設置過期時間， Redis 有此功能，過期的資料會自動消失，不需另外實作清除
* 日後也許會增加 metrics 記錄使用次數等等，可以使用 Redis 做基本的 Counting

## 後續討論

* 同 OrigURL + 同 ExpireAt 是否需要產生相同結果？ 此時 uid 是否適合用亂數，或需改用 URL hash？
* 是否需要記錄 Metrics？ 例如 locale , click times
* 超過 ExpireAt 的 URL 是否需在 DB 留存記錄？
* 如果服務遭遇大量呼叫「不存在」的短網址，可以在每一個短網址末尾加上 1~2 碼 CheckSum，只要 CheckSum 與前面 uid 內容不符合者，直接回傳 404，以減少大量查詢 DB 的行為發生，等於使用 Application 層的 CPU 消耗來換取少無謂的存取 DB 層，畢竟 Application 層的擴展是比較方便的
