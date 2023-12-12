<h1 align="center">Green Plates Back End</h1>

Backend server for the Green Plates website, another project that I'm working on with [Ben Gregory](https://github.com/bengregory23)

## API Reference

## Auth

Authentication uses JWT.

#### Signup

```http
  POST /api/v1/auth/signip
```

| Parameter  | Type     | Description                 |
|:-----------|:---------|:----------------------------|
| `username` | `string` | **Required**. Username      |
| `email`    | `string` | **Required**. Email address |
| `password` | `string` | **Required**. Password      |

#### Login

```http
  POST /api/v1/auth/login
```

| Parameter   | Type      | Description                  |
|:------------|:----------|:-----------------------------|
| `email`     | `string`  | **Required**. Email address  |
| `password`  | `string`  | **Required**. Password       |

#### Logout

```http
  POST /api/v1/auth/logout
```

## Users

#### Get all users

```http
  GET /api/v1/users
```

#### Get user by id

```http
  GET /api/v1/users/:id
```

| Parameter  | Type      | Description                   |
|:-----------|:----------|:------------------------------|
| `id`       | `string`  | **Required**. Id of the user  |

#### Delete user

*Protected route*

```http
  DELETE /api/v1/users/
```

| Parameter  | Type      | Description                   |
|:-----------|:----------|:------------------------------|
| `id`       | `string`  | **Required**. Id of the user  |

## Ingredients

#### Get all ingredients

```http
  GET /api/v1/ingredients
```

## Recipes

#### Get all recipes

```http
  GET /api/v1/recipes
```

#### Create a new recipe

```http
  POST /api/v1/recipes
```

| Parameter      | Type             | Description                                                        |
|:---------------|:-----------------|:-------------------------------------------------------------------|
| `title`        | `string`         | **Required**. Title of the recipe                                  |
| `description`  | `string`         | **Required**. Description of the recipe                            |
| `preptime`     | `uint`           | **Required**. Time in minutes                                      |
| `cooktime`     | `uint`           | **Required**. Time in minutes                                      |
| `servings`     | `uint`           | **Required**. Amount of servings                                   |
| `categoryid`   | `uint`           | **Required**. Category of the recipe                               |
| `steps`        | `[{string}...]`  | **Required**. List of strings describing each step for the recipe  |
| `ingredients`  | `check under`    | **Required**. List of all the ingredients                          |

```json
"ingredients": [
        {
            "Ingredient": {
                "ID": ...
            },
            "Quantity": ...,
            "Measurement": "..."
        },
        ...
    ]
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/NabilOuldHamou/green-plates-back
```

Go to the project directory

```bash
  cd green-plates-back
```

Install dependencies

```bash
  go get
```

Start the server

```bash
  go run .
```

## License

[MIT](https://github.com/NabilOuldHamou/green-plates-back/LICENSE)

## Authors

-   [@NabilOuldHamou](https://www.github.com/NabilOuldHamou)
-   [@BenGregory](https://www.github.com/BenGregory23)