# gorc: An IRC-like chat app built in Go

```text
░█▀▀▀░▄▀▀▄░█▀▀▄░█▀▄░
░█░▀▄░█░░█░█▄▄▀░█░░░
░▀▀▀▀░░▀▀░░▀░▀▀░▀▀▀░
```

An IRC-like chat app built in Go that is small, simple and secure. And yes, it's pronounced "gork". Inspired by the IRC servers of the 90s and the BBS scene of
the 80s, `gorc` aims to recapture the simple yet engaging exerience of chatting
with other anonymouse users online, featuring rooms catering to different
discussions and topics. The app will be small, self-contained and offer both the
ability to serve a gorc instance or join a gorc instance froma single binary.

![Follow the white rabbit.](https://c.tenor.com/IcbRC3jvTkcAAAAC/rabbit-follow-the-white-rabbit.gif)

## Project goals

- [x] Minimalist TUI interface 👨‍💻
- [x] Single binary for both client and server 🖥️
- [ ] Load config from a TOML file ⚙️
- [ ] Very small memory footprint 🤏
- [x] Users are anonymised with randomly generated IDs and gorc logs nothing 🥷
- [ ] Messages in transit are encrypted 🔐
- [ ] Messages are stored in-memory only and gone forever when the server stops 🧹
- [ ] ｂｌａｚｉｎｇｌｙ ｆａｓｔ 🔥

## How gorc will work

Launch in client-mode, optionally specifying a server to connect to:

`./gorc` or `./gorc chat.gorc.fake`

Launch in server-mode, optionally specifying a port to serve on:

`./gorc server` or `./gorc server 1234`

By default, gorc will look for client and server options in `settings.toml`. 

List additional options and help:

`./gorc help`

Show app version:

`./gorc version`
