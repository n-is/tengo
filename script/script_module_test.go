package script_test

import (
	"strings"
	"testing"

	"github.com/n-is/tengo/assert"
	"github.com/n-is/tengo/objects"
	"github.com/n-is/tengo/script"
)

func TestScriptSourceModule(t *testing.T) {
	// script1 imports "mod1"
	scr := script.New([]byte(`out := import("mod")`))
	mods := objects.NewModuleMap()
	mods.AddSourceModule("mod", []byte(`export 5`))
	scr.SetImports(mods)
	c, err := scr.Run()
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, int64(5), c.Get("out").Value())

	// executing module function
	scr = script.New([]byte(`fn := import("mod"); out := fn()`))
	mods = objects.NewModuleMap()
	mods.AddSourceModule("mod", []byte(`a := 3; export func() { return a + 5 }`))
	scr.SetImports(mods)
	c, err = scr.Run()
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, int64(8), c.Get("out").Value())

	scr = script.New([]byte(`out := import("mod")`))
	mods = objects.NewModuleMap()
	mods.AddSourceModule("mod", []byte(`text := import("text"); export text.title("foo")`))
	mods.AddBuiltinModule("text", map[string]objects.Object{
		"title": &objects.UserFunction{Name: "title", Value: func(args ...objects.Object) (ret objects.Object, err error) {
			s, _ := objects.ToString(args[0])
			return &objects.String{Value: strings.Title(s)}, nil
		}},
	})
	scr.SetImports(mods)
	c, err = scr.Run()
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, "Foo", c.Get("out").Value())
	scr.SetImports(nil)
	_, err = scr.Run()
	if !assert.Error(t, err) {
		return
	}
}
