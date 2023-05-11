# ac-snapshot

Go の練習として作成した API です。AtCoder 上のレーティングや精進状況をいつでも自分のスナップショットとして保存でき、自分のスナップショット一覧を取得できます。

スナップショットは次の情報を含みます。
- AtCoder のユーザ名
- レーティング ([atcoder.jp](https://atcoder.jp) から取得)
- AC 数 (AtCoder Problems API から取得)
- Rated Point Sum (AtCoder Problems API から取得)
- スナップショットの作成日時

## 目的

色変記事ではよく、記録目的を兼ねて AtCoder Problems のスクリーンショットが載せられます。また、rating-history.herokuapp.com (現在は停止) から取得した各コンテストサイトでの AC 数を定期的にツイートして記録する競プロerもよく見かけました。こういった記録をより手軽に、またより詳細に行えれば、自身の精進状況の把握やモチベーションの維持に役立つと考えました。

## 実行方法

```sh
$ go run .
```

## API 仕様

**GET** `/user/{username}`
* 概要：スナップショットを timestamp の降順に返す。
* パラメータ：
  * username (string, required): AtCoderのユーザ名
* レスポンス例 (200)： ```json
  [
    {
      "name": "shino16",
      "rating": 2177,
      "ac_count": 2365,
      "rps": 64300,
      "timestamp": "2023-05-11T19:09:02Z"
    }
  ]
```

**POST** `/user/{username}`
* 概要：スナップショットを作成する。
* パラメータ：
  * username (string, required): AtCoderのユーザ名
* レスポンス例 (200)：```json
  {
    "name": "shino16",
    "rating": 2177,
    "ac_count": 2365,
    "rps": 64300,
    "timestamp": "2023-05-11T19:09:02Z"
  }
```
