# Woka Server

<img src="https://avatars.githubusercontent.com/u/131611724?v=4" width="64px" height="64px" alt="Losti profile avatar" />

##### Developed By - [**Lostish**](https://github.com/lostish)

#### English

Welcome to the woka project server.

This is an API built with Gin, SurrealDB, MinIO, and a local CDN with MinIO.

#### Spanish

Bienvenido al servidor del proyecto [woka](https://github.com/lostish/woka-site/).

Esta es una API hecha con Gin, SurrealDB, MinIO
y CDN local con MinIO.

## Purpose

#### English

We want to create a server decoupled from the front-end, giving the project a single purpose.

> "An easy-to-scale and debug API, with direct and secure connections to the database and the local CDN"

#### Spanish

Queremos crear un servidor desacoplado del front-end dando un solo propisito al proyecto.

> "Una API facil de escalar y depurar, con conexiones directas y seguras a la base de datos como al CDN local"

## Structure

```bash
└── 📁src
    └── 📁cmd
        └── 📁add-services-here
        └── 📁server
            └── **/*.go
    └── 📁internal
        └── 📁config
            └── **.go
        └── 📁db
            └── **/*.go
        └── 📁handlers
            └── 📁payments
            └── 📁service
            └── 📁user
        └── 📁storage
            └── **/*.go
    └── 📁pkg
        └── 📁auth
        └── 📁cache
        └── 📁events
    └── 📁scripts
        └── **/*.sh
```
