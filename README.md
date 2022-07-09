# Personal Website

## Install Golang

```bash
brew install go
```

## CLI for Golang

```
brew install golang-migrate
```

##

## API
- Blog
  - [GET] `api/blogs` [Done]
    - Get blogs
    http://127.0.0.1:8080/api/blogs?limit=10&page=1
    
    {"status": string, "message": string, 
      data:
      {"blogs": [
        {
          "id": number,
          "title": number,
          "link": number,
          "visitor": number,
          "img_link": string,
          "created_at": datetime,
          "updated_at": datetime,
          "tags"[
            {
              "id": number,
              "name": string,
              "created_at": datetime
            }
          ]
        }
      ]},
      "meta":{"limit": number,"page": number,"order_by":"id","total_page": number}
    }

  - [GET] `api/blogs?tagId=id` [TODO]
    - Get blogs by tag_id


  - [GET] `api/blog/:id` [Done]
    - Get detail blog by blog_id

    http://127.0.0.1:8080/api/blog/1

    {
      "status":string,
      "message":string,
      "data":{
        "id":number,
        "title":string,
        "link":string,
        "visitor": number,
        "img_link":string,
        "created_at":"2022-06-17T15:57:43Z",
        "updated_at":"2022-06-17T15:57:43Z",
        "tags"[
            {
              "id": number,
              "name": string,
              "created_at": datetime
            }
        ],
        "emoji": {"id": number,"funny":number,"sad":number,"wow":number,"clap":number,"perfect":number,"love":number,"hard":number,"good":number,"mad":number,"created_at":datetime,"updated_at":datetime"},
        "html": string
      }
    }
  - [POST] `api/blog`
    - Create blog
    - need header 
  []

  - [PUT] `api/blog/:id` 
    - Update detail blog by blog_id
  
  - [GET] `api/visitor` [Done]
  http://127.0.0.1:8080/api/visitor?blogId=1
   {
      "status":string,
      "message":string,
  }

- Tag
  - [GET] `api/tags` [Done]
    - Get tags
    http://127.0.0.1:8080/api/tags
    {
      "status":string,
      "message":string,
      "data":{
        "tags": [
          {"id": number,"name": string,"created_at": datetime}
        ],
        "meta":{"limit":number,"page":number,"order_by": string,"total_page":number}
      }
    }

  - [GET] `api/tag/:id` [Done]
    - Get detail tag by tag_id
    http://127.0.0.1:8080/api/tag/1
    {
      "status":string,
      "message":string,
      "data":{ "id": number,"name":string,"created_at": datetime}
    }

- Emoji
  - [PUT] `api/emoji/:id` [Done]
    - Update detail emoji by blog_id
    http://127.0.0.1:8080/api/emoji/:id
    {
        funny: number,
        sad: number,
        wow: number,
        clap: number,
        perfect: number,
        love: number,
        hard: number,
        good: number,
        mad: number
    }
    
    {
      "status":string,
      "message":string,
      "data":{
        id: 1,
        funny: number,
        sad: number,
        wow: number,
        clap: number,
        perfect: number,
        love: number,
        hard: number,
        good: number,
        mad: number
        created_at: datetime,
        updated_at: datetime
      }
    }

- Comment

  - [GET] `api/comments?blogId=1` [DONE]
  http://127.0.0.1:8080/api/comments?blogId=1
  {
      "status":string,
      "message":string,
      "data":{
        "comments":[
          {
            "id":number,
            "name":string,
            "comment":string,
            "character":number,
            "created_at":datetime,
            "updated_at":datetime,
          }
        ],
      "meta":{"limit":number,"page":number,"order_by":string,"total_page":number}}
    }
  - [POST] `api/comment` [DONE]
  http://127.0.0.1:8080/api/comment
  {
    blog_id: number,
    name: string,
    comment: string,
  }
  {
    status: string,
    message: string,
    data: {
      id: 6,
      name: string,
      comment: string,
      character: number,
      created_at: datetime,
      updated_at: datetime
    }
  }
  blog_id: 1,
    name: "test1",
    comment: "test123",

- Auth
  - Login
  {
    account: string,
    password: string
  }
  {
    status: string,
    message: string,
    data: {
      user: {
        id: number,
        uuid: string,
        account: string,
        first_name: string,
        last_name: string,
        email: string,
        role: string,
        created_at: datetime,
        updated_at: datetime,
        login_time: datetime
      },
      access_token: string,
      refresh_token: 'string,
      expired_in: int,
      token_type: string
  }

