# Golang API Server Sample

## Abstract

- Many to Many Web API Server Sample
- Web Framework : Gin (https://github.com/gin-gonic/gin)
- ORM : GORM (https://github.com/jinzhu/gorm)

## Getting Started

```sh
git pull https://github.com/naoki-kishi/go-api-sample
cd go-api-sample
docker-compose up
```

## Start Server

```sh
docker compose exec api go run main.go
```

## API docments

### Entry

#### Response example

```json
{
  "id": 1,
  "title": "entry1",
  "tags": [
    {
      "id": 1,
      "name": "tag1"
    }
  ]
}
```

#### Get all entries

`GET /entries`

#### Get an entry by id

`GET /entries/:id`

#### Create an entry

`POST /entries`

If you want to create relations with tags, you don't have to include tag id.

Request example

```json
{
  "id": 1,
  "title": "entry1",
  "tags": [
    {
      "name": "tag1"
    }
  ]
}
```

#### Update an entry

`PUT /entries/:id`

#### Delete an entry

`DELETE /entries/:id`

### Tag

#### Response example

```json
{
  "id": 1,
  "name": "tag1"
}
```

#### Get all tags

`GET /tags`

#### Get a tag by id

`GET /tags/:id`

#### Create a tag

`POST /tag`

#### Update a tag

`PUT /tags/:id`

#### Delete a tag

`DELETE /tags/:id`

## References

- [クリーンアーキテクチャの書籍を読んだのでAPIサーバを実装してみた](https://qiita.com/yoshinori_hisakawa/items/f934178d4bd476c8da32)
- [【SOLID原則】依存性逆転の原則 - DIP](https://zenn.dev/chida/articles/e46a66cd9d89d1)
- [ボトムアップドメイン駆動設計](https://nrslib.com/bottomup-ddd/)
- [クリーンアーキわからんかった人のためのクリーンじゃないけどクリーンみたいなオニオンに見せかけたSOLIDの話](https://zenn.dev/streamwest1629/articles/no-clean_like-clean_its-onion-solid)
- [go-onion-architecture-sample](https://github.com/nanamen/go-onion-architecture-sample)

## Architecture

<img src="https://user-images.githubusercontent.com/57553474/179379482-571b0913-3b1e-4908-b6a5-71e0444f38ed.jpg" width="60%" />
