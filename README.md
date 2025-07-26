# Проект для паркинга CSV



------------------------------------------------------------
# Примеры:

## Примеры 1:

Входные данные:
```bash
go run cmd/main.go -file ./test.csv \
   -header=false  -validate-type "col_3:int" -filter "col_3>=30" -sort "col_3:asc" -verbose 
```
Выходные данные:
```bash
⚠️ Запись 1: поле 'col_3' должно быть целым числом
⚠️ Запись 2 не соответствует фильтру
Запись 1:
  col_1: Alice
  col_2: alice@example.com
  col_3: 30
Запись 2:
  col_1: Charlie
  col_2: charlie@example.com
  col_3: 40
Запись 3:
  col_2: charlie2@example.com
  col_3: 45
  col_1: Charlie2222
```


## Примеры 2:

Входные данные:
```bash
go run cmd/main.go -file ./test.csv \
   -header=true -required "Name,Age" -range "Age:18-70" -validate-type "Age:int" -filter "Age>=30" -sort "Age:asc" -verbose 
```
Выходные данные:
```bash
⚠️ Запись 2 не соответствует фильтру
Запись 1:
  Name: Alice
  Email: alice@example.com
  Age: 30
Запись 2:
  Name: Charlie
  Email: charlie@example.com
  Age: 40
Запись 3:
  Name: Charlie2222
  Email: charlie2@example.com
  Age: 45
```

------------------------------------------------------------