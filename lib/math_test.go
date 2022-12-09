package lib

import "testing"

func TestVector2(t *testing.T) {
	expected := true
	actual := Vector2{}.IsTouching(Vector2{})
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}

	expected = true
	actual = Vector2{}.IsTouching(Vector2{X: 1})
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}

	expected = true
	actual = Vector2{}.IsTouching(Vector2{X: 1, Y: 1})
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}

	expected = true
	actual = Vector2{}.IsTouching(Vector2{X: -1, Y: -1})
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}

	expected = true
	actual = Vector2{}.IsTouching(Vector2{Y: -1})
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}

	expected = false
	actual = Vector2{}.IsTouching(Vector2{X: -2, Y: -1})
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}

	expected = false
	actual = Vector2{}.IsTouching(Vector2{X: -2, Y: -2})
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}

	expected = false
	actual = Vector2{}.IsTouching(Vector2{X: 2})
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}

	expected = false
	actual = Vector2{}.IsTouching(Vector2{X: 2, Y: -1})
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}
}
