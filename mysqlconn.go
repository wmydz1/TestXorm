package main
import (
    _"github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "fmt"
    "net/http"
    "time"
    "encoding/json"
    "strconv"
)
var engine *xorm.Engine
var output string
var NameData string

var page int64
var pageSize int64


type Xorm struct {
    Id   int64 `pk`
    Name string  `xorm:"varchar(25) notnull unique 'usr_name'"`
    CreatedAt time.Time `xorm:"created"`
    ImageUrl string `xorm:"varchar(255) notnull"`
}
func initHelper() {
    var err error
    engine, err =xorm.NewEngine("mysql", "root:chen1994@/test?charset=utf8")
    if err!=nil {
        fmt.Println(err)
    }
    isTrue, _ := engine.IsTableExist("xorm")
    if !isTrue {
        engine.CreateTables(&Xorm{})
    }

}
func insert(name string, url string) {
    initHelper()
    user := new(Xorm)
    user.Name=name
    user.ImageUrl =url
    user.CreatedAt=time.Now()
    affect, err := engine.Insert(user)
    if err!=nil {
        fmt.Println(err.Error())
    }
    fmt.Println(affect)

}



// handle the data from browser


func HandleData(w http.ResponseWriter, r *http.Request) {
    if r.Method =="GET" {
        r.ParseForm()
        if len(r.Form)>0 {
            if len(r.Form["page"])>0 {
                if r.Form["page"][0]!="" {
                    mpage, err := strconv.ParseInt(r.Form["page"][0], 10, 64)
                    if err!=nil {
                        fmt.Println(err.Error())
                    }

                    page=mpage

                }

            }
            if len(r.Form["pageSize"])>0 {
                if r.Form["pageSize"][0]!="" {
                    mpageSize, err := strconv.ParseInt(r.Form["pageSize"][0], 10, 64)
                    if err!=nil {
                        fmt.Println(err.Error())
                    }

                    //                    fmt.Println(mpageSize)
                    pageSize=mpageSize

                }

            }
        }
        if page!=0&&pageSize!=0 {
            find(page, pageSize)
        }

        page=0
        pageSize=0
    }
    if len(output)>0{
        w.Write([]byte(output))
    }

}

func find(mypage int64, mypageSize int64) {
    fmt.Println("page", mypage)
    fmt.Println("pageSize", mypageSize)

    //query the table
    initHelper()

    // 1 2 3
    // 11 21
    //pageSize=13
    userlist := make([]Xorm, 0)
    var i int64

    for i=((mypage-1)*mypageSize)+1; i<(((mypage-1)*mypageSize)+1)+pageSize; i++ {
        var user Xorm
        engine.Id(i).Get(&user)
        if user.Name !=""{
            userlist= append(userlist, user)
        }

    }


    str, _ := json.Marshal(userlist)
    output=string(str)
    fmt.Println(output)

}

func main() {
    /*
        lang := []string{"Java", "Golang", "PHP", ".Net", "C#", "C++", "C", "Python", "Google", "Facebook", "D", "Nim",
            "Snapchat", "Wechat", "QQ", "Oracle", "IBM", "Mirosoft", "Apple", "Gmail", "Hotmail", "Twitter", "Yii", "Baidu", "Sina",
            "Lenovo", "Acer", "XiaoMi", "iPhone", "Taobao", "Huawei", "Macbook", "Google Glass", "Contana", "Logoocc",
        }

        img_list := []string{"http://img01.taobaocdn.com/imgextra/i1/63297854/TB1.WWEHFXXXXXnXVXXXXXXXXXX_!!63297854-0-tstar.jpg_620x10000.jpg",
            "http://img01.taobaocdn.com/imgextra/i1/1128690374/TB1lrhHHFXXXXauXXXXXXXXXXXX_!!1128690374-0-tstar.jpg_620x10000.jpg",
            "http://img04.taobaocdn.com/imgextra/i4/1128690374/TB1hq0BHFXXXXb0XpXXXXXXXXXX_!!1128690374-0-tstar.jpg",
            "http://img.taobaocdn.com/imgextra/http://img03.taobaocdn.com/imgextra/i3/762004848/T2vL7WXa4XXXXXXXXX_!!762004848.jpg",
            "http://img.taobaocdn.com/imgextra/http://img01.taobaocdn.com/imgextra/i1/762004848/T2bZsWXdhXXXXXXXXX_!!762004848.jpg",
            "http://img.taobaocdn.com/imgextra/http://img04.taobaocdn.com/imgextra/i4/762004848/T2aP.TXehaXXXXXXXX_!!762004848.jpg",
            "http://img01.taobaocdn.com/imgextra/i1/1658448401/TB1c65CGXXXXXb8XXXXXXXXXXXX_!!1658448401-0-tstar.jpg",
            "http://img04.taobaocdn.com/imgextra/i4/1658448401/T195UKFJhXXXXXXXXX_!!1658448401-0-tstar.jpg",
            "http://img02.taobaocdn.com/imgextra/i2/321560743/T1VHWwFJRaXXXXXXXX_!!321560743-0-tstar.jpg",
            "http://img04.taobaocdn.com/imgextra/i4/10743034091143492/T15jWgFpJcXXXXXXXX_!!321560743-0-tstar.jpg_620x10000.jpg",
            "http://img03.taobaocdn.com/imgextra/i3/10743023573695259/T1w6SNXvXfXXXXXXXX_!!321560743-0-tstar.jpg",
            "http://img02.taobaocdn.com/imgextra/i2/1062991248/TB1IGlMHVXXXXcLXXXXXXXXXXXX_!!1062991248-0-tstar.jpg_620x10000.jpg",
            "http://img01.taobaocdn.com/imgextra/i1/1062991248/TB1kkIPHFXXXXcjXFXXXXXXXXXX_!!1062991248-0-tstar.jpg_620x10000.jpg",
            "http://img03.taobaocdn.com/imgextra/i3/1062991248/TB14Q.iHFXXXXaFXVXXXXXXXXXX_!!1062991248-0-tstar.jpg_620x10000.jpg",
            "http://img03.taobaocdn.com/imgextra/i3/1062991248/TB1d6.jHFXXXXXJXVXXXXXXXXXX_!!1062991248-0-tstar.jpg_620x10000.jpg",
            "http://img04.taobaocdn.com/imgextra/i4/1062991248/TB1lZNKHVXXXXXRXpXXXXXXXXXX_!!1062991248-0-tstar.jpg_620x10000.jpg",
            "http://img03.taobaocdn.com/imgextra/i3/1062991248/TB1_82LHpXXXXXFXpXXXXXXXXXX_!!1062991248-0-tstar.jpg_620x10000.jpg",
            "http://img03.taobaocdn.com/imgextra/i3/1062991248/TB1mEPaHpXXXXa6XVXXXXXXXXXX_!!1062991248-0-tstar.jpg_620x10000.jpg",
            "http://img04.taobaocdn.com/imgextra/i4/1062991248/TB1_bAPHpXXXXbaXFXXXXXXXXXX_!!1062991248-0-tstar.jpg_620x10000.jpg",
            "http://img01.taobaocdn.com/imgextra/i1/100352393/TB1nWrGHpXXXXaMXXXXXXXXXXXX_!!100352393-0-tstar.jpg",
            "http://img03.taobaocdn.com/imgextra/i3/100352393/TB1TTHOGVXXXXb1XpXXXXXXXXXX_!!100352393-0-tstar.jpg",
            "http://img04.taobaocdn.com/imgextra/i4/100352393/TB1Mj0yGXXXXXaNaXXXXXXXXXXX_!!100352393-0-tstar.jpg_620x10000.jpg",
            "http://img04.taobaocdn.com/imgextra/i4/100352393/TB1ZUlDGXXXXXaTaXXXXXXXXXXX_!!100352393-0-tstar.jpg_620x10000.jpg",
            "http://img02.taobaocdn.com/imgextra/i2/388656953/TB1IEypHVXXXXcIXXXXXXXXXXXX_!!388656953-0-tstar.jpg_620x10000.jpg",
            "http://img04.taobaocdn.com/imgextra/i4/388656953/TB1xd5OHFXXXXauXVXXXXXXXXXX_!!388656953-0-tstar.jpg",
            "http://img02.taobaocdn.com/imgextra/i2/921620593/TB1Jq7NGXXXXXaOXVXXXXXXXXXX_!!921620593-0-tstar.jpg_620x10000.jpg",
            "http://img03.taobaocdn.com/imgextra/i3/921620593/TB1_kV7GpXXXXcmXXXXXXXXXXXX_!!921620593-0-tstar.jpg",
            "http://img03.taobaocdn.com/imgextra/i3/921620593/TB1SY47GpXXXXclXXXXXXXXXXXX_!!921620593-0-tstar.jpg",
            "http://img04.taobaocdn.com/imgextra/i4/1769268675/TB1Uh8kGFXXXXXTXXXXXXXXXXXX_!!1769268675-0-tstar.jpg_620x10000.jpg",
            "http://img01.taobaocdn.com/imgextra/i1/1769268675/TB1AdFfGFXXXXchXXXXXXXXXXXX_!!1769268675-0-tstar.jpg_620x10000.jpg",
            "http://img02.taobaocdn.com/imgextra/i2/1769268675/TB1YKQHGpXXXXbuXVXXXXXXXXXX_!!1769268675-0-tstar.jpg_620x10000.jpg",
            "http://img01.taobaocdn.com/imgextra/i1/1769268675/TB1EcpkGFXXXXaGXXXXXXXXXXXX_!!1769268675-0-tstar.jpg_620x10000.jpg",
            "http://img04.taobaocdn.com/imgextra/i4/1769268675/TB1KKdXGFXXXXaqXpXXXXXXXXXX_!!1769268675-0-tstar.jpg_620x10000.jpg",
            "http://img01.taobaocdn.com/imgextra/i1/1769268675/TB1uINeGFXXXXctXXXXXXXXXXXX_!!1769268675-0-tstar.jpg_620x10000.jpg",
            "http://img04.taobaocdn.com/imgextra/i4/1769268675/TB17tsJGpXXXXawXVXXXXXXXXXX_!!1769268675-0-tstar.jpg_620x10000.jpg",

        }
        //    fmt.Println(img_list[0], lang[0])
        //
        //    for k, _ := range lang {
        //        insert(lang[k], img_list[k])
        //    }
    */


    http.HandleFunc("/", HandleData)
    http.ListenAndServe(":10000", nil)
}