package hugo

import (
	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/hugolib"
	"github.com/spf13/viper"
)

func New(config *viper.Viper) (*hugolib.HugoSites, error) {
	var cfgDeps = deps.DepsCfg{
		Cfg: config,
	}
	site, err := hugolib.NewHugoSites(cfgDeps)
	if err != nil {
		panic(err)
	}

	err = site.Build(hugolib.BuildCfg{SkipRender: true})

	return site, err
}
