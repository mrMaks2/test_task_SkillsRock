# Приложение по управлению задачами

## Описание

Это проект на Go, который реализует API для управления задачами. Он использует библиотеку `fiber` для упрощенной работы с HTTP, а также `pgx` для работы с PostgreSQL.


## Используемые технологии

*   **Golang:** Язык программирования.
*   **Fiber:** HTTP-фреймворк.
*   **PostgreSQL:** БД для хранения задач.
*   **Docker и Docker Compose:** Для контейнеризации.

## Инструкция по запуску приложения

1.  **Убедитесь, что у вас установлен Docker и Docker Compose.**

2.  **Склонируйте репозиторий:**

    ```bash
    git clone https://github.com/mrMaks2/test_task_SkillsRock/
    ```

3.  **Запустите проект с помощью команды:**

    ```bash
    docker-compose up --build
    ```

4.  **Для остановки проекта используйте команду:**

    ```bash
    docker-compose down
    ```


## Переменные окружения

В файле .env имеются следующие переменные (желательно в файле .env заменить на свои, но и без этого будет работать, так как файл не стал добавлять в .gitignore):

*   **POSTGRES_USER:**  Имя пользователя для PostgreSQL.
*   **POSTGRES_PASSWORD:** Пароль пользователя для PostgreSQL.
*   **POSTGRES_PORT:** Порт для PostgreSQL.
*   **POSTGRES_NAME:** Имя базы данных PostgreSQL.
*   **POSTGRES_USE_SSL:** Использовать ли SSL для PostgreSQL.
*   **POSTGRES_HOST:** Хост для PostgreSQL.
*   **CONN_HOST:** Порт для сервиса приложения.

## API

*   **POST /tasks:** Создание новой задачи.
*   **PUT /tasks/:id** Обновить конкретную задачу.
*   **DELETE /tasks/:id** Удалить конкретную задачу.
*   **GET /tasks:** Получить все задачи.