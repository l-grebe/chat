package chat

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// YoudaoTranslate 翻译接口
const YoudaoTranslate = "http://openapi.youdao.com/api"
const YoudaoTimeout = time.Second * 10 // 超时时间

type youdaoResponse struct {
	ErrorCode   string   `json:"errorCode"`
	Query       string   `json:"query"`
	Translation []string `json:"translation"`
}

func generateSign(query string, salt string) string {
	h := md5.New()
	h.Write([]byte(DefaultSetting.YouDaoAppKey + query + salt + DefaultSetting.YouDaoAppSecret))
	return hex.EncodeToString(h.Sum(nil))
}

func Translate(query string, form string, to string) string {
	salt := strconv.Itoa(int(time.Now().UnixNano())) // 随机数
	sign := generateSign(query, salt)                // 签名
	//salt := strconv.FormatInt(time.Now().UnixNano(), 10)
	//sign := md5.Sum([]byte(DefaultSetting.YouDaoAppKey + query + salt + DefaultSetting.YouDaoAppSecret))
	params := url.Values{}
	params.Add("q", query)
	params.Add("from", form)
	params.Add("to", to)
	params.Add("appKey", DefaultSetting.YouDaoAppKey)
	params.Add("salt", salt)
	params.Add("sign", sign)
	// params.Add("signType", "v3")

	client := http.Client{Timeout: YoudaoTimeout}
	resp, err := client.PostForm(YoudaoTranslate, params)
	if err != nil {
		fmt.Println("Request failed:", err)
		return query
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Parsing the response failed:", err)
		return query
	}

	data := new(youdaoResponse)
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Parsing the response failed:", err)
		return query
	}

	return data.Translation[0]
}

// TranslateNotDoubleBackticks Translation does not include text surrounded by back quotes
func TranslateNotDoubleBackticks(query string, form string, to string) string {
	record := make(map[string]string)
	key := ""
	backtickCnt := 0
	newQuery := ""
	for _, ch := range query {
		if ch == '`' {
			backtickCnt += 1
			if backtickCnt == 0 {
				key += string('`')
				if _, ok := record[key]; !ok {
					record[key] = "NX" + strconv.Itoa(backtickCnt/2)
				}
				newQuery += record[key]
				key = ""
				continue
			}
		}
		if backtickCnt == 1 {
			key += string(ch)
		} else {
			newQuery += string(ch)
		}
	}

	res := Translate(newQuery, form, to)
	for k, v := range record {
		res = strings.Replace(res, v, fmt.Sprintf(" %s ", k), -1)
	}
	return res
}

// TranslateNotCode Translation does not include the text of the code
func TranslateNotCode(query string, form string, to string) string {
	queryList := strings.Split(query, "```")
	res := ""
	for idx, q := range queryList {
		if idx%2 == 0 {
			res += TranslateNotDoubleBackticks(q, form, to)
		} else {
			res += fmt.Sprintf("\n\n```%s```", q)
		}
	}
	return res
}
