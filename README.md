![gopher](./img/GoStudyKLD.jpg "gopher in KLD")

`Курс по Go организованный и проведенный: GDG Kalinigrad, Innoseti, AvitoTech  13.05.2020 - 24.06.2020`

# Go для начинающих
Этот курс основан на [книге](http://golang-book.com/) и её [русской](http://golang-book.ru/) версии, а также других источниках.

## Для кого курс
Курс предназначен для тех, кто хочет познакомиться с языком Go. Уровень не важен. Подходит как для тех кто имеет большой опыт и просто хочет познакомиться с языком, так и для тех кто только начинает погружаться в мир программирования.

# План курса
[План](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v1/plan.pdf)

## Занятие 1
#### Содержание
  * История языка
  * Подготовка к работе
  * Ваша первая программа 
#### Материалы
[Слайды урока](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v1/l1.pdf)  
[Практика](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v1/main.go)  
[Ссылка на домашнее задание](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v1/t1.pdf)  
[Видео](https://www.youtube.com/watch?v=0mdBIpRqJvU) [урока](https://www.youtube.com/watch?v=UdJ-kYcyUZk)   

## Занятие 2
#### Содержание
  * Типы данных, конвертация типов 
  * Переменные, константы, арифметические операции 
  * Области видимости 
#### Материалы
[Слайды урока](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v2/l2.pdf)  
[Практика](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v2/main.go)  
[Ссылка на домашнее задание часть 1](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v2/t2.pdf)  
[Ссылка на домашнее задание часть 2](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v2/t2_1.pdf)  
[Видео](https://www.youtube.com/watch?v=amrNZtYdgTQ) [урока](https://www.youtube.com/watch?v=a7lulyfl0Ms)   

## Занятие 3
#### Содержание
  * Управляющие конструкции (условия, конструкции, циклы)
  * Срезы, массивы, мапы 
#### Материалы
[Слайды урока](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v3/l3.pdf)  
[Практика](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v3/main.go)  
[Ссылка на домашнее задание](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v3/t3.pdf)  
[Видео](https://www.youtube.com/watch?v=FOzoPWC7PLU) [урока](https://www.youtube.com/watch?v=h-ycr3D_fb0)   

## Занятие 4
#### Содержание
  * Функции
  * Указатели 
#### Материалы
[Слайды урока](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v4/l4.pdf)  
[Практика](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v4/main.go)  
[Ссылка на домашнее задание](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v4/t4.pdf)  
[Видео](https://www.youtube.com/watch?v=q7Dfd93zTNA) [урока](https://www.youtube.com/watch?v=tU1XlijY3T0)   

## Занятие 5
#### Содержание
  * Структуры и интерфейсы
  * Конкурирование или многопоточность? 
#### Материалы
[Слайды урока](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v5/l5.pdf)  
[Практика](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v5/src.zip)  
[Ссылка на домашнее задание](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v5/t5.pdf)  
[Видео](https://www.youtube.com/watch?v=FG45mfQi2uI) [урока](https://www.youtube.com/watch?v=tvtJrUwmON0)  
[*Никита Кондратьев* - Tips & Trics](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v5/l5_1.pdf)

#### Tips & Tricks. Структуры и интерфейсы
**Никита Кондратьев (Разработчик, юнит SafeDeal, Авито)**

1. Принимайте интерфейсы, возвращайте структуры.
2. В Go отсутствует наследование как в ООП. Используйте интерфейсы.
3. Не злоупотребляйте interface{}.
4. Не перегружайте интерфейсы.
5. Если структура нужна только для того, чтобы реализовать интерфейс, не экспортируйте её.
6. Помните, что nill удовлетворяет любому интерфейсу.
7. Если нужно изменить данные структуры в методе, определите метод на указателе на структуру.
8. При объявлении нового типа на основе существующего, вы не наследуете методы. Для этого нужно использовать "встраивание" типов.
9. Держите в голове, что encoding/json может работать только с "публичными" полями структур.
10. Учитывайте, что тип и указатель на тип имеют разные наборы методов.

## Занятие 6
#### Содержание
  * Пакеты и повторное использование кода 
  * Документирование кода 
  * Обработка ошибок 
#### Материалы
[Слайды урока](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v6/l6.pdf)  
[Практика](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v6/src.zip)  
[Ссылка на домашнее задание](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v6/t6.pdf)  
[Видео](https://www.youtube.com/watch?v=t9tpBCF-Nvc) [урока](https://www.youtube.com/watch?v=zDXjAMrzAwU)  
[*Вячеслав Бобик* - Tips & Tricks](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v6/l6_1.pdf)

#### Tips & Tricks. Пакеты и повторное использование кода
**Вячеслав Бобик (Разработчик, юнит Messenger, Авито)**  
Именование пакетов
 - В чем смысл пакета
 - Примеры из стандартной библиотеки
 - ? Примеры реального проекта(как надо, как НЕ надо)
 - Несколько  best practices по именованию пакета
Организация пакетов
 - Подходы в организации пакетов
   - Монолит
   - Рельсы/Джанга/
   - Разбиение по модулям
   - Удобный вариант где расскажу про группировку основанную на доменных типах(а так же что это такое) и сервисах.
   - Где черпать вдохновение: пример на репозиторий go-package-layout
   - Луше маленькая копипаста, чем зависимость.


## Занятие 7
#### Содержание
  * Тестирование
  * Стандартная библиотека
  * Следующие шаги 
#### Материалы
[Слайды урока](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v7/l7.pdf)  
[Практика](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v7/src.zip)  
[Ссылка на домашнее задание](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v7/t7.pdf)  
[Видео](https://www.youtube.com/watch?v=uEVMwra0TrM) [урока](https://www.youtube.com/watch?v=fJI7WdDcn_A)  
[*Сергей Иваненко* - Tips & Tricks](https://github.com/MaJloe3Jlo/go_begginers_course/releases/download/v7/l7_1.pdf)


#### Tips & Tricks. Unit тестирование в Go
**Сергей Иваненко (Support Systems developer, Avito)**  
Рассмотрены будут примеры unit тестинга простых ф-ий, файлов, http клиентов, sql запросов. Будет проиллиюстированно как стороннипакеты(facker, mockery, spf13/afero) помогают в этом вопросе, поговорим о внедрении зависимостей и как это помогает при тестировании.  
Встроенные инструменты code coverace, benchmark.

План:
 - unit testing табличне тесты, просто тесты, приватные функции    Примеры File, websocket, httpClient, etc
 - вспомогательные инстурменты (facker, mockery)
 - TDD
 - Слоеная архитектура (циклическая зависимость) +/- различных подходов размещения unit_tests
 - code coverace
 - code race condition, parrallel
 - benchmark

## Преподаватели
[Дима Клестов (Инносети)](https://t.me/MaJloe_3Jlo)  

## Tips & Tricks
[Вячеслав Бобик (Avito)](https://t.me/victor3)  
[Никита Кондратьев (Avito)](https://t.me/nskondratev)    
[Сергей Иваненко (Avito)](https://t.me/SergeyWh1te)  

## Организация
[Мария Круглова (GDG Kaliningrad)](https://t.me/mkruglova)  
[Никита Русин (GDG Kaliningrad)](https://t.me/rusinikita)  
