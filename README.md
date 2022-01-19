# otp-service

## 需求

[requirements](./requirements.md)

## 包版

[build](./build.md)

## 部署

[deploy](./deploy.md)

## API

| API     | Method | Description                      |
| :------ | :----- | :------------------------------- |
| /create | POST   | 產生 OTP 並寄出 email 或寄發簡訊 |
| /verify | POST   | 驗證 OTP                         |

### API Key

HTTP POST **querystring** or **header** 需埋入指定 API Key 以驗證為合法服務  
若未依規定埋入對應 API Key ，會收到 `403` forbidden 代碼  

測試 apikey 如下：  

| Service | Key                                  |
| :-----: | :----------------------------------- |
|  test   | 6e4d8aab-1ef5-487d-a72e-f6e5ed3655ee |

ex:  

   ```txt
   [POST] /create?apikey=6e4d8aab-1ef5-487d-a72e-f6e5ed3655ee
   [POST] /verify?apikey=6e4d8aab-1ef5-487d-a72e-f6e5ed3655ee
   ```

### [POST] /create

* request

   ```json
   [POST] /create
   {
       "email": "abc@mitake.com.tw",
       "mobile": "0912345678",
       "uid": "hasheduidstring",
       "metrics": "platform=ios&device=phone&os=15.0.0"
   }
   ```

   | Props   | Type   | Must | Description     |
   | :------ | :----- | :--- | :-------------- |
   | email   | string | N    | 寄發 email 目標 |
   | mobile  | string | N    | 寄發 簡訊 目標  |
   | uid     | string | Y    | 登入帳號        |
   | metrics | string | N    | 裝置資訊        |

*※ email & mobile 可以擇一輸入或都輸入，若都沒有輸入則會回傳錯誤訊息*

* response

   ```json
   {
       "id": "1E5Pc",
       "result": true
   }
   ```

   ```json
   {
       "result": false,
       "reason": "MAXTIMES"
   }
   ```

   ```json
   {
       "result": false,
       "reason": "ERROR"
   }
   ```

   | Props  | Type   | Must | Description                                          |
   | :----- | :----- | :--- | :--------------------------------------------------- |
   | result | bool   | Y    | 寄發結果                                             |
   | id     | string | N    | OTP ID，亂數大小寫英數字，長度 4~6 <br> 執行成功才會產生， |
   | reason | string | N    | 失敗原因                                             |

   | Reason    | Description                                                     |
   | :-------- | :-------------------------------------------------------------- |
   | ERROR     | 系統錯誤                                                        |
   | FORBIDDEN | APIKey 有誤                                                     |
   | BADREQ    | 輸入參數內容有誤 <br> 例如 email & mobile 均未傳入，或者 uid 未傳入 |
   | MAXTIMES  | 用戶已達每日次數上限                                            |

### [POST] /verify

* request

   ```json
   {
       "id": "1E5Pc",
       "otp": "123456"
   }
   ```

   | Props | Type   | Must | Description                        |
   | :---- | :----- | :--- | :--------------------------------- |
   | id    | string | Y    | OTP ID，亂數大小寫英數字，長度 4~6 |
   | otp   | string | Y    | OTP 答案字串                       |

* response

   ```json
   {
       "id": "1E5Pc",
       "result": true
   }
   ```

   ```json
   {
       "id": "1E5Pc",
       "result": false,
       "reason": "FAILED"
   }
   ```

   ```json
   {
       "id": "1E5Pc",
       "result": false,
       "reason": "EXPIRED"
   }
   ```

   ```json
   {
       "id": "1E5Pc",
       "result": false,
       "reason": "USED"
   }
   ```

   | Props  | Type   | Must | Description         |
   | :----- | :----- | :--- | :------------------ |
   | id     | string | Y    | 回送呼叫時的 OTP ID |
   | result | bool   | Y    | 驗證結果            |
   | reason | string | N    | 失敗原因            |

   | Reason    | Description               |
   | :-------- | :------------------------ |
   | ERROR     | 系統錯誤                  |
   | FORBIDDEN | APIKey 有誤               |
   | BADREQ    | 輸入參數內容有誤          |
   | FAILED    | 驗證失敗： OTP 答案錯誤   |
   | NOTFOUND  | 驗證失敗： 找不到 OTP ID  |
   | USED      | 驗證失敗： OTP 已被使用過 |
   | EXPIRED   | 驗證失敗： OTP 已過期     |
