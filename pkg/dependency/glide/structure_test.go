package glide

import "testing"

func TestLoadGlideFile(t *testing.T) {
	path := "testdata/glide.yaml"

	glide, err := LoadGlideFile(path)
	if err != nil {
		t.Error("could not load glide.yaml. Error: ", err)
	}
	if glide.Package != "glide-test" {
		t.Error("wrong pacakge name: ", glide.Package)
	}
	if len(glide.Import) != 2 {
		t.Error("wrong number import: ", len(glide.Import))
	}
	if glide.Import[0].Package != "import/first/package" {
		t.Error("wrong package from import")
	}
	if glide.Import[0].Version != "1.0.0" {
		t.Error("wrong version from import")
	}
	if glide.Import[1].Package != "import/second/package" {
		t.Error("wrong package from import")
	}
	if glide.Import[1].Version != "2.0.0" {
		t.Error("wrong version from import")
	}
}
