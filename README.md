# LeetCode 之 Golang 题解

该项目的初衷是用于个人算法学习。算法题自[leetcode-cn](https://leetcode-cn.com)获取。目标使用go语言解题。

[在线查看](https://leetcode.xulingming.cn)

## 生成器
```shell
go build -o generator generator/*.go
./generator  --config <配置文件> --output <题库输出目录>
 ```
参数:
>
    --config    指定不同语言题库GQL数据端点
    --output   题库md文件输出目录
    --lang      指定算法题语言,默认zh使用中文源
    --title       默认true将在MD文件开始添加特定头

## 题解
位于./solutions