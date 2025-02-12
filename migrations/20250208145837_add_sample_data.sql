-- +goose Up
-- +goose StatementBegin

-- golang --
INSERT INTO course_category (id, name, created_by, created_at, updated_at)
VALUES
    ('a1d4a83e-3d65-4df5-b890-1b1d87b24a70', 'Backend', gen_random_uuid(), now(), now()),
    (gen_random_uuid(), 'Frontend', gen_random_uuid(), now(), now()),
    (gen_random_uuid(), 'Data Science', gen_random_uuid(), now(), now()),
    (gen_random_uuid(), 'DevOps', gen_random_uuid(), now(), now()),
    (gen_random_uuid(), 'Cybersecurity', gen_random_uuid(), now(), now());

INSERT INTO courses (id, title, image_url, category_id, description, created_by)
VALUES
    ('e9c3d8b5-4a71-4f3e-b99e-8d5f6a4b7c32', 'Голанг для бэкенда', 'golang.jpg', 'a1d4a83e-3d65-4df5-b890-1b1d87b24a70', 'Разработка серверных приложений на Go', gen_random_uuid());

-- 1 --
INSERT INTO topics (id, course_id, title, order_number, description)
VALUES
    ('f1d4a83e-3d65-4df5-b890-1b1d87b24a70', 'e9c3d8b5-4a71-4f3e-b99e-8d5f6a4b7c32', 'Основы синтаксиса Go', 1, 'Изучение базового синтаксиса языка.');

INSERT INTO topic_contents (id, topic_id, content, additional_resources, video_urls, image_urls)
VALUES
    (
        gen_random_uuid(), 'f1d4a83e-3d65-4df5-b890-1b1d87b24a70',
        '# Основы синтаксиса Go\n\nGo - это компилируемый, статически типизированный язык программирования.\n\n## Объявление переменных\n```go\nvar a int = 10\nb := 20 // короткое объявление\nfmt.Println(a + b)\n```\n\n## Условия и циклы\n```go\nif a > b {\n    fmt.Println("a больше b")\n} else {\n    fmt.Println("b больше a")\n}\n\nfor i := 0; i < 5; i++ {\n    fmt.Println(i)\n}\n```',
        ARRAY['https://go.dev/', 'https://golang.org/doc/'],
        ARRAY['https://www.youtube.com/embed/YS4e4q9oBaU'],
        ARRAY['https://example.com/golang_syntax.png']
    );


INSERT INTO quizzes (id, topic_id, question, options, correct_answers, multiple_choice)
VALUES
    (gen_random_uuid(), 'f1d4a83e-3d65-4df5-b890-1b1d87b24a70', 'Как объявить переменную в Go?', ARRAY['var x int', 'x := 10', 'int x = 10', 'let x = 10'], ARRAY[true, true, false, false], true),
    (gen_random_uuid(), 'f1d4a83e-3d65-4df5-b890-1b1d87b24a70', 'Какой оператор используется для создания коротких переменных?', ARRAY[':=', '=', 'let', 'var'], ARRAY[true, false, false, false], false),
    (gen_random_uuid(), 'f1d4a83e-3d65-4df5-b890-1b1d87b24a70', 'Как в Go объявить константу?', ARRAY['const x = 10', 'let x = 10', 'x := 10', 'var x = 10'], ARRAY[true, false, false, false], false),
    (gen_random_uuid(), 'f1d4a83e-3d65-4df5-b890-1b1d87b24a70', 'Какой пакет используется для вывода текста в консоль?', ARRAY['fmt', 'io', 'print', 'os'], ARRAY[true, false, false, false], false),
    (gen_random_uuid(), 'f1d4a83e-3d65-4df5-b890-1b1d87b24a70', 'Как объявить цикл for в Go?', ARRAY['for i := 0; i < 10; i++', 'while i < 10', 'loop i < 10', 'foreach i in range(10)'], ARRAY[true, false, false, false], false);



INSERT INTO practical_tasks (id, topic_id, description, difficulty_level, order_number, starter_code, expected_output)
VALUES
    (gen_random_uuid(), 'f1d4a83e-3d65-4df5-b890-1b1d87b24a70', 'Выведите "Hello, World!"', 'easy', 1,
     'package main\n\nimport "fmt"\n\nfunc main() {\n    fmt.Println("Hello, World!")\n}',
     'Hello, World!'),

    (gen_random_uuid(), 'f1d4a83e-3d65-4df5-b890-1b1d87b24a70', 'Создайте переменную x и присвойте ей значение 42, затем выведите её.', 'easy', 2,
     'package main\n\nimport "fmt"\n\nfunc main() {\n    x := 42\n    fmt.Println(x)\n}',
     '42'),

    (gen_random_uuid(), 'f1d4a83e-3d65-4df5-b890-1b1d87b24a70', 'Реализуйте цикл, который выводит числа от 1 до 5.', 'medium', 3,
     'package main\n\nimport "fmt"\n\nfunc main() {\n    for i := 1; i <= 5; i++ {\n        fmt.Println(i)\n    }\n}',
     '1\n2\n3\n4\n5');


-- 2 --

INSERT INTO topics (id, course_id, title, order_number, description)
VALUES
    ('f2a5e2d8-4e67-4f0c-a17c-3e5f77b2e1cd', 'e9c3d8b5-4a71-4f3e-b99e-8d5f6a4b7c32', 'Горутины и конкурентность', 2, 'Работа с потоками и конкурентностью в Go.');


INSERT INTO topic_contents (id, topic_id, content, additional_resources, video_urls, image_urls)
VALUES
    (
        gen_random_uuid(), 'f2a5e2d8-4e67-4f0c-a17c-3e5f77b2e1cd',
        '# Горутины и конкурентность в Go\n\nГорутины - это лёгкие потоки выполнения, которые позволяют запускать конкурентный код.\n\n## Запуск горутины\n```go\npackage main\n\nimport (\n    "fmt"\n    "time"\n)\n\nfunc hello() {\n    fmt.Println("Привет из горутины!")\n}\n\nfunc main() {\n    go hello()\n    time.Sleep(time.Second)\n}\n```',
        ARRAY['https://go.dev/doc/effective_go#goroutines', 'https://golang.org/pkg/sync/'],
        ARRAY['https://www.youtube.com/embed/LvgVSSpwND8'],
        ARRAY['https://example.com/goroutines_example.png']
    );


INSERT INTO quizzes (id, topic_id, question, options, correct_answers, multiple_choice)
VALUES
    (gen_random_uuid(), 'f2a5e2d8-4e67-4f0c-a17c-3e5f77b2e1cd', 'Как запустить горутину в Go?', ARRAY['go function()', 'goroutine function()', 'thread function()', 'start function()'], ARRAY[true, false, false, false], false),

    (gen_random_uuid(), 'f2a5e2d8-4e67-4f0c-a17c-3e5f77b2e1cd', 'Что делает пакет sync?', ARRAY['Синхронизация потоков', 'Вывод в консоль', 'Запускает HTTP-сервер', 'Работает с файловой системой'], ARRAY[true, false, false, false], false),

    (gen_random_uuid(), 'f2a5e2d8-4e67-4f0c-a17c-3e5f77b2e1cd', 'Как избежать гонки данных в конкурентном коде?', ARRAY['Использовать мьютексы', 'Использовать глобальные переменные', 'Не использовать горутины', 'Запускать горутины без ожидания'], ARRAY[true, false, false, false], false),

    (gen_random_uuid(), 'f2a5e2d8-4e67-4f0c-a17c-3e5f77b2e1cd', 'Как завершить горутину?', ARRAY['Выход из функции', 'Использовать канал', 'Закрыть процесс', 'Удалить переменную'], ARRAY[false, true, false, false], false),

    (gen_random_uuid(), 'f2a5e2d8-4e67-4f0c-a17c-3e5f77b2e1cd', 'Что делает time.Sleep(time.Second) в Go?', ARRAY['Приостанавливает выполнение на 1 секунду', 'Останавливает горутину', 'Закрывает поток', 'Перезапускает процесс'], ARRAY[true, false, false, false], false);


INSERT INTO practical_tasks (id, topic_id, description, difficulty_level, order_number, starter_code, expected_output)
VALUES
    (gen_random_uuid(), 'f2a5e2d8-4e67-4f0c-a17c-3e5f77b2e1cd', 'Запустите горутину, которая выводит "Hello, Goroutine!"', 'easy', 1,
     'package main\n\nimport (\n    "fmt"\n    "time"\n)\n\nfunc hello() {\n    fmt.Println("Hello, Goroutine!")\n}\n\nfunc main() {\n    go hello()\n    time.Sleep(time.Second)\n}',
     'Hello, Goroutine!'),

    (gen_random_uuid(), 'f2a5e2d8-4e67-4f0c-a17c-3e5f77b2e1cd', 'Создайте две горутины, которые выводят "First" и "Second".', 'medium', 2,
     'package main\n\nimport (\n    "fmt"\n    "time"\n)\n\nfunc first() {\n    fmt.Println("First")\n}\n\nfunc second() {\n    fmt.Println("Second")\n}\n\nfunc main() {\n    go first()\n    go second()\n    time.Sleep(time.Second)\n}',
     'First\nSecond'),

    (gen_random_uuid(), 'f2a5e2d8-4e67-4f0c-a17c-3e5f77b2e1cd', 'Используйте канал для синхронизации двух горутин.', 'hard', 3,
     'package main\n\nimport (\n    "fmt"\n    "time"\n)\n\nfunc worker(ch chan string) {\n    ch <- "Горутина завершена"\n}\n\nfunc main() {\n    ch := make(chan string)\n    go worker(ch)\n    fmt.Println(<-ch)\n}',
     'Горутина завершена');

-- 3 --
INSERT INTO topics (id, course_id, title, order_number, description)
VALUES
    ('f3c3d9b8-567a-4f62-9b31-2a61f5b1e9f3', 'e9c3d8b5-4a71-4f3e-b99e-8d5f6a4b7c32', 'Работа с базами данных', 3, 'Подключение и работа с PostgreSQL в Go.');


INSERT INTO topic_contents (id, topic_id, content, additional_resources, video_urls, image_urls)
VALUES
    (
        gen_random_uuid(), 'f3c3d9b8-567a-4f62-9b31-2a61f5b1e9f3',
        '# Работа с базами данных в Go\n\nВ Go чаще всего используются PostgreSQL и MySQL.\n\n## Подключение к PostgreSQL\n```go\npackage main\n\nimport (\n    "database/sql"\n    _ "github.com/lib/pq"\n    "fmt"\n)\n\nfunc main() {\n    connStr := "user=username dbname=mydb sslmode=disable"\n    db, err := sql.Open("postgres", connStr)\n    if err != nil {\n        panic(err)\n    }\n    defer db.Close()\n    fmt.Println("Успешное подключение к базе данных!")\n}\n```',
        ARRAY['https://pkg.go.dev/database/sql', 'https://www.postgresql.org/docs/current/index.html'],
        ARRAY['https://www.youtube.com/embed/wzR7ZWE3fOo'],
        ARRAY['https://example.com/postgres_connection.png']
    );


INSERT INTO quizzes (id, topic_id, question, options, correct_answers, multiple_choice)
VALUES
    (gen_random_uuid(), 'f3c3d9b8-567a-4f62-9b31-2a61f5b1e9f3', 'Какой пакет используется для работы с базами данных в Go?',
     ARRAY['database/sql', 'pq', 'gorm', 'mongo'], ARRAY[true, false, false, false], false),

    (gen_random_uuid(), 'f3c3d9b8-567a-4f62-9b31-2a61f5b1e9f3', 'Какие базы данных поддерживает Go с помощью стандартного пакета `database/sql`?',
     ARRAY['PostgreSQL', 'MySQL', 'SQLite', 'MongoDB'], ARRAY[true, true, true, false], true),

    (gen_random_uuid(), 'f3c3d9b8-567a-4f62-9b31-2a61f5b1e9f3', 'Какие методы предоставляет интерфейс `sql.DB`?',
     ARRAY['Query()', 'Exec()', 'Connect()', 'Ping()'], ARRAY[true, true, false, true], true),

    (gen_random_uuid(), 'f3c3d9b8-567a-4f62-9b31-2a61f5b1e9f3', 'Как правильно закрыть соединение с базой данных?',
     ARRAY['db.Close()', 'db.Terminate()', 'db.Disconnect()', 'db.Quit()'], ARRAY[true, false, false, false], false),

    (gen_random_uuid(), 'f3c3d9b8-567a-4f62-9b31-2a61f5b1e9f3', 'Как выполнить SQL-запрос в Go?',
     ARRAY['db.Query()', 'db.Execute()', 'db.Exec()', 'db.Run()'], ARRAY[true, false, true, false], true);


INSERT INTO practical_tasks (id, topic_id, description, difficulty_level, order_number, starter_code, expected_output)
VALUES
    (gen_random_uuid(), 'f3c3d9b8-567a-4f62-9b31-2a61f5b1e9f3', 'Подключитесь к PostgreSQL и выведите "Connected!" при успешном подключении.', 'easy', 1,
     'package main\n\nimport (\n    "database/sql"\n    _ "github.com/lib/pq"\n    "fmt"\n)\n\nfunc main() {\n    connStr := "user=username dbname=mydb sslmode=disable"\n    db, err := sql.Open("postgres", connStr)\n    if err != nil {\n        panic(err)\n    }\n    defer db.Close()\n    fmt.Println("Connected!")\n}',
     'Connected!'),

    (gen_random_uuid(), 'f3c3d9b8-567a-4f62-9b31-2a61f5b1e9f3', 'Создайте таблицу `users` с полями `id`, `name` и `email`.', 'medium', 2,
     'package main\n\nimport (\n    "database/sql"\n    _ "github.com/lib/pq"\n    "fmt"\n)\n\nfunc main() {\n    db, _ := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")\n    defer db.Close()\n    _, err := db.Exec("CREATE TABLE users (id SERIAL PRIMARY KEY, name TEXT, email TEXT);")\n    if err != nil {\n        panic(err)\n    }\n    fmt.Println("Table created!")\n}',
     'Table created!'),

    (gen_random_uuid(), 'f3c3d9b8-567a-4f62-9b31-2a61f5b1e9f3', 'Добавьте нового пользователя в таблицу `users` и получите его ID.', 'hard', 3,
     'package main\n\nimport (\n    "database/sql"\n    _ "github.com/lib/pq"\n    "fmt"\n)\n\nfunc main() {\n    db, _ := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")\n    defer db.Close()\n    var id int\n    err := db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", "John Doe", "john@example.com").Scan(&id)\n    if err != nil {\n        panic(err)\n    }\n    fmt.Println("Inserted ID:", id)\n}',
     'Inserted ID: 1');


-- 4 --

INSERT INTO topics (id, course_id, title, order_number, description)
VALUES
    ('f4f8d4c9-4b76-4e1b-99e8-5d4a2f3b6c71', 'e9c3d8b5-4a71-4f3e-b99e-8d5f6a4b7c32', 'Создание REST API', 4, 'Разработка API с использованием Gin.');


INSERT INTO topic_contents (id, topic_id, content, additional_resources, video_urls, image_urls)
VALUES
    (
        gen_random_uuid(), 'f4f8d4c9-4b76-4e1b-99e8-5d4a2f3b6c71',
        '# Создание REST API с Gin\n\nGin - это популярный веб-фреймворк для Go.\n\n## Установка Gin\n```sh\ngo get -u github.com/gin-gonic/gin\n```\n\n## Простой сервер с Gin\n```go\npackage main\n\nimport (\n    "github.com/gin-gonic/gin"\n)\n\nfunc main() {\n    r := gin.Default()\n    r.GET("/ping", func(c *gin.Context) {\n        c.JSON(200, gin.H{"message": "pong"})\n    })\n    r.Run() // Запуск сервера\n}\n```',
        ARRAY['https://github.com/gin-gonic/gin', 'https://pkg.go.dev/github.com/gin-gonic/gin'],
        ARRAY['https://www.youtube.com/embed/wCZ5TJCBWMg'],
        ARRAY['https://example.com/gin_example.png']
    );



INSERT INTO quizzes (id, topic_id, question, options, correct_answers, multiple_choice)
VALUES
    (gen_random_uuid(), 'f4f8d4c9-4b76-4e1b-99e8-5d4a2f3b6c71', 'Какой фреймворк используется для создания REST API в Go?',
     ARRAY['Gin', 'Fiber', 'Echo', 'Django'], ARRAY[true, true, true, false], true),

    (gen_random_uuid(), 'f4f8d4c9-4b76-4e1b-99e8-5d4a2f3b6c71', 'Какие методы HTTP поддерживает Gin?',
     ARRAY['GET', 'POST', 'PUT', 'DELETE'], ARRAY[true, true, true, true], true),

    (gen_random_uuid(), 'f4f8d4c9-4b76-4e1b-99e8-5d4a2f3b6c71', 'Как создать маршрут в Gin?',
     ARRAY['r.GET()', 'r.POST()', 'r.PUT()', 'r.DELETE()'], ARRAY[true, true, true, true], true),

    (gen_random_uuid(), 'f4f8d4c9-4b76-4e1b-99e8-5d4a2f3b6c71', 'Как вернуть JSON-ответ в Gin?',
     ARRAY['c.JSON()', 'c.String()', 'c.HTML()', 'c.Render()'], ARRAY[true, false, false, false], false),

    (gen_random_uuid(), 'f4f8d4c9-4b76-4e1b-99e8-5d4a2f3b6c71', 'Как запустить сервер Gin?',
     ARRAY['r.Run()', 'r.Start()', 'r.Listen()', 'r.Serve()'], ARRAY[true, false, false, false], false);


INSERT INTO practical_tasks (id, topic_id, description, difficulty_level, order_number, starter_code, expected_output)
VALUES
    (gen_random_uuid(), 'f4f8d4c9-4b76-4e1b-99e8-5d4a2f3b6c71', 'Создайте сервер, который отвечает "pong" на GET-запрос `/ping`.', 'easy', 1,
     'package main\n\nimport (\n    "github.com/gin-gonic/gin"\n)\n\nfunc main() {\n    r := gin.Default()\n    r.GET("/ping", func(c *gin.Context) {\n        c.JSON(200, gin.H{"message": "pong"})\n    })\n    r.Run()\n}',
     '{"message": "pong"}'),

    (gen_random_uuid(), 'f4f8d4c9-4b76-4e1b-99e8-5d4a2f3b6c71', 'Добавьте маршрут POST `/user`, который принимает JSON с именем пользователя.', 'medium', 2,
     'package main\n\nimport (\n    "github.com/gin-gonic/gin"\n)\n\nfunc main() {\n    r := gin.Default()\n    r.POST("/user", func(c *gin.Context) {\n        var user struct { Name string `json:"name"` }\n        if err := c.BindJSON(&user); err != nil {\n            c.JSON(400, gin.H{"error": err.Error()})\n            return\n        }\n        c.JSON(200, gin.H{"message": "User created", "name": user.Name})\n    })\n    r.Run()\n}',
     '{"message": "User created", "name": "John"}'),

    (gen_random_uuid(), 'f4f8d4c9-4b76-4e1b-99e8-5d4a2f3b6c71', 'Добавьте обработку параметра `id` в маршруте `/user/:id`.', 'hard', 3,
     'package main\n\nimport (\n    "github.com/gin-gonic/gin"\n)\n\nfunc main() {\n    r := gin.Default()\n    r.GET("/user/:id", func(c *gin.Context) {\n        id := c.Param("id")\n        c.JSON(200, gin.H{"user_id": id})\n    })\n    r.Run()\n}',
     '{"user_id": "123"}');

-- 5 --
INSERT INTO topics (id, course_id, title, order_number, description)
VALUES
    ('f5d4a83b-7c5a-4f67-8e3c-1b5f9b2d4a6f', 'e9c3d8b5-4a71-4f3e-b99e-8d5f6a4b7c32', 'Работа с Docker и деплой', 5, 'Запуск Go-приложений в контейнерах и деплой.');


INSERT INTO topic_contents (id, topic_id, content, additional_resources, video_urls, image_urls)
VALUES
    (
        gen_random_uuid(), 'f5d4a83b-7c5a-4f67-8e3c-1b5f9b2d4a6f',
        '# Работа с Docker и деплой в Go\n\nDocker позволяет упаковать приложение в контейнер для удобного развертывания.\n\n## Установка Docker\n```sh\nsudo apt update && sudo apt install docker.io\n```\n\n## Dockerfile для Go-приложения\n```dockerfile\nFROM golang:1.18\nWORKDIR /app\nCOPY . .\nRUN go build -o main .\nCMD ["./main"]\n```',
        ARRAY['https://docs.docker.com/', 'https://hub.docker.com/'],
        ARRAY['https://www.youtube.com/embed/6fVj5pNCxps'],
        ARRAY['https://example.com/docker_example.png']
    );

INSERT INTO quizzes (id, topic_id, question, options, correct_answers, multiple_choice)
VALUES
    (gen_random_uuid(), 'f5d4a83b-7c5a-4f67-8e3c-1b5f9b2d4a6f', 'Что такое Docker?',
     ARRAY['Контейнеризация приложений', 'Виртуальная машина', 'Операционная система', 'Система контроля версий'], ARRAY[true, false, false, false], false),

    (gen_random_uuid(), 'f5d4a83b-7c5a-4f67-8e3c-1b5f9b2d4a6f', 'Какие команды используются в Docker?',
     ARRAY['docker build', 'docker run', 'docker commit', 'docker delete'], ARRAY[true, true, true, false], true),

    (gen_random_uuid(), 'f5d4a83b-7c5a-4f67-8e3c-1b5f9b2d4a6f', 'Как создать образ Docker?',
     ARRAY['docker build -t myapp .', 'docker create myapp', 'docker start myapp', 'docker compile myapp'], ARRAY[true, false, false, false], false),

    (gen_random_uuid(), 'f5d4a83b-7c5a-4f67-8e3c-1b5f9b2d4a6f', 'Как запустить контейнер Docker?',
     ARRAY['docker run myapp', 'docker execute myapp', 'docker start myapp', 'docker deploy myapp'], ARRAY[true, false, false, false], false),

    (gen_random_uuid(), 'f5d4a83b-7c5a-4f67-8e3c-1b5f9b2d4a6f', 'Какие инструменты используются для деплоя Go-приложений?',
     ARRAY['Docker', 'Kubernetes', 'Heroku', 'PostgreSQL'], ARRAY[true, true, true, false], true);


INSERT INTO practical_tasks (id, topic_id, description, difficulty_level, order_number, starter_code, expected_output)
VALUES
    (gen_random_uuid(), 'f5d4a83b-7c5a-4f67-8e3c-1b5f9b2d4a6f', 'Создайте `Dockerfile` для Go-приложения.', 'easy', 1,
     'FROM golang:1.18\nWORKDIR /app\nCOPY . .\nRUN go build -o main .\nCMD ["./main"]',
     'Dockerfile создан'),

    (gen_random_uuid(), 'f5d4a83b-7c5a-4f67-8e3c-1b5f9b2d4a6f', 'Создайте и запустите контейнер с Go-приложением.', 'medium', 2,
     'docker build -t myapp .\ndocker run -p 8080:8080 myapp',
     'Контейнер запущен'),

    (gen_random_uuid(), 'f5d4a83b-7c5a-4f67-8e3c-1b5f9b2d4a6f', 'Настройте деплой Go-приложения с Docker на сервер.', 'hard', 3,
     'scp myapp.tar user@server:/home/user/\ndocker load -i myapp.tar\ndocker run -d -p 80:8080 myapp',
     'Приложение развернуто');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM courses;
DELETE FROM topics;
DELETE FROM topic_contents;
DELETE FROM quizzes;
DELETE FROM practical_tasks;
DELETE FROM course_category;
-- +goose StatementEnd
