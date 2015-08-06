package vod

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/astaxie/beego/httplib"
)

type Channels struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	DirConfig   string    `json:"conf"`
	DirTs       string    `json:"ts"`
	DirSrc      string    `json:"src"`
	DirVideoVod string    `json:"vod"`
	Files       []string  `json:"files"`
	Transcode   []string  `json:"transcode"`
	Count       int       `json:"count"`
	CountTrans  int       `json:"count_trans"`
	Errors      ErrorMess `json:"errors"`
	ApiUrl      string    `json:"apiUrl"`
	Format      string    `json:"format"`
	Size        string    `json:"size"`
}

type VodServer struct {
	Name        string     `json:"name"`
	DirVideoVod string     `json:"vod"`
	DirConvTs   string     `json:"src"`
	DirTmp      string     `json:"tmp"`
	channels    []Channels `json:"channels"`
}

type PlaylistJson struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Name         string `json:"name"`
	Show         bool   `json:"show"`
	Picture      string `json:"picture"`
	Format       string `json:"format"`
	Link         string `json:"link"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Size         string `json:"-"`
	Duration     int    `json:"-"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"error"`
}

type ErrorMess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// func init() {
// }

// IsExitsSrc
func (this *Channels) IsExitsSrc(filename string) (string, error) {
	dir := filepath.Dir(this.DirSrc)
	fname := filepath.Join(dir, filename)
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		return fname, err
	} else {
		return fname, nil
	}
}

// IsExitsVod
func (this *Channels) IsExitsVod(filename string) (string, error) {
	dir := filepath.Dir(this.DirVideoVod)
	fname := filepath.Join(dir, filename)
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		return fname, err
	} else {
		return fname, nil
	}
}

// isExitsTs Retun Full Path
func (this *Channels) isExitsTs(filename string) (string, error) {
	dir := filepath.Dir(this.DirTs)
	fname := filepath.Join(dir, filename)
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		return fname, err
	} else {
		return fname, nil
	}
}

func (this *Channels) SetTranscode(name string) (err error) {
	if _, err := this.isExitsTs(name); err != nil {
		return err
	} else {
		this.Count++
		this.Files = append(this.Files, name)
		return nil
	}
}

func (this *Channels) SetPlayList(name string) (err error) {
	if _, err := this.isExitsTs(name); err != nil {
		return err
	} else {
		this.Count++
		this.Files = append(this.Files, name)
		return nil
	}
}

func (c *Channels) GetPlayList() (err error) {
	var p []PlaylistJson
	u := fmt.Sprintf("%s/pl/%d/%s/noresize", c.ApiUrl, c.Id, c.Format)
	err = httplib.Get(u).SetTimeout(100*time.Second, 30*time.Second).ToJson(&p)
	// req := httplib.Get(u).SetTimeout(100*time.Second, 30*time.Second)
	// res, err := req.Bytes()
	// json.Unmarshal(res, &p)
	for _, f := range p {
		c.Files = append(c.Files, f.Picture)
	}

	// fmt.Print("---------", c.Files)
	return err
}

func (c *Channels) SavePlayList(name string) (err error) {
	// dir := filepath.Dir(c.DirConfig)
	fname := filepath.Join(c.DirConfig, name)
	// fmt.Println("Write file:", fname)
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, i := range c.Files {
		// _, err := w.WriteString(fmt.Sprintln(i))
		fmt.Println("SavePlayList:", i)
		if _, err := os.Stat(i); err == nil {
			if _, err := w.WriteString(filepath.Base(i) + "\n"); err != nil {
				return err
			}
		}
	}
	w.Flush()
	return err
}

func (c *Channels) InitChannel() {
	if _, err := os.Stat(c.DirConfig); err != nil {
		if err := os.MkdirAll(c.DirConfig, 0777); err != nil {
			c.Errors.Message = fmt.Sprint("mkdir: ", c.DirConfig, " Error: ", err)
		}
	}
	if _, err := os.Stat(c.DirTs); err != nil {
		if err := os.MkdirAll(c.DirTs, 0777); err != nil {
			c.Errors.Message = fmt.Sprint("mkdir: ", c.DirTs, " Error: ", err)
		}
	}
}
