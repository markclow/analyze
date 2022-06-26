# analyze
Command line utility to analyze go source files.

## installation:
go get github.com/markclow/analyze

## examples:
### isImported
Was a package imported?
analyze isImported --fileName /Users/marcusclow/_projects/_work/analyze/testdata/functions_1.go --package fmt

### getFunctionParameters
analyze getFunctionParameters --fileName /Users/marcusclow/_projects/_work/analyze/testdata/functions_4.go --functionName computeMarsYears
Get the parameters (arguments) to a function as a comma-separated list.

### getFunctionResults
analyze getFunctionResults --fileName /Users/marcusclow/_projects/_work/analyze/testdata/functions_4.go --functionName computeMarsYears
Get the results (returned) from a function as a comma-separated list.

### getFunctionInvokedBy
analyze getFunctionInvokedBy --fileName /Users/marcusclow/_projects/_work/analyze/testdata/functions_1.go --functionName eatTacos
Get a comma-separated list of the functions that invoke the function.
