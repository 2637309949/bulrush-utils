## bulrush-utils

### func

#### Until
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