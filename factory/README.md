# Factory Pattern
---
## Overview
* 객체 생성 로직이 너무 복잡할 경우, 한 객체가 많은 field들을 가지고 있을 경우
* 객체에 Default 값을 주고 필요한 값들만 구현하고 싶은 경우
* Builder 패턴과 다른 점은 한번에 생성하느냐 (Factory) 아니면 단계적으로 생성하느냐 (Builder)
---
## Interface factory (interfacefactory.go)
* 객체를 생성 후 인터페이스를 리턴하여 struct private으로 유지하고 캡슐화 한다.
* 인터페이스를 리턴함으로써 생성자(New~)의 인자에 따라 여러 객체를 리턴 할 수 있다.

## Factory generator (factorygenerator.go)
* Factory pattern을 생성하는 두가지 접근법(functional, structural)

## Prototype factory (prototypefactory.go)
* 고정된 prototype이 있는 경우 factory pattern 구현
---
## Summary
* Facotry pattern은 객체 초기화시 유연성을 제공한다.
    * Default 값
    * Validation
    * 조건에 따른 여러 객체 리턴