package dimple

import (
	"bytes"
	"testing"
)

func TestDimple(t *testing.T) {
	d := NewDimple()

	buf := bytes.NewBufferString("Let it Goooooooooo")

	// Simplest example where we set a value to be reused.
	d.SetValue("buf", buf)

	if d.Get("buf") != buf {
		t.Error("! Value from Dimple not the same as one injected.")
	}

	// Test a generator that creates a new response each time.
	d.SetGenerator("elsa", func(nd *Dimple) interface{} {
		return bytes.NewBufferString("Elsa")
	})

	d.SetValue("part", "nose")
	d.SetGenerator("olof", func(d *Dimple) interface{} {
		return "carrot " + d.Get("part").(string)
	})

	if d.Get("olof").(string) != "carrot nose" {
		t.Error("! Reuse of other elements on container failed.")
	}

	a := d.Get("elsa")
	b := d.Get("elsa")

	if a == b {
		t.Error("! Generator is not producting new values")
	}

	d.SetValue("anna", "Anna")

	// Extend an existing option. Either a generator or value.
	d.Extend("anna", func(o interface{}, d *Dimple) interface{} {
		return o.(string) + " and Elsa"
	})

	if d.Get("anna").(string) != "Anna and Elsa" {
		t.Error("! Extends is not extending.")
	}

	d.Extend("olof", func(o interface{}, d *Dimple) interface{} {
		return o.(string) + " and it's cold"
	})

	if d.Get("olof").(string) != "carrot nose and it's cold" {
		t.Error("! Extending a generator failed.")
	}
}
