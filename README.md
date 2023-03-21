# ObsGen Go Package
How to use the client:
```go
client, err := NewClient(apiKey)

// You can replace this with the columns in your ObsGen Airtable for now
event := map[string]interface{}{
  "session_id":   "<session_id>",
  "user_id":  "<user_id>",
  "type":  "test",
  "commit_hash": "<commit_hash>",
}
err = client.LogEvent(record)

if err != nil {
  fmt.Println("Error uploading record:", err)
} else {
  fmt.Println("Record uploaded successfully!")
}
```