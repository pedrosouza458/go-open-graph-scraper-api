## Golang Open Graph Scraper Api

This is a wrapper to use my package Golang Open Graph Scraper in api, to let users access the functions without needing to use Go or to install the package, docs and deploy are coming soon.

### Run this code Locally
```git
git clone https://github.com/pedrosouza458/go-open-graph-scraper-api
```

```go

go mod tidy

```

```go

go run .

```
Now the server is running in port 8080.
```
http://localhost:8080/websites
```
## Request Guide:

  

### Request

```bash

curl 'http://localhost:8080/websites?url=https://github.com/pedrosouza458/go-open-graph-scraper'

```
### Response

```json
{
  "Descpription": "Go-Open-Graph-Scraper is a go package to get metadata from websites easily, also returning formatted logos and names for any compatible website. - pedrosouza458/go-open-graph-scraper",
  "Img": "https://opengraph.githubassets.com/50e5a16d77647cd090085c7d7daacdac42b22d8109d2095fe22cce2606e345f1/pedrosouza458/go-open-graph-scraper",
  "Logo": "https://utfs.io/f/503dc239-2ada-4caf-9f5a-ef92a778c194-n8l1zp.png",
  "Name": "github",
  "Title": "GitHub - pedrosouza458/go-open-graph-scraper: Go-Open-Graph-Scraper is a go package to get metadata from websites easily, also returning formatted logos and names for any compatible website."
}
```