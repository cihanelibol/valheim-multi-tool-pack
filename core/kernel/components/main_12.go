package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.1" )

func fro6COHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXCmHl8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ql9Lpg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UcwFZ3rbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GrkHg27TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJUvsr0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5NZwvO8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTcgdlJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yAKoAPYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zO284jCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohHcK1OQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxqUQrEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIgRkmLrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUYtr1hHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3kr3xEYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqN8JyGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJViEXeqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Za7WJIijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbT9fTHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3CGFDD1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vThDUW1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IIPOQLO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRbKA1iLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zr35n5SKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wDQmoIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rP7pE48BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFgUx0XtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWIr8tFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRWz3EPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vrEAKNucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dA6P1opDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TuJLKntfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kA4l9Zr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjBlNFzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Us3fp2nHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3z5yQ8VxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dv35jFnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vBStuyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2FZ0rABmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yvdnG0l2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxECxAe2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I22Bz2RWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ipQldr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7s8VFY5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iXH3ewCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kchb24GFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcsMDC5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcCxkX7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mlIxehcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6hzMDBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DX3x6F6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHXQ9evhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PKjq2LLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWjdTrzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdpYU1CkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OvID6vnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAK2FP9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AUXK40iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOquzXxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dpxzlIucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWTbsdd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKTgdOKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWfqnuvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u1r7P9O7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dGDkunseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UhgbSGXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7W4OZugUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ctqmjDljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxBN77fkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CxaaKezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5wGlTRwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zratfWEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zbZd1N91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OvthnPZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRNf0ZgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y46qkxXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVmFGCB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z398JQCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPfCoQ1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cbSxXwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGEOFHy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIY5moF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QF3GX7EoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGWubP8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vHQVbuu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9doqQSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gM4TiXb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0ppRtCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FeuX9LyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LrX1v1pHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1xBMvBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fv1uNEAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHCVwKglWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4KUzamBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0osOqzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UXZcdwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MC8PJIuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glUmvPOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Bq201EAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzFoLdzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFvG8pkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYQPdhDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGp6SpHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSSdokgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yCLEXYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func un1EpzKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXxNuW1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6j4pcAhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AmiRosU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pGn5wsI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTR3ktNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7lgzLQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFsOFmIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XcCTvo5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5yY8qz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ra2h1Xl9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CIeLvx2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLIqNWJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rl5ILdrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CueHJtjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThSaMQT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K2pPhoftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ObKqF2xZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iXOgLNSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yt7jvv1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func irAP0uaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDGXnzNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3XZuzaHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VULl52fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSQxiGNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yplGdrGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwcDnhNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJMR91iOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v72GFVjMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jkZPXk7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udSFHpmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WihyyQpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DI2SMI6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4yyTgI3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fn8WntK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3nPv2cREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9OdqI9CRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fRax0xsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvd8aScWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kVDNtpdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpBxRvREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBRdMWbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HGG7xr3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mN8pD27XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGBWcIRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdXP6fRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwm2ulU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywnITA09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNibYFklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXUNnVPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2n8BLKWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QlKrHNs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouSI1yw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7yk37inWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2MSXtm52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1vIyNHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQ7HoV33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mNWOEO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gDJqow6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func scmlvAqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wqRl6VhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0f1sf7QdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJANqTWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BK1odwwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2raZPMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZlXuzKNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mvj5sdZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4sf7vm00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptTV2M6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7NPZDK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2wx5el1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1Qttmu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvMKwbqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUlLjeLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XeKX4u4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbZiWibjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfxfyGcMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4iXJksdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwqDyTNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgNWHKkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3HkFsK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwOWeojCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gazcfpEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZ51yRgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwWmrnoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRM5x542Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydg48hSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mvobYHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8aB7r9IuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hyLwp1YWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKJZ29roWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHJsRpjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISnUh4z0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mroXjzSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjKe6XpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUAr4r1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKsvVj7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QvbPxNcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3snzX9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blYzN6NfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ahUg9gXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONkm0ntUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func comnZlEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nXdajj0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToIrovvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oam8PPW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2oNP7KQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqLg3DSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V330JRGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEf4QqcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XN01f7frWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9WCRDNLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRA1uxM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XBo0aKTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEeEN52PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3qcTlr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0ozAHnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IualSxLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N69p0aPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tz4v6pa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVkeh9yVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjpfYKtBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKvM77PYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVQLBHGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLAIDknxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9Wv51OMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tk8qs6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmxky9AwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q08OuFtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLh1b2HtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bphOdXmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pn1U5bvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 162K2chPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdJT40ndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHNvqjLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjqcVsHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BczAnGRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2wmIajsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzwzL2mXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQj1rq4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdcceM4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsy6h0nAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOz6EaboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1upnI0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ER1NRJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLzuhyU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEUmKqeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHVqxrIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftXcyqiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5QM2KBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7HI0FLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBAbOD5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIZECiJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhfJs61gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sA10dIDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zIxloiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldyRtUVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mERvv5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmssXjBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GejLJiWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lv1PvtwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QvWGOxfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yvNtsojOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TrBNfTncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCCa2Ob0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhfqg0JhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfkbDUsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func crahhYMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbfCLMDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MwfRxXj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gf9HgRg0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xrzs5IA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NGMy7cnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TX9SeuXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKkMmXGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bSTutKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLLvjh7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92Bhn9EJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7PSslEWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8Gv32t7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v4cgdECrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMhdUh6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxN6FgjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3hcDdyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7jRW7w2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uTx5AFnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9LaCoGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2yxjTVYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ktp3fANFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwKAPhLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pa2uuu68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoxht0cYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxKn0rXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNM9VFbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VaiX2MSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4xNlwpSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NptqMXkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5wO6syNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZRgcLHdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vtfitz6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPDUL8wqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8FD1KnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXxB57sYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaCjsqftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2047MCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7cLAZlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRHQwynkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQHag7kUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YZWjNiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uBUCqnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAVb1Ak7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QU07sO4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUL6KycvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gikyXm6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHyShOQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBRBmXATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUlauXrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ovhn040QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kq16M59fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLKawXT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXAwjGGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFGLoQF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ks6T1pnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RqGTqEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSXiFu7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IuA8uKI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9O5hDeEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mtZWx4VuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJbZqn34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjOHA2pAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzTvc5cXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b55aiCuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJ64MHxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xaIdx05lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tYxNrwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjBcReNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HkPV6XKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9bbDN1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHAB2hWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YpmcXfNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2CTH50bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDiIuhbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Yx9MhQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lpMRVwuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIZwGMtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZsOezvJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYkGmtfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VtUxBgt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpzNVOYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCo6BOCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tr6IUw4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFWWxi5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFv24bn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CX5AtglBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owQJcwNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 808mOAVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdKcwOtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ObxjykaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qf5QRLGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7kujLmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUDVhjpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvU3ONE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MO0Ao4qGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PY5PmESlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YarBMD3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bY55Pj6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TsoYyDwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1LZviV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xL76qZn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u0o3irVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCSKZDSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lah5knMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsidAc8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwByzV3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39Ok5kBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EIAyEeiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6fr4wWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TwhUlUrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Afj1nY7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ng9dNOYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o03HMVhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13lz183HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQsOUAWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qzwgF0KQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWolhn76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hmRR1NCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etF9PIspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6GR2PCK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3f5aAHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dIz8hssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xu0gjBriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8UJZDqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SoVGFm3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D75BlAPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Jogsf53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mMZkiqmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocvpWpKWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fBtZLgTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSdndq6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93yiNp6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlcpiCxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBa1NGdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uic63p0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vet7Z5xYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5jqy19iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ns73w5g0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEIZqlQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y84aCGoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bRW8H8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3kciM06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYpkxIZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNiSkc6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KVA2iEE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v820n6ubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 218WvSVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tytIk7S3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbQYTBaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlmBtPulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5iWi9zjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59zOj07ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nGOxcW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8pQ9qnkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dB2brU1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5Te6NrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7B9sCW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkugTqmUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Freutxz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ait31MfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7e163jW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToOPUr7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAJDEZHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GMIMwWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EU9UnOhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VswZpl1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qg8oqm9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yN4hok7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VA2WeocaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOZxYDHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0UslS2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3haUizejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRmy5c4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFivXUwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4RriOnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OTYyxemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTXAXaQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EtguAhDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rupw0yDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SaQvpaStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNBolCHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DyWZKsSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXJTTZw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ot7u5uZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjGahWflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbJtybn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func we3cPGGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KgUqknbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdQUfpUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DwBBZKFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUCBBIiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7sILgYMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNOsW3NXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bYYxBlfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func huiUzb9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nh9ydgc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqVZT7YXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69j14kGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tb3CXuOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLzKH7PmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mGbQ9eE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfdq20UMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVCwyGqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mq0GRdZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ePAy6rMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71aTL65TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmJVblC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 77tyMmuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RncYCHyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func or9EMDjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eECsXk2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBECYZ3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkURqAOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AdAKASTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

