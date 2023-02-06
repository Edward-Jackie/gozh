package ip

import "github.com/Edward-Jackie/gotool/pkg/ghostbot"

type IpInfo struct {
	Country   string `json:"country"`
	IspId     string `json:"isp_id"`
	QueryIp   string `json:"queryIp"`
	City      string `json:"city"`
	Ip        string `json:"ip"`
	Isp       string `json:"isp"`
	RegionId  string `json:"region_id"`
	Region    string `json:"region"`
	CountryId string `json:"country_id"`
	CityId    string `json:"city_id"`
}

type GetIpReply struct {
	Code int     `json:"code"`
	Data *IpInfo `json:"data"`
	Msg  string  `json:"msg"`
}

func GetIpInfo(ip string) (*IpInfo, error) {
	reply := &GetIpReply{}
	err := ghostbot.NewClient().Get("https://ip.taobao.com/service/getIpInfo.php?ip=" + ip).ToStruct(reply)
	if err != nil {
		return nil, err
	}
	return reply.Data, nil
}
