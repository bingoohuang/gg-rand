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
SillyName       : Knavesaber
Email           : andrewanderson864@test.com
IP v4           : 223.247.194.2
IP v6           : d2d7:e355:b1e5:7f27:2f05:394a:46d8:a116
UserAgent       : Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Safari/602.1.50
Password        : uZq4f-cxlwT-5FMwg-hXGO
Password Easy   : udsqq-8gJ3N-S52Rr-kbSM
Numbers         : 226575933607625554263151723594611489412507264130487465203199
Letters         : qNOPGgeqKMmYCLqgLrSbmzRVjDYcdsCUmQqJlIxAdDoYUrrzpuzLUXyUmxIM
JSON            : {"bumpy":{},"Adelops":null,"columellate":33.94,"nonplacental":true,"chloroform":true}
String          : WRgQVhLPCPFDhjXYKlHVXUkYXiBIADdjUQNJVVjeROWHIMQADcYYllBPQMGK
Captcha         : A5wzI /var/folders/8h/9150s4l16q5bbjq1j6cllvch0000gp/T/126066539/rand210760822.png
KSUID           : 20ZrviFHnDPlqBIApHXZpMbLKPV (Time: 2021-11-07 13:33:02 +0800 CST, Timestamp: 236263182, Payload: D2438FB3B450995ED6231526352D56A5) 
UUID            : 2d0c64c6-3f8c-11ec-be1e-acde48001122
姓名            : 苗蝰坐
性别            : 男
地址            : 四川省南充市諺婘路6889号镒癛小区12单元2496室
手机            : 18956311131
身份证          : 110870197802037914
有效期          : 19880901-20080901
发证机关                : 杭州市公安局某某分局
邮箱            : iefccafq@crbbe.com.cn
银行卡          : 6229351116740827553
日期            : 2007年01月31日
Generative art Canva    : /var/folders/8h/9150s4l16q5bbjq1j6cllvch0000gp/T/1161453087/rand3133874368.png
Unsplash                : /var/folders/8h/9150s4l16q5bbjq1j6cllvch0000gp/T/1161453087/rand691563348.png
```

## captcha

`$ gg-rand -tag captcha -n 10` to generate 10 captcha images.

![image](https://user-images.githubusercontent.com/1940588/140700928-1fd794a7-21b2-4bda-81c0-b705bd632f88.png)

## images

`$ GG_IMG_SIZE=640X320 GG_IMG_FORMAT=JPG GG_IMG_FAST=N GG_IMG_FILE_SIZE=2MiB  gg-rand -tag image -n 10`
