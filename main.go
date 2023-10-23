package main

import (
	"bufio"
	"fmt"
	"os"
	"bytes"
	"os/exec"
	"strings"
	"runtime"
	"github.com/c-bata/go-prompt"
)

var enableDebug = false

func check(e error) {
	if e!= nil {
		panic(e)
	}
}

func debug(input string) {
	if enableDebug {
		fmt.Println(input)
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

func getAWSCredentials(awsProfile string) string {
	var osFormat string
	if runtime.GOOS == "windows" {
		osFormat = "windows-cmd"
	} else {
		osFormat = "env"
	}
	debug(fmt.Sprintf("Operating System: %s", osFormat))

	cmd := exec.Command("aws", "configure", "export-credentials", "--profile", awsProfile, "--format", osFormat)
	//cmd := exec.Command("aws", "configure", "export-credentials", "--profile", awsProfile, "--format", "process") // TODO: convert to using process to get json/struct. Also look to use sts call instead
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Result: " + out.String())
	}
	
	check(err)
	return out.String()
}

func formatOutput(awsRegion string, awsCreds string) string {
	var exportWord string
	if runtime.GOOS == "windows" {
		exportWord = "set"
	} else {
		exportWord = "export"
	}
	regionCommand := fmt.Sprintf("%s AWS_REGION=%s", exportWord, awsRegion)
	output := fmt.Sprintf("%s\n%s", regionCommand, awsCreds)
	return output
}

func writeToFile(data string, tmpFile string) {
	outFile, err := os.Create(tmpFile)
	check(err)
	defer outFile.Close()
	w := bufio.NewWriter(outFile)
	_, err = w.WriteString(data)
	check(err)
	w.Flush()
}

func main() {
	fmt.Println("Select AWS Profile")
	awsProfile := prompt.Input("> ", selectAwsProfile)
	debug(fmt.Sprintf("Using AWS Profile: %s", awsProfile))

	awsRegion := getAWSRegion()
	debug(fmt.Sprintf("Using AWS Region: %s", awsRegion))

	awsCreds := getAWSCredentials(awsProfile)

	outFile := "/tmp/asrscmds"
	outText := formatOutput(awsRegion, awsCreds)
	fmt.Println(outText)
	writeToFile(outText, outFile)
	fmt.Println(fmt.Sprintf("Commands set. Run `source %s`", outFile))
}