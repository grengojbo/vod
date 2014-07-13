package vod

type Channels struct {
	Name      string `json:"name"`
	DirConfig string `json:"conf"`
	DirTs     string `json:"ts"`
}

type VodServer struct {
	Name        string     `json:"name"`
	DirVideoVod string     `json:"vod"`
	DirConvTs   string     `json:"src"`
	DirTmp      string     `json:"tmp"`
	channels    []Channels `json:"channels"`
}

// func init() {
// }

func (this *Channels) SetPlayList(name string) (err error) {
	return nil
}
