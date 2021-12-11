# Test Flash Coffee

## Setup explained

### Tooling

-   docker-compose:

    -   `docker-compose up --build` - running dengan perintah build atau rebuild applikasi.
    -   `docker-compose up -d` - running app di background dengan settingan terakhir kali build.
    -   `docker-compose down` - stop dan remove resource.


### Included sample packages

-   -   **app**
    -   [Gin](https://github.com/gin-gonic/gin) untuk route aplikasi.
    -   Listens on http://localhost:8080 .

## dev

kita bisa menjalan langsung semua app sekaligus dengan menggunakan syntax berikut:
```
docker-compose up --build
```

### api collection
untuk api collection menggunakan postman 