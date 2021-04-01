/**
 * @Author: gaoshiyao
 * @Description: serial_number_test
 * @File:  serial_number_test
 * @Date: 2021/03/31 19:52
 */

package tools

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"strings"
	"testing"
)

func TestSnum(t *testing.T) {
	redis, err := GetRedisClient()
	if err != nil {
		logrus.Error(err.Error())
	}
	client := NewSerialClient(redis, "gsy", 2)
	num, err := client.GetDayNumber()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(num)
	}
}

//GetRedisClient 获取connect
func GetRedisClient() (rcon *redis.Client, e error) {
	addr := "localhost:1128"
	if strings.HasPrefix(addr, "redis://") {
		addr = strings.TrimPrefix(addr, "redis://")
	}
	rcon = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	if _, e := rcon.Ping().Result(); e != nil {
		logrus.Errorf(`connect redis error :%s`, e.Error())
		return rcon, e
	}
	fmt.Println("connect redis success!")
	return
}
