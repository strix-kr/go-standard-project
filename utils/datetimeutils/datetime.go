package datetimeutils

import (
    "time"
)

const (
    FormatYMDHMS = time.RFC3339 // 2006-01-02T15:04:05Z07:00
    FormatYMD    = "2006-01-02"
)

// datetime utils 에서 사용할 timezone 입니다.
var loc *time.Location

// application 에서 사용할 timezone 을 설정합니다.
func init() {
    var err error
    if loc, err = time.LoadLocation("Asia/Seoul"); err != nil {
        panic(err)
    }
}

// RFC3339 를 KST 로 변환합니다.
func Parse(input string) (time.Time, error) {
    return time.ParseInLocation(FormatYMDHMS, input, loc)
}

// 현재 시각을 조회합니다.
func Now() time.Time {
    return time.Now().In(loc)
}
