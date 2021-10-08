# Decorator Pattern
---
## Overview
* 객체에 기능을 추가하고 싶을때
* 기존 코드를 변경하거나 다시 작성하고 싶지 않을 때 (OCP)
* 새로운 기능을 분리하고 싶을때 (SRP)
---
## Multiple inheritance problem (multipleinheritanceproblem.go)
* 한 객체에 여러 객체를 상속받는 것은 쉽지 않다.
* 한 객체에 여러 객체를 상속 받을때 생기는 문제점
## Multiple inheritance solution (multipleinheritancesol.go)
* 위 문제의 해결 방법
## Decorator (decorator.go)
* Decorator pattern 구현
---
## Summary
* Decorated object를 embed 하고 새로운 기능을 추가하여 해결