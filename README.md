
# VaccinationAvailableslot
Get information about Covid 19 Vaccination slots from cowin app and alert you via sms

This application will check the cowin web portal for available slot for next 7 days. If the slot is available it will display the center name and other information regarding the vacciation center.

To run the application : 

1. Goto assets folder then userninfo.json . Add the pincode area you want to search and minimum age and max age ,hours you want to check again after and your phone number for vaccination center availability.

2. If you have a Go environment set : 
    -Clone the project into local
    - Inside the folder run : 
        go run main.go 
        

P.S. : Result will check if vaccine is available and for vaccination center is available for age between minage(include) and max age(exclude).
