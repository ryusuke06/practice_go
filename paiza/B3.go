/*
PAIZAタウンでは、公共交通の手段としてタクシーが発達しています。
様々な料金のタクシーがあるため、初乗りが安くても最終的な運賃が高いタクシーを拾ってしまうかもしれません。

タクシーの運賃は初乗り距離と初乗り運賃、加算距離、加算運賃で決まります。
タクシーの乗車距離が初乗り距離未満の場合は、初乗り運賃だけで移動することができます。
初乗り距離と同じ距離だけ乗車した段階で、運賃に加算運賃が追加され、以後加算距離を移動する毎に加算運賃が追加されていきます。

あなたは今いる場所から X メートル離れた場所へ、 1 台のタクシーで移動しようとしています。
利用可能なタクシー N 台の料金のパラメータが与えられるので、タクシーで X メートル移動した際の最安値と最高値を計算してください。

例えば、 入力例 1 の場合は以下のように 600円 が最安値となり、 800円 が最高値となります。
1 番目のタクシーは初乗り距離以上なので加算距離 1 つ分が追加され 600円 かかります。
2 番目のタクシーは初乗り距離までで到着するので、初乗り運賃のみで 800円 かかります。

[入力される値]
N X
a_1 b_1 c_1 d_1
a_2 b_2 c_2 d_2
...
a_N b_N c_N d_N
・1 行目にタクシーの種類数 N、目的地までの距離 X がこの順に整数で半角スペース区切りで与えられます。
・2 行目から続く N 行には i 番目 (1 ≦ i ≦ N) のタクシーの 初乗り距離 a_i、 初乗り運賃 b_i、 加算距離 c_i、 加算運賃 d_i が整数でこの順にスペース区切りで与えられます。
・入力は合計 N + 1 行であり、最終行の末尾に改行が1つ入ります。
*/
package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
    sc.Scan()
    return sc.Text()
}

func betterTaxi(x int) (goodDeal int) {
    s := nextLine()
    arr := strings.Split(s, " ")
    start, _ := strconv.Atoi(arr[0])
    firstPayment, _ := strconv.Atoi(arr[1])
    addMeter, _ := strconv.Atoi(arr[2])
    addPayment, _ := strconv.Atoi(arr[3])

    // fmt.Println(start,firstPayment,addMeter,addPayment)
    switch {
    case x < start:
        goodDeal = firstPayment
    case x == start:
        goodDeal = firstPayment + addPayment
    case x > start:
        var roundUp int // int型は切り下げになる
        roundUp = (x - start) / addMeter + 1
        goodDeal = roundUp * addPayment + firstPayment
    }

    //fmt.Println(goodDeal)
    return
}

func main() {
    s := nextLine()
    arr := strings.Split(s, " ")
    n, _ := strconv.Atoi(string(arr[0]))
    x, _ := strconv.Atoi(string(arr[1]))

    //fmt.Println(n,x)

    highestDeal := betterTaxi(x)
    lowestDeal := highestDeal

    for i:=1;i<n;i++{
        payment := betterTaxi(x)
        switch {
        case payment >= highestDeal:
            highestDeal = payment
            //fallthroughでなぜかlowestDeal<=paymentが通ってしまう謎
        case payment <= lowestDeal:
            lowestDeal = payment
        }
    }

    fmt.Println(lowestDeal, highestDeal)
}