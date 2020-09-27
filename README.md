# ShorkInk

ShortInk - специальный HTTP сервис предназначенный для сокращения ссылок.

# Требуемое программное обеспечение
 - [MySql] (https://dev.mysql.com/downloads/)
 - GoLand 2020.2.2 [https://www.jetbrains.com/go/promo/?gclid=Cj0KCQjwzbv7BRDIARIsAM-A6-2OQr1jyKcsbO5anp7-3vmwF2G0aPFHUQddpiDI5YceIWusZk6kYFQaAvv9EALw_wcB]
 - Postman (для отправление запросов на сервер) [https://www.postman.com/downloads/]
 - Git [https://git-scm.com/downloads]
 - Go Compiler [https://golang.org/doc/install]

# Инструкция по применению
  - Скачать данный проект из репозитория с помощью `git clone` 
  - Открыть MySql и создать базу данных, скрипт которой находится в файле `sqlScript.sql`
  - В функции `main.go` найти строку подклюичения к базе данных (`globalVars.DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/shortink")`) и изменить поля `root`, `password` и `:3306` на логин пользователя базы данных, его пароль и порт, который использует база данных
  - Запустить проект командой `go run main.go`
  - Подключиться к `localhost:8000`

Для `localhost:8000/manage` доступны следущие HTTP методы
- GET: позволяет получить информацию о коротких ссылках и их полных эквивалентов
- POST: позволяет создать коротую ссылку. Ссылка создается следующим образом: в тело запроса помещается полный URL адрес. Если вы хотите, чтобы короткий URL cгерериовался сам, то больше ничего не пишете, если хотите написать свой короткий URL, то используйте ` ==> `
    **Пример**
```sh
    https://github.com/r00tm4k3r/shortInk/blob/master/main.go ==> myCustLink
```
- DELITE: удаляет URL по короткому адресу. В тело запроса нужно передать короткую ссылку

Для `localhost:8000/show/<yourShortURL>` доступны следущие HTTP методы
- GET: перенаправляет вас на сайт, который связн с вашей <yourShortURL> ссылкой
