package Config

import (
	//	"encoding/json"
	"bufio"
	"fmt"
	"os"
)

type Config struct {
	Initialized        bool
	AwsEnabled         bool
	Aws                awsConfig
	GithubEnabled      bool
	Github             githubConfig
	ServerHostName     string
	FriendlyServerName string
	ServerIP           string
	ServicePort        int
	ConfigLocation     string
}

type awsConfig struct {
	enableAWSCli  bool
	ec2InstanceID string
	accessKey     string
	secretKey     string
	region        string
}

type githubConfig struct {
}

type supportedServerEnvironments struct {
}

func (c *Config) setupConfig() {
	c.Initialized = true
	c.AwsEnabled = false
	c.configureAws()

}
func (c *Config) configureAws() {
	if !c.AwsEnabled {
		fmt.Println("AWS is not enabled, would you like to enable it? (y/n)")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		switch text {
		case "y", "Y", "yes", "Yes", "YES", "t", "T", "true", "True", "TRUE":
			c.AwsEnabled = true
		case "n", "N", "no", "No", "NO", "f", "F", "false", "False", "FALSE":
			c.AwsEnabled = false
		default:
			fmt.Println("Invalid input, please try again.")
			c.configureAws()
		}
	}
	if !c.AwsEnabled {
		return
	}
	if c.AwsEnabled {
		// if aws is enabled we should set up the AWS parameters and CLI
		c.Aws.accessKey = "test"
		c.Aws.secretKey = "test"
		c.Aws.region = "us-east-1"
		c.Aws.ec2InstanceID = "i-1234567890abcdef0"
		c.Aws.enableAWSCli = true
	}
}
