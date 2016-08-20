package redis

import (
	. "global"
	"strings"
	"time"

	"util"

	"github.com/garyburd/redigo/redis"
)

var host string
var port string
var passwd string
var connTimeout time.Duration
var readTimeout time.Duration
var writeTimeout time.Duration

func init() {
	// 默认监听 127.0.0.1:3279
	serverConf := util.Conf("db")

	host = serverConf.MustValue("redis.host", "127.0.0.1").(string)
	port = serverConf.MustValue("redis.port", "6379").(string)

	connTimeout = time.Duration(2) * time.Second
	readTimeout = time.Duration(2) * time.Second
	writeTimeout = time.Duration(2) * time.Second

	passwd = serverConf.MustValue("redis.password", "").(string)
}

const (
	// c 表示客户端
	KeyPrefix = "c:"

	// key 值，为了节约内存
	KeyWeiXin            = "w"
	KeyWeiXinAccessToken = "a"
	KeyWeiXinJsToken     = "j"

	KeyOrder = "o"
)

type RedisClient struct {
	redis.Conn
	lastKey string // 最后一次查询的 key, 用来辅助 key() 函数
	err     error
}

func NewRedisClient() *RedisClient {
	if 0 >= len(passwd) {
		// log err
		return &RedisClient{Conn: nil, err: ReidsPasswdEmpty}
	}

	// 连接不可重用 ！！！！
	conn, err := redis.DialTimeout("tcp", host+":"+port, connTimeout, readTimeout, writeTimeout)
	if err != nil {
		return &RedisClient{Conn: conn, err: err}
	}

	if _, err = conn.Do("AUTH", passwd); err != nil {
		return &RedisClient{Conn: conn, err: err}
	}

	return &RedisClient{Conn: conn, lastKey: "", err: nil}
}

// 写入 key
func (this *RedisClient) Key(key ...string) *RedisClient {
	if this.err != nil {
		return this
	}

	if 0 >= len(key) {
		this.err = ReidsKeyEmpty
	}

	this.lastKey = KeyPrefix + strings.Join(key, ":")
	return this
}

func (this *RedisClient) SET(key string, val interface{}, expireSeconds int64) error {
	if 0 < len(strings.Trim(key, "")) {
		// 设置 key
		this.Key(key)
	}
	if this.err != nil {
		return this.err
	}

	args := make([]interface{}, 2, 4)
	args[0] = this.lastKey
	args[1] = val

	if expireSeconds > 0 {
		args = append(args, "EX")
		args = append(args, expireSeconds)
	}
	_, err := redis.String(this.Conn.Do("SET", args...))
	return err
}

func (this *RedisClient) GET(key string) string {

	if 0 < len(strings.Trim(key, "")) {
		// 设置 key
		this.Key(key)
	}
	if this.err != nil {
		return ""
	}

	val, err := redis.String(this.Conn.Do("GET", this.lastKey))
	if err != nil {
		return ""
	}

	return val
}

func (this *RedisClient) DEL(key string) error {
	if 0 < len(strings.Trim(key, "")) {
		// 设置 key
		this.Key(key)
	}
	if this.err != nil {
		return this.err
	}

	_, err := redis.Int(this.Conn.Do("DEL", key))

	return err
}

func (this *RedisClient) HSET(key, field, val string) error {
	if 0 < len(strings.Trim(key, "")) {
		// 设置 key
		this.Key(key)
	}
	if this.err != nil {
		return this.err
	}

	_, err := redis.Int(this.Conn.Do("HSET", key, field, val))
	return err
}

func (this *RedisClient) HGETALL(key string) (map[string]string, error) {
	if 0 < len(strings.Trim(key, "")) {
		// 设置 key
		this.Key(key)
	}
	if this.err != nil {
		return nil, this.err
	}

	return redis.StringMap(this.Conn.Do("HGETALL", key))
}

func (this *RedisClient) INCR(key string) (int64, error) {
	if 0 < len(strings.Trim(key, "")) {
		// 设置 key
		this.Key(key)
	}
	if this.err != nil {
		return 0, this.err
	}

	return redis.Int64(this.Conn.Do("INCR", key))
}

func (this *RedisClient) HDEL(key, field string) error {
	if 0 < len(strings.Trim(key, "")) {
		// 设置 key
		this.Key(key)
	}
	if this.err != nil {
		return this.err
	}

	_, err := redis.Int(this.Conn.Do("HDEL", key, field))

	return err
}

func (this *RedisClient) Close() {
	this.Conn.Close()
}
