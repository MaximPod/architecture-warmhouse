# temperature-api

Простейший сервис температуры.
Предназначен для тестирования работы сервиса smart-house с сервисом температуры

Работает на 8081 порту

## API Endpoints

- `GET /temperature?:location=` - Health check
- `GET /temperature/:id` - Get all sensors

## Поддерживаемые id и location 
	"1": "Living Room"
	"2": "Bedroom"
	"3": "Kitchen"
	"0": "Unknown"

## Responce
Эндпойнты возвращают json объект 
```
{
    "id":"1",
    "location":"Living Room",
    "value":-25.71,
    "unit":"Celsius",
    "timestamp":"2025-06-07T19:48:13+03:00",
    "status":"ok"
}
```
