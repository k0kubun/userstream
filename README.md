# Userstream

Simple Twitter UserStream client with OAuth for Golang

## Usage

```go
package main

import "github.com/k0kubun/userstream"

func main() {
	client := &Client{
		ConsumerKey:       "CONSUMER_KEY",
		ConsumerSecret:    "CONSUMER_SECRET",
		AccessToken:       "ACCESS_TOKEN",
		AccessTokenSecret: "ACCESS_TOKEN_SECRET",
	}

	client.UserStream(func(line string) {
		// json response will be printed
		println(line)
	})
}
```

## License

public domain
