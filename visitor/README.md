# Visitor Pattern
---
* 컴포넌트(visitor)가 모든 계층화된 타입들을 traverse(visit) 하는 패턴
* Step
    1. Propagate an Accept(v *Visitor) method throughout the entire hierarchy
    2. Create a visitor VisitFoo(f Foo), VisitBar(b Bar) ... for each element
    3. Each Accept calls Visitor.VisitXXX(self)
