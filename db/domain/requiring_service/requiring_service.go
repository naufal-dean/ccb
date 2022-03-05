package requiring_service

import "github.com/naufal-dean/ccb/lib/utils"

type RequiringService struct {
	RgService string
	Service   string
	Endpoint  string
}

func (rs RequiringService) Sha256() string {
	return utils.Sha256(rs)
}
