package aggregate

type HomeStatics struct {
	TodayVisitCount     int64 `json:"todayVisitCount,string"`     // 今日访问量
	TodayPlayCount      int64 `json:"todayPlayCount,string"`      // 今日播放量
	TodayVisitUserCount int64 `json:"todayVisitUserCount,string"` // 今日访问用户数
	TodayPlayUserCount  int64 `json:"todayPlayUserCount,string"`  // 今日播放用户数
	MonthVisitCount     int64 `json:"monthVisitCount,string"`     // 本月访问量
	MonthPlayCount      int64 `json:"monthPlayCount,string"`      // 本月播放量
	MonthVisitUserCount int64 `json:"monthVisitUserCount,string"` // 本月访问用户数
	MonthPlayUserCount  int64 `json:"monthPlayUserCount,string"`  // 本月播放用户数
}
