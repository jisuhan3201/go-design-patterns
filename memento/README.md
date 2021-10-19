# Memento Pattern
---
* State를 가지는 시스템, 롤백을 가능하게 해야하는 시스템
* Object or system goes through changes
* Record every change (like command and teach a command to 'undo' itself)
* Memento read only object -> No change methods
* 객체 생성시 memento 가 같이 생기기 때문에 시스템이 복잡해질수록 메모리 문제에 취약하다.
* Memento VS Flyweight
    * Memento 
        * feedback into system, No public / mutable state
        * No public/mutable state
        * No methods
    * Flyweight
        * Can mutable state
        * Can provide additional functionality (fields / methods)