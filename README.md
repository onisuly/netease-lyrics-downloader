# netease-lyrics-downloader
网易云音乐歌词下载器

功能描述：
通过扫描文件夹下所有指定扩展名的音频文件的元数据来获取歌曲名和艺术家，再使用网易云音乐的API搜索歌曲并下载歌词；因为使用的API只是模糊搜索，所以可能出现少量错误匹配或者匹配不上的情况

使用方法：
```shell
Usage of netease-lyrics-downloader:
  -dir string
        input dir (default ".")
  -extensions string
        music extensions (default "mp3,flac")
  -proxy string
        network proxy
  -thread int
        thread number (default 20)
  -timeout int
        timeout second (default 10)
```
使用方法很简单，第一个参数是包含歌曲文件的路径地址；第二个参数是要扫描的音频文件的扩展名，多个扩展名用英文逗号分割；第三个参数是代理地址，如果你需要使用代理的话；第四个参数是并发解析和下载线程数；第五个参数是请求超时的时间，单位是秒，根据你当前的网络质量来设定

