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