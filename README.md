# MicroserviceOnGo
Микросервис с моделью машинного обучения RandomTree

Для запуска:

```sh
go run ./MicroserviceOnGo/*.go
```

## Проверка транзакции

`POST /transactions` вернет результат проверки подозрительности транзакции.

### Общая информация

В качестве тела запроса необходимо передавать JSON с информацией о транзакции.

### Пример тела запроса

```json
{
    "Sum": 0.01,
    "YearEnd": 2021,
    "Hour": 4.5,
    "WeekDay": 4,
    "HaveDeviceId": 0,
    "EuropeAsiaCountryIp": 0
}
```

### Поля запроса

Путь | Тип | Описание
---- | -------- | --------
Sum | float | Сумма транзакции
YearEnd | int  | Год окончания карты
Hour | float | Время суток (в часах)
WeekDay | int  | День недели [0:6]
HaveDeviceId | int | Наличие Id устройства входа [0,1]
EuropeAsiaCountryIp | int  | Страна, из который пользвоатель выполняет запрос, из Евразии [0,1]


### Результат запроса

* `201 Created` - успешная проверка.

```json
{
    "Sum": 0.01,
    "YearEnd": 2021,
    "Hour": 4.5,
    "WeekDay": 4,
    "HaveDeviceId": 0,
    "EuropeAsiaCountryIp": 0,
    "Refund": 1
}
```

Путь | Тип | Описание
---- | -------- | --------
Sum | float | Сумма транзакции
YearEnd | int  | Год окончания карты
Hour | float | Время суток (в часах)
WeekDay | int  | День недели [0:6]
HaveDeviceId | int | Наличие Id устройства входа [0,1]
EuropeAsiaCountryIp | int  | Страна, из который пользвоатель выполняет запрос, из Евразии [0,1]
Refund | int | Флаг подозрительности транзакции [0,1]

### Ошибки

* `400 Bad Request` – ошибка в параметрах запроса. В теле подробности.
* `422 Unprocessable Entity` – ошибка заполнения данных в JSON (отсутствие полей, неправильный формат).

```json
{
    "code": 422,
    "text": "unexpected end of JSON input",
}
```
