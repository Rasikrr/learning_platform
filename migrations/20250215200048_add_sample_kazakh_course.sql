-- +goose Up
-- +goose StatementBegin
INSERT INTO courses (id, title, image_url, category_id, description, created_by, created_at, updated_at)
VALUES (
           'aacf4fdd-88ee-44d5-9868-a03ae847ca3a',
           'Бэкендке арналған Python',
           'https://hr-profi.ru/img/podbor-it/python/image_python_01.jpg',
           'a1d4a83e-3d65-4df5-b890-1b1d87b24a70', -- ID санаты Backend
           'Python көмегімен серверлік қосымшаларды әзірлеу',
           'da00c311-1203-4cc0-99c9-4104f31e482e', -- Өз user_id-іңді қой
           NOW(),
           NOW()
       );

INSERT INTO topics (id, course_id, title, order_number, created_at, updated_at, description)
VALUES
    ('89d1bb78-aeab-4a8d-b63d-68ac28f49aca', 'aacf4fdd-88ee-44d5-9868-a03ae847ca3a', 'Айнымалылар мен деректер түрлері', 1, NOW(), NOW(), 'Python-дағы негізгі деректер түрлері: int, float, string, list, tuple, dict.'),
    ('1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd', 'aacf4fdd-88ee-44d5-9868-a03ae847ca3a', 'Шартты операторлар', 2, NOW(), NOW(), 'if, elif, else операторларын қолдану. Логикалық операциялар.'),
    ('b4009b69-1c08-4c31-98a4-e4965bd90852', 'aacf4fdd-88ee-44d5-9868-a03ae847ca3a', 'Циклдер', 3, NOW(), NOW(), 'for және while циклдерін пайдалану, тізімдермен және диапазондармен жұмыс.'),
    ('5d1a896d-c822-4250-bd2f-9a0ff71a8834', 'aacf4fdd-88ee-44d5-9868-a03ae847ca3a', 'Функциялар', 4, NOW(), NOW(), 'def, параметрлер, аргументтер және қайтару мәндері.'),
    ('6c3f328c-6a03-4072-8ee7-11796ddc2676', 'aacf4fdd-88ee-44d5-9868-a03ae847ca3a', 'Қателер мен ерекшеліктер', 5, NOW(), NOW(), 'try-except конструкциясы және қателерді өңдеу тәсілдері.');

INSERT INTO topic_contents (id, topic_id, content, additional_resources, video_urls, image_urls, created_at, updated_at)
VALUES
    -- Айнымалылар мен деректер түрлері
    (gen_random_uuid(), '89d1bb78-aeab-4a8d-b63d-68ac28f49aca',
     '# Айнымалылар мен деректер түрлері\n\nPython-да айнымалылар мәндерді сақтау үшін қолданылады.\n\n## Айнымалыны анықтау\n```python\nx = 10\ny = "Python"\nprint(x, y)\n```',
     '{"https://docs.python.org/3/tutorial/introduction.html"}',
     '{"https://www.youtube.com/embed/Y54qPoB..."}',
     '{"https://example.com/python_variables.png"}',
     NOW(), NOW()),

    -- Шартты операторлар
    (gen_random_uuid(), '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd',
     '# Шартты операторлар\n\nPython-да if-else операторлары шарттарды тексеру үшін қолданылады.\n\n## Мысал\n```python\nx = 10\nif x > 5:\n    print("Үлкен")\nelse:\n    print("Кіші")\n```',
     '{"https://docs.python.org/3/tutorial/controlflow.html"}',
     '{"https://www.youtube.com/embed/UqV59wS..."}',
     '{"https://example.com/python_if_else.png"}',
     NOW(), NOW()),

    -- Циклдер
    (gen_random_uuid(), 'b4009b69-1c08-4c31-98a4-e4965bd90852',
     '# Циклдер\n\nPython-да for және while циклдері қайталанатын операцияларды орындау үшін қолданылады.\n\n## For циклы мысалы\n```python\nfor i in range(5):\n    print(i)\n```',
     '{"https://docs.python.org/3/tutorial/controlflow.html"}',
     '{"https://www.youtube.com/embed/wRZWES..."}',
     '{"https://example.com/python_loops.png"}',
     NOW(), NOW()),

    -- Функциялар
    (gen_random_uuid(), '5d1a896d-c822-4250-bd2f-9a0ff71a8834',
     '# Функциялар\n\nPython-да функциялар кодты қайта пайдалану үшін қолданылады.\n\n## Функцияны анықтау\n```python\ndef hello():\n    print("Сәлем, әлем!")\n\nhello()\n```',
     '{"https://docs.python.org/3/tutorial/controlflow.html#defining-functions"}',
     '{"https://www.youtube.com/embed/WCZTIJ..."}',
     '{"https://example.com/python_functions.png"}',
     NOW(), NOW()),

    -- Қателер мен ерекшеліктер
    (gen_random_uuid(), '6c3f328c-6a03-4072-8ee7-11796ddc2676',
     '# Қателер мен ерекшеліктер\n\nPython-да try-except блогы қателерді өңдеу үшін қолданылады.\n\n## Мысал\n```python\ntry:\n    num = int(input("Сан енгізіңіз: "))\n    print(num * 2)\nexcept ValueError:\n    print("Қате! Сан енгізіңіз.")\n```',
     '{"https://docs.python.org/3/tutorial/errors.html"}',
     '{"https://www.youtube.com/embed/6VpJ5n..."}',
     '{"https://example.com/python_errors.png"}',
     NOW(), NOW());


INSERT INTO quizzes (id, topic_id, question, options, correct_answers, multiple_choice, created_at, updated_at)
VALUES
    -- Айнымалылар мен деректер түрлері (89d1bb78-aeab-4a8d-b63d-68ac28f49aca)
    (gen_random_uuid(), '89d1bb78-aeab-4a8d-b63d-68ac28f49aca', 'Python-да айнымалыны қалай анықтаймыз?',
     '{"x = 5", "let x = 5", "var x = 5", "define x = 5"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '89d1bb78-aeab-4a8d-b63d-68ac28f49aca', 'Python-да жолдық (string) мәнді қалай көрсетуге болады?',
     '{"Мәтін", "Мәтін", "Мәтін", "Барлық нұсқалар дұрыс"}', '{false, false, false, true}', false, NOW(), NOW()),

    (gen_random_uuid(), '89d1bb78-aeab-4a8d-b63d-68ac28f49aca', 'Python-да бүтін сан түрі қалай аталады?',
     '{"int", "float", "string", "boolean"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '89d1bb78-aeab-4a8d-b63d-68ac28f49aca', 'Python-да сөздік (dictionary) қалай құрылады?',
     '{"{}", "{}", "()", "None"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '89d1bb78-aeab-4a8d-b63d-68ac28f49aca', 'Python-да логикалық мәндер қандай?',
     '{"True және False", "1 және 0", "Yes және No", "Барлық нұсқалар дұрыс"}', '{true, false, false, false}', false, NOW(), NOW()),

    -- Шартты операторлар (1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd)
    (gen_random_uuid(), '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd', 'if шарты қандай мән қабылдайды?',
     '{"True немесе False", "1 немесе 0", "Кез келген сан", "Кез келген жол"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd', 'if-else шартында else бөлімі міндетті ме?',
     '{"Иә", "Жоқ", "Кейбір жағдайларда", "Python нұсқасына байланысты"}', '{false, true, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd', 'elif не үшін қолданылады?',
     '{"Қосымша шарттарды тексеру", "Кодты жылдам орындау", "Цикл жасау", "Файлды ашу"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd', 'Қай оператор теңдік тексеру үшін қолданылады?',
     '{"=", "==", "!=", ">="}', '{false, true, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd', 'Қай оператор терістеу үшін қолданылады?',
     '{"not", "!", "<>", "inverse"}', '{true, false, false, false}', false, NOW(), NOW()),

    -- Циклдер (b4009b69-1c08-4c31-98a4-e4965bd90852)
    (gen_random_uuid(), 'b4009b69-1c08-4c31-98a4-e4965bd90852', 'Python-да қай циклдер қолданылады?',
     '{"for, while", "do-while, for", "repeat, until", "switch, case"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), 'b4009b69-1c08-4c31-98a4-e4965bd90852', 'while циклы қашан орындалады?',
     '{"Шарт true болғанда", "Әрқашан", "Шарт false болғанда", "Кодтың соңында"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), 'b4009b69-1c08-4c31-98a4-e4965bd90852', 'for циклы негізінен не үшін қолданылады?',
     '{"Итерация үшін", "Массив құру үшін", "Файлды оқу үшін", "Базаға қосылу үшін"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), 'b4009b69-1c08-4c31-98a4-e4965bd90852', 'break операторы не істейді?',
     '{"Циклді тоқтатады", "Циклді қайталайды", "Кодты іске қосады", "Айнымалы жасайды"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), 'b4009b69-1c08-4c31-98a4-e4965bd90852', 'continue операторы не істейді?',
     '{"Циклдің келесі итерациясына өтеді", "Циклді тоқтатады", "Айнымалы мәнін өзгертеді", "Тек while циклында қолданылады"}', '{true, false, false, false}', false, NOW(), NOW()),

    -- Функциялар (5d1a896d-c822-4250-bd2f-9a0ff71a8834)
    (gen_random_uuid(), '5d1a896d-c822-4250-bd2f-9a0ff71a8834', 'Python-да функция қалай анықталады?',
     '{"def функция_аты():", "func функция_аты():", "function функция_аты()", "fn функция_аты()"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '5d1a896d-c822-4250-bd2f-9a0ff71a8834', 'Функция не үшін қолданылады?',
     '{"Кодты қайта пайдалану үшін", "Жылдамдықты арттыру үшін", "Тек деректерді сақтау үшін", "Python нұсқасына байланысты"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '5d1a896d-c822-4250-bd2f-9a0ff71a8834', 'Python-да параметрсіз функция қалай құрылады?',
     '{"def my_function():", "def my_function(x):", "function my_function():", "func my_function()"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '5d1a896d-c822-4250-bd2f-9a0ff71a8834', 'Функция қай операторды пайдаланып мән қайтара алады?',
     '{"return", "yield", "output", "break"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '5d1a896d-c822-4250-bd2f-9a0ff71a8834', 'Lambda функциясы деген не?',
     '{"Анонимді функция", "Тек бір жолдан тұратын функция", "Python-ның негізгі функциясы", "Барлық жауаптар дұрыс"}', '{true, true, false, false}', true, NOW(), NOW()),


    -- Қателер мен ерекшеліктер (6c3f328c-6a03-4072-8ee7-11796ddc2676)
    (gen_random_uuid(), '6c3f328c-6a03-4072-8ee7-11796ddc2676', 'Қандай оператор қателерді өңдеуге арналған?',
     '{"try-except", "catch", "error-handler", "debug"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '6c3f328c-6a03-4072-8ee7-11796ddc2676', 'Қандай қате орындалу кезінде (runtime) пайда болады?',
     '{"SyntaxError", "RuntimeError", "ImportError", "ValueError"}', '{false, true, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '6c3f328c-6a03-4072-8ee7-11796ddc2676', 'try-except блогында finally бөлімі не үшін қолданылады?',
     '{"Әрқашан орындалу үшін", "Қатені ұстап қалу үшін", "Кодтың жылдамдығын арттыру үшін", "Цикл жасау үшін"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '6c3f328c-6a03-4072-8ee7-11796ddc2676', 'raise операторы не үшін қолданылады?',
     '{"Қолмен қате шақыру үшін", "Қателерді түзету үшін", "Кодтың орындалуын тоқтату үшін", "Қосымша модуль жүктеу үшін"}', '{true, false, false, false}', false, NOW(), NOW()),

    (gen_random_uuid(), '6c3f328c-6a03-4072-8ee7-11796ddc2676', 'except блогында қате түрін көрсету міндетті ме?',
     '{"Иә, міндетті", "Жоқ, міндетті емес", "Python нұсқасына байланысты", "Тек SyntaxError үшін міндетті"}', '{false, true, false, false}', false, NOW(), NOW());



INSERT INTO practical_tasks (id, topic_id, description, difficulty_level, starter_code, expected_output, created_at, updated_at, order_number, test_cases, language)
VALUES
    -- Айнымалылар мен деректер түрлері (89d1bb78-aeab-4a8d-b63d-68ac28f49aca)
    (gen_random_uuid(), '89d1bb78-aeab-4a8d-b63d-68ac28f49aca', 'Айнымалыға мән тағайындап, оны экранға шығарыңыз.', 'easy',
     'x = 10\nprint(x)', '10', NOW(), NOW(), 1, false, 'python'),

    (gen_random_uuid(), '89d1bb78-aeab-4a8d-b63d-68ac28f49aca', 'Жолдық айнымалы құрыңыз және оны басып шығарыңыз.', 'easy',
     'text = "Python"\nprint(text)', 'Python', NOW(), NOW(), 2, false, 'python'),

    (gen_random_uuid(), '89d1bb78-aeab-4a8d-b63d-68ac28f49aca', 'Пайдаланушыдан сан сұрап, оны көбейтіп, экранға шығарыңыз.', 'medium',
     'num = int(input())\nprint(num * 2)', '20', NOW(), NOW(), 3, true, 'python'),

    -- Шартты операторлар (1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd)
    (gen_random_uuid(), '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd', 'Берілген санның тақ немесе жұп екенін анықтаңыз.', 'easy',
     'num = int(input())\nif num % 2 == 0:\n    print("Жұп")\nelse:\n    print("Тақ")', 'Жұп', NOW(), NOW(), 1, true, 'python'),

    (gen_random_uuid(), '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd', 'Берілген сан оң немесе теріс екенін анықтаңыз.', 'easy',
     'num = int(input())\nif num > 0:\n    print("Оң")\nelif num < 0:\n    print("Теріс")\nelse:\n    print("Нөл")', 'Оң', NOW(), NOW(), 2, true, 'python'),

    (gen_random_uuid(), '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd', 'Пайдаланушы енгізген жылдың кібісе жыл (високосный) екенін тексеріңіз.', 'medium',
     'year = int(input())\nif (year % 4 == 0 and year % 100 != 0) or (year % 400 == 0):\n    print("Кібісе жыл")\nelse:\n    print("Кібісе жыл емес")', 'Кібісе жыл', NOW(), NOW(), 3, true, 'python'),

    -- Циклдер (b4009b69-1c08-4c31-98a4-e4965bd90852)
    (gen_random_uuid(), 'b4009b69-1c08-4c31-98a4-e4965bd90852', '1-ден 10-ға дейінгі сандарды шығару үшін for циклын пайдаланыңыз.', 'easy',
     'for i in range(1, 11):\n    print(i)', '1\n2\n3\n4\n5\n6\n7\n8\n9\n10', NOW(), NOW(), 1, false, 'python'),

    (gen_random_uuid(), 'b4009b69-1c08-4c31-98a4-e4965bd90852', '1-ден 100-ге дейінгі барлық жұп сандардың қосындысын есептеңіз.', 'medium',
     'sum_even = sum(i for i in range(1, 101) if i % 2 == 0)\nprint(sum_even)', '2550', NOW(), NOW(), 2, true, 'python'),

    (gen_random_uuid(), 'b4009b69-1c08-4c31-98a4-e4965bd90852', 'Қолданушы енгізген сөзді 5 рет басып шығарыңыз.', 'easy',
     'word = input()\nfor i in range(5):\n    print(word)', 'Python\nPython\nPython\nPython\nPython', NOW(), NOW(), 3, true, 'python'),

    -- Функциялар (5d1a896d-c822-4250-bd2f-9a0ff71a8834)
    (gen_random_uuid(), '5d1a896d-c822-4250-bd2f-9a0ff71a8834', 'Екі санды қабылдап, олардың қосындысын қайтаратын функция жазыңыз.', 'easy',
     'def sum_two(a, b):\n    return a + b\nprint(sum_two(3, 5))', '8', NOW(), NOW(), 1, true, 'python'),

    (gen_random_uuid(), '5d1a896d-c822-4250-bd2f-9a0ff71a8834', 'Жолдық аргумент қабылдайтын және оның ұзындығын қайтаратын функция жазыңыз.', 'easy',
     'def length(text):\n    return len(text)\nprint(length("Python"))', '6', NOW(), NOW(), 2, true, 'python'),

    (gen_random_uuid(), '5d1a896d-c822-4250-bd2f-9a0ff71a8834', 'Lambda функциясын пайдаланып, екі санның көбейтіндісін есептеңіз.', 'medium',
     'mul = lambda x, y: x * y\nprint(mul(4, 5))', '20', NOW(), NOW(), 3, true, 'python'),

    -- Қателер мен ерекшеліктер (6c3f328c-6a03-4072-8ee7-11796ddc2676)
    (gen_random_uuid(), '6c3f328c-6a03-4072-8ee7-11796ddc2676', 'Қатені өңдеу үшін try-except блогын пайдаланып, сан енгізуді сұраңыз.', 'medium',
     'try:\n    num = int(input("Сан енгізіңіз: "))\n    print(num * 2)\nexcept ValueError:\n    print("Қате! Сан енгізіңіз.")', '10', NOW(), NOW(), 1, true, 'python'),

    (gen_random_uuid(), '6c3f328c-6a03-4072-8ee7-11796ddc2676', 'Файлды оқуға тырысып, қатені өңдеңіз.', 'hard',
     'try:\n    with open("file.txt", "r") as f:\n        print(f.read())\nexcept FileNotFoundError:\n    print("Файл табылмады")', 'Файл табылмады', NOW(), NOW(), 2, true, 'python'),

    (gen_random_uuid(), '6c3f328c-6a03-4072-8ee7-11796ddc2676', 'raise көмегімен өзіңіздің қателігіңізді тастаңыз.', 'hard',
     'def check_age(age):\n    if age < 18:\n        raise ValueError("Кіру рұқсат етілмеген!")\n    print("Рұқсат етілді")\ntry:\n    check_age(16)\nexcept ValueError as e:\n    print(e)', 'Кіру рұқсат етілмеген!', NOW(), NOW(), 3, true, 'python');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM courses
WHERE id = 'aacf4fdd-88ee-44d5-9868-a03ae847ca3a';

DELETE FROM topics
WHERE id IN ('89d1bb78-aeab-4a8d-b63d-68ac28f49aca',
             '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd',
             'b4009b69-1c08-4c31-98a4-e4965bd90852',
             '5d1a896d-c822-4250-bd2f-9a0ff71a8834',
             '6c3f328c-6a03-4072-8ee7-11796ddc2676');

DELETE FROM quizzes
WHERE topic_id IN ('89d1bb78-aeab-4a8d-b63d-68ac28f49aca',
                   '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd',
                   'b4009b69-1c08-4c31-98a4-e4965bd90852',
                   '5d1a896d-c822-4250-bd2f-9a0ff71a8834',
                   '6c3f328c-6a03-4072-8ee7-11796ddc2676');

DELETE FROM practical_tasks
WHERE topic_id IN ('89d1bb78-aeab-4a8d-b63d-68ac28f49aca',
                   '1dc2b89e-a4bb-48e5-a1b1-3763fa7784fd',
                   'b4009b69-1c08-4c31-98a4-e4965bd90852',
                   '5d1a896d-c822-4250-bd2f-9a0ff71a8834',
                   '6c3f328c-6a03-4072-8ee7-11796ddc2676');
-- +goose StatementEnd
