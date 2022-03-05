package required_service

import "github.com/naufal-dean/ccb/lib/utils"

type RequiredService struct {
	Endpoint   string
	RdService  string
	RdEndpoint string
}

func (rs RequiredService) Sha256() string {
	return utils.Sha256(rs)
}
