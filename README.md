# RESTful-API-with-Golang
RESTful API with Golang

## 🏃 Getting Started
### Requirements
The program must be already installed on your machine:
1. Golang
2. Docker
3. Postman
4. Code Editor (recommended : GoLand from Jetbrain)

Let start setup your machine for the `development environment` ⚙.

In this section, you need to install several programs which is needed on the development phase.
Just follow the instruction, and you will in set!
**Note : I will write this instruction for Linux Machine**

Install golang (`min version is 1.14`):
```shell script
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install golang-go
```

### IDE & Testing Tools
For the better experiences, we are recommend you to use these IDE and testing tools:
1. [Goland](https://www.jetbrains.com/go/) for main code editor.
2. [Postman](https://www.postman.com) for testing RESTful HTTP API.
3. [DBeaver](https://dbeaver.io) for sql database management, or
4. [DataGrip](https://www.jetbrains.com/datagrip/) for sql database management.


### Set the env
Before you start to run this service, you should set configs with yours.
Create `.env` file. Take a look at: `.env.example`.
```shell script
cp .env.example .env
```
### Config the env
- Set `configuration of your database postgres`


# Fast Running App With Docker Compose
```shell script
 docker-compose up --build
```
# Running Unit Test
```shell script
go test -v ./controllers/... --cover
```

## Testing With Postman
### CRUD Table Employees
**URL Request :**

Method | URL                         | Note 
--- |-----------------------------|  --- 
GET | /employees                  | get all employees
GET | /employees/:id              | get detail of employee
POST | /employees                  | add new employee
PUT | /employees/:id              | update detail of employee
DELETE | /employees/:id      | delete (remove) employee


#### GET EMPLOYEES
    url : http://localhsot:8081/employees
    
    method : GET

<img src="https://i.ibb.co/qdYtrb2/Screenshot-from-2023-10-15-15-15-33.png" alt="get-employees" border="0">

#### FIND EMPLOYEE
    url : http://localhsot:8081/employees/:id
    
    method : GET

<img src="https://i.ibb.co/S3BP78T/Screenshot-from-2023-10-15-15-17-32.png" alt="find-employees" border="0">

#### SAVE EMPLOYEE

    url : http://localhsot:8081/employees
    
    method : POST

Fill Body Request Json with :

Key | Value                         |
--- |-------------------------------| 
first_name | string                        
last_name | string                        |
email | string                        
hire_date | date with format : DD/MM/YYYY |


<img src="https://i.ibb.co/rZxNJkz/Screenshot-from-2023-10-15-15-21-26.png" alt="save-employee" border="0">

#### UPDATE EMPLOYEE
    url : http://localhsot:8081/employees/:id
    
    method : PUT

Fill Body Request Json with :

Key | Value                         |
--- |-------------------------------| 
first_name | string                        
last_name | string                        |
email | string                        
hire_date | date with format : DD/MM/YYYY |


<img src="https://i.ibb.co/RCZ7SS4/Screenshot-from-2023-10-15-15-23-22.png" alt="update-employee" border="0">


#### DELETE EMPLOYEE

    url : http://localhsot:8081/employees/:id
    
    method : DELETE


<img src="https://i.ibb.co/3ryKMcY/Screenshot-from-2023-10-15-15-28-47.png" alt="delete-employee" border="0">


## ERROR HANDLING
In this API, is equipped with error handling ,for example:

### VIEW EMPLOYEE BY ID NOT FOUND
If ID Employee not found in this DB, there is got 404 not found
<img src="https://i.ibb.co/25L7nLN/Screenshot-from-2023-10-15-15-32-30.png" alt="employee-not-found" border="0">

### SAVE EMPLOYEE , EMAIL ALREADY EXIST
When save new employee with email is already exist in databasem there is gor error with 422 error code
<img src="https://i.ibb.co/tK7wNx9/Screenshot-from-2023-10-15-22-03-58.png" alt="save-employee-already-exist" border="0">


## Created by
Ibnu Faisal for Mekari Case Study Test
Thanks
