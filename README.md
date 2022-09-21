# go-coupang-partners

> 쿠팡 파트너스 상품리스트 링크 생성&반환 프로그램

## ⚠️ 경고
이 프로젝트는 간단한 작업을 진행하기 위해 진행된 프로젝트 입니다. 또한 허가받지 않고 진행하는 프로젝트로 언제든지 중단될 수 있습니다.

쿠팡 파트너스에 가입하여 상품을 검색하고, 광고를 할 수 있는 링크를 생성할 수 있다.
대량으로 생성하기 위해서는 API가 필요한 상황인데 방금 가입한 사용자의 경우 API사용이 불가능하다.
사이드 프로젝트의 진행을 위해 많은데이터가 필요한 상황인데 이를 수작업으로 진행할 수 없으므로 약간의 시간을 투자하여 데이터를 추출할 수 있도록 프로그램을 만들었다.

## 프로젝트 목표
쿠팡 Id/Pw/Keyword 입력 시 검색된 항목중 첫 페이지에 해당하는 상품리스트를 반환하고, 해당 리스트의 Partners Link를 전달.
- [x] 쿠팡 파트너스 로그인
- [x] 쿠팡 파트너스 상품 검색및 반환
- [ ] 파트너스 쿠팡 상품 링크 생성및 반환

## Usage
```sh
go run main.go --id "아이디" --pw "비밀번호" --keyword "키워드(ex: 양파)"
```
기본적으로 아래 데이터로 출력이 되는데 필요한 데이터만 가져와서 변경될 예정이다.
```json
{
  "rCode": "0",
  "rMessage": "",
  "data": {
    "products": [
      {
        "productId": 204889338,
        "itemId": 602752284,
        "vendorItemId": 3750883707,
        "deliveryChargeType": [
          "FREE"
        ],
        "originPrice": 8500,
        "salesPrice": 8500,
        "vendorIdList": [
          "A00163320"
        ],
        "badges": [
          "FREE_SHIPPING"
        ],
        "retail": false,
        "type": "PRODUCT",
        "title": "미쁜스토어 덜맵고 달달한 2022년산 햇 양파 5~20kg, 1개, 양파 5kg(중/소 장아찌용)",
        "rawImage": "image/vendor_inventory/d82f/759a16c647f1225cd01f49c30d9cbc5067d0c972eb1a75891434f38c0bbc.jpg",
        "image": "https://thumbnail2.coupangcdn.com/thumbnails/remote/212x212ex/image/vendor_inventory/d82f/759a16c647f1225cd01f49c30d9cbc5067d0c972eb1a75891434f38c0bbc.jpg",
        "isAdult": false,
        "isSoldOut": false,
        "isFreeShipping": true
      },
    ],
    "total": 1000,
    "shortUrl": "https://link.coupang.com/a/추천자ID",
    "sharing": {
      "title": "쿠팡을 추천 합니다!",
      "description": "쿠팡은 로켓배송",
      "imageUrl": "http://image15.coupangcdn.com/image/mobile/v3/img_fb_like.png"
    }
  }
}
```
