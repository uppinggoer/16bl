package global

import (
	"os"
	"os/exec"
	"path/filepath"
)

var (
	BASE_PATH        string // eg /XX/XX/XX/16bl/app/weixin
	LOG_PATH         string // eg /XX/XX/XX/16bl/log
	VAR_PATH         string // eg /XX/XX/XX/16bl/var
	COMMON_CONF_PATH string // eg /XX/XX/XX/16bl/conf
	CONF_PATH        string // eg /XX/XX/XX/16bl/app/weixin/conf
	APP_CONF_PATH    string // eg /XX/XX/XX/16bl/app/weixin/conf/app
	TPL_PATH         string // eg  /XX/XX/XX/16bl/app/weixin/template
	STATIC_PATH      string // eg  /XX/XX/XX/16bl/app/weixin/static
	DATA_PATH        string // eg  /XX/XX/XX/16bl/app/weixin/data
)

func init() {
	// 计算 base_path
	{
		curFilename := os.Args[0]
		binaryPath, err := exec.LookPath(curFilename)
		if err != nil {
			panic(err)
		}

		binaryPath, err = filepath.Abs(binaryPath)
		if err != nil {
			panic(err)
		}

		BASE_PATH = filepath.Dir(filepath.Dir(binaryPath)) + "/"
	}
	BASE_PATH = "/Users/yanghongzhi/work/16bl/app/weixin"

	homePath := filepath.Dir(filepath.Dir(BASE_PATH))
	LOG_PATH = homePath + "/log/"
	VAR_PATH = homePath + "/var/"
	COMMON_CONF_PATH = homePath + "/conf/"

	CONF_PATH = BASE_PATH + "/conf/"
	APP_CONF_PATH = BASE_PATH + "/conf/app/"
	TPL_PATH = BASE_PATH + "/template/"
	STATIC_PATH = BASE_PATH + "/static/"
	DATA_PATH = BASE_PATH + "/data/"
}
