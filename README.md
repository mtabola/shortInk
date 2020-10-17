# ShortInk

ShortInk - специальный HTTP сервис предназначенный для сокращения ссылок.

# Требуемое программное обеспечение
  + [MySql][1] 
  + [GoLand 2020.2.2] [2]
  + [Git] [3]
  + [Go Compiler] [4]

# Инструкция по применению
  - Скачать данный проект из репозитория с помощью `git clone https://github.com/r00tm4k3r/shortInk.git` 
  - Открыть MySql и создать базу данных, скрипт которой находится в файле `sqlScript.sql`
  - В функции `main.go` найти строку подключения к базе данных (`globalVars.DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/shortink")`) и изменить поля `root`, `password` и `:3306` на логин пользователя базы данных, его пароль и порт, который использует база данных
  - Запустить проект командой `go run main.go`
  - Подключиться к `localhost:8000`
  
# Возможности сервиса 
  Сервис позволяет вам добавлять, удалять и изменять ссылки, а также вставлять их в поисковую строку.  
  `localhost:8000/<ваша короткая ссылка>` в поисковой строке сразу-же перенаправит Вас на сайт, который связан с этой короткой ссылкой.  

# Инструкция по применению
  После того, как запустился сервис, Вы можете создать новую короткую ссылку  
  ![Картинка добавления](https://github.com/r00tm4k3r/shortInk/mdResources/addLink.png)
  
  Поле FullLink должно быть автоматически заполнено.
  
  При создании Вы можете оставить поле ShortLink пустым, тогда название выдастся автоматически.
  
  ![Заполняемые поля](https://github.com/r00tm4k3r/shortInk/mdResources/addLinkFields.png)
  
  После чего, эта ссылка доступна в главном меню. Вы можете:
  - Перейти на сайт, кликнув на короткую ссылку
  - Изменить данные о ссылке
  
  ![Картинка изменения](https://github.com/r00tm4k3r/shortInk/mdResources/editLink.png)
  
  В разделе "изменить" Вы можете поменять как и полную ссылку, так и изменить название короткой ссылки или удалить его вовсе (В этом случае короткое наименование будет создано автоматически).
  
  Так-же кроме изменения, вы можете удалить данные. Тогда удаляется полная и короткая ссылка.
  
  ![Картинка удаления](https://github.com/r00tm4k3r/shortInk/mdResources/deleteLink.png)
  
  **Пример**
  
  До удаления
  
  ![Картинка до удаления](https://github.com/r00tm4k3r/shortInk/mdResources/beforeDelete.png)
  
  После удаления
  
  ![Картинка после удаления](https://github.com/r00tm4k3r/shortInk/mdResources/afterDelete.png)
    
[1]: https://dev.mysql.com/downloads/
[2]: https://www.jetbrains.com/go/promo/?gclid=Cj0KCQjwzbv7BRDIARIsAM-A6-2OQr1jyKcsbO5anp7-3vmwF2G0aPFHUQddpiDI5YceIWusZk6kYFQaAvv9EALw_wcB
[3]: https://git-scm.com/downloads
[4]: https://golang.org/doc/install