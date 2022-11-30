package dto

type IPRanges struct {
	Prefixes []IPRange `json:"prefixes"`
}

type IPRange struct {
	IPPrefix           string `json:"ip_prefix"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
}
