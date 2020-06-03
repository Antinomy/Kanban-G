package conf

type BanConfig struct {
	Name             string
	Folder           string
	SupportShortMode bool
	Prefix           string
}

type Jconf struct {
	BanSize    int
	BanConfigs []BanConfig
}
