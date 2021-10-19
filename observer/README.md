# Observer Pattern
---
* Object 가 어떤 행위 (method)를 했을 때 알려줌
* 혹은 시스템 내에 어떤 이벤트 발생시 알리기 위함
* Observable (관찰해야할 객체), Observer (관찰을 하는 객체)
* Must provide a way of clients to subcribe
* Event data sent from observable to all observers
* Data represented as interface{}