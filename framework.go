package projectcfg

type FrameworkName string

type FrameworkMap map[FrameworkName]interface{}

type Frameworks []Framework

func (me FrameworkMap) Unmarshal(data []byte, v interface{}) error {
	return nil
}

func (me Frameworks) Unmarshal(data []byte, v interface{}) error {
	return nil
}

type Framework interface {
	GetName() FrameworkName
}
