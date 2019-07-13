## Directory
<!-- TOC -->

- [Directory](#directory)
- [bulrush-utils](#bulrush-utils)
    - [common](#common)
        - [Some](#some)
    - [array](#array)
        - [Append](#append)
    - [func](#func)
        - [Until](#until)
        - [Chain](#chain)
- [MIT License](#mit-license)

<!-- /TOC -->
## bulrush-utils

### common

#### Some
```go
iden.Routes.ObtainTokenRoute = utils.Some(iden.Routes.ObtainTokenRoute, "/obtainToken").(string)
iden.Routes.RevokeTokenRoute = utils.Some(iden.Routes.RevokeTokenRoute, "/revokeToken").(string)
iden.Routes.RefleshTokenRoute = utils.Some(iden.Routes.RefleshTokenRoute, "/refleshToken").(string)
iden.Routes.IdenTokenRoute = utils.Some(iden.Routes.IdenTokenRoute, "/idenToken").(string)
```

### array

#### Append
```go
r.POST(routePrefixs.Create(name), utils.Append(func(c *gin.Context) {
    handler := func(c *gin.Context) {
        create(name, c, ai.gorm, opts)
    }
    h1 := createHooks(ai.gorm, handler)
    h1.Pre(h.pre)
    h1.Post(h.post)
    h1.Auth(h.auth)
    h1.R(c)
}, handlers).([]gin.HandlerFunc)...)
```
### func

#### Until
```go
func(c *gin.Context) {
    token := utils.Until(
        c.Query(tokenKey),
        c.PostForm(tokenKey),
        c.Request.Header.Get(tokenKey),
        func() interface{} {
            value, _ := c.Cookie(tokenKey)
            return value
        },
    ).(string)
    iden.setToken(c, &Token{AccessToken: token})
    c.Next()
}
```

#### Chain
```go
func(c *gin.Context) {
    if _, err := utils.Chain(
        func(ret interface{}) (interface{}, error) {
            return iden.Auth(c)
        },
        func(ret interface{}) (interface{}, error) {
            return iden.ObtainToken(ret)
        },
        func(ret interface{}) (interface{}, error) {
            token := ret.(*Token)
            c.SetCookie(tokenKey, token.AccessToken, 60*60*24, "/", "", false, true)
            c.JSON(http.StatusOK, map[string]interface{}{
                "AccessToken":  token.AccessToken,
                "RefreshToken": token.RefreshToken,
                "ExpiresIn":    token.ExpiresIn,
                "CreatedAt":    token.CreatedAt,
                "UpdatedAt":    token.UpdatedAt,
            })
            return nil, nil
        },
    ); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": err.Error(),
        })
        return
    }
}
```

## MIT License

Copyright (c) 2018-2020 Double

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.