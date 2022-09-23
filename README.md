# gorc: An IRC-like chat app built in Go

```
░█▀▀▀░▄▀▀▄░█▀▀▄░█▀▄░
░█░▀▄░█░░█░█▄▄▀░█░░░
░▀▀▀▀░░▀▀░░▀░▀▀░▀▀▀░
```

An IRC-like chat app built in Go that is small, simple and secure. And yes, it's pronounced "gork".

## Project goals:

- [x] Minimalist TUI interface 👨‍💻
- [x] Single binary for both client and server 🖥️
- [ ] Load config from a TOML file ⚙️
- [ ] Very small memory footprint 🤏
- [ ] Users are anonymous, UID is pseudorandom and gorc logs nothing 🥷
- [ ] Messages in transit are encrypted 🔐
- [ ] Messages are in memory only and purged when the server stops 🧹
- [ ] ｂｌａｚｉｎｇｌｙ ｆａｓｔ 🔥

## How gorc will work:

Launch in client-mode:

`./gorc` or `./gorc client`

Launch in server-mode:

`./gorc server`

List additional options and help:

`./gorc help`

Show app version and info:

`./gorc version`