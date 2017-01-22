#### Schedule your movie acting calendar to act in the most movie
- **Input** A set of movie with start date and end date
- **Output** largest amount of movies I can act in these movie

### Possible solutions
- Brutefoce by picking a each movie, then pick the rest of the movie one by one, comparing all the answers and pick the best one
  - too god damn slow
- Pick the earliest job that start , the earlier you work the better it is isn't it
  - what if the first earliest movie take entire year to finish ---> you are fucked
- Pick the shortest movie to fit all the most amount of movie in your year
  - some of these short movie is in between of2 big movie, so you pick the small movie , then block 2 big moviie ---> less  money

### Best solution
- Pick the movie with the end time to be the least in all movie
- Continue to pick the remaining movie using this concept
 You can never lose by doing this because
  - All the other movies can be shorter, but none of them block less movie than the one you pick
  - You always get the most remaining movies to pick from
 

