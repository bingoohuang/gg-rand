# gg-rand

generate random data like name, email, uuid, address, images and etc.

1. build and install: `make`
2. run: `gg-rand`

usage:

```sh
$ gg-rand -h    
Usage of gg-rand:
  -dir,d   string   In which dir to generate random files. (default temp dir)
  -input,i string   Input file name. 
  -n       int      How many random values to generate. (default 1)
  -len,l   int      Length.
  -cost,c  bool     Time costed.
  -tag     string   Which type to generate, like uuid, art, id, email and etc. (empty for prompt, all or all)
```

```sh
$ gg-rand -t all
2022/02/18 12:51:28 blake3hash-zeebo: NHCQdMI36d87iJEQdlKAzhe-Hj-L-IhKRisBf_fNiPQ (len: 43), cost 81.024754ms
2022/02/18 12:51:28 blake3hash-luke: tHA_FGtStLU7OF9t_kAj0Rho4Xc8-7kh9faQxIRp-wU2K34EMWpgNm-_wYn3BStK07GqufPhZ31zLvsdYZXVVo2ZAzO7sEgEUU41VZHF4SvryAMDeNFxCKBTJFOY6RRrQGynH_5qAmA5flM7pDX5xGzUOrUqYviD-YjN8fxUk0IfmxN-b8KRw_ZZ7NlFIr7iAzD3IE8ekV92OfakOX0E1fbHDMb45dECta9kEFvsv8otiNGFAdkZsw-vZJ7TPCAI9q0nGscKpEPr84AzwLSjN285vPnbvOxAd1u-z3kYbAcell56jMjry4ogoYWvt4lwwKs43qLPVANdPfq9iaD_wQ (len: 342), cost 88.356272ms
2022/02/18 12:51:28 xxhash: 9_wX8E03ujk (len: 11), cost 88.075279ms
2022/02/18 12:51:28 md5-hash: J6dkyz7sTx4oElml0ubWuA (len: 22), cost 100.482214ms
2022/02/18 12:51:28 sha256-hash: dJuerNOhrUDg1EwcHriZiB0qfS0ppbKfsMMgFMm6KAY (len: 43), cost 105.760105ms
2022/02/18 12:51:28 murmur3-32-hash: SLlhBQ (len: 6), cost 88.320255ms
2022/02/18 12:51:28 murmur3-64-hash: KHHKVUc_G_Q (len: 11), cost 86.119748ms
2022/02/18 12:51:28 murmur3-128-hash: 17tY7TSO6t4KNeyXR_bY6Q (len: 22), cost 80.836139ms
2022/02/18 12:51:28 imo-hash: gICABV-2lI5hWj5vdgf-4w (len: 22), cost 80.966643ms
2022/02/18 12:51:28 Base64Std: HZbpDZr2vwthi+i4obr8Wq+eHxcSh18u3ADAO/dHZQ8jZ8uzbbKu2vIniRAbsiE7IvcMsHLabR82L9cTg0h4C2dc9PA2C9fKRLG5OZb5oAyc0aGAtNVt+X5VQwNl4dhppdOSgg== (len: 136), cost 2.053µs
2022/02/18 12:51:28 Base64RawStd: 1Pv3+S2qRGvA0nlPju7+i8dJAcD/OLKsF764OMd7VV6k7veS9C62/h+7qKk17xytDy30N//c81EwtA9IMPg05KZXaqc28QeMyGRQAMHyPLT0+BD5zQyQMl7UdYasgXc0ONL0iA (len: 134), cost 685ns
2022/02/18 12:51:28 Base64RawURL: krmNZK4zVIds1EFms9-dJhvHZgeSXdkWQmTrE0VUxPhL1ShxxmnXBZaZgvi8M6dMUeprBG6anW1niN2cvdPbGQ9-rIhfhEGWB1xOeK66XK_HyqSyxBZ-xf0xLW0iv_hsxqFa_g (len: 134), cost 464ns
2022/02/18 12:51:28 Base64URL: pqbQSByIQ6wbIRbJPVu8iUlv1Emyi5DtVNLorGwM28UJ8hymbHxR4YVWvBVjOyk20mVgZdnXxEBG4Llm-czY5T_lHconRKXO3jlvwjA9z-fANHLo_HoYhpxTgGt8a3cJh0-9uQ== (len: 136), cost 563ns
2022/02/18 12:51:28 SillyName: Reapersouth (len: 11), cost 2.172µs
2022/02/18 12:51:28 Email: jamesmiller418@test.org (len: 23), cost 2.645µs
2022/02/18 12:51:28 IP v4: 142.98.123.92 (len: 13), cost 2.744µs
2022/02/18 12:51:28 IP v6: 535b:d234:e94d:4ef0:99f7:8c65:b319:f323 (len: 39), cost 6.589µs
2022/02/18 12:51:28 UserAgent: Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0 (len: 73), cost 586ns
2022/02/18 12:51:28 Password: !AzkY-2w((&-cX$FD-U%R2 (len: 22), cost 44.946µs
2022/02/18 12:51:28 Password Easy: HBVgT-sMtw4-B6wEh-2y8f (len: 22), cost 34.234µs
2022/02/18 12:51:28 Numbers: 055896412259566545883041596158322132654915086038663600067540 (len: 60), cost 9.744µs
2022/02/18 12:51:28 Letters: tEXlXHrIgGBXFcmKTYvwmpDrtOEcQfmeblibIiCPBtcULdzkZXqBhGXWSPgD (len: 60), cost 2.125µs
2022/02/18 12:51:28 JSON: {"ardri":4.43,"prenumber":null,"regressor":null,"unmind":null,"deformative":"cephalomotor"} (len: 91), cost 14.841µs
2022/02/18 12:51:28 String: GdHCNafIDDEkXjUZUTHFgSAFHWDSGhldPdOEeBAalkQQTTigeBAgFWVMhHTC (len: 60), cost 943ns
2022/02/18 12:51:28 Captcha DirectString: 8JkrB /var/folders/c8/ft7qp47d6lj5579gmyflxbr80000gn/T/3865434016/25Giag2ItPNhuVSZUoT3CUXfqRQ.png (len: 97), cost 3.271063ms
2022/02/18 12:51:28 formula: [11 × 8 = ?]
2022/02/18 12:51:28 Captcha SimpleMathFormula: 88 /var/folders/c8/ft7qp47d6lj5579gmyflxbr80000gn/T/3865434016/25GiahEJyhsQH3ZT2sDqu32RBo0.png (len: 94), cost 3.25333ms
2022/02/18 12:51:28 sony/sonyflake: uV2a2BnNzoWTkB (len: 14), cost 84.08µs
2022/02/18 12:51:28 max: 
 int64: 9223372036854775807 (len: 19), uint64: 18446744073709551615 (len: 20) (len: 78) [
 int32: 2147483647 (len: 10), uint32: 4294967295 (len: 10) 
 int16: 32767 (len: 5), uint16: 65535 (len: 5) 
 int8: 127 (len: 3), uint8: 255 (len: 3) 
 float64: 179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368.000000, float32 340282346638528859811704183484516925440.000000], cost 19.074µs
2022/02/18 12:51:28 oklog/ulid: 01FW5JQE1RRARKVCYBY7R519KV (len: 26) [base32固定26位，48位时间(ms)+80位随机], cost 7.139µs
2022/02/18 12:51:28 chilts/sid: 1RKn3FRk91d-5bOymEA1zqH (len: 23) [32位时间(ns)+64位随机], cost 2.57µs
2022/02/18 12:51:28 kjk/betterguid: -MwA9vVszx6JReLZiqJB (len: 20) [32位时间(ms)+72位随机], cost 1.18µs
2022/02/18 12:51:28 segmentio/ksuid: 25GiafVbOsNMzg15wWwHVKRl75o (len: 27) [32位时间(s)+128位随机，20字节，base62固定27位，优选], cost 5.224µs
2022/02/18 12:51:28 customized/ksuid base64: 25GiabtZKlp5SMyM9A4dYQjT6Tw (len: 27) [25Giafr7LrbCRIbMx9xSnCMv9XQ 25GiajgdMz5RXqe7tJrmLJqio63 a<=b: true b<=c: true], cost 11.876µs
2022/02/18 12:51:28 lithammer/shortuuid: aghKjPHdBcwJnk2XU8kvq4 (len: 22) [UUIDv4 or v5, 紧凑编码], cost 26.285µs
2022/02/18 12:51:28 google/uuid v4: 242e763c-2aa5-4360-8fa0-62e9418d3e4c (len: 36) [128位随机], cost 1.929µs
2022/02/18 12:51:28 satori/go.uuid v4: f939e15a-abe4-453d-ab0a-c2e18ed3e4e4 (len: 36) [UUIDv4 from RFC 4112 for comparison], cost 1.898µs
2022/02/18 12:51:28 aidarkhanov/nanoid/v2: ZV4uObe6DmZyvz7odmmq_ (len: 21), cost 4.011µs
2022/02/18 12:51:28 matoous/go-nanoid/v2: GkxWbCDP5fQjxqM6agYAW (len: 21), cost 2.018µs
2022/02/18 12:51:28 coolbed/mgo-oid Mongodb Object ID: 620f25d0c3666e6a0cd918d0 (len: 24) [Timestamp: 1645159888], cost 1.065µs
2022/02/18 12:51:28 rs/xid Mongo Object ID: c87ibk7m20rmk34omjmg (len: 20) [32 位Time: 2022-02-18 12:51:28 +0800 CST, 24位Machine: 9hA3, Pid: 27148, , Counter: 10007789 4B time(s) + 3B machine id + 2B pid + 3Brandom], cost 14.914µs
2022/02/18 12:51:28 BSON Object ID: 620f25d0c3666e6a0c000001 (len: 24), cost 14.85µs
2022/02/18 12:51:28 snowflake ID: 1494535013352345600 (len: 19) [41位 Time: 1645159888958, 10位 Node: 1, 12位 Step:0], cost 1.401µs
2022/02/18 12:51:28 Random ID with fixed length 12: 101335860980 (len: 12), cost 7.671µs
2022/02/18 12:51:28 customized snowflake ID with fixed length 12: 100754487808 (len: 12), cost 782ns
2022/02/18 12:51:28 customized snowflake ID with uint32: 5894424 (len: 7), cost 555ns
2022/02/18 12:51:28 姓名: 唐棍蚘 (len: 9), cost 7.795µs
2022/02/18 12:51:28 性别: 男 (len: 3), cost 298ns
2022/02/18 12:51:28 地址: 安徽省宿州市訊攔路4057号婐徿小区15单元380室 (len: 60), cost 1.586µs
2022/02/18 12:51:28 手机: 18278881764 (len: 11), cost 2.644µs
2022/02/18 12:51:28 身份证: 428929197706083004 (len: 18), cost 4.625µs
2022/02/18 12:51:28 有效期: 20030623-20230623 (len: 17), cost 1.617µs
2022/02/18 12:51:28 发证机关: 乐山市公安局某某分局 (len: 30), cost 451ns
2022/02/18 12:51:28 邮箱: lltauocc@nzuig.cloud (len: 20), cost 966ns
2022/02/18 12:51:28 银行卡: 6214869843396646 (len: 16), cost 2.188µs
2022/02/18 12:51:28 日期: 2012年03月18日 (len: 17), cost 1.121µs
2022/02/18 12:51:28 Generative art: Junas: /var/folders/c8/ft7qp47d6lj5579gmyflxbr80000gn/T/3865434016/25GiaikJX26mWeNlGBZTBj7a5dP.png (len: 98), cost 15.272492ms
2022/02/18 12:51:33 Unsplash: /var/folders/c8/ft7qp47d6lj5579gmyflxbr80000gn/T/3865434016/25GibCS3FbU3QJMO5zNEKqSLMy9.png (len: 91), cost 4.043413278s
2022/02/18 12:51:33 Image: 17KiB 1433886362_640x320.jpg (len: 28), cost 15.307691ms
2022/02/18 12:51:33 PBE Encrypt: <nil> , cost 451ns
2022/02/18 12:51:33 PBE Decrypt: <nil> , cost 232ns
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

## Unique IDs

```sh
$ gg-rand -t id       
2022/02/18 12:53:06 oklog/ulid: 01FW5JTCZWVRTAGCGJ4B6E0DNT (len: 26) [base32固定26位，48位时间(ms)+80位随机], cost 4.938µs
2022/02/18 12:53:06 chilts/sid: 1RKn4ezVuGt-6a4qNe_bDsZ (len: 23) [32位时间(ns)+64位随机], cost 4.572µs
2022/02/18 12:53:06 kjk/betterguid: -MwAAIEw6eMXaJwRyDl9 (len: 20) [32位时间(ms)+72位随机], cost 995ns
2022/02/18 12:53:06 segmentio/ksuid: 25Gin0uvPFI2YB0EAWW9Pwo0Xs8 (len: 27) [32位时间(s)+128位随机，20字节，base62固定27位，优选], cost 5.711µs
2022/02/18 12:53:06 customized/ksuid base64: 25GimvRwpjg07Gzp6dkWpYCI2ka (len: 27) [25GimzuwbAddDE0nyVjmlVheS2U 25Gin3F0rwwMCkfaqnXfcRJXkMh a<=b: true b<=c: true], cost 7.841µs
2022/02/18 12:53:06 lithammer/shortuuid: CGrL8k8ciS7A3NAqES35cB (len: 22) [UUIDv4 or v5, 紧凑编码], cost 26.529µs
2022/02/18 12:53:06 google/uuid v4: b3c9b1be-c2e2-4d7e-98d6-215a5c540390 (len: 36) [128位随机], cost 1.94µs
2022/02/18 12:53:06 satori/go.uuid v4: 077430aa-a5dc-48f3-9d03-c9788c303721 (len: 36) [UUIDv4 from RFC 4112 for comparison], cost 2.006µs
2022/02/18 12:53:06 aidarkhanov/nanoid/v2: I1ddyW-pL7i0lW219B2ve (len: 21), cost 1.798µs
2022/02/18 12:53:06 matoous/go-nanoid/v2: qxDCUAGlpGbkuZkcpSOEb (len: 21), cost 1.993µs
2022/02/18 12:53:06 coolbed/mgo-oid Mongodb Object ID: 620f2632c3666e6a714e0afd (len: 24) [Timestamp: 1645159986], cost 4.259µs
2022/02/18 12:53:06 rs/xid Mongo Object ID: c87iccnm20rmks9t7c00 (len: 20) [32 位Time: 2022-02-18 12:53:06 +0800 CST, 24位Machine: 9hA3, Pid: 27249, , Counter: 4012800 4B time(s) + 3B machine id + 2B pid + 3Brandom], cost 7.879µs
2022/02/18 12:53:06 BSON Object ID: 620f2632c3666e6a71000001 (len: 24), cost 9.101µs
2022/02/18 12:53:06 snowflake ID: 1494535421097414656 (len: 19) [41位 Time: 1645159986172, 10位 Node: 1, 12位 Step:0], cost 1.244µs
2022/02/18 12:53:06 Random ID with fixed length 12: 101349623713 (len: 12), cost 6.027µs
2022/02/18 12:53:06 customized snowflake ID with fixed length 12: 100754587648 (len: 12), cost 1.673µs
2022/02/18 12:53:06 customized snowflake ID with uint32: 5895208 (len: 7), cost 665ns
```

## Resources

1. [Generating good unique ids in Go](https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html)
2. [A brief history of the UUID](https://segment.com/blog/a-brief-history-of-the-uuid/)
