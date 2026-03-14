# Log-Linter
Линтер для проверки стиля логов в коде. Позволяет выявлять нарушения стиля логов, обеспечивая единообразие логов в проекте.

## Особенности
Проверка стиля логов на соответствие заданным правилам:

Логи должны начинаться со строчной буквы.

Логи должны быть только на английском языке (латиница).

Логи не должны содержать спецсимволы и эмодзи.

Логи не должны содержать чувствительные данные (пароли, токены, ключи API и т.д.).
## Поддерживаемые логгеры:
log/slog

go.uber.org/zap

## Установка как плагин для golangci-lint
1. Создайте файл `.custom-gcl.yml`.
2. Добавьте в него следующее содержимое:
```
version: XXX
plugins:
  - module: "linter"
    path: "plugin"
```
3. Соберите бинарник golangci-lint:
```
golangci-lint custom
```
4. В конфигурационном файле .golangci.yml укажите:
```
linters-settings:
  custom:
    gochecklog:
      type: "plugin"
linters:
  - enable:
    - gochecklog
```
5. Запустите проверку:
```
./custom-gcl run --config=.golangci.yml ./...
```
## Примеры
Примеры можно найти в [testdata](https://github.com/sollvey/Log-Linter/tree/main/testdata/src)
