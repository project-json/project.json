package tests

//
//import (
//	"github.com/projectcfg/projectcfg/util/test"
//	"testing"
//)
//
//type testDefaults struct {
//	group string
//	class string
//	name  string
//	ver   string
//	style byte
//}
//
//const EmptyString = ""
//
//var integer = testDefaults{"wordpress", EmptyString, EmptyString, "1", IntegerVersionStyle}
//var dotted = testDefaults{"gearbox", EmptyString, EmptyString, "0.0.0", DottedVersionStyle}
//
//type identityData struct {
//	defaults testDefaults
//	in       string
//	out      string
//	class    string
//	group    string
//	name     string
//}
//
//func (ld identityData) Input() string {
//	return ld.in
//}
//func (ld identityData) Output() string {
//	return ld.out
//}
//
//var identityTests = []identityData{
//	{integer, EmptyString, test.ErrExpected, EmptyString, EmptyString, EmptyString},
//	{integer, "webserver:abc", test.ErrExpected, EmptyString, EmptyString, EmptyString},
//	{integer, ":1", test.ErrExpected, EmptyString, EmptyString, EmptyString},
//	{integer, "webserver:1:0", test.ErrExpected, EmptyString, EmptyString, EmptyString},
//	{integer, "world/gearbox.org/wordpress/webserver:1", test.ErrExpected, EmptyString, EmptyString, EmptyString},
//	{integer, "gearbox.org/lxmp/webserver", "gearbox.org/lxmp/webserver:1", "gearbox.org", "lxmp", "webserver"},
//	{integer, "gearbox.org/wordpress/webserver:1", "gearbox.org/wordpress/webserver:1", "gearbox.org", "wordpress", "webserver"},
//	{integer, "gearbox.org/wordpress/webserver", "gearbox.org/wordpress/webserver:1", "gearbox.org", "wordpress", "webserver"},
//	{integer, "wordpress.org/jetpack/webserver", "wordpress.org/jetpack/webserver:1", "wordpress.org", "jetpack", "webserver"},
//	{integer, "wordpress/webserver", "gearbox.org/wordpress/webserver:1", "gearbox.org", "wordpress", "webserver"},
//	{integer, "webserver", "gearbox.org/wordpress/webserver:1", "gearbox.org", "wordpress", "webserver"},
//	{integer, "webserver:2", "gearbox.org/wordpress/webserver:2", "gearbox.org", "wordpress", "webserver"},

//	{dotted, EmptyString, test.ErrExpected, EmptyString, EmptyString, EmptyString},
//	{dotted, "nginx:abc", test.ErrExpected, EmptyString, EmptyString, EmptyString},
//	{dotted, ":1.14.0", test.ErrExpected, EmptyString, EmptyString, EmptyString},
//	{dotted, "nginx:1.14.0:0", test.ErrExpected, EmptyString, EmptyString, EmptyString},
//	{dotted, "world/github.com/gearbox/nginx:1.14.0", test.ErrExpected, EmptyString, EmptyString, EmptyString},
//	{dotted, "nginx:2", "github.com/gearbox/nginx:2.0.0", "github.com", "gearbox", "nginx"},
//	{dotted, "github.com/gearbox/nginx:1.14.0", "github.com/gearbox/nginx:1.14.0", "github.com", "gearbox", "nginx"},
//	{dotted, "github.com/lxmp/nginx", "github.com/lxmp/nginx:0.0.0", "github.com", "lxmp", "nginx"},
//	{dotted, "github.com/gearbox/nginx", "github.com/gearbox/nginx:0.0.0", "github.com", "gearbox", "nginx"},
//	{dotted, "git.gearbox.org/gearbox/nginx", "git.gearbox.org/gearbox/nginx:0.0.0", "git.gearbox.org", "gearbox", "nginx"},
//	{dotted, "gearbox/nginx", "github.com/gearbox/nginx:0.0.0", "github.com", "gearbox", "nginx"},
//	{dotted, "nginx", "github.com/gearbox/nginx:0.0.0", "github.com", "gearbox", "nginx"},

//}
//
//func TestIdentity(t *testing.T) {
//	for _, ld := range identityTests {
//		d := ld.defaults
//		l := NewIdentity(d.style)
//		l.SetDefaults(d.group, d.class, d.name, d.ver)
//		th := test.NewHarness(t, ld, l)
//		th.Run(func() {
//			if parseTest(th) == nil {
//				nameTest(th)
//				groupTest(th)
//				classTest(th)
//				locationTest(th)
//			}
//		})
//	}
//}
//
//func getIdentity(th *test.Harness) *PreviousIdentity {
//	return th.Item.(*PreviousIdentity)
//}
//
//func parseTest(th *test.Harness) error {
//	err := getIdentity(th).Parse(th.Input())
//	if th.Output() == test.ErrExpected {
//		if err == nil {
//			th.T.Errorf("wanted error %q, did not get", th.Input())
//		}
//	}
//	return err
//}
//
//func locationTest(th *test.Harness) {
//	l := th.Item.(*PreviousIdentity)
//	ls := l.GetIdentity()
//	if ls != th.Output() {
//		th.T.Errorf("wanted %q, got %q", th.Output(), ls)
//	}
//}
//
//func classTest(th *test.Harness) {
//	l := th.Item.(*PreviousIdentity)
//	d := th.InOut.(identityData)
//	c := l.GetClass()
//	if c != d.class {
//		th.T.Errorf("wanted %q, got %q", d.class, c)
//	}
//}
//
//func groupTest(th *test.Harness) {
//	l := th.Item.(*PreviousIdentity)
//	d := th.InOut.(identityData)
//	g := l.GetGroup()
//	if g != d.group {
//		th.T.Errorf("wanted %q, got %q", d.group, g)
//	}
//}
//func nameTest(th *test.Harness) {
//	l := th.Item.(*PreviousIdentity)
//	d := th.InOut.(identityData)
//	g := l.GetName()
//	if g != d.name {
//		th.T.Errorf("wanted %q, got %q", d.name, g)
//	}
//}
