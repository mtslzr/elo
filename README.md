# Elo

Elo is a Go-based Elo Ratings calculator.

_Current Version: 0.1.1_

| Master | Develop |
|:-:|:-:|
|[![Master](https://travis-ci.com/mtslzr/elo.svg?branch=master)](https://travis-ci.com/mtslzr/elo)|[![Develop](https://travis-ci.com/mtslzr/elo.svg?branch=develop)](https://travis-ci.com/mtslzr/elo)|

## Getting Started

```bash
go get github.com/mtslzr/elo
```

```go
import "github.com/mtslzr/elo"
```

## Functions

### Calculate New Rating

Takes a team rating, win probability, and result (win = 1, loss = 0, draw = 0.5), and returns the new rating. Based on kFactor of 32.

newR = kF * (Result - winP)

```go
newRating := elo.NewRating(oldRating, winProbability, result)
```

```go
newRating := elo.NewRating(1500, 0.450, 1)
// 1518
```

### Calculate Win Probability

Takes a team and opponent rating, and returns the win probability as a float64.

winProb = 1 / (10^-((teamR - oppR)/400) + 1)

```go
winProb := elo.WinProbability(teamRating, opponentRating)
```

```go
winProb := elo.WinProbability(1640, 1535)
// 0.647 (or 64.7%)
```
