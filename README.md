# Advent of Code 2023
These are my solutions for the [Advent of Code 2023](https://adventofcode.com/2023/about), written in go.

My goal for this year is to get around learning some go. I didn't write a single line of go (other than a "hello, world", which is also included in this repo) before i started on day 1. And didn't really do any programming for several years (i do write a lot of bash scripts and the occasional python script to automate things, but i don't consider that to be real programming)

So far, it's been a pretty nice experience.

Here are some thoughts on those days (spoiler free of course).

## Day 1
Starting without some proper basics was a bit annoying as i had to google more or less everything. I did start reading a book (The Go Programming Language by Alan A. A. Donovan & Brian W. Kernighan) about go before is started, but i only reached page 30 and had no opportunity even try the simplest things. Luckily, through some curiosity (looking through the command palette in vscode), i found how to create simple unit tests and tried to use them to get through the first puzzle. It worked quite well, but of course i ran into several edge cases when trying to go for the actual solution. But one edge case after the other, i managed to solve all those cases and had a solution i was quite happy with.

## Day 2
The basics started to work a bit better. But since i used some new things, it was still a bit tedious (thank god i'm halfway decent at searching and filtering out the usefulnes).
Again, i used unit tests to constanly see if my code works as expected. But i did spend a bit more time thinking about my code and i actually completed both puzzles without running into edge cases that i didn't cover, so both worked on the first try.

## Day 3
The puzzle today was a bit more tricky, so i decided to first concentrate on parsing the input into a "proper" data structure. It took more time than i expected. That was partly due to some knots in my head and partly due to misunderstandings about how things work, memorywise (i didn't know, that "range" always creates copies and thus had problems updating some values while iterating through the parts list again). But again, the time was well spent, as reading and parsing the input was basically the whole puzzle. I took about 2 hours for the first puzze (including parsing input and creating the final data structure), but only about 1 minutes for the second puzzle, because i already had all the data available. Also, both solutions worked in the first try again. Yay!
I'm not fully satisfied with my solution, because of the deeply nested code, but since it doesn't need to go into production, i'll just leave it at that.

## Day 4
Instead putting the values into a useful data structure, i started out doing it the simple way, just generating the value on a line by line basis. When i started the second part, i immediately knew that i should have gone a different route. So i changed everything and suddenly, the second part was way easier. I should have trusted my instincts and gone with that right away.

## Day 5
That took a while. Time to compute that is. The final solution now only covers the second puzzle, because i was running late and didn't have time to properly include both (and i've been to lazy to do it now). The final solution takes a lot of time because there are >2 billion seeds to check. It could be improved vastly by parallel processing, which i didn't cover so far. Maybe i will add it later. 

## Day 6
That was surprisingly quick and simple. Nice change of pace/complexity. I did "overprepare" for the second puzzle, but it took not much time, so i'll just leave it in there. At least i was prepared enough to make the second puzzle super easy again.
I used an iterative process to calculate the min and max values, even though i'm sure that it can be reasoned out just using maths. Not the most efficient, but i'm not here for math. 

## Day 7
There's not much to say today. I really liked the idea. I obviously question my decision to convert a full hand to a single number, but since it worked, i can't really complain ;-)
What i can complain about is the fact that i didn't spend any time making the code a bit nicer to read (i.e. i used variable names like k and v). Maybe i'll come back to that. 

## Day 8
Phew. That was a wild ride. The first puzzle was quite easy and easily done. The second though... That was a the sort of puzzle i was fearing the most. In the end, i think it's more of a math puzzle, than a programming puzzle. Yes, programming is necessary to get actual values (in a reasonable time), but the overall problem needed to be solved from a mathematical standpoint. Also, there were some apparent regularities in the given examples, but so far, most of the previous examples left several uncovered edge cases. And today, the apparent regularity needed to be taken as granted for the full puzzle input as well. That is where the puzzle lost me, and it took quite some time to find out what was going on. The naive solution could get an answer, but since it would take roughly 14 trillion (!) steps, it would take _A LOT_ of time. In the end, i was quite frustrated. 

## Day 9
Recursion time! What a pleasant surprise. I kinda expected, that challenges would be getting way more difficult on the weekend, as one would have more time to spend. At least for this day, that wasn't the case (and also it was back to a "proper" programming solution!). I'm quite happy with my solution. It uses recursion and the code is super readable. Everything is quite compact without loosing any clarity. 

## Day 10
I don't know that was going on, but for some reason i had huge knots in my head that i couldn't loosen up. In principle, the solutions are both not that difficult, but it took ages to figure out how to do it properly. For the first puzzle, i was kinda torn between a recursive and an iterative solution and somehow i couldn't really concentrate on one of them (and thus constantly mixed up both approaches). When i finally had a proper working solution (iterative, taversing only 1 direction instead of both at the same time), i didn't expect the second puzzle to be much more difficult. And in the end, i must admit that it really isn't. But again, i couldn't wrap my head around my problems. I think i was hovering around the actual solution for about an hour without making any progress. And then it suddenly hit me, and i was done about 5 minutes later. What a demoralizing progress.
I hope it was just my accumulated lack of sleep from the last week.