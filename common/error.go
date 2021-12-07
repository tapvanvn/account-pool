package common

import "errors"

var ErrAPInvalidForm = errors.New("Invalid form")
var ErrAPCampaignExisted = errors.New("Campaign existed")
var ErrAPCampaignNotExisted = errors.New("Campaign is not existed")
var ErrAPGeneratorNotExisted = errors.New("Generator is not existed")
