## foodTinder
Create a session and start rating our products!

## Overview API

| Method | URL | Action |
|--------|-----|--------|
| `GET` | `/v1/session` | Generate a new and unique session | 
| `GET` | `/v1/session/<:id>` | Get all the votes from a given session `<:id>` | 
| `GET` | `/v1/product` | Get all the available products |
| `POST` | `/v1/product/<:id>` | Vote for or against a given product `<:id>` |

## curl commands

### Vote for or against a given product
* The path after product in the URL must be the `product_id` for which a vote will be cast.

* The body of the request must be valid JSON, with two fields: '`vote`' (boolean) and '`session_id`' (string).
If you like the product set the boolean to `true`, if you dislike it set the boolean to `false`.
The `session_id` must be a valid session that you have previously received, otherwise the system will not accept your request.
The same goes to the `product_id` in the URL.

```
curl --header "Content-Type: application/json" \
  --request GET \
  --data '{"vote":true,"session_id":"y7anpvouHGf84UNlAUt5bqHatjB2cZpKe9LrMQ2BR39S0uztJ22Ii7n2ihAtSGra"}' \
  http://5.75.165.31:8000/v1/product/3aba3a59-fd44-45e8-80db-7d4771b8f822
```

### Get all the available products

```
curl --header "Content-Type: application/json" \
  --request GET \
  http://5.75.165.31:8000/v1/product
```


## Limitations
* The system retrieves all the products available at an endpoint only once after it boots.
Afterwards, the system does not communicate further with the endpoint that provides the products.
To see all the available products, send a `GET` request to `/v1/product`

* The functionality to retrieve the aggregated score of a product accross all sessions was not implemented due to a time constraint.

All rights reserved. Eduardo Rodriguez (c) 2023 
