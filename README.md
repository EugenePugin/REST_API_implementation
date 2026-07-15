# User's Guide

How to run:

- at terminal-1, used to manage the server role:
  - launch the program, e.g.:  
  _.\rest-demo.exe_
  
- at terminal-2, used to manage the client role:
   - use *curl.exe* to make the API calls, e.g.:   
   _curl.exe http://localhost:8080/api/v1/products/_  
   _curl.exe http://localhost:8080/api/v1/products/41_  
   _curl.exe http://localhost:8080/api/v1/users_  
   _curl.exe -H "Authorization: secret-token" http://localhost:8080/api/v1/users/7/orders_  

# ToDo:
- to implement API first approach (for now it is Code First)
- to implement throttling middleware
- to adopt Prism for integration tests development