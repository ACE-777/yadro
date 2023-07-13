# Запуск
Входной файл с данными для программы размещается по следующей директории: ```.\cmd\yadro\testFiles\```.

Запуск осуществляется из корня репозитория, с указанием названия файла в командной строке:
```
cd .\cmd\yadro\ | go run main.go testFiles\test_file_1.txt
```
# docker build
Указать путь до файла, на котором требуется запустить программу с использованием docker образа на 9 строке Dockerfile: ```CMD ["./main", "cmd/yadro/testFiles/test_file_1.txt"]``` 

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
