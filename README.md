#Analyze
Command line utility to analyze go source files using go/ast.

## Installation:
    go get github.com/markclow/analyze

## Help
    analyze --help
    
## Available Commands:
### isImported
Was a package imported in the go file?

    analyze isImported --fileName ~/_projects/_work/analyze/testdata/functions_1.go --package fmt

### getFunctionParameters
Get a comma-separated list of the parameters passed to the function in the go file.

Useful to check if the function was written as expected (ie expects the correct parameters).

    analyze getFunctionParameters --fileName ~/_projects/_work/analyze/testdata/functions_4.go --functionName computeMarsYears


### getFunctionResults
Get a comma-separated list of the results returned by the function in the go file.

Useful to check if the function was written as expected (ie returns the correct results).

    analyze getFunctionResults --fileName ~/_projects/_work/analyze/testdata/functions_4.go --functionName computeMarsYears


### getFunctionInvokedBy
Get a comma-separated list of the functions that invoke the function in the go file.

Useful to check if the function is called anywhere.

    analyze getFunctionInvokedBy --fileName ~/_projects/_work/analyze/testdata/functions_1.go --functionName eatTacos

