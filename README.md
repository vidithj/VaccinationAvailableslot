# VaccinationAvailableslot
Get information about Covid 19 Vaccination slots from cowin app and alert you via sms

This application will check the cowin web portal for available slot for next 7 days. If the slot is available it will display the center name and other information regarding the vacciation center.

To run the application : 
Option 1 : If you have go installed in local then : 
1. Goto assets folder then userninfo.json . Add the pincode area you want to search and minimum age and max age ,hours you want to check again after and your phone number for vaccination center availability.
2. If you have a Go environment set : 
    -Clone the project into local
    - Inside the folder run : 
        go run main.go 
Option 2 : If you dont have go installed in local then : 
1. If you use a mac download binary "main" and run in terminal at the binary location : 
    - chmod + x main
    - ./main {pincode} {minage} {maxage} {hrs to check again} {phonenumber}
    - example :  ./main 226010 18 50 1 2345678921
2. If you use a linux download binary "mainLinux" and run in terminal at the binary location : 
    - chmod + x mainLinux
    - ./mainLinux {pincode} {minage} {maxage} {hrs to check again} {phonenumber}
    - example :  ./mainLinux 226010 18 50 1 2345678921

3. If you use a windows download binary "mainWindows" and run in terminal at the binary location : 
    - chmod + x mainWindows
    - ./mainWindows {pincode} {minage} {maxage} {hrs to check again} {phonenumber}
    - example :  ./mainWindows 226010 18 50 1 2345678921

P.S. : Result will check if vaccine is available and for vaccination center is available for age between minage(include) and max age(exclude).
