# iter
Iterator based on empty interface.
- Make an iterator:
    - `Iter(slice)`
    - `From(func() (item any))`
    - `Range(from, to int)`
- Map an iterator:
    - `Map(func(a any) (b any))`
    - `Scan(func(state any, v any) (nextState any))`
- Take part of an iterator:
    - `Filter(prediction func(v any) bool)`
    - `Take(count int)`
    - `Skip(count int)`
- Consume an iterator:
    - `Reduce(func(a, b any) (result any))`
    - `Fold(func(state any, v any) -> (nextState any))`
    - `Last()`
    - `At(position int)`
- returning `nil` for short-circuit, except for map and reduce.
