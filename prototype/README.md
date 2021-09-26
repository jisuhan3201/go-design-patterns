# Prototype Pattern
---
## Overview
* 복잡한 객체는 밑그림부터 시작하지 않는다.
* 기존 디자인이 Prototype이고, 복사(DeepCopy) 후 커스터마이즈 하여 사용한다.
---
## Prototype Deep Copy (prototypedeepcopy.go)
* Shallow Copy의 문제점
* Deep Copy의 방법을 설명

## Prototype Copy Method (prototypecopymethod.go)
* 완벽하지 않은 Deep Copy Method 구현

## Prototype Copy Through Serialization (prototypeserialization.go)
* 완벽한 Deep Copy Method 구현

## Prototype Factory (prototypefactory.go)
* Prototype + Factory Pattern을 이용한 Example
---
## Summary
* prototype, 부분적인 객체 구성 -> Deep Copy 구현 -> 객체 Customize -> prototype factory 구현으로 편리한 API 제공