/*
Paiza 国の選挙シーズンがやってきました。 今回の選挙では M 人の立候補者と N 人の有権者がいます。
Paiza 国の人々は政治に無関心なので、最初、すべての有権者はどの立候補者も支持していません。

Paiza 国には大きな広場が一つあります。 選挙活動の期間中、立候補者たちはこの広場で一人ずつ演説をします。
（同じ人が複数回演説することもあれば、1 回も演説しないこともあります。）

演説が終わるたびに、「他のそれぞれの立候補者の支持者から 1 人ずつ」および「誰も支持していない有権者から 1 人」が演説をした人を支持するようになります。

// 入力例
2 100 4
2
2
2
1
// 出力例
1
2
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

func winner(m int, n int, k int) (winner []int) {
    candidate := make([]int, m)
    for i:=0;i<k;i++{
        vote, _ := strconv.Atoi(nextLine())
        vote -= 1

        voter := n

        if voter > 0 {
            candidate[vote] += 1
            voter -= 1
        }

        for j:=0;j<m;j++{
            if candidate[j] > 0 {
                candidate[vote] += 1
                candidate[j] -= 1
            }
        }
    }

    winner = make([]int, 1) // 素のままだとlen=0,cap=0になる謎(naked returnのときは多分こうなる仕様？)
    winRate := 0
    for x:=0;x<m;x++{
        // fmt.Printf("nowWinner=%v,vote=%v,candidate=%v\n",winner[0],candidate[x],x+1)
        if winRate < candidate[x] {
            winner[0] = x+1
            winRate = candidate[x]
        }
    }


    for y:=0;y<m;y++{
        if winRate == candidate[y] && y+1 != winner[0] {
            winner = append(winner,y+1)
        }
    }
    // fmt.Printf("len=%v,cap=%v,%v\n",len(winner),cap(winner),winner)
    return
}

func main() {
    s := nextLine()
    // s = strings.Replace(s, " ", "",-1) "2 100 3"とかのときにm=2 n=1 k=0とかになってしまうw
    arr := strings.Split(s, " ")
    m, _ := strconv.Atoi(string(arr[0]))
    n, _ := strconv.Atoi(string(arr[1]))
    k, _ := strconv.Atoi(string(arr[2]))

    // fmt.Printf("candidate=%v%T,voter=%v%T,voteNum=%v%T\n", m, m, n, n, k, k)

    for _, v := range winner(m, n, k) {
        fmt.Println(v)
    }
}