# Development Setup
1. Install and setup Go 1.9.2+
2. Install Go's dep tool (go get -u github.com/golang/dep/cmd/dep)
3. Run `dep ensure` from inside this repository to fetch the dependencies
4. Run `go install`
5. Run the produced executable

# Accessing the API
The SNC API is currently hosted on a [Heroku hobby plan](https://snc-api.herokuapp.com/).
