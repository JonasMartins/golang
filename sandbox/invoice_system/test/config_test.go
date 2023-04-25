package test

import (
	"project/src/pkg/utils"
	"project/src/services/financial/configs"
	"reflect"
	"testing"
)

func TestGetFilePath(t *testing.T) {

	u := utils.New()
	f, err := u.GetFilePath(&[]string{"configs", "base.yaml"})
	if err != nil {
		t.Logf("Err %v", err)
	}

	t.Log(f)
	got, want := 1, 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetRoot(t *testing.T) {

	u := utils.New()

	got, err := u.GetRootDir()
	if err != nil {
		t.Logf("Err %v", err)
	}

	want := "/home/jonas/Documentos/dev/golang/sandbox/invoice_system/src/"
	if *got != want {
		t.Errorf("got %s, wanted %s", *got, want)
	}

}

func TestLoadConfig(t *testing.T) {

	var want configs.Config
	got, err := configs.LoadConfig()
	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
	if reflect.TypeOf(got) != reflect.TypeOf(&want) {
		t.Errorf("got %q, wanted %q", *got, want)
	}
	got = nil
}
