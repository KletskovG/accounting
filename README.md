# Accounting

Expense record
date, amount, category, note

# Scenarios

* Add transaction
* Get last x transactions
* Remove by id
* Update by id
* Generate CSV for range of dates, with default values

### Cmd

### HTTP

### TODO
    * DB module should not log fatal errors to prevent app from crash
    * Swithc to plain body HTTP handler
    * Previous bullet wil allow to reuse code in CMD and HTTP handlers
    * Merge transactions in one day and one category in one record. Leave them traceble 
    

### Scenarios
CMD
* [x] Add 
* [x] List
* [x] Update
* [x] Remove
* [x] Report
* [x] Global config
* [x] Init command

DB
* [x] Add 
* [x] List
* [x] Update
* [x] Remove
* [x] Report


HTTP Server
* [x] Add 
* [x] List
* [] Update
* [x] Remove
* [] Report
