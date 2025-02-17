-- +goose Up
-- +goose StatementBegin
INSERT INTO faq_categories (id, name, created_at, updated_at) VALUES
('41466d9a-769f-4ac5-8df4-fca433333cf6', 'Data Science', NOW(), NOW()),
('46475485-39d1-480d-b999-b8ac3d0d4c45', 'Frontend Development', NOW(), NOW()),
('01c8cdd2-5a06-4e46-9fe8-efe3cd9c7f69', 'Backend Development', NOW(), NOW()),
('0299259c-058b-4f12-9bc3-a86766e893a9', 'DevOps & Cloud', NOW(), NOW()),
('97421703-4157-47a2-86a7-9f693a684315', 'Machine Learning & AI', NOW(), NOW()),
('9f23bdfc-ec05-4c89-a258-b8142e7e5312', 'Cybersecurity', NOW(), NOW()),
('525d7793-125e-4b45-9291-179a83cc5ab9', 'Mobile Development', NOW(), NOW()),
('f53877cd-7614-447c-8920-a03484fc0fae', 'Game Development', NOW(), NOW()),
('edb4e5f9-f1e4-4cc8-9e61-04d52ce3d496', 'Database Management', NOW(), NOW()),
('1c298e99-22db-4a26-83bc-b42622881612', 'Software Testing & QA', NOW(), NOW());


INSERT INTO faq_questions (id, category_id, user_id, title, body, created_at, updated_at) VALUES
-- Data Science
('a683c117-d30a-4625-9305-6a6abdfecdbd', '41466d9a-769f-4ac5-8df4-fca433333cf6', 'f315df8e-b8cd-4c3e-b2eb-c287acaf8c31', 'Python-да деректерді талдау үшін қандай кітапханалар қолданылады?', 'Python тілінде деректерді талдау үшін ең жиі қолданылатын кітапханалар қандай? Олардың негізгі ерекшеліктері мен айырмашылықтары қандай?', NOW(), NOW()),
('7fadacc0-d2c8-46d5-b5ee-a1c73a16b1dd', '41466d9a-769f-4ac5-8df4-fca433333cf6', '0a49a92a-694c-429f-b77a-6994ff4441e0', 'Machine Learning моделін бағалау метрикалары', 'Машиналық оқыту модельдерін бағалау үшін қандай негізгі метрикалар қолданылады? Қай жағдайда қандай метриканы таңдау керек?', NOW(), NOW()),
('4bcd9a99-2db5-4c9b-9a6d-6f4f17cec0d5', '41466d9a-769f-4ac5-8df4-fca433333cf6', '639fd406-6af2-4e50-a090-0d92d428749d', 'Нейрондық желілер мен шешімдер ағашының айырмашылығы', 'Нейрондық желілер мен шешімдер ағашының арасындағы негізгі айырмашылықтар қандай? Әр әдіс қандай есептерді шешу үшін қолданылады?', NOW(), NOW()),

-- Frontend Development
('31872be6-c209-4d25-a0da-52dbc29542ad', '46475485-39d1-480d-b999-b8ac3d0d4c45', 'b9cf3d64-2b19-4d02-812f-cbebbb437bbc', 'React пен Vue айырмашылығы неде?', 'React пен Vue фреймворктарының арасындағы негізгі айырмашылықтар қандай? Қай жағдайда қайсысын қолданған дұрыс?', NOW(), NOW()),
('7e18d023-7711-4bb1-8ab2-a29ebfd6510f', '46475485-39d1-480d-b999-b8ac3d0d4c45', '2de42bcb-a2ea-40cf-8d86-6498c09b3d1f', 'CSS Grid пен Flexbox қайсысы жақсы?', 'CSS Grid пен Flexbox арасындағы айырмашылықтар қандай? Қайсысын қандай жағдайда қолдану тиімді?', NOW(), NOW()),
('0a7eafb3-c0f6-4a06-986d-ed17c8a4cc02', '46475485-39d1-480d-b999-b8ac3d0d4c45', '1dbf9a0e-aa30-4002-ac80-6bb6d8e19495', 'JavaScript-тің асинхронды функциялары қалай жұмыс істейді?', 'JavaScript-те асинхронды функциялар қалай орындалады? async/await және Promise механизмдерінің рөлі қандай?', NOW(), NOW()),

-- Backend Development
('151e0c9a-4e8d-4aac-85c2-62a2e3879769', '01c8cdd2-5a06-4e46-9fe8-efe3cd9c7f69', 'ec9f58d5-82d5-442e-a664-ad21f2d0c17b', 'REST пен GraphQL айырмашылығы қандай?', 'REST пен GraphQL архитектураларының негізгі айырмашылықтары қандай? Қайсысын қандай жағдайларда қолдану тиімді?', NOW(), NOW()),
('bed8e315-421f-41c7-94b7-df239273bcce', '01c8cdd2-5a06-4e46-9fe8-efe3cd9c7f69', '3c5456ff-eef4-4c30-9f37-e52566fc7032', 'Golang-та көп ағынды өңдеу қалай іске асырылады?', 'Golang-та көп ағынды бағдарламалау қалай жұмыс істейді? Goroutines және Channels қалай қолданылады?', NOW(), NOW()),
('77571461-b2b5-4689-bd7e-e957611eca65', '01c8cdd2-5a06-4e46-9fe8-efe3cd9c7f69', '13e42502-d51a-42c8-a8a8-b92e1104bdd4', 'JWT (JSON Web Token) қалай жұмыс істейді?', 'JWT пайдаланушы аутентификациясында қалай жұмыс істейді? Оның негізгі компоненттері қандай?', NOW(), NOW()),

-- DevOps & Cloud
('85588286-2425-49b3-af00-5d3c7344200c', '0299259c-058b-4f12-9bc3-a86766e893a9', '964b13c7-3798-4551-8f5e-db7c181e8320', 'Docker мен Kubernetes айырмашылығы қандай?', 'Docker мен Kubernetes технологиялары қалай ерекшеленеді? Оларды бірге қолдануға бола ма?', NOW(), NOW()),
('0519f252-c6de-4730-909a-ca0dd90308f0', '0299259c-058b-4f12-9bc3-a86766e893a9', '106ddee2-e2a9-4035-921f-bd5cbe7a628e', 'CI/CD дегеніміз не және оны қалай орнатуға болады?', 'CI/CD дегеніміз не? Оны іске асыру үшін қандай құралдар пайдаланылады?', NOW(), NOW()),
('07a3354a-78f9-4a07-8883-2fdbc1336473', '0299259c-058b-4f12-9bc3-a86766e893a9', 'c992013c-5cde-48cf-8af0-00184338baf9', 'AWS пен Google Cloud қайсысы жақсырақ?', 'AWS пен Google Cloud арасындағы негізгі айырмашылықтар қандай? Қайсысы қай жағдайда тиімді?', NOW(), NOW()),

-- Machine Learning & AI
('6e952333-d0b9-4b64-8f9e-0fe4b0516b35', '97421703-4157-47a2-86a7-9f693a684315', '9fa1f519-d805-42c7-b9b1-a752365d1c56', 'Нейрондық желілерді оқытудың негізгі қадамдары қандай?', 'Нейрондық желілерді оқытудың негізгі қадамдары қандай? Кері тарату алгоритмі қалай жұмыс істейді?', NOW(), NOW()),
('a6da7bc0-a04c-47c3-9132-f6bfa0851539', '97421703-4157-47a2-86a7-9f693a684315', 'd131a22f-0b3d-4767-9806-6a277dc0744d', 'TensorFlow мен PyTorch айырмашылығы қандай?', 'TensorFlow мен PyTorch кітапханаларының арасындағы негізгі айырмашылықтар қандай? Қайсысы зерттеу үшін, қайсысы өндірістік қолданыс үшін тиімді?', NOW(), NOW()),

-- Cybersecurity
('f317f5c9-efc1-4778-81d3-994ba74686e7', '9f23bdfc-ec05-4c89-a258-b8142e7e5312', '30bd5baf-d603-4b23-936b-c3633df709fe', 'SQL Injection дегеніміз не және оны қалай алдын алуға болады?', 'SQL Injection қалай орындалады және оны болдырмау үшін қандай әдістер қолданылады?', NOW(), NOW()),
('33bb30e8-5272-4ac5-bfaf-6592caadf377', '9f23bdfc-ec05-4c89-a258-b8142e7e5312', '636e0dc6-ce94-4f26-bd39-b8fb32b4b378', 'Хэштеу мен шифрлау арасындағы айырмашылық қандай?', 'Хэштеу мен шифрлау арасындағы негізгі айырмашылықтар қандай? Әрқайсысы қандай мақсатта қолданылады?', NOW(), NOW()),

-- Mobile Development
('82e87700-e4c0-422d-a743-fef120f9233a', '525d7793-125e-4b45-9291-179a83cc5ab9', 'cb2d6442-ec58-4fb6-aea6-23dcd2a6b574', 'Flutter мен React Native қайсысы жақсы?', 'Flutter мен React Native арасындағы негізгі айырмашылықтар қандай? Қайсысын қандай жобаларда қолданған дұрыс?', NOW(), NOW()),
('7fd5573a-1d73-409a-bae6-2a3e20cfbd56', '525d7793-125e-4b45-9291-179a83cc5ab9', '13d450c5-6e2c-4bd8-bc5a-9096f59497b2', 'Android пен iOS үшін нативті қолданба жасаудың айырмашылықтары қандай?', 'Android пен iOS үшін нативті қолданба әзірлеудің негізгі ерекшеліктері қандай? Қай тілдер қолданылады?', NOW(), NOW()),

-- Game Development
('bc7201c7-91d2-42cb-90a2-1a58b14bcf21', 'f53877cd-7614-447c-8920-a03484fc0fae', 'a58335b7-8e7c-4276-8c15-e4e16ca894ec', 'Unity мен Unreal Engine айырмашылығы қандай?', 'Unity мен Unreal Engine ойын қозғалтқыштарының негізгі айырмашылықтары қандай? Қайсысы қандай ойындарға лайықты?', NOW(), NOW()),
('81e157cf-c04e-44c8-9279-8d6691942409', 'f53877cd-7614-447c-8920-a03484fc0fae', '65148bce-f576-4e3a-b30b-4951cc1e3621', 'Ойын физикасы қалай жүзеге асырылады?', 'Ойындарда физикалық модельдеу қалай іске асырылады? Қандай физикалық қозғалтқыштар қолданылады?', NOW(), NOW()),

-- Database Management
('44168e76-80be-4019-9eea-481dbaf0f8d4', 'edb4e5f9-f1e4-4cc8-9e61-04d52ce3d496', 'cfd3b85c-2e4a-4fe6-8bfe-4eafe6612759', 'SQL мен NoSQL айырмашылығы қандай?', 'SQL мен NoSQL дерекқорларының негізгі айырмашылықтары қандай? Әрқайсысы қандай жағдайларда қолданылады?', NOW(), NOW()),
('24cf8351-c647-4010-ac00-3f22bf56bc2a', 'edb4e5f9-f1e4-4cc8-9e61-04d52ce3d496', '8d44290e-c670-42ee-bbc1-bf5917dd80de', 'PostgreSQL мен MySQL қайсысы жақсы?', 'PostgreSQL мен MySQL арасындағы айырмашылықтар қандай? Қай дерекқор жүйесі қандай сценарийлерге тиімді?', NOW(), NOW());



INSERT INTO faq_answers (id, question_id, user_id, body, created_at, updated_at) VALUES
-- Data Science
('a106bc74-7072-4cdb-845e-1e7d3d258caf', 'a683c117-d30a-4625-9305-6a6abdfecdbd', '74f5e34d-a966-420a-8427-f3abb843ee0f', 'Python тілінде деректерді талдау үшін Pandas, NumPy және SciPy кеңінен қолданылады.', NOW(), NOW()),
('57183ac9-c06f-4fbf-9a8b-da1f04899cb9', '7fadacc0-d2c8-46d5-b5ee-a1c73a16b1dd', '8b864659-a44a-41a6-8bfe-0821d5dbd0dc', 'Machine Learning моделін бағалау үшін дәлдік (Accuracy), F1-score, Precision және Recall қолданылады.', NOW(), NOW()),
('77d7f414-d6bc-4072-8464-247404cb01cb', '4bcd9a99-2db5-4c9b-9a6d-6f4f17cec0d5', '4ad4a8f2-5743-4c51-ba05-79147ed01a14', 'Нейрондық желілер үлгілерді тануға, ал шешімдер ағаштары ережелерге негізделген шешім қабылдауға жарамды.', NOW(), NOW()),

-- Frontend Development
('bbe4cd54-d06c-41da-96e2-42b1e3596a03', '31872be6-c209-4d25-a0da-52dbc29542ad', 'a0f389a1-f340-42da-9d8f-e54bbb01a83d', 'React – компоненттік құрылымға негізделген, Vue – қарапайым және жеңіл үйренуге болатын framework.', NOW(), NOW()),
('f84c3b9f-e471-49a6-b47f-d44fa6b855fd', '7e18d023-7711-4bb1-8ab2-a29ebfd6510f', '8a9508d6-17ff-4582-b108-d5f03d906f3a', 'Flexbox бір өлшемді орналасу үшін жақсы, ал CSS Grid екі өлшемді макеттер үшін тиімді.', NOW(), NOW()),
('96f4606e-4e7f-436e-ad35-3cee56f192e5', '0a7eafb3-c0f6-4a06-986d-ed17c8a4cc02', 'b2026668-3ed4-49cf-adde-89eba5f827e4', 'Асинхронды JavaScript функциялары async/await және Promise көмегімен орындалады.', NOW(), NOW()),

-- Backend Development
('6ce1cd46-fbf1-4afe-b879-3f49f7b4f587', '151e0c9a-4e8d-4aac-85c2-62a2e3879769', 'f17b3266-2cfc-4973-9b22-b058582df736', 'REST деректерді URL арқылы қайтарады, ал GraphQL сұранысты оңтайландырып қажетті деректерді ғана қайтарады.', NOW(), NOW()),
('03af2d04-2b8b-4e4e-adf7-cef7f672ba03', 'bed8e315-421f-41c7-94b7-df239273bcce', '5b268d42-dc4a-42fb-ad63-8d0c1422ae59', 'Golang-та goroutine және channels механизмдері арқылы көп ағынды өңдеу жүзеге асырылады.', NOW(), NOW()),
('da99333d-5f39-4a38-8bdc-1375d3092cb5', '77571461-b2b5-4689-bd7e-e957611eca65', 'da99333d-5f39-4a38-8bdc-1375d3092cb5', 'JWT пайдаланушы аутентификациясы үшін қолданылады және ол Header, Payload және Signature бөліктерінен тұрады.', NOW(), NOW()),

-- DevOps & Cloud
('61576177-bc94-4f77-a0d4-bb3278b1d871', '85588286-2425-49b3-af00-5d3c7344200c', '5236b356-aad4-45ea-91d5-023a5917e148', 'Docker – контейнерлеу платформасы, ал Kubernetes – контейнерлерді басқару және оркестрациялау құралы.', NOW(), NOW()),
('f1993d7c-79c0-4839-a871-d2ac288050ed', '0519f252-c6de-4730-909a-ca0dd90308f0', 'c3e99151-af51-4ce4-9d30-fcd976162ed7', 'CI/CD процесін орнату үшін Jenkins, GitHub Actions немесе GitLab CI қолданылады.', NOW(), NOW()),
('48e605fc-90b3-4d6b-befa-9426cd4f6678', '07a3354a-78f9-4a07-8883-2fdbc1336473', '62b61f53-3d68-481c-9894-767e34b9d831', 'AWS ауқымды инфрақұрылымға ие, ал Google Cloud мәліметтерді өңдеуге жақсы.', NOW(), NOW()),

-- Machine Learning & AI
('4ccfdf4b-6796-4a8a-8937-acaec6e5d85b', '6e952333-d0b9-4b64-8f9e-0fe4b0516b35', '53b7c283-5ca4-4f70-8dc2-fd0c69479d12', 'Нейрондық желілерді оқыту деректерді дайындаудан, модель құрудан, шығын функциясын анықтаудан, кері таратудан және параметрлерді реттеуден тұрады.', NOW(), NOW()),
('a55dc36a-5c27-4cb6-9d70-7e9efc945016', 'a6da7bc0-a04c-47c3-9132-f6bfa0851539', '58e83dd0-8806-434f-8c49-86fd028129e3', 'TensorFlow – өндірістік қолдануға, ал PyTorch – зерттеу мен прототиптеуге ыңғайлы.', NOW(), NOW()),

-- Cybersecurity
('73231892-0fd3-4142-abc3-99eebc32da54', 'f317f5c9-efc1-4778-81d3-994ba74686e7', 'bce6d0e5-2d07-4a61-8d24-7b67840eac24', 'SQL Injection сұраныс ішіне зиянды SQL кодын енгізуді білдіреді, оны болдырмау үшін параметрленген сұраныстарды қолдану керек.', NOW(), NOW()),
('58e686e8-b72e-459d-8827-e8969918c37d', '33bb30e8-5272-4ac5-bfaf-6592caadf377', 'd873c134-9762-46e1-969a-32b5b9387ea1', 'Хэштеу деректерді біржақты шифрлайды, ал шифрлау кері шифрлауды қарастырады.', NOW(), NOW()),

-- Mobile Development
('c0455593-c298-473d-b77a-a8f3f38326ba', '82e87700-e4c0-422d-a743-fef120f9233a', 'a558bfdd-6cee-445f-b7fd-a2215c4e1c6a', 'Flutter Dart тілін қолданады, ал React Native JavaScript негізінде жұмыс істейді.', NOW(), NOW()),
('e50c5eff-5a16-4eea-8764-0b0696970e11', '7fd5573a-1d73-409a-bae6-2a3e20cfbd56', 'd850d3ae-3875-4d62-a205-e57031c2a341', 'Android үшін Java/Kotlin, ал iOS үшін Swift/Objective-C пайдаланылады.', NOW(), NOW()),

-- Game Development
('5e583c0d-f7e6-42d2-bf22-7837b1404691', 'bc7201c7-91d2-42cb-90a2-1a58b14bcf21', '082d9f06-c8c9-4977-9d3e-52fb16a57667', 'Unity мобильді және 2D ойындар үшін жақсы, ал Unreal AAA деңгейіндегі ойындарға арналған.', NOW(), NOW()),
('0a1befc8-f21a-4ada-8d23-a7d7ce4ffa21', '81e157cf-c04e-44c8-9279-8d6691942409', 'f6f55eba-e94c-47e5-b30a-0aa4cfc477d3', 'Ойын физикасын жүзеге асыру үшін PhysX немесе Havok секілді қозғалтқыштар қолданылады.', NOW(), NOW()),

-- Database Management
('6bc731ef-17bc-4edc-b7b4-9dfccf4c2dc4', '44168e76-80be-4019-9eea-481dbaf0f8d4', '26ac60f0-8fe5-471f-a489-5902741f51ac', 'SQL құрылымдалған деректер үшін, ал NoSQL икемді сақтау үшін пайдаланылады.', NOW(), NOW()),
('9f4877db-a6fd-4cb8-b899-d8c86f851ba5', '24cf8351-c647-4010-ac00-3f22bf56bc2a', '4712ce26-7052-4a6c-9877-1b745063c94e', 'PostgreSQL күрделі сұраныстар үшін жақсы, ал MySQL жылдам әрі жеңіл жобаларға арналған.', NOW(), NOW());


INSERT INTO users (id, name, last_name, email, password, created_at, updated_at, account_role) VALUES
('74f5e34d-a966-420a-8427-f3abb843ee0f', 'Adil', 'Zhumashev', 'adil.zhumashev@mail.com', '$2a$10$ABC12345xyz', NOW(), NOW(), 'user'),
('8b864659-a44a-41a6-8bfe-0821d5dbd0dc', 'Aigul', 'Serikbayeva', 'aigul.serikbayeva@mail.com', '$2a$10$DEF67890uvw', NOW(), NOW(), 'user'),
('4ad4a8f2-5743-4c51-ba05-79147ed01a14', 'Bauyrzhan', 'Myrzaliev', 'bauyrzhan.myrzaliev@mail.com', '$2a$10$GHI13579rst', NOW(), NOW(), 'user'),
('a0f389a1-f340-42da-9d8f-e54bbb01a83d', 'Erlan', 'Sagyndykov', 'erlan.sagyndykov@mail.com', '$2a$10$JKL24680mno', NOW(), NOW(), 'user'),
('8a9508d6-17ff-4582-b108-d5f03d906f3a', 'Alina', 'Tulegenova', 'alina.tulegenova@mail.com', '$2a$10$MNO97531xyz', NOW(), NOW(), 'user'),
('b2026668-3ed4-49cf-adde-89eba5f827e4', 'Damir', 'Alimkhanov', 'damir.alimkhanov@mail.com', '$2a$10$PQR86420uvw', NOW(), NOW(), 'user'),
('f17b3266-2cfc-4973-9b22-b058582df736', 'Samat', 'Orazbayev', 'samat.orazbayev@mail.com', '$2a$10$STU15937rst', NOW(), NOW(), 'user'),
('5b268d42-dc4a-42fb-ad63-8d0c1422ae59', 'Aliya', 'Ramazanova', 'aliya.ramazanova@mail.com', '$2a$10$VWX26849mno', NOW(), NOW(), 'user'),
('da99333d-5f39-4a38-8bdc-1375d3092cb5', 'Nursultan', 'Bekmuratov', 'nursultan.bekmuratov@mail.com', '$2a$10$YZA35768xyz', NOW(), NOW(), 'user'),
('5236b356-aad4-45ea-91d5-023a5917e148', 'Aruzhan', 'Beisekova', 'aruzhan.beisekova@mail.com', '$2a$10$BCD12345uvw', NOW(), NOW(), 'user'),
('c3e99151-af51-4ce4-9d30-fcd976162ed7', 'Madina', 'Yesenzhanova', 'madina.yesenzhanova@mail.com', '$2a$10$EFG67890rst', NOW(), NOW(), 'user'),
('62b61f53-3d68-481c-9894-767e34b9d831', 'Zhandos', 'Togzhanov', 'zhandos.togzhanov@mail.com', '$2a$10$HIJ13579mno', NOW(), NOW(), 'user'),
('53b7c283-5ca4-4f70-8dc2-fd0c69479d12', 'Dana', 'Tuleshova', 'dana.tuleshova@mail.com', '$2a$10$KLM24680xyz', NOW(), NOW(), 'user'),
('58e83dd0-8806-434f-8c49-86fd028129e3', 'Zhanbolat', 'Serikov', 'zhanbolat.serikov@mail.com', '$2a$10$NOP97531uvw', NOW(), NOW(), 'user'),
('bce6d0e5-2d07-4a61-8d24-7b67840eac24', 'Gulnaz', 'Utegenova', 'gulnaz.utegenova@mail.com', '$2a$10$QRS86420rst', NOW(), NOW(), 'user'),
('d873c134-9762-46e1-969a-32b5b9387ea1', 'Aibek', 'Karabayev', 'aibek.karabayev@mail.com', '$2a$10$TUV15937mno', NOW(), NOW(), 'user'),
('a558bfdd-6cee-445f-b7fd-a2215c4e1c6a', 'Ruslan', 'Talgatov', 'ruslan.talgatov@mail.com', '$2a$10$WXY26849xyz', NOW(), NOW(), 'user'),
('d850d3ae-3875-4d62-a205-e57031c2a341', 'Kamila', 'Zhandarova', 'kamila.zhandarova@mail.com', '$2a$10$ZAB35768uvw', NOW(), NOW(), 'user'),
('082d9f06-c8c9-4977-9d3e-52fb16a57667', 'Miras', 'Sagynov', 'miras.sagynov@mail.com', '$2a$10$CDE12345rst', NOW(), NOW(), 'user'),
('f6f55eba-e94c-47e5-b30a-0aa4cfc477d3', 'Serik', 'Bekbolatov', 'serik.bekbolatov@mail.com', '$2a$10$FGH67890mno', NOW(), NOW(), 'user'),
('26ac60f0-8fe5-471f-a489-5902741f51ac', 'Aruzhan', 'Saparova', 'aruzhan.saparova@mail.com', '$2a$10$IJK13579xyz', NOW(), NOW(), 'user'),
('4712ce26-7052-4a6c-9877-1b745063c94e', 'Amina', 'Tuleshova', 'amina.tuleshova@mail.com', '$2a$10$LMN24680uvw', NOW(), NOW(), 'user');


INSERT INTO users (id, name, last_name, email, password, created_at, updated_at, account_role) VALUES
('f315df8e-b8cd-4c3e-b2eb-c287acaf8c31', 'Dias', 'Nurzhanov', 'dias.nurzhanov@mail.com', '$2a$10$XyZT8aF1L93J8A1KwZUPuOTIVb1oPZlUZrR3FG9d7zFyyMKo1P5Qe', NOW(), NOW(), 'user'),
('0a49a92a-694c-429f-b77a-6994ff4441e0', 'Kamila', 'Omarova', 'kamila.omarova@mail.com', '$2a$10$P6VvJ6/91KQYf9L2FlpdZuqoxIQ5/p13AqHEgPfbMRZsGR9RXN8GC', NOW(), NOW(), 'user'),
('639fd406-6af2-4e50-a090-0d92d428749d', 'Rustem', 'Alpysov', 'rustem.alpysov@mail.com', '$2a$10$ePoOeC6VzIZRpLfeg2ft/u1tPpFb/kbYXe7m9sErT6S1XyY35OuG6', NOW(), NOW(), 'user'),
('b9cf3d64-2b19-4d02-812f-cbebbb437bbc', 'Madina', 'Bekzhanova', 'madina.bekzhanova@mail.com', '$2a$10$zJ72eQdZdUfwW3RsOZptm.r5buvP/PXMGpDeHpI7S7UjLKQ7fCdmW', NOW(), NOW(), 'user'),
('2de42bcb-a2ea-40cf-8d86-6498c09b3d1f', 'Sanzhar', 'Tursynbayev', 'sanzhar.tursynbayev@mail.com', '$2a$10$DE3G7T5V/hJvP5Q./MllhWflFwGrECq9v6m0spUETADuE9CU3m0qC', NOW(), NOW(), 'user'),
('1dbf9a0e-aa30-4002-ac80-6bb6d8e19495', 'Laura', 'Karimova', 'laura.karimova@mail.com', '$2a$10$PCjLfdf8ae4AB6ZqG9hBVuoU3izQFQ/hYOA4tcX2g7p2XhTS.B1Qq', NOW(), NOW(), 'user'),
('ec9f58d5-82d5-442e-a664-ad21f2d0c17b', 'Zhanbolat', 'Ryskulov', 'zhanbolat.ryskulov@mail.com', '$2a$10$8z/PoOTXbHdkeGyOC6tUGeX33IUp69dw82G9VvbhqHGJx9nnl1y96', NOW(), NOW(), 'user'),
('3c5456ff-eef4-4c30-9f37-e52566fc7032', 'Aliya', 'Tulegenova', 'aliya.tulegenova2@mail.com', '$2a$10$NjF7iE3VkgY3UdF/k/TNQu96YZsDO7XHzF1E9kq4IHyCZ9kOkKJQG', NOW(), NOW(), 'user'),
('13e42502-d51a-42c8-a8a8-b92e1104bdd4', 'Nursultan', 'Bektayev', 'nursultan.bektayev2@mail.com', '$2a$10$V.Ro9dtR/Va/FJGzHHi/MuPbw3bnMNyTXzOYKNrQZsH6sXTLiEFJq', NOW(), NOW(), 'user'),
('964b13c7-3798-4551-8f5e-db7c181e8320', 'Samat', 'Esenov', 'samat.esenov2@mail.com', '$2a$10$A96LQoFQ7Gk0v0PxWGP/sTg1fHNkAP10XVrAcR0Ry9OyF9r0HheF6', NOW(), NOW(), 'user'),
('106ddee2-e2a9-4035-921f-bd5cbe7a628e', 'Aruzhan', 'Amankeldi', 'aruzhan.amankeldi2@mail.com', '$2a$10$67LudOeQH3XPi/T34P3TZeVh6lqNOYO06d41xDD/WwZVcAq8zy9Bi', NOW(), NOW(), 'user'),
('c992013c-5cde-48cf-8af0-00184338baf9', 'Yerlan', 'Bekturov', 'yerlan.bekturov2@mail.com', '$2a$10$3DTSkJLV/zs/n/JD9Vj5r9J1Dkp76xv1cQ8aVFYBTeFvTVAfnVuV6', NOW(), NOW(), 'user'),
('9fa1f519-d805-42c7-b9b1-a752365d1c56', 'Marat', 'Bekmuratov', 'marat.bekmuratov2@mail.com', '$2a$10$xQMT9T0KpdM/Rd7BnvNfOQpkZmHyXAVq8H5xPe28ldP0cRQT6UWxS', NOW(), NOW(), 'user'),
('d131a22f-0b3d-4767-9806-6a277dc0744d', 'Dana', 'Tuleshova', 'dana.tuleshova2@mail.com', '$2a$10$yJ63ImfI0pSp5VpMncZk8nE5zyD2iK1WJ7vPPUuXzgfVRMlDl6J62', NOW(), NOW(), 'user'),
('30bd5baf-d603-4b23-936b-c3633df709fe', 'Alisher', 'Zhanabayev', 'alisher.zhanabayev2@mail.com', '$2a$10$5NL6WvlKq/RFGl1FQvS18.eCYy/rhVbg1MvnZxzzPImwP8AlyJd5a', NOW(), NOW(), 'user'),
('636e0dc6-ce94-4f26-bd39-b8fb32b4b378', 'Kamila', 'Sadyk', 'kamila.sadyk2@mail.com', '$2a$10$h9HDebQzCdt/ZD2KLU9G2u7PbvVhKzLFLW8cA4pPQMQVqFBw/VF.G', NOW(), NOW(), 'user'),
('cb2d6442-ec58-4fb6-aea6-23dcd2a6b574', 'Nurbek', 'Tulegenov', 'nurbek.tulegenov2@mail.com', '$2a$10$4azRWtZKquc3FaLp0Tk/FjShuJde0E6tFQ4/zJ/KD5S68J9kOHffCa', NOW(), NOW(), 'user'),
('13d450c5-6e2c-4bd8-bc5a-9096f59497b2', 'Amina', 'Ramazanova', 'amina.ramazanova2@mail.com', '$2a$10$wCD2/PQXNk8q7G6jP0KaMC7eN12LoYIuS88MbXM1UhZtI5nMl2C1a', NOW(), NOW(), 'user'),
('a58335b7-8e7c-4276-8c15-e4e16ca894ec', 'Miras', 'Togzhanov', 'miras.togzhanov2@mail.com', '$2a$10$DUO6B6Z/ltHcBaWoXFn84U3deV6ohB2xWy34A3ZRfbcVll8CpXdfm', NOW(), NOW(), 'user'),
('65148bce-f576-4e3a-b30b-4951cc1e3621', 'Serik', 'Bekbolatov', 'serik.bekbolatov2@mail.com', '$2a$10$7eOxj2oI.pCf7lSMYIT57mRQ2X0XHlfK4l.XF9W0LzPb8AzVh04lC', NOW(), NOW(), 'user'),
('cfd3b85c-2e4a-4fe6-8bfe-4eafe6612759', 'Dias', 'Nurtas', 'dias.nurtas@mail.com', '$2a$10$ABCD12345xyz', NOW(), NOW(), 'user'),
('8d44290e-c670-42ee-bbc1-bf5917dd80de', 'Aruzhan', 'Omarova', 'aruzhan.omarova@mail.com', '$2a$10$EFGH67890uvw', NOW(), NOW(), 'user');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM faq_categories;
DELETE FROM faq_questions;
DELETE FROM faq_answers;
-- +goose StatementEnd
