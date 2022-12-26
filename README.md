# GOMessage
Сервис обмена сообщений

# API documentation


**Загрузить сообщения между двумя пользователями**
Запрос: GET /message 
Параметры запроса (Query): 
user (обязательный) логин первого пользователя
companion (обязательный) логин второго пользователя
starttime (необязательный) начальное время в формате unix. Будут загружены только сообщения, после указанного времени

Структура ответа
[
    {"id": идентификатор сообщения,
    "body": текст,
    "sender": отправитель,
    "recipient":получатель,
    "attached_path":имя файла в базе данных,
    "attached_name": имя файла, которое увидит пользователь при скачивании,
    "time": время в unix формате}
]
Ответ на запрос - массив объектов (сообщений)
Пример ответа
```go
[
    {"id":1,"body":"Привет!","sender":"Bob","recipient":"Molly","attached_path":"","attached_name":"","time":1672057862},
    {"id":2,"body":"Hello","sender":"Molly","recipient":"Bob","attached_path":"Molly_1672058763577305","attached_name":"screen.png","time":1672058763}
]
```


**Опубликовать сообщение**
Запрос: POST /message 
Параметры запроса (Query): 
отсутствуют

Прикрепленная HTML форма:
```go
    user - текст, логин отправителя
    body - текст, тело сообщения
    recipient - текст, логин получателя
    attachedfile (необязательный) - файл, прикрепленный, к сообщению, файл
   
``` 

Ответ на запрос - id добавленного сообщения
Структура ответа
```go
id
```

Пример ответа
```go
1842
```


**Удалить сообщение**
Запрос: GET /message/delete
Параметры запроса (Query): 
id (обязательный) id сообщения
user (обязательный) логин получателя или отправителя

Ответ на запрос - количество удаленных сообщений
Структура ответа
```go
deleted_messages
```

Пример ответа
```go
1
```

**Удалить все сообщения между пользователями**
Запрос: GET /message/delete
Параметры запроса (Query): 
id (обязательный) id сообщения
user (обязательный) логин первого пользователя
companion (обязательный) логин второго пользователя

Ответ на запрос - количество удаленных сообщений
Структура ответа
```go
deleted_messages
```

Пример ответа
```go
25
```

**Скачать файл**
Запрос: GET /file
Параметры запроса (Query): 
path (обязательный) адрес файла в локальном хранилище, который назначается программно при загрузке файла на сервер. В структуре сообщения соответствует "attached_path"
name (обязательный) Имя, которое будет назначено скаченному файлу. Можно использовать имя, которое использовал пользователь при загрузке файла на сервер. В структуре сообщения "attached_name". Либо можно назначить свое имя (имя необходимо указывать с расширением).

Ответ на запрос - скачивание файла


# Структура модуля MessageGO
Модуль разделен на пакеты:
transport
controller
models
repository
ObjectStorage

**transport**
Пакет, который содержит всю транспортную логику. При необходимости можно заменить его на сокеты или любой другой транспортный протокол

**controller**
Управление всей логикой приложения

**ObjectStorage**
Пакет управления локальным хранилищем для файлов. В этот пакет вынесена логика работы с загружаемыми файлами. Файлы записываются в память, /static/objstorage в папке приложения. При необходимости можно изменить на удаленное хранилище

**repository**
Пакет управления базой данных. В модуле подключена база данный SQLite, так как не требует зависимостей при развертывании и считается приемлимой для проектов, у которых меньше 100 тыс посетителей в сутки. База храниться как файл base.db в папке приложения. При необходимости пакет можно заменить на более мощное решение хранения данных

**models**
Вспомогательный пакет для хранения структур данных