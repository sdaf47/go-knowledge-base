## Прототип микросервиса (shortlinkserver)
Прототип микросервиса на вход которому поступает URL а он генерирует короткую ссылку.

### Запуск
```bash
make
```
Соберется и запустится контейнер docker-compose с redis и двумя сервисами на портах 8080 и 8000.


### Инструменты
- [go-kit](https://github.com/go-kit/kit) - предоставляет готовую архитектуру для микросервисов
- [gRPC](https://grpc.io/) - во-первых, очень удобно декларировать сущности сразу для всех языков, во-вторых, на мой взгляд, удобнее REST.
- [redis](https://redis.io/) - что может быть удобнее и быстрее key-value хранилища, когда у нас есть только key и value?
- [Docker](https://www.docker.com/) - это же микросервис =)
- [highwayhash](https://github.com/minio/highwayhash) - очень быстро формирует хэш.

Сервис `linkshort` получает методом `Encode()` ссылку, и возвращает код, по коротому она будет доступна.
Методом `Decode()` сервису передается код, после чего он возвращает ссылку.

За доменное имя и интерфейс отвечает сервис `shortlinkweb`

### Web-интерфейс (shortlinkweb)
[Web-интерфейс](https://github.com/sdaf47/go-knowledge-base/blob/master/shortlink/cmd/shortlinkweb) написан на коленке. 
Ведет общение с микросервисом по формированию коротких ссылок посредством gRPC.

### Что не учтено в прототипе?
- Отслеживание количества ссылок, созданных одним пользователем.
- Ссылки также могут содержать и иные данные (дата создания, истечение срока хранения...)
- Валидация ссылок на корректность (а также на отсутствие XSS-атак, например).
- Возможно, есть какие-то **не технические проблемы**, например, фишинг. Для технической стороны это означает, что нужна какая-то база для запрещенных/разрешенных доменов и соответствующие обработки в валидации.
- . . . 