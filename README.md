# kervan-scout-go

```go
client := kervanscout.New(kervanscout.Config{
  Host:         "...",
  ProjectToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.....",
})

type Data struct {
  Agent string `json:"agent"`
  Json  string `json:"json"`
}

data := Data{
  Agent: c.GetHeader("User-Agent"),
  Json:  "{\"test\": \"test\"},{05327776655,5555 1111 1111 1233,foo@example.com}",
}

defer client.CatchPanicError()

client.SendIssue("foo", "bar", data)
```
