/*
 * @Author: lisheng
 * @Date: 2023-01-16 10:24:18
 * @LastEditTime: 2023-01-16 10:53:12
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /jf-go-kit/common/weather/weather_calc.go
 */
package weathercalc

type weather struct {
	HourPrecipitation float64
	Visibility        float64
	Humidity          float64
	Cloud             float64
}

func NewWeather(hourPrecipitation, visibility, humidity, cloud float64) *weather {
	return &weather{
		HourPrecipitation: hourPrecipitation,
		Visibility:        visibility,
		Humidity:          humidity,
		Cloud:             cloud,
	}
}

/**
 * @description: 计算天气状况
 * @return {*}
 * @author: liqiyuWorks
 */
func (w *weather) Condition() string {
	var condition string
	if 0.1 <= w.HourPrecipitation && w.HourPrecipitation <= 9.9 {
		condition = "小雨"
	} else if 10 <= w.HourPrecipitation && w.HourPrecipitation <= 24.9 {
		condition = "中雨"
	} else if 25 <= w.HourPrecipitation && w.HourPrecipitation <= 49.9 {
		condition = "大雨"
	} else if 50 <= w.HourPrecipitation && w.HourPrecipitation <= 99.9 {
		condition = "暴雨"
	} else if 100 <= w.HourPrecipitation && w.HourPrecipitation <= 249.9 {
		condition = "大暴雨"
	} else if 250 <= w.HourPrecipitation {
		condition = "特大暴雨"
	} else {
		if w.Visibility/1000 < 1 && w.Humidity > 90 {
			condition = "雾"
		} else if w.Visibility/1000 < 10 && w.Humidity < 80 {
			condition = "霾"
		} else {
			if 0 <= w.Cloud && w.Cloud <= 2 {
				condition = "晴"
			} else if 3 <= w.Cloud && w.Cloud <= 5 {
				condition = "少云"
			} else if 6 <= w.Cloud && w.Cloud <= 8 {
				condition = "多云"
			} else if 9 <= w.Cloud && w.Cloud <= 10 {
				condition = "阴"
			} else {
				condition = "晴"
			}
		}
	}

	return condition
}
