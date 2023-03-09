package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//go:embed db/npa_report.csv
var npareport []byte

//go:embed db/5XX_report.txt
var fivexx []byte

//go:embed db/special_handling.csv
var specialhandling []byte

func main() {

	// banner
	fmt.Println(" _______   _ _       _   ")
	fmt.Println("|__   __| | (_)     | |  ")
	fmt.Println("   | | ___| |_ _ __ | |_ ")
	fmt.Println("   | |/ _ \\ | | '_ \\| __|")
	fmt.Println("   | |  __/ | | | | | |_ ")
	fmt.Println("   |_|\\___|_|_|_| |_|\\__|")
	fmt.Println("")
	fmt.Println("The street smart telephone bookworm.")
	fmt.Println("====================================")
	fmt.Println("")

	// require phone number
	if len(os.Args) < 2 {
		fmt.Println("Phone number argument required")
		fmt.Println("example: telint 1234567890")
		return
	}
	pn := os.Args[1]

	// handle help flag
	if pn == "--help" {
		printHelp()
		return
	}

	// validate phone number
	vpnmsg := ValidPhoneNumber(pn)
	if vpnmsg != "" {
		fmt.Println(vpnmsg)
		return
	}

	cc := "0"
	// if the phone number includes a country code number
	// (this only works for single digit country codes, but this script only supports the US and Canada for now so that is fine)
	if len(pn) == 11 {
		// country code
		// take off the first number and determine the country
		cc = pn[0:1]
		// phone number
		// trim off the first number
		pn = pn[1 : len(pn)-1]
	}
	// area code
	// grab the fist 3 numbers
	ac := pn[0:3]

	// Informational notes
	fmt.Println("")
	fmt.Println("Analyzing " + pn)
	if cc != "0" {
		fmt.Println("Country Code: " + cc)
	}
	fmt.Println("Area Code: " + ac)

	fmt.Println("Checking NANP database for " + ac)

	ACdata := getACdata(ac)
	if len(ACdata) == 0 {
		fmt.Println("Area code not found in NANP database.")
		fmt.Println("This phone number may not exist, or be a foreign number unknown by this tool.")
	} else {
		fmt.Println("NANP area code match found")
		fmt.Println("")
		fmt.Println("Reserved: " + ACdata[4])
		fmt.Println("Assigned: " + ACdata[5])
		fmt.Println("Assignment Date: " + ACdata[6])
		fmt.Println("Location: " + ACdata[8])
		fmt.Println("Country: " + ACdata[9])
		fmt.Println("In Service: " + ACdata[10])
		fmt.Println("Service Date: " + ACdata[11])
		fmt.Println("Extra Info: " + ACdata[3])
		if len(ACdata[8]) == 2 {
			fmt.Println("Area Code Map: https://www.nationalnanpa.com/area_code_maps/display.html?" + strings.ToLower(ACdata[8]))
		}
		// Check 5XX area code report database
		if pn[0:1] == "5" {
			ac5xx := getAC5XXdata(pn)
			if len(ac5xx) > 0 {
				fmt.Println("")
				fmt.Println("Record found for " + ac5xx[0] + "XXXX numbers")
				fmt.Println("OCN: " + ac5xx[1])
				fmt.Println("Company: " + ac5xx[2])
			}
		}
		// if reservation, assigment, and service status are all set to NO
		if ACdata[4]+ACdata[5]+ACdata[10] == "NoNoN" {
			fmt.Println("")
			fmt.Println("WARNING:")
			fmt.Println("This number is not currently valid. If someone is using it, it is most likely for suspicious activity.")
			fmt.Println("                                   ................                             ")
			fmt.Println("                                 ....................                           ")
			fmt.Println("                               ....,,,,,,,,,,,,,,,,,..                          ")
			fmt.Println("                               .,,,,,,,,,,,,,,,,,,,,,,.         ..,**           ")
			fmt.Println("                               ,,,,,,,,*********,,,,,,      .,,(     ,,         ")
			fmt.Println("                               .,*******************,     .//.       **         ")
			fmt.Println("                                 *******************       /        .*          ")
			fmt.Println("                 ...........,...  **///////////***        *        /*           ")
			fmt.Println("                ,*******************//////////,          .      .//             ")
			fmt.Println("                ,***,,,*////*******//****,,..             ////*/                ")
			fmt.Println("              .....,,,,***/////////****/***,,,.      .. #/                      ")
			fmt.Println("             ...,,,,****///////////,     */***,,,,...(%##                       ")
			fmt.Println("            ..,,,****//////////               ****..*/**,                       ")
			fmt.Println("           ..,,,*****/////                       ...#*/*                        ")
			fmt.Println("           .,,****////***,,,                     (#.*/.                         ")
			fmt.Println("          ...,,,*    ***,,,,,,.                 ((                              ")
			fmt.Println("          ...,,,        **,,,,,..                                               ")
			fmt.Println("          ...,,*           *,,,...          ....//                              ")
			fmt.Println("    ...,,,,,,**              ,,,....   .....,///.                               ")
			fmt.Println(" ....,****                    ,,,........,*///                                  ")
			fmt.Println(" ...,,,,,                       ,,.....,////                                    ")
			fmt.Println(" ..   ....    ,/((,              ,..,////                                       ")
			fmt.Println("  ,*,,,,,,*..        ,*****/     .,//(    *****/,                               ")
			fmt.Println("                             .......,,,,.......       .,,*,.                    ")
			fmt.Println("                                      .. .....                                  VERY INTERESTING...")

		}
	}

	// Handle special cases
	shdata := getSpecialHandlingData(pn)
	if len(shdata) > 0 {
		fmt.Println("")
		fmt.Println("NANP Special handling record found!")
		fmt.Println("AC: " + shdata[1])
		fmt.Println("NXX: " + shdata[2])
		fmt.Println("State: " + shdata[0])
		fmt.Println("Rate Center: " + shdata[3])
		fmt.Println("Notes: " + shdata[4])
	}

	fmt.Println("")
	// N11 Area Codes
	handleN11ACs(pn)

	// Toll Free Area Codes
	handleTollFreeACs(pn)
}

// Handle N11 Area Codes
// Output custom messages for nationally reserved phone numbers
func handleN11ACs(pn string) {
	if strings.HasPrefix(pn, "211") {
		fmt.Println("This is a nationally reserved number for: Community Information and Referral Services")
	}
	if strings.HasPrefix(pn, "311") {
		fmt.Println("This is a nationally reserved number for: Non-Emergency Police and Other Governmental Services")
	}
	if strings.HasPrefix(pn, "411") {
		fmt.Println("This is a nationally reserved number for: Local Directory Assistance")
	}
	if strings.HasPrefix(pn, "511") {
		fmt.Println("This is a nationally reserved number for:")
		fmt.Println("\tUSA    -> Traffic and Transportation Information")
		fmt.Println("\tCanada -> Provision of Weather and Traveller Information Services")
	}
	if strings.HasPrefix(pn, "611") {
		fmt.Println("This is a nationally reserved number for: Repair Service")
	}
	if strings.HasPrefix(pn, "711") {
		fmt.Println("This is a nationally reserved number for: Telecommunications Relay Service (TRS)")
	}
	if strings.HasPrefix(pn, "811") {
		fmt.Println("This is a nationally reserved number for:")
		fmt.Println("\tUSA    -> Access to One Call Services to Protect Pipeline and Utilities from Excavation Damage")
		fmt.Println("\tCanada -> Non-Urgent Health Teletriage Services")
	}
	if strings.HasPrefix(pn, "911") {
		fmt.Println("This is a nationally reserved number for: Emergency Services")
	}
}

// Handle Toll Free Area Codes
func handleTollFreeACs(pn string) {
	tollfreePrefixes := [...]string{"800", "833", "844", "855", "866", "877", "888"}
	for i := 0; i < len(tollfreePrefixes); i++ {
		if strings.HasPrefix(pn, tollfreePrefixes[i]) {
			fmt.Println("WARNING:")
			fmt.Println("This number is a toll free phone number, not tied to any geo-location.")
			fmt.Println("Both business and scammers will use numbers like this to avoid being tied to a location.")
		}
	}
}

// Get Area Code Data
// https://www.nationalnanpa.com/nanp1/npa_report.csv
func getACdata(ac string) []string {
	lines := strings.Split(string(npareport), "\n")
	for i := 2; i < len(lines); i++ {
		if len(lines[i]) > 0 {
			if strings.HasPrefix(lines[i], ac) {
				return strings.Split(lines[i], ",")
			}
		}
	}
	return []string{}
}

// Special Handling Data
// https://nationalnanpa.com/reports/Codes_requiring_special_handling.xlsx
func getSpecialHandlingData(pn string) []string {
	lines := strings.Split(string(specialhandling), "\n")
	for i := 1; i < len(lines); i++ {
		if len(lines[i]) > 0 {
			data := strings.Split(lines[i], ",")
			if strings.HasPrefix(pn, data[1]+data[2]) {
				var note string = ""
				if len(data[4]) > 0 {
					note = lines[i][(strings.Index(lines[i], data[3]+",") + len(data[3]) + 1):(len(lines[i]))]
				}
				return []string{data[0], data[1], data[2], data[3], note}
			}
		}
	}
	return []string{}
}

// Get Area Code 5XX Data
// https://nationalnanpa.com/nanp1/All5XXNXXCodesReport.txt
func getAC5XXdata(pn string) []string {
	lines := strings.Split(string(fivexx), "\n")
	for i := 2; i < len(lines); i++ {
		if len(lines[i]) > 0 {
			data := strings.Split(lines[i], "\t")
			prefix := data[0] + data[1]
			if strings.HasPrefix(pn, prefix) {
				return []string{prefix, data[3], data[4]}
			}
		}
	}
	return []string{}
}

func ValidPhoneNumber(pn string) string {
	// must be longer than 2 digits long
	if len(pn) < 3 {
		return "Phone number must be more that 2 numbers long"
	}
	// must be less than 12 numbers long
	if len(pn) > 11 {
		return "Phone number is too long. It must be 11 or less numbers long"
	}
	// must be only numbers
	match, _ := regexp.MatchString("^\\d+$", pn)
	if !match {
		return "Invalid phone number. Please only enter numbers."
	}
	return ""
}

func printHelp() {
	fmt.Println("Description: Analyze phone numbers using public records/reports/regulations")
	fmt.Println("Author: Xenophonsec")
	fmt.Println("License: MIT")
	fmt.Println("")
	fmt.Println("USAGE:")
	fmt.Println("")
	fmt.Println("telint PHONENUMBER")
	fmt.Println("")
	fmt.Println("Phone numbers must be numbers only: 1234567890")
	fmt.Println("Country codes are accepted. Simply append it to the beginning: 2223334444 -> 12223334444")
}
