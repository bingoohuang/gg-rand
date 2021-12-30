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
Base64Std               : s2HWywPO3Jybzi0UhEPoOZ15q8z4P15eTwjAhLGN5c80tPMn35AUf6zDIppOsymx5efWI7hi09WtCCZthr3vS5UMeOfwSCVDCxHQ4Uh2DSTK6jysPgM+r7UYSxgxEIywxqF7WQ==
Base64RawStd            : d8QXBl25KifzDTQZQ9MEw4n0bVTxTCNNToOpNrEQSwPp91Vz1RL2l4f8dTAKroYdqmk2ctc19UZVccXu146BAqN6n6hSGLbb790T+dbkmWtarARIzfkhjimO09zdArfomNY1rg
Base64RawURL            : bKI0w1H8b_FldxsjWFgJAOMmsRuEK4jHPJZnG7KASZU5pTT5Q7zJfuRqGPnLNyVjj5EUZSvBChh6F9-O-vTuy1OrScbcq5GYH2p2UrI-9K7jOhV-Lw8eeXy5Ghq2R6vPSnZCEQ
Base64URL               : VbvwFHSnTiCR5Moif4jmUrfYmfk4uYYN6z7TT9_JxiOUuvdjnPlHKnrmyCPmB8Uc47tQ-pB0AzIc8TjxJV2U99sZ_fx0cza_TJjzBJ6Eq0WMLyYI8ZDbplyGnvTsVb4uskq6PA==
SillyName               : Hidequilt
Email                   : sofiathompson531@test.net
IP v4                   : 25.170.114.230
IP v6                   : 805c:2212:f43:14d7:7c43:c136:fa6f:b9f1
UserAgent               : Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36
Password                : 9Oi{g-6hG4z-U$168-$Dp3
Password Easy           : grS62-qCHFq-3ed9K-JStK
Numbers                 : 539166067412416667334954974546947515445722139493153830556092
Letters                 : IIdnAocQuIpBSvyGAixjPvTQmtCXOtclhGOwsTARoPPbNgoGVZingkwuEppp
JSON                    : {"reattract":{"humidity":{"kirk":null},"tricksily":{"dreng":"outstrut"}},"unmind":null,"larch":true,"melliferous":true,"corrupting":[]}
String                  : MUQBNXJLXVKkUMYclPQPFGYOZFPPXAGZkZCUUkRUSgjGbCYZciYBDVPRDYTa
Captcha                 : Nkc8B /var/folders/c8/ft7qp47d6lj5579gmyflxbr80000gn/T/4016527990/22zlxsC324MLy0EK3mPFbpEiKpz.png
KSUID                   : 22zlxtPor8gLqh2Jc3FTDRpJIWw (Time: 2021-12-30 15:15:45 +0800 CST, Timestamp: 240848545, Payload: D904B7E413BA058889853B246519703A) 
UUID                    : 4a64d2b2-8931-4eb4-8324-463cbdf42cbd
Aidarkhanov Nano ID     : n4H-MgPP4F72oM-0kPPaj
Matoous Nano ID         : sXKjs8KQJ92W7duUACO7-
Mongodb Object ID       : 61cd5ca1c3666e79e0bcec93 (Timestamp: 1640848545)
Xid Mongo Object ID     : c76lp8dmk1u7jo4asln0 (Machine: tqB8, Pid: 31200, Time: 2021-12-30 15:15:45 +0800 CST, Counter: 9102702)
BSON Object ID          : 61cd5ca1c3666e79e0000001
Snowfake ID             : 1476451927121203200 (Time: 1640848545233, Node: 1, Step:0)
姓名                    : 方壆娋
性别                    : 女
地址                    : 山东省济南市銦翶路3464号渏悫小区4单元1840室
手机                    : 13890151323
身份证                  : 507814197608039922
有效期                  : 20061120-20261120
发证机关                        : 大同市公安局某某分局
邮箱                    : eodytjij@pural.net
银行卡                  : 6225684849954910204
日期                    : 1991年04月01日
Generative art  : Junas: /var/folders/c8/ft7qp47d6lj5579gmyflxbr80000gn/T/4016527990/22zlxsfODBaBZK4OjGsMewo0RtA.png
Unsplash        : /var/folders/c8/ft7qp47d6lj5579gmyflxbr80000gn/T/4016527990/22zlyakecXFevHMoUC3vFMqDehv.png
Image           : 16.8KiB 1928690884_640x320.jpg
```

## captcha

`$ gg-rand -tag captcha -n 10` to generate 10 captcha images.

![image](https://user-images.githubusercontent.com/1940588/140700928-1fd794a7-21b2-4bda-81c0-b705bd632f88.png)

## images

`$ GG_IMG_SIZE=640X320 GG_IMG_FORMAT=JPG GG_IMG_FAST=N GG_IMG_FILE_SIZE=2MiB  gg-rand -tag image -n 10`
