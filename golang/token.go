package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/google/uuid"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 数字越小，权限越大
const (
	AdminRole  = "0"
	WriterRole = "1"
	ReaderRole = "2"
)

const (
	sdkPrefix  = "NETLESSSDK_"
	roomPrefix = "NETLESSROOM_"
	taskPrefix = "NETLESSTASK_"
)

type SDKContent struct {
	Role string
}

type RoomContent struct {
	Role string
	Uuid string
}

type TaskContent struct {
	Role string
	Uuid string
}

// SDKToken 生成 sdk token
func SDKToken(accessKey string, secretAccessKey string, lifespan int64, content *SDKContent) string {
	m := map[string]string{
		"role": content.Role,
	}
	return createToken(sdkPrefix)(accessKey, secretAccessKey, lifespan, &m)
}

// RoomToken 生成 room token
func RoomToken(accessKey string, secretAccessKey string, lifespan int64, content *RoomContent) string {
	m := map[string]string{
		"role": content.Role,
		"uuid": content.Uuid,
	}
	return createToken(roomPrefix)(accessKey, secretAccessKey, lifespan, &m)
}

// TaskToken 生成 task token
func TaskToken(accessKey string, secretAccessKey string, lifespan int64, content *TaskContent) string {
	m := map[string]string{
		"role": content.Role,
		"uuid": content.Uuid,
	}
	return createToken(taskPrefix)(accessKey, secretAccessKey, lifespan, &m)
}

// bufferToBase64 buffer 转 base64
// 并格式化字符
func bufferToBase64(b []byte) string {
	str := base64.StdEncoding.EncodeToString(b)

	// 替换 "+" 到 "-"
	// 替换 "/" 到 "_"
	{
		r := strings.NewReplacer(
			"+", "-",
			"/", "_",
		)
		str = r.Replace(str)
	}

	// 移除末尾所以的 "="
	// 例如: hello== -> hello
	{
		r := regexp.MustCompile("=+$")
		str = r.ReplaceAllString(str, "")
	}

	return str
}

// getMapKeys 提取 map 里的 key
// 并且以 key 为主进行排序
func getMapKeys(m *map[string]string) []string {
	keys := make([]string, len(*m))

	i := 0
	for k := range *m {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	return keys
}

// encodeURIComponent 基于 url 编码对字符串进行编码
// 最终实现和 JavaScript 中的 encodeURIComponent 一致
func encodeURIComponent(str string) string {
	r := url.QueryEscape(str)

	// golang 的 url.QueryEscape 实现和 JavaScript 中的 encodeURIComponent 不一致的地方是对空格处理
	// golang 里会把空格转成 "+"，所以这里通过字符串替换，再把 "+" 转成 "%20"，以保证和 js 的实现一致
	r = strings.ReplaceAll(r, "+", "%20")
	return r
}

// stringify 序列化 Map
// 实现逻辑可参考: https://github.com/sindresorhus/query-string/blob/master/index.js#L284
func stringify(m *map[string]string) string {
	keys := getMapKeys(m)

	var arr []string
	for _, k := range keys {
		if (*m)[k] != "" {
			arr = append(arr, encodeURIComponent(k)+"="+encodeURIComponent((*m)[k]))
		}
	}

	return strings.Join(arr, "&")
}

// mergeMap 合并两个 map，返回新的 map 对象
// 在有相同值的情况下，m2 的优先级更高
func mergeMap(m1, m2 *map[string]string) map[string]string {
	result := make(map[string]string)

	for k, v := range *m1 {
		result[k] = v
	}

	for k, v := range *m2 {
		result[k] = v
	}

	return result
}

// jsonStringify 模范 JavaScript 中 JSON.stringify 的实现
// 因程序本身的 map 只有一层，而非嵌套 map，所以此方法没有实现嵌套 map 转 string
func jsonStringify(m *map[string]string) string {
	keys := getMapKeys(m)

	var body []string
	for _, k := range keys {
		body = append(body, "\""+k+"\""+":"+"\""+(*m)[k]+"\"")
	}

	result := strings.Join(body, ",")
	return "{" + result + "}"
}

// expireAt 根据当时时间及用户传入毫秒数，来计算过期时间
func expireAt(lifespan int64) string {
	return strconv.FormatInt(time.Now().UTC().UnixNano()/1e6+lifespan, 10)
}

// hmac256 实现 hmac 加密并转成 hex，实现可参考 nodejs: crypto/createHmac 及 digest("hex")
func hmac256(data string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)

	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}

// createToken 根据 prefix 生成相应的 token
func createToken(prefix string) func(string, string, int64, *map[string]string) string {
	return func(accessKey string, secretAccessKey string, lifespan int64, content *map[string]string) string {
		m := map[string]string{
			"ak":    accessKey,
			"nonce": uuid.Must(uuid.NewRandom()).String(),
		}
		m = mergeMap(content, &m)

		if lifespan > 0 {
			m["expireAt"] = expireAt(lifespan)
		}

		m["sig"] = hmac256(jsonStringify(&m), secretAccessKey)

		query := stringify(&m)
		return prefix + bufferToBase64([]byte(query))
	}
}
