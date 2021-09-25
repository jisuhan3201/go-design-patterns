### Builder Pattern
---
# Overview
* 많은 인자를 가진 함수로 객체를 만들 경우 생산적이지 않음
* 많은 구성들 혹은 단계적으로 객체를 만들어야 할 경우

# Builder Facets (builderfacets.go)
* 하나의 객체를 여러 양상(facet)으로 만들어야 할때

# Builder Parameter (builderparameter.go)
* API가 어떻게 builder를 참조 할지를 설명 Private struct
* Builder struct 를 Public으로 하여 Builder struct를 참조하게 한다.
* sendMailImpl이라는 private function을 두고 client로 하여금 action 함수 인자를 받는 SendEmail 함수만 사용하도록 한다.

# Functional Builder (functionalbuilder.go)
* Builder pattern을 함수형 프로그래밍으로 접근
---
# Summary
* 하나의 Builder는 한 객체를 만드는 구성품들 중 하나를 만드는데 사용된다.
* Builder를 만들때는 Fluent하게 만들어라, Receiver를 리턴하여 Chaining하게 만들어라
* 여러 양상 (Facet)을 갖는 객체의 경우 여러 개의 builder로 분리하여 구현한다.

