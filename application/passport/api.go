package passport

import (
	"fmt"
	"regexp"
	"subcenter/infra"
	"time"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	gjson "github.com/tidwall/gjson"
)

func getQRcode() (string, string) {
	api := "https://passport.bilibili.com/qrcode/getLoginUrl"
	body, err := infra.Get(api, "", nil)
	if err != nil {
		panic("getLoginUrl error")
	}
	code := gjson.Parse(string(body)).Get("code").Int()
	if code == 0 {
		qrcodeUrl := gjson.Parse(string(body)).Get("data.url").String()
		authCode := gjson.Parse(string(body)).Get("data.oauthKey").String()
		return qrcodeUrl, authCode
	} else {
		panic("getQRcode error")
	}
}

func verifyLogin(auth_code string) {
	pat := regexp.MustCompile(`^https:\/\/passport\.biligame\.com\/crossDomain\?DedeUserID=(.*)&DedeUserID__ckMd5=(.*)&Expires=.*&SESSDATA=(.*)&bili_jct=(.*)&gourl=.*`)
	for {
		api := "https://passport.bilibili.com/qrcode/getLoginInfo"
		data := map[string][]string{
			"oauthKey": {auth_code},
		}
		body, err := infra.PostFormWithCookie(api, "", data)
		if err != nil {
			continue
		}
		status := gjson.Parse(string(body)).Get("status").Bool()
		if status {
			url := gjson.Parse(string(body)).Get("data.url").String()
			fmt.Println("Login success!")
			result := pat.FindSubmatch([]byte(url))
			fmt.Printf("DedeUserID=%s; ", string(result[1]))
			fmt.Printf("DedeUserID__ckMd5=%s; ", string(result[2]))
			fmt.Printf("SESSDATA=%s; ", string(result[3]))
			fmt.Printf("bili_jct=%s;\n", string(result[4]))
			break
		}
		time.Sleep(time.Second * 3)
	}
}

func LoginBili() {
	fmt.Println("请最大化窗口，以确保二维码完整显示，回车继续")
	fmt.Scanf("%s", "")
	loginUrl, authCode := getQRcode()
	qrcode := qrcodeTerminal.New()
	qrcode.Get([]byte(loginUrl)).Print()
	fmt.Println("或将此链接复制到手机B站打开:", loginUrl)
	verifyLogin(authCode)
	fmt.Println("请将上述内容加入配置文件bili.toml中")
}
