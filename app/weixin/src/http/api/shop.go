package api

// 分类页接口
type Shop struct {
	ClassList []*Class `json:"classList"`
}

func (self *Shop) Format() {
	for _, classItem := range self.ClassList {
		classItem.Format()
	}
}
