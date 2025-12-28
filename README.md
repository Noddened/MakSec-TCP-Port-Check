# TCP Port Checker

## Сборка

```sh
make build
#или сразу с запуском
make run
```

## Запуск с параметром

```sh
./build/tcp-port-checker -config ./configs/config.toml
```

`-config` — путь к конфигурационному файлу (по умолчанию: `./configs/config.toml`)

## Выполнение запроса к прил

---

### Требования
- Go 1.18+
- Make

## Демонстрация работы

### Проверка IP в локальной сети (порт открыт)
![Локальный IP — порт открыт](images/local_ip_open.png)

### Проверка внешнего IP
![Внешний IP — порт открыт](images/external_ip_open.png)

![Внешний IP — порт закрыт](images/external_ip_closed.png)

### Логи роутера в терминале
![Пример вывода в терминале](images/terminal_output.png)


### Обработка неправильного запроса
Если пользователь отправляет некорректный запрос, то сервер возвращает ошибку с пояснением:

![Некорректный запрос](images/bad_request.png)
