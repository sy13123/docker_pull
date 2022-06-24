package timetool

import (
	"fmt"
	"time"
	"regexp"
	"go_pull/pkgs/util/logtool"
)


func Timetostr() time.Time {
	return time.Now()
}

func Strtorime(st string,location string) time.Time {
	if location==""{
		location = "Asia/Shanghai"
	}
	_t := "2006-01-02 15:04:05"
	r, _ := regexp.Compile(`T.*Z`)
	if r.MatchString(st){
		_t = "2006-01-02T15:04:05Z"
	}
	loc, _ := time.LoadLocation(location)
	tt, _ := time.ParseInLocation(_t, st, loc)
	return tt
}


func Time_add(m time.Time, _s string) time.Time {
	_t, err := time.ParseDuration(_s)
	logtool.Fatalerror(err)
	return  m.Add(_t)
}


func test_gettime() {
	//获取当前时间
	t := time.Now() //2018-07-11 15:07:51.8858085 +0800 CST m=+0.004000001
	fmt.Println(t)
  
    // 当前时间加三秒
	//m, _ := time.ParseDuration("300s")
    a := 300
	m1 := t.Add(time.Duration(a) * time.Second).Format("2006-01-02 15:04:05")
	
	fmt.Println(m1)

	m2 := t.Add(-300 * time.Second).Format("2006-01-02 15:04:05")
	
	fmt.Println(m2)

	
//	//获取当前时间戳10位
//	fmt.Println(t.Unix()) //1531293019
//    //获取当前时间戳13位
//	fmt.Println(t.UnixNano()/ 1e6)
//	//获得当前的时间
//	fmt.Println(t.Format("2006-01-02 15:04:05"))  //2018-7-15 15:23:00
//    
//	fmt.Println(time.Now().Local().GoString())
//
//	//时间 to 时间戳
//	loc, _ := time.LoadLocation("Asia/Shanghai")        //设置时区
//	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-07-11 15:07:51", loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
//	fmt.Println(tt.Unix())                             //1531292871
//  
//	//时间戳 to 时间
//	tm := time.Unix(1531293019, 0)
//	fmt.Println(tm.Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19
//  
//
   b:=Strtorime("2022-06-09T00:56:33.372346283Z","UTC")
   c:=Strtorime("2022-06-09T00:56:32.372346283Z","UTC")
   d:=time.Now().UTC() 

   fmt.Println(c.Before(b))
   fmt.Println(b.Before(c))
   fmt.Println(d.After(b))
   
   fmt.Println(d.Format("2006-01-02 15:04:05") )
//	//fmt.Println(b)
//    
//
//
//	//获取当前年月日,时分秒
//	//y := t.Year()                 //年
//	//m := t.Month()                //月
//	//d := t.Day()                  //日
//	//h := t.Hour()                 //小时
//	//i := t.Minute()               //分钟
//	//s := t.Second()               //秒
//	//fmt.Println(y, m, d, h, i, s) //2018 July 11 15 24 59
 }
 
