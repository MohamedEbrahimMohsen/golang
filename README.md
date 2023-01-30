# golang
This repository will include multiple golang examples.

### Proxy
A proxy which redirect every request from server localhost:8080 to another server localhost:9090

### QRCode
Generate QR Code for a URL.

```
1- go run main.go
2- send a POST request "localhost:8000/qr"
3- add payload to the request body as following:
{
    "url" : "https://www.google.com"
}

-> response should be like that:
{
    "image": "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAABW0lEQVR42uyYMc7sIAyEHVG45AgchavlaByFI1C6QMyTgexLVtr0P2aKCEVfZdlmGNra2tp6E7qqEyIuRD6H8edcC2j6cdVB5o+Q9RvNAYeWqfazR1EgAMkq4IQhvpB1YBYKyIsCYy6IiIV8+Tk4qwOfPcnCxZeXRbo2MDUBIP+4HVcH2oHmgMrCGIXKFFPEuRYAoI0TS18RSlBMZAwgOppDdWMais9zQVgD2nXQjgFKQKaIZA44tGFcJYZ0m6RrMsXbnlwDuEwxhGXYA2gVojXggBolVBrXok5FSPeGMAI03Q/VVRbq9kArlZ4+ygSglaLeL/o68Li0GjBHQ5EeDpD2Q4rWgMsV479NCo89aQS4QpLuitUlaaEo2gPmq5m7S6IxOelmF5cCeh0UGOEAcBoFdC50MALy17VoBPiEJDMi7cC3bV4feGTmIxx8CdX/LLC1tWVR/wIAAP//NsU3kT5oyjoAAAAASUVORK5CYII="
}
```
