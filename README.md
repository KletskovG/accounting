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
    * Last time finished on Update command implementation - check how DB works and implement CLI verison of command
    * Add global config to CLI
    * Add init command to CLI
    * Swithc to plain body HTTP handler
    * Previous bullet wil allow to reuse code in CMD and HTTP handlers
    * Finish HTTP Server
    

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
