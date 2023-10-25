package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	list, err := split("A")
	assert.Nil(t, err)
	assert.Equal(t, []string{"A"}, list)

	list, err = split("zeusCloud")
	assert.Nil(t, err)
	assert.Equal(t, []string{"zeus", "Cloud"}, list)

	list, err = split("Zeuscloud")
	assert.Nil(t, err)
	assert.Equal(t, []string{"Zeuscloud"}, list)

	list, err = split("zeus_cloud")
	assert.Nil(t, err)
	assert.Equal(t, []string{"zeus", "cloud"}, list)

	list, err = split("talZeus_cloud")
	assert.Nil(t, err)
	assert.Equal(t, []string{"tal", "Zeus", "cloud"}, list)

	list, err = split("ZEUSCLOUD")
	assert.Nil(t, err)
	assert.Equal(t, []string{"Z", "E", "U", "S", "C", "L", "O", "U", "D"}, list)

	list, err = split("zeuscloud")
	assert.Nil(t, err)
	assert.Equal(t, []string{"zeuscloud"}, list)

	list, err = split("")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(list))

	list, err = split("a_b_CD_EF")
	assert.Nil(t, err)
	assert.Equal(t, []string{"a", "b", "C", "D", "E", "F"}, list)

	list, err = split("_")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(list))

	list, err = split("__")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(list))

	list, err = split("_A")
	assert.Nil(t, err)
	assert.Equal(t, []string{"A"}, list)

	list, err = split("_A_")
	assert.Nil(t, err)
	assert.Equal(t, []string{"A"}, list)

	list, err = split("A_")
	assert.Nil(t, err)
	assert.Equal(t, []string{"A"}, list)

	list, err = split("welcome_to_zeus_cloud")
	assert.Nil(t, err)
	assert.Equal(t, []string{"welcome", "to", "zeus", "cloud"}, list)
}

func TestFileNamingFormat(t *testing.T) {
	testFileNamingFormat(t, "zeuscloud", "welcome_to_zeus_cloud", "welcometozeuscloud")
	testFileNamingFormat(t, "_zeus#cloud_", "welcome_to_zeus_cloud", "_welcome#to#zeus#cloud_")
	testFileNamingFormat(t, "Zeus#cloud", "welcome_to_zeus_cloud", "Welcome#to#zeus#cloud")
	testFileNamingFormat(t, "Zeus#Cloud", "welcome_to_zeus_cloud", "Welcome#To#Zeus#Cloud")
	testFileNamingFormat(t, "Zeus_Cloud", "welcome_to_zeus_cloud", "Welcome_To_Zeus_Cloud")
	testFileNamingFormat(t, "zeus_Cloud", "welcome_to_zeus_cloud", "welcome_To_Zeus_Cloud")
	testFileNamingFormat(t, "zeusCloud", "welcome_to_zeus_cloud", "welcomeToZeusCloud")
	testFileNamingFormat(t, "ZeusCloud", "welcome_to_zeus_cloud", "WelcomeToZeusCloud")
	testFileNamingFormat(t, "ZEUSCloud", "welcome_to_zeus_cloud", "WELCOMEToZeusCloud")
	testFileNamingFormat(t, "ZeusCLOUD", "welcome_to_zeus_cloud", "WelcomeTOZEUSCLOUD")
	testFileNamingFormat(t, "ZEUSCLOUD", "welcome_to_zeus_cloud", "WELCOMETOZEUSCLOUD")
	testFileNamingFormat(t, "ZEUS*CLOUD", "welcome_to_zeus_cloud", "WELCOME*TO*ZEUS*CLOUD")
	testFileNamingFormat(t, "[ZEUS#CLOUD]", "welcome_to_zeus_cloud", "[WELCOME#TO#ZEUS#CLOUD]")
	testFileNamingFormat(t, "{zeus###cloud}", "welcome_to_zeus_cloud", "{welcome###to###zeus###cloud}")
	testFileNamingFormat(t, "{zeus###cloudzeus_cloud}", "welcome_to_zeus_cloud", "{welcome###to###zeus###cloudzeus_cloud}")
	testFileNamingFormat(t, "ZeuszeusCloudcloud", "welcome_to_zeus_cloud", "WelcomezeusTozeusZeuszeusCloudcloud")
	testFileNamingFormat(t, "前缀ZeusCloud后缀", "welcome_to_zeus_cloud", "前缀WelcomeToZeusCloud后缀")
	testFileNamingFormat(t, "ZeusCloud", "welcometozeuscloud", "Welcometozeuscloud")
	testFileNamingFormat(t, "ZeusCloud", "WelcomeToZeusCloud", "WelcomeToZeusCloud")
	testFileNamingFormat(t, "zeuscloud", "WelcomeToZeusCloud", "welcometozeuscloud")
	testFileNamingFormat(t, "zeus_cloud", "WelcomeToZeusCloud", "welcome_to_zeus_cloud")
	testFileNamingFormat(t, "Zeus_Cloud", "WelcomeToZeusCloud", "Welcome_To_Zeus_Cloud")
	testFileNamingFormat(t, "Zeus_Cloud", "", "")
	testFileNamingFormatErr(t, "zeus", "")
	testFileNamingFormatErr(t, "gOCloud", "")
	testFileNamingFormatErr(t, "cloud", "")
	testFileNamingFormatErr(t, "zeusZEro", "welcome_to_zeus_cloud")
	testFileNamingFormatErr(t, "zeusZERo", "welcome_to_zeus_cloud")
	testFileNamingFormatErr(t, "cloudzeus", "welcome_to_zeus_cloud")
}

func testFileNamingFormat(t *testing.T, format, in, expected string) {
	format, err := FileNamingFormat(format, in)
	assert.Nil(t, err)
	assert.Equal(t, expected, format)
}

func testFileNamingFormatErr(t *testing.T, format, in string) {
	_, err := FileNamingFormat(format, in)
	assert.Error(t, err)
}
