package conf

type BanConfig struct {
	Name             string
	Folder           string
	SupportShortMode bool
}

type Jconf struct {
	BanSize    int
	BanConfigs []BanConfig
}
