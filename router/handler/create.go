package handler

import (
	"fmt"
	"net/http"
	"time"
	"url-shortener/common/log"
	util "url-shortener/common/util"
	"url-shortener/config"
	model "url-shortener/model/redis"

	"github.com/gin-gonic/gin"
)

type ReqCreate struct {
	URL       string    `json:"url" binding:"required"`
	ExpiredAt time.Time `json:"expireAt"`
}

type RespCreate struct {
	ID       string `json:"id"`
	ShortURL string `json:"shortUrl"`
}

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var in ReqCreate
		if err := c.ShouldBind(&in); err != nil {
			log.TraceLog.Infof("[Create] body: %s", err.Error())
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// TODO: check in.URL is a real URL ?
		// TODO: check in.URL is an infine-redirect-loop URL ?

		// generate new uid
		idLen := config.Root.Application.IDLen
		uid := util.GenerateID(idLen)

		// set to db if id not exists
		if in.ExpiredAt.Unix() <= time.Now().Unix() {
			err := model.URL.SetNX(uid, in.URL)
			if err != nil {
				log.TraceLog.Infof("[Create] redis err: %s when in= %+t", err.Error(), in)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		} else {
			err := model.URL.SetNXWithExpireTime(uid, in.URL, in.ExpiredAt)
			if err != nil {
				log.TraceLog.Infof("[Create] redis err: %s when in= %+t", err.Error(), in)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}

		c.JSON(http.StatusOK, RespCreate{
			ID:       uid,
			ShortURL: fmt.Sprintf("%s%s", config.Root.Application.Domain, uid),
		})
	}
}
