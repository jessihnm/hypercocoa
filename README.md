# A Gateway for Hypercocoa

## CLI
### Add Farmer
```
./hypercocoa add farmer -y Escuintla -c Guatemala -e iamfarmer@example.com -f Miguel -n Cortez -u 12 -p +502123456789 -s "Ent. Colonia Las Trozas" -z 0

{"ID":"0579a4ed-debd-4fe0-a8bb-2209d84f00f4","Surname":"Cortez","Firstname":"Miguel","Address":{"Number":"12","Street":"Ent. Colonia Las Trozas","City":"Escuintla","ZipCode":"0","Country":"Guatemala"},"Phone":"+502123456789","Email":"iamfarmer@example.com"}
Submit Transaction: CreateAsset, creates new asset on the ledger 
2022/07/09 17:14:47 *** Transaction committed successfully
```

## Web Gateway
### Get All Assets
```
curl localhost:1337/hypercocoa/assets -XGET -i

HTTP/1.1 200 OK
Date: Sat, 09 Jul 2022 15:24:36 GMT
Content-Length: 552
Content-Type: text/plain; charset=utf-8

[
  {
    "ID": "8afe3b91-7885-4154-b860-e89aba020050",
    "AssetType": "cocoabag",
    "Refs": {
      "farm": "f15e5956-dc37-4c52-a28c-8fefa1913cb1"
    },
    "Data": {
      "currency": "USD",
      "price": "20",
      "type": "trinitario",
      "weight": "20"
    }
  },
  {
    "ID": "c7c31b83-cd7f-42b3-ab5d-77d43e6e516f",
    "AssetType": "cocoabag",
    "Refs": {
      "farm": "e74fe0eb-1249-412e-8041-eb0718966ee1"
    },
    "Data": {
      "currency": "USD",
      "price": "12",
      "type": "criollo",
      "weight": "5"
    }
  }
```

### Get Single Assets
```
curl localhost:1337/hypercocoa/asset/c7c31b83-cd7f-42b3-ab5d-77d43e6e516f -XGET -i

HTTP/1.1 200 OK
Date: Sat, 09 Jul 2022 15:32:21 GMT
Content-Length: 245
Content-Type: text/plain; charset=utf-8

{
  "ID": "c7c31b83-cd7f-42b3-ab5d-77d43e6e516f",
  "AssetType": "cocoabag",
  "Refs": {
    "farm": "e74fe0eb-1249-412e-8041-eb0718966ee1"
  },
  "Data": {
    "currency": "USD",
    "price": "12",
    "type": "criollo",
    "weight": "5"
  }
}
```

### Add Asset 
```
 curl localhost:1337/hypercocoa/assets -i -XPOST -d '  {
    "ID": "c7c31b83-c34f-42b3-ab5d-77asdasdasdf",
    "AssetType": "cocoabag",
    "Refs": {
      "farm": "e74fe0eb-1249-412e-8041-eb0718966ee1"
    },
    "Data": {
      "currency": "USD",
      "price": "25",
      "type": "criollo",
      "weight": "15"
    }
  }'

HTTP/1.1 200 OK
Date: Sat, 09 Jul 2022 15:32:54 GMT
Content-Length: 0

```

### Delete Asset 
```
curl localhost:1337/hypercocoa/assets/c7c31b83-c34f-42b3-ab5d-77asdasdasdf -XDELETE -i

HTTP/1.1 200 OK
Date: Sat, 09 Jul 2022 15:33:23 GMT
Content-Length: 39
Content-Type: text/plain; charset=utf-8

{
   "type": "Buffer",
   "data": []
}
```

### Update Asset
```
curl localhost:1337/hypercocoa/assets/3e780024-c96b-4b7b-9840-7d64c03c676d -i -XPATCH -d ' {    "ID": "3e780024-c96b-4b7b-9840-7d64c03c676d",
    "AssetType": "cocoabag",
    "Refs": {
      "farm": "66356f82-433a-4bc0-8f8e-19bc914f8c3e",
      "factory": "367ebbb2-5ade-4096-9efb-037fa6007c03"
    },
    "Data": {
      "currency": "GTQ",
      "price": "67",
      "type": "Criollo",
      "weight": "8"
    }
  }'

HTTP/1.1 200 OK
Access-Control-Allow-Headers: Origin, Content-Type, X-Auth-Token
Access-Control-Allow-Methods: HEAD, GET, POST, PUT, PATCH, DELETE
Access-Control-Allow-Origin: *
Date: Sat, 09 Jul 2022 21:00:44 GMT
Content-Length: 0
```