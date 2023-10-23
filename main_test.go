package main

import (
	"testing"
)

// Test if AWS_DEFAULT_REGION environment variable works
func TestGetAWSRegionDefaultSet(t *testing.T) {
	awsDefaultRegion := "us-east-1"
	t.Setenv("AWS_DEFAULT_REGION", awsDefaultRegion)
	//os.Setenv()
	expect := awsDefaultRegion
	result := getAWSRegion()
	if result != expect {
		t.Fatalf("FAILED.\nExpect: '%s'\nResult: '%s'", expect, result)
	}
}

// Test if AWS_REGION environment variable works
func TestGetAWSRegionSet(t *testing.T) {
	awsDefaultRegion := "us-east-1"
	t.Setenv("AWS_REGION", awsDefaultRegion)
	//os.Setenv()
	expect := awsDefaultRegion
	result := getAWSRegion()
	if result != expect {
		t.Fatalf("FAILED.\nExpect: '%s'\nResult: '%s'", expect, result)
	}
}

// TODO: Test if setting aws region with interactive prompt works

// Test if formatting output to run in terminal works
func TestFormatOutput(t *testing.T) {
	awsRegion := "us-east-1"
	awsCreds := ""
	expect := "export AWS_REGION=us-east-1\n" //TODO: need to update test when I get credentials working
	result := formatOutput(awsRegion, awsCreds)
	if result != expect {
		t.Fatalf("FAILED.\nExpect: '%s'\nResult: '%s'", expect, result)
	}
}

// TODO: Test if writing to file works
// TODO: Test if parsing aws config file works
// TODO: Test if converting to suggestions works
// TODO: Test if selecing AWS Profile works - interactive
// TODO: Test if selecting AWS Region works - interactive
// TODO: Test if getting AWS credentials works
