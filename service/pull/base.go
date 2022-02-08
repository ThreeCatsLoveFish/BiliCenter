package pull

type Pull struct {
	Type  string `config:"type"`
	URL   string `config:"url"`
	Token string `config:"token"`
}

// Obtain data from source and write to pushdata
func (pull Pull) Obtain() (string, string) {
	return "", ""
}
