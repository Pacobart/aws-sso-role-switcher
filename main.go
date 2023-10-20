package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"runtime"
	"github.com/c-bata/go-prompt"
)

func check(e error) {
	if e!= nil {
		panic(e)
	}
}

func parseAwsConfigForProfiles() ([]string) {
	homeDir, err := os.UserHomeDir()
	check(err)
	readFile, err := os.Open(homeDir + "/.aws/config")
	check(err)

	var awsProfiles []string
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			line = strings.Replace(line, "[", "", 1)
			line = strings.Replace(line, "]", "", 1)
			line = strings.Replace(line, "profile ", "", 1)
			awsProfiles = append(awsProfiles, line)
		}
	}
	readFile.Close()
	return awsProfiles
}

func convertToSuggestions(input []string) []prompt.Suggest {
	suggests := make([]prompt.Suggest, len(input))
	for i := range input {
		suggests[i] = prompt.Suggest{Text: input[i]}
	}
	return suggests
}

func selectAwsProfile(d prompt.Document) []prompt.Suggest {
	awsProfiles := parseAwsConfigForProfiles()
	s := convertToSuggestions(awsProfiles)
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func selectAWSRegion(d prompt.Document) []prompt.Suggest {
	awsRegions := []string{
		"af-south-1",
    	"ap-east-1",
    	"ap-northeast-2",
    	"ap-northeast-1",
    	"ap-south-1",
    	"ap-south-2",
    	"ap-southeast-1",
    	"ap-southeast-2",
    	"ap-southeast-3",
    	"ap-southeast-4",
    	"ca-central-1",
    	"eu-central-1",
    	"eu-north-1",
    	"eu-west-1",
    	"eu-west-2",
    	"eu-west-3",
    	"sa-east-1",
    	"us-east-1",
    	"us-east-2",
    	"us-west-1",
    	"us-west-2",
}
	s := convertToSuggestions(awsRegions)
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func getAWSRegion() string {
	awsDefaultRegion := os.Getenv("AWS_DEFAULT_REGION")
	awsRegion := os.Getenv("AWS_REGION")

	if awsRegion != "" {
		return awsRegion
	} else if awsDefaultRegion != "" {
		return awsDefaultRegion
	} else {
		fmt.Println("AWS region environment variables not set. Select AWS Region")
		region := prompt.Input("> ", selectAWSRegion)
		return region
	}
}

func getAWSCredentials(awsProfile string) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command(fmt.Sprintf("aws configure export-credentials --profile %s --format windows-cmd", awsProfile), os.Getenv("PATH"))
		err := cmd.Run()
		check(err)
	} else {
		cmd := exec.Command(fmt.Sprintf("aws configure export-credentials --profile %s --format env", awsProfile), os.Getenv("PATH"))
		err := cmd.Run()
		check(err)
	}
}

func main() {
	//fmt.Println("Select AWS Profile")
	//awsProfile := prompt.Input("> ", selectAwsProfile)
	//fmt.Println(fmt.Sprintf("Using AWS Profile: %s", awsProfile))

	awsRegion := getAWSRegion()
	//fmt.Println(fmt.Sprintf("Using AWS Region: %s", awsRegion))

	if runtime.GOOS == "windows" {
		fmt.Println(fmt.Sprintf("set AWS_REGION=%s", awsRegion))
	} else {
		fmt.Println(fmt.Sprintf("export AWS_REGION=%s", awsRegion))
	}

	//getAWSCredentials(awsProfile)
}