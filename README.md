Study fiber framework websocket middleware
===
### Make chat using fiber framework middleware websocket
<br/>

How to use
---
run main file (go run main.go)
and open client.html
<br/>

Usage
---
Name|Addr
-|-
fiber|https://github.com/gofiber/fiber
fiber websocket|https://github.com/gofiber/websocket
<br/>

Chat URL
---
- take __http://localhost:3000/ws/:user_id__ and upgrade http to websocket
<br/>


Client.html
---
- client.html make random socket connection url
```js
    new WebSocket("ws://localhost:3000/ws/".concat(Math.round(Math.random()*10000)))
```
- display 
```
    random_number(which is user_id) : message | eg. 2478 : Hello World!
```
