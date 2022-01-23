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
2022/01/23 18:36:17 blake3hash-zeebo: sH7FuTntW4R-4H49ny7SpbsS9RODOS0XfH4-L9SGYmc
2022/01/23 18:36:18 blake3hash-luke: 8cZHx4Ed094YEz6kDtv8obqjjSwK1D-xRGyzuyQyQz5I-xDQpKIp--TVUQW6NW3yxwjzbZFMhYCDb2b-Xh6SnZ7fMb4ud9fNRk1HPKSEZMD9cpVBDvU9I8voh-d2SpP7rAV4V4dKsALm4rnPjdFUVcKCC0tcYsTxbcVNU5Gb2Eid9vIgOP82muuBK_XOVVhGz0YRvjbN2kE_0vLm4ffz4NRAARfj2e0L3dW6lmqp15I6sHi6NfK2_k_EzpoFdgq03vmDhBhWmV7_ykgU11Gm72Jt-kdCW0qfzZ_MxmIxzyPel_Qg87up5IZ4UCxlskgzE77UKu-nJD9M6Q-jzl6LgQ
2022/01/23 18:36:18 xxhash: 1gLJAfAMyzY
2022/01/23 18:36:18 md5-hash: nqxuBva_ESrytDrllC7HYA
2022/01/23 18:36:18 sha256-hash: 8fZDKRcsu_7diaFKQ_w93gP1CLHdRu5PvamDT4JIqO0
2022/01/23 18:36:18 imo-hash: gICABbgQ-JzBSNwYdFb9TQ
2022/01/23 18:36:18 Base64Std: XJzNIIOiV8Gqya5J5BtIruoovDJvGKlZO8OeGkFODXN4aR+8vRt0RnLUS7fCjN1Nv6ywgSFb+sbd6PXV5PXlHjKAHZPCjy9nHJHqjdbwCMRTybiemUpTn20+8a1Z0tQ+GrSnoQ==
2022/01/23 18:36:18 Base64RawStd: PQb7Grt5YIrVPS+GnD5IZCTEg9s8AcbvkwducmLqlD6PunOPCSVGJvghkoStFIDfIVThNjewwWiITEr6mtGi6smWr8oKa8gZeq1uXorBXmttx5DjAtdLtzxrqCmYqUdluWl+nA
2022/01/23 18:36:18 Base64RawURL: 0Qy5cEjIw5tkuIDRBJkXy9Eygj3LKJzD_3GWwt9O1LaGFc9Ik9iqHXVJjf-F3qVNxZOz8FIvJgmFZjlT1qNOMhOwrl0MPr9p424VyfVuRRpYoRI6US1TK3uv55YKVWIGf6oC2A
2022/01/23 18:36:18 Base64URL: uzI06fsX_t2h4EsMYSChmXXa6isev-u68sdbWkUFLk99U6NfJCOeveF54JNJOPQxzsWIHNyuYHDWccCW0wtb8lWtgQwCZt2lRP5pbkPfDhYU7KpwSDP8CPArrkjuTCVHS4Kk9g==
2022/01/23 18:36:18 SillyName: Owllava
2022/01/23 18:36:18 Email: jaydenrobinson465@example.org
2022/01/23 18:36:18 IP v4: 184.14.214.141
2022/01/23 18:36:18 IP v6: 200c:934:349d:8d83:5ce0:6968:f3bd:5a9b
2022/01/23 18:36:18 UserAgent: Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/28.0.1469.0 Safari/537.36
2022/01/23 18:36:18 Password: CZ[AB-)Vp$O-2uv9v-}CQ5
2022/01/23 18:36:18 Password Easy: kvWEm-f7g58-kz9aA-ecjr
2022/01/23 18:36:18 Numbers: 838889063003503321340946875831892589355399053895788007094338
2022/01/23 18:36:18 Letters: WjEUPmfOLWDnvYHLjTeYSZVilYFRxzyLzDXVpAghovBGDDxgKvNUIvleGnZg
2022/01/23 18:36:18 JSON: {"bitterful":[],"simool":null,"rememberably":null,"reagent":false,"cantoner":null}
2022/01/23 18:36:18 String: CBJKSbdJFKZAVNIkMSNXKkAXRRUQHBFUKjWWICkYQQEDhbIjQZJPPXLWSjIS
2022/01/23 18:36:18 Captcha: YphdK /var/folders/8h/9150s4l16q5bbjq1j6cllvch0000gp/T/3172545281/245xJk4JuHszyUQ21vrlSdfwzDW.png
2022/01/23 18:36:18 KSUID: 245xJgU15IDbcF5FH24Cj2AWGvf (Time: 2022-01-23 18:36:18 +0800 CST, Timestamp: 242934178, Payload: 474AF519DBF343F259EEC1560790C6D7) 
2022/01/23 18:36:18 UUID: 0eee0de7-769a-4854-812f-109f89d33e19
2022/01/23 18:36:18 Aidarkhanov Nano ID: VQaGChORcl5XnvxynLbSN
2022/01/23 18:36:18 Matoous Nano ID: sA8xMuxGQxM3ejC2hYDZb
2022/01/23 18:36:18 Mongodb Object ID: 61ed2fa2c3666e0b596e54aa (Timestamp: 1642934178)
2022/01/23 18:36:18 Xid Mongo Object ID: c7miv8nm20rgmmbmje60 (Machine: 9hA3, Pid: 2905, Time: 2022-01-23 18:36:18 +0800 CST, Counter: 7773068)
2022/01/23 18:36:18 BSON Object ID: 61ed2fa2c3666e0b59000001
2022/01/23 18:36:18 Snowfake ID: 1485199706316345344 (Time: 1642934178319, Node: 1, Step:0)
2022/01/23 18:36:18 姓名: 宗玤鴀
2022/01/23 18:36:18 性别: 女
2022/01/23 18:36:18 地址: 云南省昆明市腪媦路3370号泙蚕小区16单元1588室
2022/01/23 18:36:18 手机: 17087415977
2022/01/23 18:36:18 身份证: 424784201901307375
2022/01/23 18:36:18 有效期: 20050910-20250910
2022/01/23 18:36:18 发证机关: 成都市公安局某某分局
2022/01/23 18:36:18 邮箱: ntcupnwz@vgdvk.com.cn
2022/01/23 18:36:18 银行卡: 6223242660845284299
2022/01/23 18:36:18 日期: 1988年12月10日
2022/01/23 18:36:18 Generative art: Junas: /var/folders/8h/9150s4l16q5bbjq1j6cllvch0000gp/T/3172545281/245xJkJKGN4TH0fsIhKGreAVYOu.png
2022/01/23 18:36:19 Unsplash: /var/folders/8h/9150s4l16q5bbjq1j6cllvch0000gp/T/3172545281/245xJqkpC7MW4y7yuprKykOiaqK.png
2022/01/23 18:36:19 Image: 16.9KiB 2104742387_640x320.jpg
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
