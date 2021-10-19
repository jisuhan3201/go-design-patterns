# Mediator Pattern
---
* in out system 에서 사용하면 좋다. ex. chat room, Player in MMORPG
* No direct reference (포인터가 죽어도 서로 영향이 없음)
* Communication between other component without being aware of each other or having direct access to each other
* Chat room (chatroom.go)
    * Mediator == ChatRoom struct
        * Engage bidrectional communication with its connected component
    * Component == Person struct
        * Have methods the mediator can call
    * Person struct 는 서로의 포인터에 접근하지 않는다.
