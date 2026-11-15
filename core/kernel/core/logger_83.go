package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.5" )

func CWTNRCMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxa4pBQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUbYu6HqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOXKaQYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c074H16PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0EgiWQJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YU3C0YbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rs6eYF3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEjGhkhnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCYAlZtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfKbxxGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v90EixcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAga5VEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAgkhZzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fuKEmLBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iDd5G99OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOZfuEIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkHfhBhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WntbKk0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trrFCs06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRI2OVMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3tkzkuKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u675HovfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJt7kCAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YVeTi813Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCCPHptZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1EeJsTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36GffkkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAXMkEkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func brJGptnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZnoMF952Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCDszVzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XrDMWSDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RvFlnVUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMDUXnMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2U29LQ2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHxk2ONMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWpyFLKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WnD7WTcJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 00kzpKqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qeuOCGmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJIuXrYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKeFh0CrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func alPC7xM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zazI59ueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTpB9dwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yIvFNmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n3PhpHB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPmeDPQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MKngv63mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyW7pvvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nXdBJ7HIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvfpkkTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57cFDFxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MUS1EQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8InXD1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drFmL8xAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2g0ZJTz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hyRXxYzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QO4doaIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNYmLtJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5Umyr1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sieZSYcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zFhqrInDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0jMKUUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H2XbCPXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9OCEOampWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVWoOl9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9IIxwK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func okE8sWg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljgqAlvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsZkMWaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6v3uUwz3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbz0n1SeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnbKZEC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nd2H1DChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q60Ef0wiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXRrjVCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7DyyoJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 32NjC4vrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHfa06EvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcDjAmF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lHH08lcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJ4rmsrhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwvwIZlnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func arHoKCC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JM5IpteYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMZuXFHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0UpXLVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ePOpCd8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func phBr6VXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86hSbtfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJVLiHvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vr4QRSW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4bP6xxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTUwmCzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8UPcE62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haIoqa7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkS7xXyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MoBEcE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFCNoWTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYrnB6MBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RAspqWanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4SQKk9R2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbnIuGvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKoSOLVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWtnA00dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1FY0Y77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PfGne1BYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoxE33RHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wCvYFdAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdCrrErEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rLq6PdAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhbMvS0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oStW3YmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLzhRYylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NY4uqY3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3BDbGheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6TJKEsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogkv7buaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FKqjxcPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAOFgUm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19Wk8vI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qByh3K8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t82D7NzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EELzqbCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TmvGjLSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwH7chlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8RTsP2tgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgJHlZHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mP0YZnq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gl2Fm4GfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4rz5nCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwWJD4LJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPDfIR3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8FG54aVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eMXvM8tdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ee8CbyGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfrinclxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kKaNsm7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2I9q07oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7tqH0U4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ontROovLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7AAn3XaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jNx0osIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbbFnes0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtBrZom3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ATuByPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func feCXVZ1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RcQoabsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgXgFpHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OyZJlKwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfI37bZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TO0MofuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fqgq9w1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwOBHMjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJQ58A0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbq98YISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWqS6CpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GcQOA9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xE0oq3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EhrnuVeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xq0rKlmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWcMzE9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNxcWvW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYy7QHwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0D2fIbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQolsZQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OvqCwK1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wHGoG51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KDb9BRGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRgBOHYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kD2tmu38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3VDMr62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3IIR87pEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gaqhc81kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1yoSmtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXQVDQwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYuHJVBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXAIhgsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXPJqJDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ShFAaIAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uf9ar9KPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXx67GOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rRLnvQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esHAlJNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func irXior7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQWKmAdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvdUM3cSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AOnWZj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gE1fHwxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHiEWc23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lB32azw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fTSU39alWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTvzZXVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ly58PTeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nHJppsbFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70eCY3UkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IcnM9d49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hlIRrW51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8z9bbQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9xqVKAyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2bs8H84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmcMvIlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJa50wGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XbVZ1W9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKRf3RLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qjAqds0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ndw87EAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zIMTc3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2w3TfOYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dHOPSmUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImCuFQuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2Lei2w1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmQUqWEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AT5SvfWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDit4c5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7lLx4SlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func km9fJQPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAys0JBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vcDXkYCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNc20GmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PB9DUUaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QCVxbJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NamtlU4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func panL7mKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BkUn70LnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OvD3Dm7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPD9chTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5l9LxNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sq4CeaiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQX5wxEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ENLalPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpOiv2bdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYXzgeoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCo4ASErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyhCFk0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucMV3qIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2ZUhWIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7PKCt81OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJVynOB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNFADwkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Knx941IXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kr6FYKb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9v6rZjWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6k1TggsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNYd8RNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtvxK7UhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qbqhf2JlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBEfV3jpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yWqR5iAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQMfp0HIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efxuBWAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swyMSkhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4GfuYziCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSHWHwk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qaFmgssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func deU9RzvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZVvIwKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGMVtnPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygyqyAbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQ4brtrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAXNiWnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QkUlp1d1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5PMzPsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYZ0KWkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PETEgCm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYWT4MA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vo8vwWK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 73PmA3X8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f3qUnFhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBCQn86cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYSkR8Y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Xwq8EjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8XdYkM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jS6VNZDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6Q64v8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6r3eBe9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VD4DbIpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qcqq0tBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8H9CQBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3s2D4HAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0TCTO2D9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aFYkppYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1T7Gb19MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWoIhz8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwLggI3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jkQONlLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ny1OzNOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8RQ4OUJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mh5D6oXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ytVK1imtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLtYhHbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

