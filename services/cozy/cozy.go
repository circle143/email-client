package cozy

import "circledigital.in/api/utils/common"

type cozyService struct{}

func CreateCozyService(_ common.IApp) common.IService {
	return &cozyService{}
}
