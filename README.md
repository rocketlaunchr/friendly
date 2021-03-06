friendly - The FRiENDly webserver for front-end developers (For local development)
======


<p align="center">
<img src="https://github.com/rocketlaunchr/friendly/raw/master/logo.png" alt="friendly" />
</p>


⭐ **the project to show your appreciation.**

When you are developing your website locally, you may be encountering **CORS** issues.
In my case, I had an `iframe` that was calling `parent` to access the host site.
All major browsers such as Chrome, Safari and Firefox were blocking the call.

You can use `friendly` by placing the server in the same directory as your project to
run your website as if it was run on an actual production server.

It's that simple and easy to use.

It supports:
- http and https
- custom ports
- custom paths (so you don't need to place it in the same directory as your project)


**[Download here](https://github.com/rocketlaunchr/friendly/releases)**

## Usage

```bash
./friendly -d "<path>"  -b --save -r -s
```

or just place the application in your project path and run it without flags.

The recommended way is to install it globally (add to $PATH). Then you can run the server from anywhere without setting a `path` (which defaults to the current working directory).

## Flags

### port (p)

Set a custom port. By default, it is `8080` for http and `4430` for https.

### path (d)

Point to the directory of your project.

### https (s)

Automatically create a self-signed SSL certificate. The browser will ask whether you trust the certificate. Allow it.

### browser (b)

Open the project automatically on your default browser the moment the server starts up.

### save

In https mode, everytime the server starts, it will create a new self-signed certificate.
The browser will repetitively ask if you trust the certificate. This can be annoying.
Use this setting to reuse the same certificate.

### remove (r)

Delete a certificate you may have saved in the past.

### quiet (q)

Don't show any logs of the incoming requests.


## Installation

Just download the prebuilt executables from the [Releases](https://github.com/rocketlaunchr/friendly/releases). It is available for **Windows**, **macOS** and **Linux**.

If you want to customize the project to your needs, then clone this repo. You will need to know how to build Go projects after downloading the dependencies.


```bash
GITCOMMIT=$(git rev-parse --short HEAD) && \
VERSION=$(git describe --always) && \
env GOOS=darwin GOARCH=amd64  go build -ldflags "-X main.GITCOMMIT=$GITCOMMIT -X main.VERSION=$VERSION -s -w" .
```

**NOTE:** Replace GOOS with `darwin`(macOS), `windows` or `linux`.



Other useful packages
------------

- [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Statistics and data manipulation
- [dbq](https://github.com/rocketlaunchr/dbq) - Zero boilerplate database operations for Go
- [electron-alert](https://github.com/rocketlaunchr/electron-alert) - SweetAlert2 for Electron Applications
- [igo](https://github.com/rocketlaunchr/igo) - A Go transpiler with cool new syntax such as fordefer (defer for for-loops)
- [mysql-go](https://github.com/rocketlaunchr/mysql-go) - Properly cancel slow MySQL queries
- [react](https://github.com/rocketlaunchr/react) - Build front end applications using Go
- [remember-go](https://github.com/rocketlaunchr/remember-go) - Cache slow database queries


## Legal Information

The license is a modified MIT license. Refer to the `LICENSE` file for more details.

**© 2020 PJ Engineering and Business Solutions Pty. Ltd.**

## Final Notes

Feel free to enhance features by issuing pull-requests.