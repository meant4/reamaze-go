# reamaze-go: Re:amaze Golang API Client

![testing](https://github.com/meant4/reamaze-go/actions/workflows/tests.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/meant4/reamaze-go)](https://goreportcard.com/badge/github.com/meant4/reamaze-go)
[![codecov](https://codecov.io/gh/meant4/reamaze-go/graph/badge.svg?token=R6888TMAL6)](https://codecov.io/gh/meant4/reamaze-go)
[![License](https://img.shields.io/github/license/meant4/reamaze-go)](https://github.com/meant4/reamaze-go/blob/main/LICENSE)
[![godoc](https://godoc.org/github.com/meant4/reamaze-go?status.svg)](https://pkg.go.dev/github.com/meant4/reamaze-go@v0.0.0-20240116210523-dc1b94da3bce/reamaze)
[![GitHub release](https://img.shields.io/github/release/meant4/reamaze-go.svg)](https://github.com/meant4/reamaze-go/releases/latest)
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)


## Overview

This repository houses a Golang package designed to serve as a client for interacting with the Re:amaze REST API. 
Re:amaze is providing businesses with a unified platform for customer communication, support, and engagement. With features such as real-time messaging, automated responses, and comprehensive reporting, Re:amaze empowers organizations to deliver exceptional customer experiences.

Re:amaze exposes a robust API that this client encapsulates and simplifies for seamless integration into Golang applications.

## Endpoint Coverage

**reamaze-go** covers all [https://www.reamaze.com/api](https://www.reamaze.com/api) REST API methods:

- **ARTICLES:** Retrieving Articles, Get Article, Creating Articles, Updating Articles
- **CHANNELS:** Retrieving Channels, Retrieving Channels
- **CONTACTS:** Retrieving Contacts, Create Contacts, Update Contacts, Get Contact Identities, Create Identities
- **CONTACT NOTES:** Retrieving Notes, Create Note, Update a note, Delete a note
- **CONVERSATIONS:** Retrieving Conversations, Get Conversation, Creating Conversations, Updating Conversations
- **MESSAGES:** Retrieving Messages, Creating Messages
- **REPORTS:** Volume, Response Time, Staff, Tags, Channel Summary
- **RESPONSE TEMPLATES:** Retrieving Response Templates, Get Response Templates, Creating Response Templates, Updating Response Templates
- **STAFF:** Retrieving Staff, Create Staff User
- **STATUS PAGE:** Retrieving Incidents, Get Incident, Create Incident, Update Incident, Retrieving Systems

## Getting Started

### Prerequisites
Before getting started, ensure you have the following:

- Re:amaze Account: You need a Re:amaze account with a valid email address.

- API Token: Generate an API token from your Re:amaze account settings. This token will be used for authentication when making API requests.

- Re:amaze Brand: Associate your Re:amaze account with the relevant brand. The brand serves as a context for API operations, and you must specify it when initializing the client.

### Installation

To incorporate the Re:amaze Golang API client into your project, use the following `go get` command:

```bash
go get -u github.com/meant4/reamaze-go/reamaze
```

## Example Usage

### Retrieving Articles

```go
package main

import (
	"fmt"
	"github.com/meant4/reamaze-go/reamaze"
)


func main() {
    // Replace these values with your Re:amaze credentials
    email := "your-email@example.com"
    apiToken := "your-api-token"
    brand := "your-brand"

    // Initialize Re:amaze client
    reamazeClient, err := reamaze.NewClient(email, apiToken, brand)
    if err != nil {
        log.Println(err)
    }

    // Example: Get a list of articles
    articles, err := reamazeClient.GetArticles()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Process the list of articles as needed
    fmt.Println("Articles:", articles)
}
```

Refer to the documentation for detailed information on each endpoint and usage examples.

Also please visit [godoc](https://pkg.go.dev/github.com/meant4/reamaze-go@v0.0.0-20240116210523-dc1b94da3bce/reamaze) for all available methods and types in this package

## Contribution Guidelines

We welcome contributions, bug reports, and feature requests. Fork and make a Pull Request, or create an Issue if you see any problem.

## Sponsorship

The development of this package was sponsored by [Meant4.com Software House](https://meant4.com/?utm_source=github_reamaze)

## License

This project is licensed under the [MIT License](LICENSE).

Thank you for considering and contributing to the Re:amaze Golang API client. We look forward to your involvement in enhancing its functionality and reliability.