# API 文件

| Method | API          | Description    |
| :----- | :----------- | :------------- |
| POST   | /api/v1/urls | 建立新的短網址 |
| GET    | /:uid        | 呼叫短網址     |


## [POST] /api/v1/urls

Request Body  
*application/json*  

| Prop     | Must | Type   | Description                                                                          |
| :------- | :--- | :----- | :----------------------------------------------------------------------------------- |
| url      | Y    | string | 待縮的長網址                                                                         |
| expireAt | N    | string | 短網址使用期限 <br> 可傳入日期時間字串 <br> 未指定或指定失敗將使用系統預設 (60 days) |

Response Body  
*application/json*  

| Prop     | Must | Type   | Description      |
| :------- | :--- | :----- | :--------------- |
| id       | Y    | string | 短網址 Unique ID |
| shortUrl | Y    | string | 短網址           |

Status Code

| Code | Description  |
| :--- | :----------- |
| 200  | OK           |
| 400  | 傳入參數有誤 |
| 500  | 系統異常     |

## [GET] /:uid

Request

```text
GET http://localhost:10005/:uid
```

Status Code

| Code | Description         |
| :--- | :------------------ |
| 302  | Redirect to OrigURL |
| 400  | uid empty           |
| 500  | 系統異常            |