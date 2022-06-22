```json
{
    "tags": true,
    "outs": [
        {
            "colors": "raw|no|auto",
            "output": "io.Writer.(*os.File)",
            "levels": "Trace|Debug|Info|Warning|Error|Fatal|Panic|Text",
            "format": "logrus",
        }
    ]
}
```

# Tests
Run all tests
```bash
go test ./...
```