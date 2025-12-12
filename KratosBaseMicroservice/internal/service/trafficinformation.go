package service

import (
	"context"
	"fmt"
	"math"
	
	v1 "KratosBaseMicroservice/api/trafficinformation/v1
)

type TrafficInformation struct {
	bandwithUsage float64
	dailyTraffic  float64
}

func (s *TrafficInformation) GetBandwidth(ctx context.Context, req *v1.BandwithRequest) (string, error) {
	var data *TrafficInformation = &TrafficInformation{
		bandwithUsage: 0.0,
		dailyTraffic:  0.0,
	}
	data.bandwithUsage = 1.02354
	return fmt.Sprintf(" banwidth usage is %.2f GB/sec", math.Round(data.bandwithUsage*100)/100), nil
}

func (s *TrafficInformation) GetDailyTraffic(ctx context.Context, req *v1.DailyTrafficRequest) (string, error) {
	var data *TrafficInformation = &TrafficInformation{
		bandwithUsage: 0.0,
		dailyTraffic:  0.0,
	}
	data.dailyTraffic = 12.56789 
	return fmt.Sprintf(" daily traffic is %.2f Terabyte", math.Round(data.dailyTraffic*100)/100), nil
}
