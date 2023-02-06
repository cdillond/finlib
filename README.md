finlib aims to provide a collection of functions useful for financial calculations. It takes inspiration from the ta-lib library (https://ta-lib.org) and from several Excel functions, but it doesn't provide all of the same functionality, and the implementations here are original and written entirely in Go.
While writing this library, I tried to remain cognizant of the problems inherent to floating point arithmetic. Therefore, where possible, I allow for user control over the precision of the functions. There are four precision levels:
* 0: Default - (recommended) returns a result that has been corrected for floating point instability. This is generally more accurate than the naive implementation, but it may have a very slight performance penalty.
* 1: Naive - returns a result that is not error-corrected. The implementation is designed with avoidance of precision errors in mind, but they are not corrected for. This is fast and accurate enough for most use cases.
* 2: Exact - uses the math.Big package to return an exact result; this is not recommended because it has a large performance penalty and is rarely significantly more accurate than the ErrorCorrected implementation.
* 3: Fast - a naive implementation optimized for speed, without consideration for precision errors. This is not recommended, because although it will normally be the fastest option, it may return unacceptably inaccurate results.


This project is a work in progress. It is likely to change significantly in the future, and should not be relied upon in a production environment.