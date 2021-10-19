# Go-Design-Patterns

---

* Creational Design Patterns
    * Solid
    * Builder Pattern
        * 객체 구조화(construction)이 복잡할 때 유용.
        * 빌더를 sub 빌더로 나누는 방법도 있다.
        * fluent interface -> chaining 을 가능하게 하여 가독성을 높인다.
    * Fatories
        * 생성자를 두는 패턴
    * Prototype
        * 기존 객체에서 새로 객체를 만들 때 유용
        * Deep copy 혹은 serialization을 이용한 copy
        * copy()
    * Singleton
        * 단 하나의 객체 생성만을 보장해야 할 때 (이런 경우는 주로 한 객체가 리소스를 많이 차지한다.)
        * thread safe 해야하고 lazy하게 인스턴스화 (sync.Once()) 하는게 좋다.
        * 추출(extracting interface)과 의존성 주입(dependency injection에 주의하라.

---

* Structural Design Patterns
    * Adapter
        * 기존 인터페이스를 가져와 새로운 인터페이스로 변환
    * Bridge
        * 추상화와 구현을 분리한다(Decouple) -> 결합도를 낮춘다.
    * Composite
        * 비슷한 behavior(method)를 가진 각각의 객체들을 일관되게 client에게 제공할때
    * Decorator
        * 객체에 추가적인 책임(resposibilities)을 붙일 때 (open close principal을 유지하며)
        * 기존 객체를 변경하기 보다 Aggregate 객체를 만들어 새로운 behavior(method)와 기존 객체를 포함한다.
        * embedding or pointers
    * Facade
        * 복잡한 시스템을 숨기고 간결한 인터페이스만 제공한다.
    * Flyweight
        * 많은 수의 비슷한 객체를 중복없이 효율적으로 메모리상에 올린다.
    * Proxy
        * Decorator 패턴과 유사하며 기존 객체를 통합(aggregate) 하기보다 복사(copy)하는 식이다.

---

* Behavioral Design Patterns
    * Chain of Responsibility
        * 여러 컴포넌트들이 정보나 이벤트를 Chaining하여 처리해야할 때 유용하다.
        * 체인 안의 각 element들은 next element를 참조한다(주로 포인터로).
    * Command
        * 하나의 요구(request)를 여러 객체로 캡슐화한다.
    * Interpreter
        * textual input을 structure로 변환할때 유용.
        * interpreter, compiler 등을 만들 때 잘 사용된다.
    * Iterator
        * Aggregate object에 접근하여 element를 navigate 하기 위한 인터페이스 제공시에 유용하다.
        * Visitor 와 유사하다.
    * Mediator
        * 여러 객체간 중재자 역할이 필요할 때 유용하다.
    * Memento
        * 직접적인 조작을 하지 않지만 시스템의 상태를 나타내야할 때 유용하다.
        * 객체 생성시 Memento가 같이 생성되는 경우가 많아 메모리 관리에 취약할 수 있다.
    * Observer
        * 한 객체가 발생 시킨 이벤트(event)를 다른 객체가 듣고(listen) 알려줘야(notify) 할 때
    * State
        * 상태를 가지고 변경되어야 될 경우들
    * Strategy & Template Method
        * 두가지는 비슷한 패턴
        * 알고리즘의 뼈대를 만들고 구체적인 부분을 구현으로 채우는 경우
    * Visitor
        * 계층적인 구조에서 모든 element들에 접근할 때