package services

import (
    "github.com/strix-kr/go-standard-project/external/weather"
    "log"
)

//
func GetWeaterByLocation(lon, lat float32) {
    
    // 외부 서비스를 호출합니다.
    w, err := weather.GetWeather(lon, lat)
    if err != nil {
        log.Fatal(err)
        return
    }
    
    log.Println(w)
    
}
