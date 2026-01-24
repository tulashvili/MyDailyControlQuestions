## Описание

Этот проект — небольшое консольное приложение на Go для ежедневной саморефлексии.

Приложение содержит набор вопросов, сгруппированных по категориям (например, Biology, Psychology).
При запуске main.go пользователю по очереди предлагаются эти вопросы.
На каждый вопрос нужно ответить, оценив своё текущее состояние по шкале от 1 до 5 по итогам прошедшего дня.

Цель проекта — помочь зафиксировать субъективное состояние, сформировать привычку регулярной оценки дня и создать основу для дальнейшего анализа динамики.

## Change log

- ~~Добавить дни ответов на вопросы~~ ([MR](https://github.com/tulashvili/MyDailyControlQuestions/commit/afebfc49fb06369f0fcac6386ee9b87399327ffe))
- ~~Добавить проверку на диапазон ответа~~ ([MR](https://github.com/tulashvili/MyDailyControlQuestions/commit/8a169e58148d5a09547f7b5ae9cb491ba0a18585))
- ~~Изменение архитектуры приложения в соответствии с приницами чистой архитектуры (как я их понял)~~ ([MR](https://github.com/tulashvili/MyDailyControlQuestions/commit/f9c80d8b62388cb9dd3c6eb5dbd50bd4a8876e6c))
- ~~Сохранение результатов ответов в базу данных вместо json-файла~~ ([MR](https://github.com/tulashvili/MyDailyControlQuestions/commit/f9c80d8b62388cb9dd3c6eb5dbd50bd4a8876e6c))
- ~~Добавить отображение статистики за несколько дней/недель/месяцев~~ ([MR](https://github.com/tulashvili/MyDailyControlQuestions/commit/8ab1a00f8e538a3402b10f1871017500bec7e11c))
- Добавить возможность импорта данных из EXCEL, который:
    - разбит по неделям
    - каждая неделя - один лист
- ~~Создать makefile и прикрутить линтер~~ ([MR](https://github.com/tulashvili/MyDailyControlQuestions/commit/b381af895f580a186e9eb7a07be5a0b26dc70f4f))
- Добавить возможность настройки бэкапа данных по расписанию
