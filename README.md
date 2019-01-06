# susun [![CircleCI](https://circleci.com/gh/faizhasim/susunplease.svg?style=svg)](https://circleci.com/gh/faizhasim/susunplease)

## Quickstart

```bash
go get -u github.com/faizhasim/susunplease/cmd/susun
cat $(susun rules showpath)
# vim $(susun rules showpath)
# ## -or-
# susun rules open

susun process '/tmp/inbox/***' /tmp/outbox
```
