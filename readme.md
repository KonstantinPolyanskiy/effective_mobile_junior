<h1 align="center">
<a href="https://disk.yandex.ru/i/OFSTEypvwZTxqA" target="_blank">Тестовое задание</a> на должность Go Junior Developer
</h1>
<h2>Описание проекта</h2>
<p>Сервис написанный на языке Go, предоставляющий REST API интерфейс для обогащения наиболее вероятной информацией (гендер, страна, пол) о людях по имени.</p>
<p>Использует сторонние сервисы:</p>
<ul>
        <li><a href="https://api.agify.io/?name=Example">Agify</a></li>
        <li><a href="https://api.genderize.io/?name=Example">Genderize</a></li>
        <li><a href="https://api.nationalize.io/?name=Example">Nationalize</a></li>
</ul>
   <h2>Реализован функционал:</h2>

   <ul>
        <li>
            <h3>Добавление персоны:</h3>
            <pre>
                <code>
                    POST localhost:8080/person
                    {
                        "name": "Example Name",
                        "surname": "Example Surname",
                        "patronymic": "Optional Patronymic"
                    }
                </code>
            </pre>
        </li>
        <li>
            <h3>Удаление (мягкое) персоны:</h3>
            <pre>
                <code>
                    DELETE localhost:8080/person?id=1
                </code>
            </pre>
        </li>
        <li>
            <h3>Получение персон/ы:</h3>
            <pre>
                <code>
                    GET localhost:8080/person?name=Konstantin&country=RU&older=20&gender=male&limit=5&offset=0
                </code>
            </pre>
        </li>
        <li>
            <h3>Изменение данных сущности по id:</h3>
            <pre>
                <code>
                    PATCH localhost:8080/person?id=1
                    {
                        "name": "other name",
                        "surname": "other surname",
                        "patronymic": "other patronymic",
                        "age": 99,
                        "gender_name": "female",
                        "country_code": "FR"
                    }
                </code>
            </pre>
        </li>
    </ul><h2>Использованные библиотеки:</h2>
<ul>
        <li><a href="https://github.com/jackc/pgx">pgx</a> - драйвер PostgresSQL</li>
        <li><a href="https://github.com/uber-go/zap">zap</a> - логгирование</li>
        <li><a href="https://github.com/joho/godotenv">godotenv</a> - работа с .ENV файлами</li>
        <li><a href="https://github.com/pressly/goose">goose</a> - работа с миграциями</li>
        <li><a href="https://github.com/gin-gonic/gin">gin</a> - роутер</li>
        <li><a href="https://github.com/Masterminds/squirrel">squirrel</a> - SQL билдер</li>
        <li><a href="https://github.com/vektra/mockery/v2">mockery</a> - генерация моков</li>
</ul>
<ol>
        <li>Клонировать репозиторий:<br>
            <code>https://github.com/KonstantinPolyanskiy/effective_mobile_junior.git</code>
        </li>
        <li>Перейти в папку с проектом:<br>
            <code>cd effective_mobile_junior</code>
        </li>
        <li>Запустить тесты:<br>
            <code>make test</code>
        </li>
        <li>Запустить базу данных:<br>
            <code>docker compose up</code>
        </li>
        <li>Применить миграции:<br>
            <code>goose postgres "postgres://konstantin:publicPassword@localhost:5432/postgres?sslmode=disable" up</code>
        </li>
        <li> Запустить сервис:<br>
            <code>make run</code>
        </li>
</ol>
    