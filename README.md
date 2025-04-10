## Auth-Service
#### Build container
- `make docker-build`  
#### Start container
- `make docker-up`  
***
#### To run local development  
- `make run`  
#### To run test  
- `make test`  
***
#### The service will serve as `http://localhost:1300` which contain path below  
- *`POST /login` for login with userId and passcode*  
- *`POST /verify` for verify user passcode*  
- *`GET /banner` for getting user banner*  
- *`GET /health` for service health checking*  
***
##### *See Makefile for other command*