# Telint
## The street smart telephone bookworm.

This tool uses public records, reports, and regulations to analyze telephone numbers.
It currently covers any North American phone number, including both the USA and Canada.
Telint has NANP databases (tens of thousands of records) baked directly into its binary, so you can get lighting fast offline lookups without taking up a bunch of storage (less than 4MB).

```
$ telint 7038577488
 _______   _ _       _   
|__   __| | (_)     | |  
   | | ___| |_ _ __ | |_ 
   | |/ _ \ | | '_ \| __|
   | |  __/ | | | | | |_ 
   |_|\___|_|_|_| |_|\__|

The street smart telephone bookworm.
====================================


Analyzing 7038577488
Area Code: 703
Checking NANP database for 703
NANP area code match found

Reserved: No
Assigned: Yes
Assignment Date: 01-JAN-47
Location: VA
Country: US
In Service: Y
Service Date: 01-JAN-47
Extra Info: 
```

Features:
- Zero dependencies
- Zero internet connections
- Lighting fast phone number lookup
  - Geo-location data
    - State
    - Country
  - Reservation confirmation
  - Service status
  - Area code assignment info
- [National Number Recognition](#national-numbers)
- [Warnings and notes](#warnings)
- [Checks special case reports by the NANP](#special-handling)
- [Company Ownership Identification](#company-ownership-identification)

## Installation

Go v19 is the only prerequisite.

### Install using go

```
go install github.com/xenophonsec/telint
```

### Build from source

```
git clone https://github.com/xenophonsec/telint.git
cd telint
go build
chmod +x telint
./telint phonenumber
```

## National Numbers

Telint knows about

```
$ telint 811
 _______   _ _       _   
|__   __| | (_)     | |  
   | | ___| |_ _ __ | |_ 
   | |/ _ \ | | '_ \| __|
   | |  __/ | | | | | |_ 
   |_|\___|_|_|_| |_|\__|

The street smart telephone bookworm.
====================================


Analyzing 811
Area Code: 811
Checking NANP database for 811
NANP area code match found

Reserved: No
Assigned: No
Assignment Date: 
Location: 
Country: 
In Service: N
Service Date: 
Extra Info: N11 Code

WARNING:
This number is not currently valid. If someone is using it, it is most likely for suspicious activity.

This is a nationally reserved number for:
        USA    -> Access to One Call Services to Protect Pipeline and Utilities from Excavation Damage
        Canada -> Non-Urgent Health Teletriage Services
```

## Warnings

Telint is street smart. It knows when a number looks potentially phishy.
Warnings will alert you to potentially suspicious details.

```
$ telint 8884038499
 _______   _ _       _   
|__   __| | (_)     | |  
   | | ___| |_ _ __ | |_ 
   | |/ _ \ | | '_ \| __|
   | |  __/ | | | | | |_ 
   |_|\___|_|_|_| |_|\__|

The street smart telephone bookworm.
====================================


Analyzing 8884038499
Area Code: 888
Checking NANP database for 888
NANP area code match found

Reserved: No
Assigned: Yes
Assignment Date: 25-MAY-95
Location: NANP AREA
Country: 
In Service: Y
Service Date: 01-MAR-96
Extra Info: 

WARNING:
This number is a toll free phone number, not tied to any geo-location.
Both business and scammers will use numbers like this to avoid being tied to a location.
```

```
$ telint 2434850399
 _______   _ _       _   
|__   __| | (_)     | |  
   | | ___| |_ _ __ | |_ 
   | |/ _ \ | | '_ \| __|
   | |  __/ | | | | | |_ 
   |_|\___|_|_|_| |_|\__|

The street smart telephone bookworm.
====================================


Analyzing 2434850399
Area Code: 243
Checking NANP database for 243
NANP area code match found

Reserved: No
Assigned: No
Assignment Date: 
Location: 
Country: 
In Service: N
Service Date: 
Extra Info: 

WARNING:
This number is not currently valid. If someone is using it, it is most likely for suspicious activity.
```

## Special Handling
```
$ telint 8132008504
 _______   _ _       _   
|__   __| | (_)     | |  
   | | ___| |_ _ __ | |_ 
   | |/ _ \ | | '_ \| __|
   | |  __/ | | | | | |_ 
   |_|\___|_|_|_| |_|\__|

The street smart telephone bookworm.
====================================


Analyzing 8132008504
Area Code: 813
Checking NANP database for 813
NANP area code match found

Reserved: No
Assigned: Yes
Assignment Date: 01-JAN-53
Location: FL
Country: US
In Service: Y
Service Date: 01-JAN-53
Extra Info: 

NANP Special handling record found!
AC: 813
NXX: 200
State: FL
Rate Center: TAMPA
Notes: "Green highlight - The TAMPA rate center in NPAs 813 and 941 no longer exists.   Central office codes can be requested under the TAMPACEN, TAMPAEST, TAMPANTH, TAMPASTH, or TAMPAWST rate centers (FL PSC Order No. PSC-01-2113-FOF-TP)."
```

## Company Ownership Identification

Telint ill tell you if a record connecting a company to a range of phone numbers is available.
```
$ telint 5002035968
 _______   _ _       _   
|__   __| | (_)     | |  
   | | ___| |_ _ __ | |_ 
   | |/ _ \ | | '_ \| __|
   | |  __/ | | | | | |_ 
   |_|\___|_|_|_| |_|\__|

The street smart telephone bookworm.
====================================


Analyzing 5002035968
Area Code: 500
Checking NANP database for 500
NANP area code match found

Reserved: No
Assigned: Yes
Assignment Date: 19-MAY-94
Location: NANP AREA
Country: 
In Service: Y
Service Date: 19-MAY-94
Extra Info: 

Record found for 500203XXXX numbers
OCN: 9148
Company: CELLCO PARTNERSHIP DBA VERIZON WIRELESS 
```

## Sources

- NPA Database: https://www.nationalnanpa.com/nanp1/npa_report.csv
- 5XX Report: https://nationalnanpa.com/nanp1/All5XXNXXCodesReport.txt
- Area Codes with Special Handling Database: https://nationalnanpa.com/reports/Codes_requiring_special_handling.xlsx

