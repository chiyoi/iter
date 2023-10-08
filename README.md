# iter
Iterator based on generic.
- Interface.
  ```go
  Iterator[T].Next() (T, bool)
  ```
  The second return value indicates whether iteration is stopped.
  ```go
  Continue = true
  Break    = false
  ```
- Make an iterator.
  - `IteratorFunc`, function as an iterator.
  - `Iter` on a slice.
  - Iterate over a `Range` with step.
  - `Empty` iterator stops immediately.
  - `Repeat` a certain value forever.
  - `Chain` two iterators sequentially.
  - `Zip` two iterators to iterate them simultaneously.
- Map an iterator.
  - `Map` with mapping function.
  - `Scan` with initial state and mapping function.
- Take part of an iterator.
  - `Take` some elements.
  - `Skip` some elements.
  - `Filter` with prediction function.
- Consume an iterator.
  - Get the `Last` element.
  - Get the element `At` the specified position.
  - `Reduce` with reducing function.
  - `Fold` with initial state and reducing function.
  - `Collect` an iterator into a slice.
  - Get the `Min`, or `Max` element for an iterator with ordered type, or calculate the `Sum`.
