# WebCrawler

A Go application that uses concurrency to crawl a base URL and counts internal links, up to the max page count.

### Usage:
``` ./crawler url maxConcurrency maxPages ```

#### Example:
``` ./crawler https//news.ycombinator.com/ 3 25  ```

##### Report output:
``` 
=============================
  REPORT for https://news.ycombinator.com/
=============================
Found 128 internal links to news.ycombinator.com/item
Found 76 internal links to news.ycombinator.com/
Found 75 internal links to news.ycombinator.com/user
Found 75 internal links to news.ycombinator.com/vote
Found 41 internal links to news.ycombinator.com/from
Found 32 internal links to news.ycombinator.com/reply
Found 31 internal links to news.ycombinator.com/hide
Found 6 internal links to news.ycombinator.com/front
Found 4 internal links to news.ycombinator.com
Found 4 internal links to news.ycombinator.com/ask
Found 4 internal links to news.ycombinator.com/jobs
Found 4 internal links to news.ycombinator.com/login
Found 4 internal links to news.ycombinator.com/newcomments
Found 4 internal links to news.ycombinator.com/newest
Found 4 internal links to news.ycombinator.com/news
Found 4 internal links to news.ycombinator.com/newsfaq.html
Found 4 internal links to news.ycombinator.com/show
Found 4 internal links to news.ycombinator.com/submit
Found 3 internal links to news.ycombinator.com/lists
Found 3 internal links to news.ycombinator.com/newsguidelines.html
Found 3 internal links to news.ycombinator.com/security.html
Found 1 internal links to news.ycombinator.com/fave
Found 1 internal links to news.ycombinator.com/highlights
Found 1 internal links to news.ycombinator.com/invited
Found 1 internal links to news.ycombinator.com/pool
```