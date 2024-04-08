package model

type UsageDTO struct {
	Bandwidth int      `json: bandwidth`
	Devices   []Device `json: devices`
}
