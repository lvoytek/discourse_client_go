# Discourse API Client for Go
Library for interacting with a Discourse site via the [Discourse API](https://docs.discourse.org/)

## Installation
Download the library to your project with `go get`:

```bash
go get github.com/lvoytek/discourse_client_go
```
## Usage
### Initialization
All interactions with a site using this library require a `Client` variable which can be created with either `NewClient` or `NewAnonymousClient`. An API key and username are needed to upload data or access admin data, which can be provided with `NewClient`:

```go
discourseClient := discourse.NewClient("http://127.0.0.1:3000/", "714552c...", "system")
```
Most downloading can be done without an API key. If you would just like to get data from the Discourse site, then you can create an anonymous client:

```go
discourseClient := discourse.NewAnonymousClient("https://discourse.ubuntu.com")
```
### Access

Functions that access the Discourse site are meant to match [Discourse API](https://docs.discourse.org/) calls 1:1. The input will include the `Client` variable, and a variable with included fields that match the required data to upload if needed. The return will be either a success/fail, or a variable with fields matching the site's output.

