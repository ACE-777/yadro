# запуск
Запуск осуществляется из корня репозитория
```
cd .\cmd\yadro\ | go run main.go testFiles\test_file_1.txt
```
# docker build
Собрать контейнер:
```
docker build . -t yadro 
```
Проверить, что контейнер собрался:
```
docker images  
```
Сгененрировать docker контейнер:
```
docker run --name yadro  yadro 
```
