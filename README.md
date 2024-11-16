<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px style="border-radius:50%" src="https://kappa.lol/wfFBr" alt="Project logo"></a>
</p>

<h3 align="center">doodocs-days-backend</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/exoneges/doodocs-days-backend.svg)](https://github.com/exoneges/doodocs-days-backend/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/exoneges/doodocs-days-backend.svg)](https://github.com/exoneges/doodocs-days-backend/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> Few lines describing your project.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
<!-- - [TODO](../TODO.md) -->
<!-- - [Contributing](../CONTRIBUTING.md) -->
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

RESP API for zipping/unzipping files and sending archives via SMTP.

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 
<!-- See [deployment](#deployment) for notes on how to deploy the project on a live system. -->

### Prerequisites

```
go version 1.21+
```

### Installing

A step by step series of examples that tell you how to get a development env running.

Setup the environment variables
```
$env:DOODOCS_DAYS2_BACKEND_AUTH_USERNAME = "USERNAME_FOR_MIDDLEWARE_REQUEST"
$env:DOODOCS_DAYS2_BACKEND_AUTH_PASSWORD = "PASSWORD_FOR_MIDDLEWARE_REQUEST"

$env:DOODOCS_DAYS2_BACKEND_MAIL_USERNAME = "USERNAME_FOR_SMTP_REQUEST"
$env:DOODOCS_DAYS2_BACKEND_MAIL_PASSWORD = "APP_PASSWORD_FOR_SMTP_REQUEST"
```

Install dependences

```
go get gopkg.in/mail.v2  
```

Build the project

```
go build -o server ./cmd/main.go
```

Run it

```
./server.go
```


## 🔧 Running the tests <a name = "tests"></a>

<!-- Explain how to run the automated tests for this system. -->

### Break down into end to end tests

<!--
Explain what these tests test and why

```
Give an example
```
-->

### And coding style tests

<!---
Explain what these tests test and why

```
Give an example
```
-->

## 🎈 Usage <a name="usage"></a>

```
$ ./app --help
Doodocs days-2 backend project

    Usage:
            server [--port <N>] [--dir [S]]
            server --help

    Options:
            --help          Show this screen.
            --porn N        Port number.
            --dir S         Path to the data directory.
```

REST API Endpoints:
- 
Request
```POST /api/archive/information HTTP/1.1
Content-Type: multipart/form-data; boundary=-{some-random-boundary}

-{some-random-boundary}
Content-Disposition: form-data; name="file"; filename="my_archive.zip"
Content-Type: application/zip

{Binary data of ZIP file}
-{some-random-boundary}--
```
Responce

```HTTP/1.1 200 OK
Content-Type: application/json

{
    "filename": "my_archive.zip",
    "archive_size": 4102029.312,
    "total_size": 6836715.52,
    "total_files": 2,
    "files": [
        {
            "file_path": "photo.jpg",
            "size": 2516582.4,
            "mimetype": "image/jpeg"
        },
        {
            "file_path": "directory/document.docx",
            "size": 4320133.12,
            "mimetype": "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
        }
    ]
}
```

Request
```POST /api/archive/files HTTP/1.1
Content-Type: multipart/form-data; boundary=-{some-random-boundary}

-{some-random-boundary}
Content-Disposition: form-data; name="files[]"; filename="document.docx"
Content-Type: application/vnd.openxmlformats-officedocument.wordprocessingml.document

{Binary data of file}
-{some-random-boundary}
Content-Disposition: form-data; name="files[]"; filename="avatar.png"
Content-Type: image/png

{Binary data of file}
-{some-random-boundary}--
```
Responce

```HTTP/1.1 200 OK
Content-Type: application/zip

{Binary data of ZIP file}
```

Request
```POST /api/mail/file HTTP/1.1
Content-Type: multipart/form-data; boundary=-{some-random-boundary}

-{some-random-boundary}
Content-Disposition: form-data; name="file"; filename="document.docx"
Content-Type: application/vnd.openxmlformats-officedocument.wordprocessingml.document

{Binary data of file}
-{some-random-boundary}
Content-Disposition: form-data; name="emails"

elonmusk@x.com,jeffbezos@amazon.com,zuckerberg@meta.com
-{some-random-boundary}--
```
Responce

```
HTTP/1.1 200 OK
```

## 🚀 Deployment <a name = "deployment"></a>

<!-- Add additional notes about how to deploy this on a live system. -->

## ⛏️ Built Using <a name = "built_using"></a>

- [Golang](https://go.dev/) - Language

<!-- - [MongoDB](https://www.mongodb.com/) - Database
- [Express](https://expressjs.com/) - Server Framework
- [VueJs](https://vuejs.org/) - Web Framework
- [NodeJs](https://nodejs.org/en/) - Server Environment -->

## ✍️ Authors <a name = "authors"></a>

- [![Status](https://img.shields.io/badge/github-exoneges-success?logo=github)](https://github.com/exoneges)
- [![Status](https://img.shields.io/badge/alem-igussak-red?logo=github)](https://platform.alem.school/git/igussak)
<a href="https://t.me/undefinedbro" target="_blank"><img src="https://img.shields.io/badge/telegram-@undefinedbro-blue?logo=Telegram" alt="Status" /></a>

<!-- See also the list of [contributors](https://github.com/kylelobo/ -->

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- doodocs
- Alem school
