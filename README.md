# IP-based wallet grouping service
*Submission for Brave interview*

## Software
This project is a proof of concept (written in Golang) of an implementation of a service operating an ip-grouping strategy. It groups the wallet_id via the (**sha256**) hash of their ip for privacy purposes.

### Launch the software
With go properly installed, you can run the software with `go run *.go` at the root folder of the project.

### API Specification
The treshold for the number of users per ip is stored as a const `GroupTreshold` in users/users.go

The software exposes an HTTP endpoint with two possible requests:

1. **Get user request data**


**POST:** [localhost:8080/submitUserData](http://localhost:8080/submitUserData)
with JSON body of that form :
```json
{ 
    "ip_src": "95.40.123.251",
    "wallet_id": "aaaa8ddd61886ef-2d04s-sd4145-a5sddd03-27sdsd7c1ecd0773",
    "payload": []
}
```

It will return `{"user_added": true}` or `{"user_added": false}`, depending on the success or not of the request

2. **Check if user is reputable or not**


**GET:** [localhost:8080/isUserReputable/{wallet_id}](http://localhost:8080/isUserReputable/{wallet_id})

It will return `{"isReputable": true}`, `{"isReputable": false}` or `{"error": invalid user}` respectively if the user is reputable/not reputable, or if is invalid (i.e. doesnt exist).

## Software
The report with the answers of the questions 3) and 4) is available [here](./report.pdf)