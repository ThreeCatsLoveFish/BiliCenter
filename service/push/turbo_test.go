package push

import "testing"

func TestTurboData(t *testing.T) {
	data := TurboPush{}
	data.SetTitle("# Test function")
	data.SetContent("Test ONLY")

	result := data.ToString()
	except := `title=# Test function&desp=Test ONLY`
	if result != except {
		t.Fatalf("ToString() error! get: %s, except: %s", result, except)
	}
}

func TestTurboPush(t *testing.T) {
	push := NewPush(0)
	title := "# Test turbo"
	content := "Success if you can see this info!"
	if err := push.Submit(title, content); err != nil {
		t.Fatalf("Submit failed, error: %v", err)
	}
}
