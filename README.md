# lottery

根据最近若干次开奖结果出现次数最多的球来预测下次的开奖结果

## 交互模式
取最近 n 次开奖结果，按照出现次数最多的数字来预测，n <= 100, 默认30  
`lottery n `  

 打印最近 n 次的开奖结果，n 只能取如下值 30, 50, 100  
`lottery data n`


## 后台模式
自定义时间，自动推送预测号码到飞书  

`lottery daemon`

，此模式需要`config.json`配置,  其中`web_hook_url`是飞书机器人的WebHookURL, 详见[飞书文档](https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN)

```jsonc
{
  "web_hook_url": "https://open.feishu.cn/open-apis/bot/v2/hook/xxxxxx",
  "predict_nums": 30, // 结果预测基于最近 30 次的开奖结果
  "day_of_week": [2, 4, 0], // 每周二，四，日发送通知
  "time": "18:00" // 发送通知的时间为18:00
}
```
