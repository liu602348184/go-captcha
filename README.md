<div align="center">
<img width="120" style="padding-top: 50px; margin: 0;" src="http://47.104.180.148/go-captcha/gocaptcha_logo.svg"/>
<h1 style="margin: 0; padding: 0">Go Captcha</h1>
<p>Behavior Security Captcha</p>
<a href="https://goreportcard.com/report/github.com/liu602348184/go-captcha"><img src="https://goreportcard.com/badge/github.com/liu602348184/go-captcha"/></a>
<a href="https://godoc.org/github.com/liu602348184/go-captcha"><img src="https://godoc.org/github.com/liu602348184/go-captcha?status.svg"/></a>
<a href="https://github.com/liu602348184/go-captcha/releases"><img src="https://img.shields.io/github/v/release/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/liu602348184/go-captcha/blob/master/LICENSE"><img src="https://img.shields.io/github/license/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/liu602348184/go-captcha"><img src="https://img.shields.io/github/stars/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/liu602348184/go-captcha"><img src="https://img.shields.io/github/last-commit/wenlng/go-captcha.svg"/></a>
</div>

<br/>

> English | [中文](README_zh.md)

<p><a href="https://github.com/liu602348184/go-captcha">Go Captcha</a> is a behavior security CAPTCHA, which implements the generation of random verification text and the verification of click position information.</p>

<p> ⭐️ If it helps you, please give a star.</p>

- Github：[https://github.com/liu602348184/go-captcha](https://github.com/liu602348184/go-captcha)
- Go Example：[https://github.com/liu602348184/go-captcha-example](https://github.com/liu602348184/go-captcha-example)
- Vue Example：[https://github.com/liu602348184/go-captcha-vue](https://github.com/liu602348184/go-captcha-vue)
- React Example：[https://github.com/liu602348184/go-captcha-react](https://github.com/liu602348184/go-captcha-react)
- Online Demo：[http://47.104.180.148:8081/go_captcha_demo](http://47.104.180.148:8081/go_captcha_demo)

<br/>

<div align="center"> 
    <img src="http://47.104.180.148/go-captcha/go-captcha.jpg?v=9" alt="Reward Support">
    <br/> 
    <img src="http://47.104.180.148/go-captcha/go-captcha-02.png?v=7" alt="Reward Support">
    <br/>
</div>

## Dependency Library
```shell
$ go get -u github.com/golang/freetype
$ go get -u golang.org/x/crypto
$ go get -u golang.org/x/image
```

## Install Captcha Module
```shell
$ go get -u github.com/liu602348184/go-captcha/captcha
```

## Import Captcha Module
```go
package main

import "github.com/liu602348184/go-captcha/captcha"

func main(){
   // ...
}
```

## Quick Use
```go
package main
import (
    "fmt"
    "github.com/liu602348184/go-captcha/captcha"
)

func main(){
    // Captcha Single Instances
    capt := captcha.GetCaptcha()
    
    // Generate Captcha
    dots, b64, tb64, key, err := capt.Generate()
    if err != nil {
        panic(err)
        return
    }
    
    // Main image base64 code
    fmt.Println(len(b64))
    
    // Thumb image base64 code
    fmt.Println(len(tb64))
    
    // Only key
    fmt.Println(key)
    
    // Dot data For verification
    fmt.Println(dots)
}

```

## Captcha Instances
- New Instances or Get Single Instances
```go
package main
import (
    "fmt"
    "github.com/liu602348184/go-captcha/captcha"
)

func main(){
	// Captcha Instances
    // capt := captcha.NewCaptcha() 
    
    // Captcha Single Instances
    capt := captcha.GetCaptcha()

    // ====================================================
    fmt.Println(capt)

}
```

## Set Configuration
After version v1.2.3, the default size of large drawing is 300×240px, the default size of the small drawing is 150×40px.


### Set Chars
Some fonts are attached by default. If other Chinese strings are set, you may need to import a font file.
```go
package main
import (
    "fmt"
    "github.com/liu602348184/go-captcha/captcha"
)

func main(){
    capt := captcha.GetCaptcha()
    
    // ====================================================
    // Method: SetRangChars (chars []string) error;
    // Desc: Set random char of captcha
    // ====================================================
    // Letter
    //chars := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    //_ = capt.SetRangChars(strings.Split(chars, ""))
    
    // Two Letter
    //chars := []string{"HE","CA","WO","NE","HT","IE","PG","GI","CH","CO","DA"}
    //_ = capt.SetRangChars(chars)

    // Chinese Char
    chars := []string{"你","好","呀","这","是","点","击","验","证","码","哟"}
    _ = capt.SetRangChars(chars)

    // ====================================================
    fmt.Println(capt)
}
```

### Set Font File Configuration
You can copy the resource files under the "https://github.com/liu602348184/go-captcha-example/tree/main/resources" to the directory of your project.
```go
package main
import (
    "fmt"
    "os"
    "github.com/liu602348184/go-captcha/captcha"
)

func main(){
    capt := captcha.GetCaptcha()
    
    path, _ := os.Getwd()    
    // ====================================================
    // Method: SetFont(fonts []string);
    // Desc: Set random font of captcha text
    // ====================================================
    capt.SetFont([]string{
        path + "/__example/resources/fonts/fzshengsksjw_cu.ttf",
    })
}
```


### Set Big Image Configuration
Tip: Some images are attached by default. 
```go
package main
import (
    "fmt"
    "golang.org/x/image/font"
    "os"
    "github.com/liu602348184/go-captcha/captcha"
)

func main(){
    capt := captcha.GetCaptcha()
    
    path, _ := os.Getwd()    
    // ====================================================
    // Method: SetBackground(color []string);
    // Desc: Set random image of background
    // ====================================================
    capt.SetBackground([]string{
        path + "/__example/resources/images/1.jpg",
        path + "/__example/resources/images/2.jpg",
    })

    // ====================================================
    // Method: SetImageSize(size Size);
    // Desc: Set size of captcha
    // ====================================================
    capt.SetImageSize(captcha.Size{300, 300})

    // ====================================================
    // Method: SetImageQuality(val int);
    // Desc: Set quality of captcha, The compression level ranges from 1 to 5. QualityCompressNone has no compression. The default is the lowest compression level
    // ====================================================
    capt.SetImageQuality(captcha.QualityCompressNone)

    // ====================================================
    // Method: SetFontHinting(val font.Hinting);
    // Desc: Set Hinting of font (HintingNone,HintingVertical,HintingFull)
    // ====================================================
    capt.SetFontHinting(font.HintingFull)

    // ====================================================
    // Method: SetTextRangLen(val captcha.RangeVal);
    // Desc: Set random length of font
    // ====================================================
    capt.SetTextRangLen(captcha.RangeVal{6, 7})

    // ====================================================
    // Method: SetRangFontSize(val captcha.RangeVal);
    // Desc: Set random size of font
    // ====================================================
    capt.SetRangFontSize(captcha.RangeVal{32, 42})

    // ====================================================
    // Method: SetTextRangFontColors(colors []string);
    // Desc: Set random hex color of font
    // ====================================================
    capt.SetTextRangFontColors([]string{
        "#1d3f84",
        "#3a6a1e",
    })
 
    // ====================================================
    // Method: SetImageFontAlpha(val float64);
    // Desc:Set alpha of font
    // ====================================================
    capt.SetImageFontAlpha(0.5)

    // ====================================================
    // Method: SetTextShadow(val bool);
    // Desc:Set shadow of font
    // ====================================================
    capt.SetTextShadow(true)

    // ====================================================
    // Method: SetTextShadowColor(val string);
    // Desc:Set shadow color of font
    // ====================================================
    capt.SetTextShadowColor("#101010")

    // ====================================================
    // Method: SetTextShadowPoint(val captcha.Point);
    // Desc:Set shadow point of font
    // ====================================================
    capt.SetTextShadowPoint(captcha.Point{1, 1})

    // ====================================================
    // Method: SetTextRangAnglePos(pos []captcha.RangeVal);
    // Desc:Set angle of font
    // ====================================================
    capt.SetTextRangAnglePos([]captcha.RangeVal{
        {1, 15},
        {15, 30},
        {30, 45},
        {315, 330},
        {330, 345},
        {345, 359},
    })

    // ====================================================
    // Method: SetImageFontDistort(val int);
    // Desc:Set distort of font
    // ====================================================
    capt.SetImageFontDistort(captcha.DistortLevel2)
}
```

#### Set Thumb Image Configuration
Tip: Some images are attached by default. 
```go
package main
import (
    "fmt"
    "os"
    "github.com/liu602348184/go-captcha/captcha"
)

func main(){
    capt := captcha.GetCaptcha()
    
    path, _ := os.Getwd()    
    
    // ====================================================
    // Method: SetRangCheckTextLen(val captcha.RangeVal);
    // Desc:Set random length of font
    // ====================================================
    capt.SetRangCheckTextLen(captcha.RangeVal{2, 4})

    // ====================================================
    // Method: SetRangCheckFontSize(val captcha.RangeVal);
    // Desc:Set random size of font
    // ====================================================
    capt.SetRangCheckFontSize(captcha.RangeVal{24, 30})
 
    // ====================================================
    // Method: SetThumbTextRangFontColors(colors []string);
    // Desc: Set random hex color of font
    // ====================================================
    capt.SetThumbTextRangFontColors([]string{
        "#1d3f84",
        "#3a6a1e",
    })

    // ====================================================
    // Method: SetThumbBgColors(colors []string);
    // Desc: Sets the random hex color of the background
    // ====================================================
    capt.SetThumbBgColors([]string{
        "#1d3f84",
        "#3a6a1e",
    })

    // ====================================================
    // Method: SetThumbBackground(colors []string);
    // Desc:Set random image of background
    // ====================================================
    capt.SetThumbBackground([]string{
        path + "/__example/resources/images/r1.jpg",
        path + "/__example/resources/images/r2.jpg",
    })

    // ====================================================
    // Method: SetThumbBgDistort(val int);
    // Desc:Set the distort of background
    // ====================================================
    capt.SetThumbBgDistort(captcha.DistortLevel2)

    // ====================================================
    // Method: SetThumbFontDistort(val int);
    // Desc:Set the distort of font
    // ====================================================
    capt.SetThumbFontDistort(captcha.DistortLevel2)

    // ====================================================
    // Method: SetThumbBgCirclesNum(val int);
    // Desc:Sets the number of dots
    // ====================================================
    capt.SetThumbBgCirclesNum(20)

    // ====================================================
    // Method: SetThumbBgSlimLineNum(val int);
    // Desc:Set number of lines
    // ====================================================
    capt.SetThumbBgSlimLineNum(3)
  
}
```

## Other
```go
package main
import (
    "fmt"
    "os"
    "github.com/liu602348184/go-captcha/captcha"
)

func main(){
    path, _ := os.Getwd()    
    // ====================================================
    // Method: ClearAssetCacheWithPath(paths []string) bool;
    // Desc: Clear Cache
    // ====================================================
    captcha.ClearAssetCacheWithPaths([]string{
    	path + "/__example/resources/images/1.jpg",
    	path + "/__example/resources/fonts/fonts.ttf",
    })     

    // ====================================================
    // Method: captcha.CheckPointDist(sx, sy, dx, dy, width, height int64) bool;
    // Desc: Check point position
    // ====================================================
    captcha.CheckPointDist(0, 30, 0, 30, 30, 30)    

    // ====================================================
    // Method: captcha.CheckPointDistWithPadding(sx, sy, dx, dy, width, height, padding int64) bool;
    // Desc: Check point position
    // ====================================================
    captcha.CheckPointDistWithPadding(0, 30, 0, 30, 30, 30, 5) 
}
```

## Generate Captcha Data
```go
package main
import (
    "fmt"
    "os"
    "github.com/liu602348184/go-captcha/captcha"
)

func main(){
    capt := captcha.GetCaptcha()
    
    // generate ...
    // ====================================================
    // dots:  Character position information
    //  - {"0":{"Index":0,"Dx":198,"Dy":77,"Size":41,"Width":54,"Height":41,"Text":"SH","Angle":6,"Color":"#885500"} ...}
    // imageBase64:  Verify the clicked image
    // thumbImageBase64: Thumb displayed
    // key: Only Key
    // ====================================================
    dots, imageBase64, thumbImageBase64, key, err := capt.Generate()
    if err != nil {
        panic(err)
        return
    }
    fmt.Println(len(imageBase64))
    fmt.Println(len(thumbImageBase64))
    fmt.Println(key)
    fmt.Println(dots)
}
```

## Api Params
```
// Example: Get captcha data
    Respose Data = {
        "code": 0,
        "image_base64": "...",
        "thumb_base64": "...",
        "captcha_key": "...",
    }     

// Example: Post check data
    Request Data = {
        dots: "x1,y1,x2,y2,...."
        key: "......"
    }
```
<br/>

> Buy the author coffee: [http://witkeycode.com/sponsor](http://witkeycode.com/sponsor)

<br/>

## LICENSE
    MIT
