package dateutil

import "time"

import _ "github.com/uniplaces/carbon"

var patternMap map[string]string

func init() {
	patternMap = make(map[string]string)
	patternMap["HH:mm"] = "15:04"
	patternMap["HH:mm:ss"] = "15:04:05"
	patternMap["HH时mm分"] = "15时04分"
	patternMap["HH时mm分ss秒"] = "15时04分05秒"
	patternMap["yyyy-MM"] = "2006-01"
	patternMap["yyyy-MM-dd"] = "2006-01-02"
	patternMap["yyyy-MM-dd HH:mm:ss"] = "2006-01-02 15:04:05"
	patternMap["yyyy/MM"] = "2006/01"
	patternMap["yyyy/MM/dd"] = "2006/01/02"
	patternMap["yyyy/MM/dd HH:mm:ss"] = "2006/01/02 15:04:05"
	patternMap["yyyy年MM月"] = "2006年01月"
	patternMap["yyyy年MM月dd日"] = "2006年01月02日"
	patternMap["yyyy年MM月dd日 HH时mm分ss秒"] = "2006年01月02日 15时04分05秒"
}

func FormatTime(time time.Time, pattern string) string {
	tmp := patternMap[pattern]
	if tmp == "" {
		tmp = pattern
	}
	return time.Format(tmp)
}
