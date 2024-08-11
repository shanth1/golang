#### Запуск в фоновом режиме

```sh
go run main.go &
```

#### Просмотр фоновых процессов

```sh
ps aux | grep main.go
```

или (для игнорирования grep процесса)

```sh
ps aux | grep '[m]ain.go'
```

#### Просмотр процессов через порт
```sh
sudo lsof -i :PORT_NUMBER
```

#### Остановка процессов

Стандартная:
```sh
kill PID
```

Принудительная
```sh
kill -9 PID
```

По файлу
```sh
pkill -f main.go
```
