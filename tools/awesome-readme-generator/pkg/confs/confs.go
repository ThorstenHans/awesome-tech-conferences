package confs

import (
	"sort"
	"time"

	_ "gopkg.in/yaml.v2"
)

type ReadmeModel struct {
	ConferencesByYear map[int][]Conference
	Total             int
}

type ConferenceList []Conference

func (cl ConferenceList) SortByStartDateDesc() {
	sort.Slice(cl, func(i, j int) bool {
		return cl[i].Start.After(cl[j].Start)
	})
}

func NewReadmeModel(c ConferenceList) *ReadmeModel {
	c.SortByStartDateDesc()
	cc := ReadmeModel{
		ConferencesByYear: make(map[int][]Conference),
		Total:             len(c),
	}

	for _, conf := range c {
		y := conf.Start.Year()
		if _, ok := cc.ConferencesByYear[y]; !ok {
			cc.ConferencesByYear[y] = []Conference{conf}
		} else {
			cc.ConferencesByYear[y] = append(cc.ConferencesByYear[y], conf)
		}
	}
	return &cc
}

type Conference struct {
	Name    string
	Url     string
	Country string
	City    string
	Start   time.Time
	End     time.Time
	Twitter string `yaml:"twitter,omitempty"`
	Mode    string
	Topics  []string
	Cfp     CallForPapers `yaml:"cfp`
}

type CallForPapers struct {
	Start time.Time
	End   time.Time
	Url   string
}

func (c Conference) HasCallForPapers() bool {
	return !time.Time.IsZero(c.Cfp.Start) &&
		!time.Time.IsZero(c.Cfp.End) &&
		c.Cfp.Url != ""

}
