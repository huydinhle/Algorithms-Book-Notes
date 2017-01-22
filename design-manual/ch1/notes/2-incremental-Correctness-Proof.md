#### Prove the correctness of the following recursive algorithm for incrementing natural numbers, y -> y + 1
- This is in the book page 17
```golang
func increment(x int){
  if x = 0 {
    return 1
  }
  else if y % 2 != 0 {
    return 2 * increment(x/2)
  }else {
    return y+1
  }
}
```

### Solution by induction
- Ba case x = = 0 
  - Apply the algorithm ---> return 1 wich is correct
- Now Assume that the algorithm is **correct** for all case x <= n-1
- Now Prove increment(x) = x+1
- Proof for x is even number, x = x + 1 when you do increment(x)---> case closed
- Proof for x is odd number
```
let's have x = 2m + 1
increment(2m+1)  = 2 * Increment((2m+1)/2
                 = 2 * Increment(m + 1/2)
                 = 2 * Increment(m)
                 = 2 * (m+1)
                 = 2m + 2 = 2m + 1 + 1 = x + 1
=> case closed
```
