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

- Заворачивать структуру в структуру как это происходит с контекстом.
    Таким образом повысится производительность исключив копирование
    массива. Модификация должна коснуться xfield.Field.
    xfield.Fields больше не будет, работаем по контексту.
- xtarget должен поставлять метод, производящий io.Writer.
    Метод принимающий xlevel.Level и копирующий в себя xfield.Fields.