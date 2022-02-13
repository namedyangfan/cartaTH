# How to build and run
 1. make sure `Go` has been set correctly, see official document here: https://go.dev/doc/install
 2. install dependencies from go.mod
```
go mod download
```
3. build
```
go build main.go
```
4. run with sample test file
```
./main ./test/test.csv 2020-04-01 2
```
expected output:
```
E002,Bobby Jones,NSO-001,600.00
E003,Cat Helms,NSO-002,0.00
E001,Alice Smith,ISO-001,1000.00
E001,Alice Smith,ISO-002,800.00
```
5. run test
```
go test ./...
```
# Key design decisions
1. I used golang because it provides types, easy environment setup, and the convenience of async flow
2. the input file size could potentially be big, so the data is read line by line
3. the input data source might change from `csv` file to from database in the future. 
This can be done by having as long as the new `data repository`
implement the interface `GetVestingEvents`.
4. significant figures need to be applied when reading the data and 
printing the data. I was not sure why it need to appied when reading the data, as 
for scientific computing we usually carry as much sig fits as possible to reduce 
cumulative error. As a result variable `places` are required in both 
`CmdCalcualteVestingSchedule` and `DataRepo`
5. A switch statement is used in `CalculateTotalVestedShares`, this is for the convenience to add more actions like `VEST` and `CANCEL` in the future

# Key assumptions
1. customer ID and customer name have one to one relation
2. customer name is spelled correctly
3. there is no duplication in the vesting events data
4. the vesting events data would always yield a logical result with no negative numbers
5. the program should terminate when it encounters error while parsing the data
6 . input date format follows YYYY-MM-DD format

# If I had more time
1. write integration test to test the application end to end
2. rather than process all the events in a function, break them down into smaller tasks partitioned by customer ID and award ID.
The idea is to have these smaller tasks processed concurrently by lambda functions. 