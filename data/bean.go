package data

type SrcUrl string
type DstUrl string

type PointList struct {
	Port   int               `yaml:"port"`
	Points map[SrcUrl]DstUrl `yaml:"points"`
}

type PortList struct {
	PointLists []PointList `yaml:"point_lists"`
}
