package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"

	"url-shortener/common/log"
	"url-shortener/config"
	model "url-shortener/model/redis"
	"url-shortener/router"
	"url-shortener/router/handler"
)

var isDebugMode bool = true
var targetURL string = "https://www.google.com.tw/?hl=zh_TW"

func init() {

	configFilePath := "../_config/config_test.yaml"

	log.Initialize(isDebugMode)

	// 初始化設定檔
	config.Setup("yaml", configFilePath)
}

func Test_createOK(t *testing.T) {
	router := router.SetupRouter(isDebugMode)

	// post body string
	body := handler.ReqCreate{
		URL:       targetURL,
		ExpiredAt: time.Now().Add(5 * time.Second),
	}
	j, _ := json.Marshal(body)

	// mock redis data
	db, mock := redismock.NewClientMock()
	mock.Regexp().ExpectSetNX(`.*`, `.*`, 0).SetVal(true)
	mock.Regexp().ExpectExpireAt(`.*`, body.ExpiredAt).SetVal(true)
	model.NewURL(db)

	// call api
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/urls", strings.NewReader(string(j)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var respBody handler.RespCreate
	err := json.Unmarshal(w.Body.Bytes(), &respBody)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func Test_redirectOK(t *testing.T) {
	router := router.SetupRouter(isDebugMode)

	key := "redirectok"

	// mock redis data
	db, mock := redismock.NewClientMock()
	model.NewURL(db)
	model.URL.SetNX(key, targetURL)
	mock.ExpectGet(fmt.Sprintf("URL::%s", key)).SetVal(targetURL)

	// call api
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", key), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusFound, w.Code)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func Test_redirectNotFound(t *testing.T) {
	router := router.SetupRouter(isDebugMode)

	key := "notofund"

	// mock redis data
	db, mock := redismock.NewClientMock()
	model.NewURL(db)
	mock.ExpectGet(fmt.Sprintf("URL::%s", key)).RedisNil()

	// call api
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", key), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func Test_redirectExpired(t *testing.T) {
	router := router.SetupRouter(isDebugMode)

	key := "redirectexpired"

	// mock redis data
	db, mock := redismock.NewClientMock()
	model.NewURL(db)
	model.URL.SetNXWithExpireTime(key, targetURL, time.Now().Add(3*time.Second))
	mock.ExpectGet(fmt.Sprintf("URL::%s", key)).RedisNil()

	// call api
	time.Sleep(5 * time.Second)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", key), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
