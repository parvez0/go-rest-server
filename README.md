# Go Rest server
This project contains the golang source code for the devOps assignment, you can refer the document [here](https://gist.github.com/VortoEng/53a027df8665b2bcca160b8256393f4f).

### Prerequisite
You need Golang version > v1.13 installed on your computer. you will also need a working kubernetes cluster
to deploy the these services as kubernetes pods.

### Installation

Download the source code using git 
```bash
$ git clone <youtgiturl>
```
After cloning the project you need to install the dependencies using the following command
```bash
$ cd path/to/downloaded/project
$ go mod -d -v ./...
```
Once the dependencies have been installed you can start the project using the following command
```bash
$ go run main.go
```

The project will start with the default configurations values which can be changed in the config file
present in root directory. Or you can provide the values in environment variables as well
```bash
$ export DB.HOST=localhost
$ export DB.PASSWORD=<your_password>
$ go run main.go
```

This server provides 2 end points given below

```bash
ALL /health-check
```
```json
{
  "Success": true,
  "Message": "Go server is working !!"
}
```

```bash
GET /invalid-deliveries
```
```json
[
  {
    "Id": 3,
    "SupplierId": "8",
    "DriverId": "3",
    "UpdatedAt": "2020-06-19 19:41:04.774421+00",
    "CreatedAt": "2020-06-19 19:41:04.774421+00"
  }
]
```

### Docker image
For building the docker image run the following command in the root directory of the project
```bash
$ docker build -t <tag> .
$ docker push <tag>
```
