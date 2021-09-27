# Singleton Pattern
---
## Overview
* 시스템에서 단 한번만 실행되는(혹은 만들어지는) 객체(혹은 컴포넌트) 주로 한번 실행 비용이 비싸다.
    * Database repository
    * Object factory
* 추가적인 복사를 막기 위해
* Lazy instantiation (필요로 하기 전까진 인스턴스화하지 않음)에 주의 해야한다.
---
## Singleton
* 간단한 Singleton 예제
* 테스트시 문제점과 해결 방법
---
## Summary
* sync.Once를 이용한 Lazy Instantiation
* 인터페이스에 의존하라 not concrete type
* 의존성을 줄여라
