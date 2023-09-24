# Finch
[[logo.jpg]]

## Inspiration
Tourists and travelers find themselves in potentially dangerous situations due to a lack of information. It's a scenario we've all heard of, and it's one we're dedicated to changing. Specifically the 2015 Waze app case in Brazil, where a group of tourists ended up in a favela being tragically killed. 

## What it does
Finch provides the ability to report a dangerous situation on a map, so other users can avoid the area when trying to get somewhere. 

## How we built it
We used Go's standard library to serve HTML templates over an HTTP server and the Google Maps API. We also used PostgreSQL as our database (running with docker) and TailwindCSS for our styling.  

## Challenges we ran into
Merge issues lost us time and we couldn't deploy our service on time. We still hope to finish our project after the event.

## Accomplishments that we're proud of
Providing a correctly calculated route avoiding a dangerous area. 

## What we learned
We learned Go and TailwindCSS for easy server side rendering.

## What's next for Finch
Finishing the training of our machine learning model to learn over time which areas are dangerous in what times. 

## Set up

1. Install dependencies
- Go
- modd
- goose
- docker and docker compose
- Makefile

2. Copy example env
```
cp .env.example .env
```

3. Add database env variable to source shell file
```
export FINCH_DB_DSN='host=localhost user=finch password=finch dbname=finch sslmode=disable'
```

4. Start program
```sh
make dev
```

5. View application on localhost:4000
