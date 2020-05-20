package conf

type BanConfig struct {
	Name   string
	Folder string
}

type Jconf struct {
	BanSize    int
	BanConfigs []BanConfig
}
