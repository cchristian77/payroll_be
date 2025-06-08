### Assumptions 

1. Payroll is run per payroll period, which calculates as following : \
   - Calculating hourly rate : User's base salary / (working days (20) * working hours (8)) 
   - Summarizing attendance days in the current period, ranging between payroll period's start date and date. \
     ```Attendance Pay = total attendance days * hourly rate * working hours (8)```
   - Summarizing overtime hours in the current period. \
     ```Overtime Pay = total overtime hours * hourly rate * 2``` 
   - Summarizing pending reimbursements. 
   - Calculating take home pay \
     ```Take Home Pay = Attendance Pay + Overtime Pay + Reimbursement Pay```

### Domain
1. ```Users```: Manages employee and administrator data including role-based control.
2. ```Payroll Periods```: Defines the periods for payroll processing cycles.
3. ```Payslips```: Stores detailed salary information including basic pay, deductions, and final take-home amounts for each employee per period.
4. ```Attendances```: Records attendance data including check-in and check-out times. One user should only have one attendance per day. 
5. ```Overtimes```: Tracks overtime hours worked per attendance. One attendance should only have one overtime.
6. ```Reimbursements```: Manages employee expense reimbursement claims and the approval status. 
7. ```Request Logs```: Records system activity and API request logs for monitoring purposes.

### Area of Improvements 
1. Create cronjob for check out automation in case users forget to checkout. \
   For instance, the cronjob runs daily in the midnight to update attendances which have missing check out time. 
2. Create cronjob to delete expired sessions in the database, running daily. 
3. Create more unit tests to increase code coverage. 