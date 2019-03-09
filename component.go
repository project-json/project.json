package projectcfg

//
//const (
//	//StackComponent      = iota
//	ServiceComponent = iota
//	//ExecutableComponent
//	//ScriptComponent
//	//SourceComponent
//	//DataComponent
//	//MediaComponent
//)
//
//type ComponentClass int
//
//type ComponentList []*Component
//
//type Component struct {
//	class ComponentClass
//	*ComponentType
//	*ComponentId
//}
//
//func (c *Component) GetType() string {
//	return c.ComponentType.GetType()
//}
//
//func (c *Component) GetIdentity() string {
//	return c.ComponentId.GetIdentity()
//}
//
//func NewComponent(class ComponentClass, typestr, refstr string) *Component {
//	var err error
//	ct := NewComponentType()
//	err = ct.Parse(typestr)
//	if err != nil {
//		panic(err)
//	}
//	cr := NewComponentId()
//	err = cr.Parse(refstr)
//	if err != nil {
//		panic(err)
//	}
//	return &Component{
//		class:         class,
//		ComponentType: ct,
//		ComponentId:   cr,
//	}
//}
//
//func NewServiceComponent(typestr, refstr string) *Component {
//	return NewComponent(ServiceComponent, typestr, refstr)
//}
//
////func NewStackComponent(typestr,refstr string) *Component {œ
////	return NewComponent(StackComponent,typestr,refstr)
////}
//
////func NewExecutableComponent(typestr,refstr string) *Component {
////	return NewComponent(ExecutableComponent,typestr,refstr)
////}
////
////func NewScriptComponent(typestr,refstr string) *Component {
////	return NewComponent(ScriptComponent,typestr,refstr)
////}
////
////func NewSourceComponent(typestr,refstr string) *Component {
////	return NewComponent(SourceComponent,typestr,refstr)
////}
////
////func NewDataComponent(typestr,refstr string) *Component {
////	return NewComponent(DataComponent,typestr,refstr)
////}
////
////func NewMediaComponent(typestr,refstr string) *Component {
////	return NewComponent(MediaComponent,typestr,refstr)
////}
//
