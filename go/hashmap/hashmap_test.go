package hashmap

import (
	"testing"
)

func TestHashMap(t *testing.T) {
	hm := New(8)
	hm.Set("name", "nange")
	hm.Set("age", 29)
	hm.Set("sex", "male")
	hm.Set("like", "programing")
	hm.Set("home", "github.com/nange")
	hm.debug()

	name, ok := hm.Get("name")
	if !ok {
		t.Errorf("cant't get key name")
		return
	}
	if name != "nange" {
		t.Errorf("name should == nange, but get:%v", name)
		return
	}

	hm.Delete("home")
	hm.debug()

	age, ok := hm.Get("age")
	if !ok {
		t.Errorf("cant't get key age")
		return
	}
	if age != 29 {
		t.Errorf("age should == 29, but get:%v", age)
		return
	}

	hm.Delete("like")
	hm.debug()

	sex, ok := hm.Get("sex")
	if !ok {
		t.Errorf("cant't get key sex")
		return
	}
	if sex != "male" {
		t.Errorf("sex should == male, but get:%v", sex)
		return
	}

	hm.Set("test", "test")
	hm.debug()

}
