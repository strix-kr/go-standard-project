package weather

type Weather struct {
    Lon    float32
    Lat    float32
    Status string
}

// 날씨를 조회합니다.
func GetWeather(lon, lat float32) (Weather Weather, err error) {
    // do something
    return
}
