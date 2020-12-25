# About

This is some about text

## Initial Version

The initial version will consist of 2 files/packages.
 - A server that:
   - Receives text in a request body
   - Finds the number of words in the body of text
   - Returns the number of words found
 - A client that:
   - Reads a slice of strings from a file
   - Individually sends each slice of strings to the server in a post (or gRPC call)
   - Receive the number of words in that string as response
