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
2022/01/23 22:33:09 blake3hash-zeebo: PJ5ZV3A5RRcGZfWyNRlBnZd_O4qGLpPl5u5nCARzfPE, cost 80.742381ms
2022/01/23 22:33:09 blake3hash-luke: yWZYouzuR90GTFER54ZKO1OVAjT0TuiIdum9S8HP-SXl4zMO2d4IGgdljNJzTnrSjL_U1TmH_3EGCzDS3R6kZFaWvhqKfZVdGLFfPJ3dsJbTHpfTkTopSL7yPKxNP5WsDFinBBVnapUhloiclAJd6t8tUNiMqkwIzqb68UIPFoSIh2Mt0GoDg7pdCrHdWahMTAdnfn1alJbbSdAp2rwS5g5apHJ5g9a1DvHYz76MK0VRtOav4C8UUkIaPIEd1DgU3b6QZTmXSuihagNM4miW4WYKFs7UKJF2ZH_MegWe8y94OwVKU7FujDkqL4yWz_F0biitVuefs8BXRPTokBlt-g, cost 77.227436ms
2022/01/23 22:33:09 xxhash: lnwaRNk09Y0, cost 72.403683ms
2022/01/23 22:33:09 md5-hash: akoJmwfWqUN0OYw5AmUyVA, cost 91.057546ms
2022/01/23 22:33:10 sha256-hash: bsHROpReF8eH0dTm_uOSSmDfAF6A12FlGM5L6hzTM64, cost 102.320918ms
2022/01/23 22:33:10 murmur3-32-hash: fUBT2A, cost 76.984569ms
2022/01/23 22:33:10 murmur3-64-hash: ulTNGS2lJXk, cost 74.565215ms
2022/01/23 22:33:10 murmur3-128-hash: yDWUW7EjCnVwxcZe8Ame5w, cost 72.289683ms
2022/01/23 22:33:10 imo-hash: gICABTmUgNRrggQrueZ2DQ, cost 69.669246ms
2022/01/23 22:33:10 Base64Std: kevsaZAgRRVRY9T3ZKtPqLDcfajygLx0c1H8NRMu0w5ivXQNhARJXUB3r4gAbwoj3Tg+BmgCFhCsb/otCk84ZXNLS1oqXve4ayHA6Ca9ZcFdhTtslXb00WtAQLLQ3OUWpykxBg==, cost 10.58µs
2022/01/23 22:33:10 Base64RawStd: qZrIhHmFPXfXMGPXHsMQixysSQ+AqYOXkRPprVGxPWqk4eehjx2EYeAGmi5RfITTCU1uHy4Hb6p9rFl2j0Pihuu+3DQvM+N5zIieO1Pe6FwBoTZ0gf4HSmfAJdarxG8ouzfB3w, cost 794ns
2022/01/23 22:33:10 Base64RawURL: t9W3-XUwP8Cfyr5gznvKUJHZnMtHE3OFKmUIBCzEb_DB_2dMeBgV-9kD2az-fqBzHO7RYlfICKQLxcT-YuK8wLrKgd5r9ue48Em8yLP-RB-w9EnYZFsLd8J-E3CzhcRqDfzdiA, cost 580ns
2022/01/23 22:33:10 Base64URL: 4UuhmfKFo9qucthEKWVPmnLfMIyGgt1EApeeUdILvvl4_hiR1h0L-w4GocpGSJmZSJJgmfSsFpMwwy7B4uGNcvNPLe7w3IeXYWusZVzwvE0kHHaJbR9_8UkkPAaPJVINNs0ARw==, cost 564ns
2022/01/23 22:33:10 SillyName: Chilleriridescent, cost 2.489µs
2022/01/23 22:33:10 Email: avathompson520@test.org, cost 3.127µs
2022/01/23 22:33:10 IP v4: 33.177.133.101, cost 13.176µs
2022/01/23 22:33:10 IP v6: 5802:a7c7:f419:a64:614e:418a:77c4:dad4, cost 1.871µs
2022/01/23 22:33:10 UserAgent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:52.0) Gecko/20100101 Firefox/52.0, cost 681ns
2022/01/23 22:33:10 Password: C?s7O-]]F]a-0G[3I-XObM, cost 56.475µs
2022/01/23 22:33:10 Password Easy: SWsA6-kySCb-d655v-zHtC, cost 34.863µs
2022/01/23 22:33:10 Numbers: 250166123848045435121681614516459885310105082646722299461075, cost 13.619µs
2022/01/23 22:33:10 Letters: CmfEsTxaOMfSpeYYBIVGnreLBjTPhqsecWoTqHLTHMmlpRbRVSdnsHMFelIL, cost 2.531µs
2022/01/23 22:33:10 JSON: {"stratocrat":true,"Babiana":"stability","feeler":["hamble",false],"foreseer":null,"quinton":[null,null]}, cost 9.863µs
2022/01/23 22:33:10 String: JGEYHDaFlBTUQGHTEhNHPHOlSJYQSJVMbiVXYWKXSHJcCNTYgUWfIYNAOJLT, cost 989ns
2022/01/23 22:33:10 Captcha: dv5Nf /var/folders/8h/9150s4l16q5bbjq1j6cllvch0000gp/T/24667456/246Q7g8xTMy4L993NZoPLRO17iA.png, cost 4.737323ms
2022/01/23 22:33:10 KSUID: 246Q7kCMGsekxcuLj7tJksUxuUF (Time: 2022-01-23 22:33:10 +0800 CST, Timestamp: 242948390, Payload: E2B3BD706BEC8A422964D0F7F459497B) , cost 17.303µs
2022/01/23 22:33:10 UUID: 07222c05-36f2-426f-89de-d3e7626a67e1, cost 2.403µs
2022/01/23 22:33:10 Aidarkhanov Nano ID: ywXqIAQiF8OHvSCXnVtMf, cost 4.559µs
2022/01/23 22:33:10 Matoous Nano ID: i6jEQJcEc8GzB34lUSA1o, cost 2.038µs
2022/01/23 22:33:10 Mongodb Object ID: 61ed6726421aa9239a0a3d2e (Timestamp: 1642948390), cost 2.246µs
2022/01/23 22:33:10 Xid Mongo Object ID: c7mme9nm20ri76kd87l0 (Machine: 9hA3, Pid: 9114, Time: 2022-01-23 22:33:10 +0800 CST, Counter: 9257450), cost 4.406µs
2022/01/23 22:33:10 BSON Object ID: 61ed6726421aa9239a000001, cost 12.401µs
2022/01/23 22:33:10 Snowfake ID: 1485259315827707904 (Time: 1642948390334, Node: 1, Step:0), cost 2.427µs
2022/01/23 22:33:10 姓名: 钟奥鮉, cost 1.875µs
2022/01/23 22:33:10 性别: 男, cost 774ns
2022/01/23 22:33:10 地址: 西藏自治区拉萨市挂瞫路7736号鍩亂小区7单元1701室, cost 6.756µs
2022/01/23 22:33:10 手机: 15992702119, cost 1.377µs
2022/01/23 22:33:10 身份证: 630275201202034488, cost 14.437µs
2022/01/23 22:33:10 有效期: 20191230-20391230, cost 1.669µs
2022/01/23 22:33:10 发证机关: 丽江市公安局某某分局, cost 416ns
2022/01/23 22:33:10 邮箱: uqmkkvug@lijjq.vip, cost 821ns
2022/01/23 22:33:10 银行卡: 6217281420195002112, cost 3.159µs
2022/01/23 22:33:10 日期: 1972年12月18日, cost 1.073µs
2022/01/23 22:33:10 Generative art: Junas: /var/folders/8h/9150s4l16q5bbjq1j6cllvch0000gp/T/24667456/246Q7kI5MYd1xmAmrJVVClyPxEP.png, cost 14.250519ms
```

## captcha

`$ gg-rand -tag captcha -n 10` to generate 10 captcha images.

![image](https://user-images.githubusercontent.com/1940588/140700928-1fd794a7-21b2-4bda-81c0-b705bd632f88.png)

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
