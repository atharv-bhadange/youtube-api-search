# Youtube Video Search API
API to fetch latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

## Approach
1. Get the youtube video metadata(title, description, publishing datetime, thumbnails URLs, etc.) using the Youtube Data API in a goroutine.
2. Store the metadata in a database.
3. Create an API to fetch the metadata from the database in a paginated response sorted in reverse chronological order of publishing datetime.

## Stack 
- Golang : supports concurrency and is fast
- GoFiber : fast and lightweight web framework
- Postgres : fast and reliable relational database

## How to run
- Ensure Go and Postgres are installed
- Ensure the Postgres server is running and the database that you want to use is created
- Clone the repository
- Copy the `.env.example` file to `.env` and update the values
    - Pro tip: Supply multiple API KEYs as a comma separated values to avoid hitting the quota limit
- Build and run the application
```bash
go build -o ./build/main && ./build/main 
```
- You can optionally pass the query tag as a command line argument, the default is `fampay`
```bash
go build -o ./build/main && ./build/main cricket
```

## API
- `GET /search` : Gets the paginated list of videos sorted in reverse chronological order of publishing datetime
    - Query parameters
        - `page` : page number
        - `limit` : number of videos per page

