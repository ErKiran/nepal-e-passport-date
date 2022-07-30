## E-PASSPORT-DATE-AVAILABLE

Check the avaiable date for appointment in the E-PASSPORT OFFICE.

`` Usage``
 
 ```go run main.go date ``` 
 If no args are passed you will see all the aviable address options in the terminal. 

 Like this 
 ![AllOtions](https://github.com/ErKiran/nepal-e-passport-date/blob/master/docs/allOptions.png)

 We can use arrows to move or type to filter 

 If we type to fiter the results we will see output like this 
![FilteredOptions](https://github.com/ErKiran/nepal-e-passport-date/blob/master/docs/filteredOptions.png)

Or We can pass address as argument to date to get result of specific address 
Like   
```go run main.go date kathmandu```

![WithArgs](https://github.com/ErKiran/nepal-e-passport-date/blob/master/docs/withArgs.png)

 If no dates are available you have option to get notified when dates will be available. 

![Notify](https://github.com/ErKiran/nepal-e-passport-date/blob/master/docs/notifyMe.png)
 If you choose to be notified you will have to filled either email or sms to be notified and when date will be available you will get notified on the choosed options.

If Dates are available you can see the avaiable the date.

![Date](https://github.com/ErKiran/nepal-e-passport-date/blob/master/docs/date.png)

From the available Dates you can choose the timeslots

![TimeSlot](https://github.com/ErKiran/nepal-e-passport-date/blob/master/docs/timeSlot.png)

Added cron as well