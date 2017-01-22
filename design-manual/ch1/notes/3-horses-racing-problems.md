### Horse racing
- There are 25 horses
- At most, 5 horses can race at the same time
- Find the minimum amount of races you have to do to get top three horses 

### Solution
- First 5 races:
  - Divide the 25 horses into 5 groups  A,B,C,D,E
  - Let these guy race and get out 5 winner
  - A1,B1,C1,D1,E1
- 6th race
  - Get A1, B1, C1, D1, E1 to race
  - Get the fastest 3 horses and discard the other 2( there is no chance they can be the fastest three any more)
  - We have F1,F2,F3
- 7th race
  - Take the bottom 2 from top 3 for the first 2 slots
  - The third slot will come from he second best horses that is in group of F1
  - The fourth slot will come from he second best horses that is in group of F2
