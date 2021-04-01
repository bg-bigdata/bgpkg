/**
 * @Author: gaoshiyao
 * @Description: serial_number
 * @File:  serial_number
 * @Date: 2021/03/31 19:43
 */

package tools

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

type SerialNum struct {
	RedisConn *redis.Client // redis client
	Key       string        // redis key
	Max       int           // 几位数
}

func NewSerialClient(redisClient *redis.Client, key string, max int) *SerialNum {
	if max == 0 {
		max = 5
	}
	return &SerialNum{RedisConn: redisClient, Key: key, Max: max}
}

func (s SerialNum) GetDayNumber() (no string, err error) {
	// 获取key
	var num int64
	result, err := s.RedisConn.Get(s.Key).Result()
	if result == "" {
		// 设置过期时间
		lastTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 23:59:59")
		a := time.Duration(int(lastTime.Unix()) - int(time.Now().Unix()))
		// 添加key
		s.RedisConn.Set(s.Key, 0, time.Second*a)
	}
	num, err = s.RedisConn.Incr(s.Key).Result()
	if err != nil {
		return
	}
	if float64(num) > math.Pow(10, float64(s.Max)) {
		err = errors.New("超过今日最大流水号限制！")
		return
	}
	numStr := strconv.Itoa(int(num))
	no = time.Now().Format("20060102150405")
	len := s.Max - len(numStr)
	for i := 1; i < len; i++ {
		no += "0"
	}
	no += numStr
	return
}
