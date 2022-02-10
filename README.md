# gg-rand

generate random data like name, email, uuid, address, images and etc.

1. build and install: `make`
2. run: `gg-rand`

usage:

```sh
$ gg-rand -h
Usage of gg-rand:
  -dir string   in which dir to generate random files. (default temp dir)
  -n   int      how many random values to generate. (default 1)
  -tag string   which random type to generate, like uuid, art, id, email and etc. (default all types)
```

```sh
$ gg-rand
2022/02/09 16:24:41 blake3hash-zeebo: OdbRwQm39Xg320PK3wGpScNotdT2Vepc7ytlI0dDODQ (len: 43), cost 110.262677ms
2022/02/09 16:24:41 blake3hash-luke: JHYuMdCv8ukbHhKCEkP1_yuIjTK5wkwtM9CxD-8_HPnjiZyXRWDtfY6MGnag_bVdcgg94dgYNRtAwAZmXRYKUHCg0cQcKiyygFLerPAv6QzLFnxxclU2heQzy_r0NziXiMvhyps1Gi6-UOoF6TQvrwftrQfHolJb_-pKCOK_Cm3fIWDDj40Azy5CLLVV_SmDfZ5rwRQxmLfp5K4f8rEM-ikY9EwGradqtBo_CxNIc98BBAPN4OClqK5QTVW7g4ghduJad88V9knM1zYSNVlka5Lq4L1dV4VAjQCrMevm15BsEh1csLAdQdIqky6et6UQnVcp09NmfBt6WoKYHl1pSQ (len: 342), cost 124.770632ms
2022/02/09 16:24:41 xxhash: q9h7Tuzi1qM (len: 11), cost 116.536003ms
2022/02/09 16:24:41 md5-hash: cUe18VdgGa-XFCljkq7wiA (len: 22), cost 132.357882ms
2022/02/09 16:24:41 sha256-hash: boaaniPXpXX_GFamzKj9Ea3GJTAnqiYnnuXwqQz1MqU (len: 43), cost 137.361395ms
2022/02/09 16:24:42 murmur3-32-hash: MPD03Q (len: 6), cost 106.671235ms
2022/02/09 16:24:42 murmur3-64-hash: EuQz3Y4Nnek (len: 11), cost 114.245447ms
2022/02/09 16:24:42 murmur3-128-hash: e1Fj91t7t6QMFXei13nMvw (len: 22), cost 140.784351ms
2022/02/09 16:24:42 imo-hash: gICABRY0XoE9IHAgLWqCJA (len: 22), cost 106.050856ms
2022/02/09 16:24:42 Base64Std: f8mqas/S1LGCwwK1FR320ENVuVCbHT72TiSZyjpM7PtlA0nMMGiPXzS/JPqXsFxnr2fhNZSt6F5NH6BnkoyEWGJ5fu6ueAHk7FL09TBvj8XZYP1RmyUjO4iiOfoKdEb5KEcjjA== (len: 136), cost 5.031µs
2022/02/09 16:24:42 Base64RawStd: b4pUObv7+IJBZu943mI7bpRCtZlR8MhC+IG9ABALM+21xPeP4hqCMMBP9vsRAw/0oU1gU0GxZKmAg1Vdar1IiH+A4cntm6PItMZ0D0Qo63+tJRu3qMJ/FpYTeXANXBJj7XzYcg (len: 134), cost 764ns
2022/02/09 16:24:42 Base64RawURL: ZtEyJVQZ51VALEbv3hilkwWXjXPToadvSZz1XrL31o7EKRqV4MZhMbO5Jh9NL4Qkevoy2w1bQGNXIlbvqfP0ApGAaQVU_GPFf0RIGUyVm_jU4RAyJ7BSUv5TXFg_BV-ivtyycg (len: 134), cost 534ns
2022/02/09 16:24:42 Base64URL: FlH8nbpRn4uhbkdC1dUebUY0DZEowvvlk9WK6A4ol7kDCIfqXCYLAGp0Ey0sCwhtynpB9t6Xd78oSqTYXpQRwmnAs0qgLnrZJdGZCz-R-9S1TDp9AbucCLXyqiRxKBjZhlZw-w== (len: 136), cost 516ns
2022/02/09 16:24:42 SillyName: Mooserelic (len: 10), cost 5.703µs
2022/02/09 16:24:42 Email: masonwilson107@example.org (len: 26), cost 2.198µs
2022/02/09 16:24:42 IP v4: 60.207.191.140 (len: 14), cost 8.763µs
2022/02/09 16:24:42 IP v6: 65ae:2f5a:d92d:cd01:12b8:3927:d2d0:9b12 (len: 39), cost 3.944µs
2022/02/09 16:24:42 UserAgent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3191.0 Safari/537.36 (len: 113), cost 304ns
2022/02/09 16:24:42 Password: ptsBM-kDedt-S{hzo-Lno9 (len: 22), cost 52.782µs
2022/02/09 16:24:42 Password Easy: ypu8x-DJSwt-BShh3-5J3S (len: 22), cost 33.681µs
2022/02/09 16:24:42 Numbers: 498191799037817419435967889659003612890309828314892366290145 (len: 60), cost 5.749µs
2022/02/09 16:24:42 Letters: LobfjxyUwsdGYctNwTZwKetlKFiKZxsVMeBrUUxANDKRjsenhBrsUvQUocIr (len: 60), cost 2.345µs
2022/02/09 16:24:42 JSON: {"bitterful":"aculea","permissibly":{"preliberally":[],"theorize":[]},"perilymphatic":true,"Didelphidae":true,"rigging":null} (len: 125), cost 4.4µs
2022/02/09 16:24:42 String: BTSRFQHeDDUBgQOHKIKIEIjNRTkDbBUBWYOSIUPbKEXiQNfFOVBEUSHKXjDH (len: 60), cost 1.972µs
2022/02/09 16:24:42 Captcha: K2uVB /var/folders/c8/ft7qp47d6lj5579gmyflxbr80000gn/T/3399639142/24riPYhhSEfX50mvA6gueAEgRLN.png (len: 97), cost 4.502781ms
2022/02/09 16:24:42 sony/sonyflake: olncG1UcoJ3RkB (len: 14), cost 97.309µs
2022/02/09 16:24:42 max: int64: 9223372036854775807 (len: 19), int32: 2147483647 (len: 10), int16: 32767, float64: 1.7976931348623157e+308, float32:3.4028234663852886e+38 (len: 145), cost 3.594µs
2022/02/09 16:24:42 oklog/ulid: 01FVESBCP39WCVMB54918YD9TV (len: 26) [48位时间(ms)+64位随机], cost 12.474µs
2022/02/09 16:24:42 chilts/sid: 1RI49qr~eo0-5B92iPbeSeC (len: 23) [32位时间(ns)+64位随机], cost 2.491µs
2022/02/09 16:24:42 kjk/betterguid: -MvS_QA2HBqKtC5Fwqbq (len: 20) [32位时间(ms)+72位随机], cost 956ns
2022/02/09 16:24:42 segmentio/ksuid: 24riPX4Ge1orfjiCiHn6a4QdoXA (len: 27) [32位时间(s)+128位随机，20字节，base62固定27位，优选], cost 5.044µs
2022/02/09 16:24:42 lithammer/shortuuid: fG5UJ57EYoGwvDHFtPWG5d (len: 22) [UUIDv4 or v5, 紧凑编码], cost 28.595µs
2022/02/09 16:24:42 google/uuid v4: 548ff664-15d6-4576-87d5-4695d3c7fe9a (len: 36) [128位随机], cost 2.056µs
2022/02/09 16:24:42 satori/go.uuid v4: ec10990f-0a6a-4b55-bd26-a08d00099c69 (len: 36) [UUIDv4 from RFC 4112 for comparison], cost 1.427µs
2022/02/09 16:24:42 aidarkhanov/nanoid/v2: Ex39CoboQ9cSehI11_3ln (len: 21), cost 1.706µs
2022/02/09 16:24:42 matoous/go-nanoid/v2: 0lp862X43q68DvDE3-Zdq (len: 21), cost 1.788µs
2022/02/09 16:24:42 coolbed/mgo-oid Mongodb Object ID: 62037a4ac3666e4e63d1f9cd (len: 24) [Timestamp: 1644395082], cost 976ns
2022/02/09 16:24:42 rs/xid Mongo Object ID: c81nkinm20rksoqq5g70 (len: 20) [32 位Time: 2022-02-09 16:24:42 +0800 CST, 24位Machine: 9hA3, Pid: 20067, , Counter: 5909518 4B time(s) + 3B machine id + 2B pid + 3Brandom], cost 7.494µs
2022/02/09 16:24:42 BSON Object ID: 62037a4ac3666e4e63000001 (len: 24), cost 9.635µs
2022/02/09 16:24:42 Snowfake ID: 1491327182297894912 (len: 19) [41位 Time: 1644395082436, 10位 Node: 1, 12位 Step:0], cost 1.948µs
2022/02/09 22:09:23 snowflake ID with length 12: 100000001537 (len: 12), cost 4.845µs
2022/02/09 16:24:42 姓名: 禄鋅籴 (len: 9), cost 3.888µs
2022/02/09 16:24:42 性别: 女 (len: 3), cost 318ns
2022/02/09 16:24:42 地址: 广西壮族自治区桂林市蚐珗路2174号夕啹小区16单元807室 (len: 72), cost 3.246µs
2022/02/09 16:24:42 手机: 18341190333 (len: 11), cost 1.136µs
2022/02/09 16:24:42 身份证: 367384201204211866 (len: 18), cost 10.186µs
2022/02/09 16:24:42 有效期: 19980428-20180428 (len: 17), cost 1.393µs
2022/02/09 16:24:42 发证机关: 七台河市公安局某某分局 (len: 33), cost 574ns
2022/02/09 16:24:42 邮箱: qtsmsepg@xhaee.space (len: 20), cost 1.588µs
2022/02/09 16:24:42 银行卡: 6214481128785714252 (len: 19), cost 2.465µs
2022/02/09 16:24:42 日期: 1970年01月30日 (len: 17), cost 2.443µs
2022/02/09 16:24:42 Generative art: Junas: /var/folders/c8/ft7qp47d6lj5579gmyflxbr80000gn/T/3399639142/24riPY84OtPHPyASifoc2LgC4l0.png (len: 98), cost 16.161366ms
2022/02/09 16:24:44 Unsplash: /var/folders/c8/ft7qp47d6lj5579gmyflxbr80000gn/T/3399639142/24riPpWdmZ5nTBm0dFtKClUqwVA.png (len: 91), cost 2.479014432s
2022/02/09 16:24:44 Image: 2.9KiB 1021952211_640x320.png (len: 29), cost 19.239017ms
```

## captcha

`gg-rand -tag captcha -n 10` to generate 10 captcha images.

![image](https://user-images.githubusercontent.com/1940588/140700928-1fd794a7-21b2-4bda-81c0-b705bd632f88.png)

`gg-rand -t formula -n 10`

![image](https://user-images.githubusercontent.com/1940588/153322018-57f891c5-7456-454c-a195-fa15b136082c.png)

## images

`$ GG_IMG_SIZE=640X320 GG_IMG_FORMAT=JPG GG_IMG_FAST=N GG_IMG_FILE_SIZE=2MiB  gg-rand -tag image -n 10`

## hash functions

```sh
🕙[2022-01-23 18:32:34.382] ❯ gg-rand -t hash -i '魔戒首部曲：魔戒现身.终极加长版.mkv' -c
2022/01/23 18:34:44 Started
2022/01/23 18:34:46 blake3hash-zeebo: 6p2D6TWec8knuEcmwbjw4F7_ubnBp9qYTWO787waCO8
2022/01/23 18:34:46 Completed, cost 2.418253944s
2022/01/23 18:34:46 Started
2022/01/23 18:34:49 blake3hash-luke: 6p2D6TWec8knuEcmwbjw4F7_ubnBp9qYTWO787waCO_4Ihzid_d0qYz6kU3gacZ0eYMg_Twhzm6ioui51N9Xq4mWT9Xo-kuAYiouqvwjD7sLmYGg7FeW-W6exOhRBowjcKbVJyCNnsqgucqw8Ca2EzJ4tYd27QNWuoGTsLMoSKfC08LTWMuHIhEH5kRSR6_J2d6NtnMrAUc9t9XVBNEtDPkXbLCTihr_YXTCfUuQb7WrCUnAhInzvuSfFiTR8gPss_OBpXbvVYxFEH0JAgPjtzZbBz6es_UkdCPBdUTbbZte12gaheQA1Y6bwI0bjmCUxw3h_t6OPrtNTFRL6_CYdw
2022/01/23 18:34:49 Completed, cost 2.904523985s
2022/01/23 18:34:49 Started
2022/01/23 18:34:50 xxhash: SN4WkuaIrVE
2022/01/23 18:34:50 Completed, cost 1.173249351s
2022/01/23 18:34:50 Started
2022/01/23 18:34:58 md5-hash: W2WtAp9nzHZGRDYabofufQ
2022/01/23 18:34:58 Completed, cost 8.076151521s
2022/01/23 18:34:58 Started
2022/01/23 18:35:13 sha256-hash: YEaOGH6VGtCuz4LOgC-dM-fK2N6o4HxRmKrU3pOSzV4
2022/01/23 18:35:13 Completed, cost 14.625180691s
2022/01/23 18:35:13 Started
2022/01/23 18:35:13 imo-hash: _6ClwBZqBTM3Q9wZJOFVLQ
2022/01/23 18:35:13 Completed, cost 105.436µs
```

## Unqiue IDs

```sh
$ gg-rand -t id       
2022/02/09 16:23:27 oklog/ulid: 01FVES938PC4CACQTBYK3ZFMSD (len: 26) [48位时间(ms)+64位随机], cost 14.525µs
2022/02/09 16:23:27 chilts/sid: 1RI48kqobmd-3HU5_~expUy (len: 23) [32位时间(ns)+64位随机], cost 4.276µs
2022/02/09 16:23:27 kjk/betterguid: -MvS_7oLnBCmeHRyA9FN (len: 20) [32位时间(ms)+72位随机], cost 1.02µs
2022/02/09 16:23:27 segmentio/ksuid: 24riG9AXPwasx8N0b2SaYs2xWFX (len: 27) [32位时间(s)+128位随机，20字节，base62固定27位，优选], cost 4.868µs
2022/02/09 16:23:27 lithammer/shortuuid: XXkfrGF5CtWbaoDfpw5ptH (len: 22) [UUIDv4 or v5, 紧凑编码], cost 13.349µs
2022/02/09 16:23:27 google/uuid v4: 3c3c7a15-3c2f-4669-9488-c69c19e716c9 (len: 36) [128位随机], cost 2.114µs
2022/02/09 16:23:27 satori/go.uuid v4: f5ef4dde-e9b2-4ef0-82ce-0b2c1a220c4b (len: 36) [UUIDv4 from RFC 4112 for comparison], cost 1.471µs
2022/02/09 16:23:27 aidarkhanov/nanoid/v2: M8Dgq-Mpo58yowcPCHu4C (len: 21), cost 4.498µs
2022/02/09 16:23:27 matoous/go-nanoid/v2: wW6yvGpMySMdca-TBgxVG (len: 21), cost 1.845µs
2022/02/09 16:23:27 coolbed/mgo-oid Mongodb Object ID: 620379ffc3666e4dc567e0f1 (len: 24) [Timestamp: 1644395007], cost 1.207µs
2022/02/09 16:23:27 rs/xid Mongo Object ID: c81njvvm20rkrh97sdh0 (len: 20) [32 位Time: 2022-02-09 16:23:27 +0800 CST, 24位Machine: 9hA3, Pid: 19909, , Counter: 2614114 4B time(s) + 3B machine id + 2B pid + 3Brandom], cost 6.067µs
2022/02/09 16:23:27 BSON Object ID: 620379ffc3666e4dc5000001 (len: 24), cost 8.717µs
2022/02/09 16:23:27 Snowfake ID: 1491326866961731584 (len: 19) [41位 Time: 1644395007254, 10位 Node: 1, 12位 Step:0], cost 6.335µs
2022/02/09 19:19:58 Random ID with fix length 12: 100201261364 (len: 12), cost 9.655µs
```

## Resources

1. [Generating good unique ids in Go](https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html)
2. [A brief history of the UUID](https://segment.com/blog/a-brief-history-of-the-uuid/)
