package projectcfg

type FrameworkName string

type FrameworkBag map[FrameworkName]interface{}

type FrameworkMap map[FrameworkName]Framework

type Frameworks []Framework

type Framework interface {
	GetName() string
	GetInstance() interface{}
}

//var _ json.Unmarshaler = (*FrameworkBag)(nil)
//func (me *FrameworkBag) UnmarshalJSON(data []byte) error {
//	fwb := make(map[string]interface{},0)
//	err := json.Unmarshal(data,&fwb)
//	if err != nil {
//		panic(err)
//	}
//	return err
//}
