package service

import (
    // Change alias from 'v1' to 'trafficv1'
    trafficv1 "KratosBaseMicroservice/api/trafficInformation/v1"
    "context"
)

/*type TrafficInformation struct {
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
}*/

// The struct name should be consistent and reflect the Kratos service.
// It should also satisfy the generated interface pb.TrafficinformationServer.
type TrafficinformationService struct {
    trafficv1.UnimplementedTrafficinformationServer
}

// Update all function signatures and calls to use the new alias
func (s *TrafficinformationService) GetBandwidth(ctx context.Context, req *trafficv1.BandwithRequest) (*trafficv1.BandwithResponse, error) {
    // ... logic
    return &trafficv1.BandwithResponse{
        // ...
    }, nil
}

func (s *TrafficinformationService) GetDailyTraffic(ctx context.Context, req *trafficv1.DailyTrafficRequest) (*trafficv1.DailyTrafficResponse, error) {
    // ... logic
    return &trafficv1.DailyTrafficResponse{
        // ...
    }, nil
}
