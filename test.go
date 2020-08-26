package main

import (
	"fmt"
	"os"
)

type ByteSize float64

func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}

// ————————————————
// 原文作者：Go &#25216;&#26415;&#35770;&#22363;&#25991;&#26723;&#65306;&#12298;&#39640;&#25928;&#30340; Go &#32534;&#31243; Effective Go&#65288;2020&#65289;&#12299;
// 转自链接：https://learnku.com/docs/effective-go/2020/initialization/6244
// &#29256;&#26435;&#22768;&#26126;&#65306;&#32763;&#35793;&#25991;&#26723;&#33879;&#20316;&#26435;&#24402;&#35793;&#32773;&#21644; LearnKu &#31038;&#21306;&#25152;&#26377;&#12290;&#36716;&#36733;&#35831;&#20445;&#30041;&#21407;&#25991;&#38142;&#25509;
const (
    _           = iota // 通过赋予空白标识符来忽略第一个值
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func main(){
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))


fmt.Printf("%s",KB)
}