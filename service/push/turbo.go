package push

import (
	"fmt"
)

const (
	// Data name
	TurboName = "turbo"
	// Token name
	TurboEnv = "TURBO"
)

func init() {
	registerData(TurboName, &TurboData{})
}

// Server-Turbo data type
type TurboData struct {
	title   string
	desp    string
}

// Set title of data
func (TurboData) DataName() string {
	return TurboName
}

// Set title of data
func (data *TurboData) SetTitle(title string) {
	data.title = title
}

// Set body of data
func (data *TurboData) SetContent(content string) {
	data.desp = content
}

// Marshal the data and obtain json string
func (data *TurboData) ToString() string {
	return fmt.Sprintf("title=%s&desp=%s",
		data.title,
		data.desp,
	)
}
