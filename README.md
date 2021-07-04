# telegram-bot-gosample
Пример создания Telegram бота на Go

## Порядок что делать для запуска

1. получить токен своего бота от @BotFather
2. скопировать settings.example.json в settings.json
3. заполнить поле "token" в файле settings.json
4. запустить бота

## как запустить бота

Вариант 1. Выполнить (пример для linux):

```bash
go build -o bot
./bot
```

для windows соответственно:

```bash
go build -o bot.exe
./bot.exe
```

Вариант 2. Быстро, но не рекомендуется

```bash
go run *.go
```
