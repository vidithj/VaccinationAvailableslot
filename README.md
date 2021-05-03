# VaccinationAvailableslot
Get information about Covid 19 Vaccination slots from cowin app

This application will check the cowin web portal for available slot for next 7 days. If the slot is available it will display the center name and other information regarding the vacciation center.

To run the application : 
1. Goto assets folder then userninfo.json . Add the pincode area you want to search and minimum age and max age for vaccination center availability.
2. If you have a Go environment set : 
    -Clone the project into local
    - Inside the folder run : 
        go run main.go 
3. If you use a mac download binary "main" and run in terminal at the binary location : 
    - chmod + x main
    - . /main
4. If you use a linux download binary "mainLinux" and run in terminal at the binary location : 
    - chmod + x mainLinux
    - . /main

4. If you use a windows download binary "mainWindows" and run in terminal at the binary location : 
    - chmod + x mainWindows
    - . /main