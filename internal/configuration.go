package internal

import (
	"bytes"
	"encoding/json"
	"github.com/rivo/tview"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
)

func InitConfig() *Configuration {
	if _, err := os.Stat("/etc/gopherBuilder/config.json"); os.IsExist(err) {
		config := LoadConfig()
		return config
	} else {
		aws, github, defaultPath, name, path, port := getSetupConfigArgs()
		config := &Configuration{}
		initializedConfig, err := config.SetupConfig(aws, github, defaultPath, name, path, port)
		if err != nil {
			panic(err)
		}
		return initializedConfig
	}
}

func getSetupConfigArgs() (bool, bool, bool, string, string, int) {
	var (
		aws, github, defaultPath bool
		name, path               string
		port                     int
	)
	app := tview.NewApplication()
	form := tview.NewForm().
		AddCheckbox("Enable AWS Integration", false, nil).
		AddCheckbox("Enable Github Integration", false, nil).
		AddInputField("Friendly Server Name", "", 20, nil, nil).
		AddCheckbox("Use Default Config Path", true, nil).
		AddInputField("Server Port", "", 20, nil, nil)
	form.AddButton("save", func() {
		aws = form.GetFormItemByLabel("Enable AWS Integration").(*tview.Checkbox).IsChecked()
		github = form.GetFormItemByLabel("Enable Github Integration").(*tview.Checkbox).IsChecked()
		name = form.GetFormItemByLabel("Friendly Server Name").(*tview.InputField).GetText()
		defaultPath := form.GetFormItemByLabel("Use Default Config Path").(*tview.Checkbox).IsChecked()
		if !defaultPath {
			form.AddInputField("Config Path", "", 20, nil, nil)
			path = form.GetFormItemByLabel("Config Path").(*tview.InputField).GetText()

		} else {
			path = ""
		}
		portStr := form.GetFormItemByLabel("Server Port").(*tview.InputField).GetText()
		port, _ = strconv.Atoi(portStr)
		app.Stop()
	})
	form.SetBorder(true).SetTitle("Setup Configuration")
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	return aws, github, defaultPath, name, path, port
}

type Configuration struct {
	Initialized        bool         `json:"initialized"`
	AwsEnabled         bool         `json:"awsEnabled"`
	Aws                AwsConfig    `json:"awsSettings"`
	GithubEnabled      bool         `json:"github_enabled"`
	Github             GithubConfig `json:"github"`
	ServerHostName     string       `json:"server_host_name"`
	FriendlyServerName string       `json:"friendly_server_name"`
	ServerIP           string       `json:"server_ip"`
	ServicePort        int          `json:"service_port"`
	ConfigLocation     string       `json:"config_location"`
}

type AwsConfig struct {
	Ec2InstanceID string `json:"ec_2_instance_id"`
	AccessKey     string `json:"access_key"`
	SecretKey     string `json:"secret_key"`
	Region        string `json:"region"`
}

type GithubConfig struct {
}

// SetupConfig is used to setup the configuration directory and settings for the application when initialized.
// This function expects the following parameters:
// awsEnable: bool - This is used to enable or disable the AWS integration.
//
// githubEnabled: bool - This is used to enable or disable the Github integration.
//
// defaultConfigPath: bool - This is used to determine if the default config path should be used, or if a custom path should be used.
//
// customConfigPath: string - This is used to set a custom path for the configuration directory. (Pass an empty string
// if defaultConfigPath is true)
//
// friendlyServerName: string - This is used to set the friendly server name for the application.
//
// serverPort: int - This is used to set the port that the application will listen on. (Default is 9292; should be passed if not set)
func (c *Configuration) SetupConfig(awsEnable, githubEnabled, defaultConfigPath bool, customConfigPath, friendlyServerName string, serverPort int) (*Configuration, error) {
	c.AwsEnabled = awsEnable
	err := error(nil)
	/* if awsEnable {
		configureAws()
	}*/
	c.GithubEnabled = githubEnabled
	/*	if githubEnabled {
		configureGithub()
	}*/
	c.ServerHostName = getHostName()
	c.FriendlyServerName = friendlyServerName
	c.ServerIP = getOutboundIP()
	c.ServicePort = serverPort
	if defaultConfigPath {
		configLocation := "/etc/gopherBuilder"
		c.ConfigLocation = configLocation
		err := os.MkdirAll(configLocation, 0755)
		if err != nil {
			panic(err)
		}
		if _, err := os.Stat(configLocation); os.IsNotExist(err) {
			panic("Failed to create config directory. Please check permissions and try again.")
		}
		createFile(configLocation, "config", "json")
	} else {
		configLocation := customConfigPath + "/gopherBuilder"
		c.ConfigLocation = configLocation
		err := os.MkdirAll(configLocation, 0755)
		if err != nil {
			panic(err)
		}
		if _, err := os.Stat(configLocation); os.IsNotExist(err) {
			panic("Failed to create config directory. Please check permissions and try again.")
		}
		createFile(configLocation, "config", "json")
	}
	c.Initialized = true
	jsonData, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(c.ConfigLocation+"/config.json", jsonData, 0644)
	if err != nil {
		return nil, err
	}
	return c, err
}

func LoadConfig() *Configuration {
	configLocation := "/etc/gopherBuilder"
	configFile, err := os.ReadFile(configLocation + "/config.json")
	if err != nil {
		panic(err)
	}
	var config Configuration
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}
	return &config
}

func createFile(dir, name, filetype string) {
	file, err := os.Create(dir + "/" + name + "." + filetype)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func getHostName() string {
	cmd := exec.Command("hostname")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}
