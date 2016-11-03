package parser

import (
	"strings"
	"testing"
)

func TestParser1(t *testing.T) {
	p, err := New("test_data/template.txt", "test_data/variables.yaml")
	if err != nil {
		t.Fatal(err)
	}

	out, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	expected := `A b c Hola!
The quick brown Fox likes:
- apple
- orange
- banana
Whenever you are asked if you can do a job,
tell them, 'Certainly I can!'. Then get busy and find out how to do it.
- Theodore Roosevelt`

	if strings.TrimSpace(expected) != strings.TrimSpace(out) {
		t.Error("not matching")
	}
}

func TestParser2(t *testing.T) {
	p, err := New("test_data/template2.txt", "test_data/variables.yaml")
	if err != nil {
		t.Fatal(err)
	}

	out, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	expected := `A b c Hola!
The $quick brown Fox likes:
- $apple
- $orange
- $banana
Whenever you are asked if you can do a job,
tell them, 'Certainly I can!'. Then get busy and find out how to do it.
- Theodore Roosevelt`

	if strings.TrimSpace(expected) != strings.TrimSpace(out) {
		t.Error("not matching")
	}
}

func TestParser3(t *testing.T) {
	p, err := New("test_data/template3.txt", "test_data/variables.yaml")
	if err != nil {
		t.Fatal(err)
	}

	out, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	expected := `A b c Hola!
The $quick brown Fox likes:
- $apple
- $orange
- $banana
Whenever you are asked if you can do a job,
tell them, 'Certainly I can!'. Then get busy and find out how to do it.
- Theodore Roosevelt`

	if strings.TrimSpace(expected) != strings.TrimSpace(out) {
		t.Error("not matching")
	}
}

func TestParser4(t *testing.T) {
	p, err := New("test_data/template4.txt", "test_data/variables.yaml")
	if err != nil {
		t.Fatal(err)
	}

	_, err = p.Parse()
	if err == nil {
		t.Error("Expecting error")
	}
}

func TestParser5(t *testing.T) {
	p, err := New("test_data/template5.txt", "test_data/variables.yaml")
	if err != nil {
		t.Fatal(err)
	}

	_, err = p.Parse()
	if err == nil {
		t.Error("Expecting error")
	}
}

func TestParser6(t *testing.T) {
	p, err := New("test_data/template6.txt", "test_data/variables.yaml")
	if err != nil {
		t.Fatal(err)
	}

	_, err = p.Parse()
	if err == nil {
		t.Error("Expecting error")
	}
}

func TestParser7(t *testing.T) {
	p, err := New("test_data/template7.txt", "test_data/variables.yaml")
	if err != nil {
		t.Fatal(err)
	}

	_, err = p.Parse()
	if err == nil {
		t.Error("Expecting error")
	}
}

func TestParser8(t *testing.T) {
	p, err := New("test_data/template8.txt", "test_data/variables.yaml")
	if err != nil {
		t.Fatal(err)
	}

	_, err = p.Parse()
	if err == nil {
		t.Error("Expecting error")
	}
}
