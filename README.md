Hello gbatis

## Example

```go
import (
    "github.com/mallbook/gbatis"
)

func main() {
    
    s, err := gbatis.OpenSession()
    if err != nil {
        
    }
    defer s.Close()
    stu, err := s.SelectOne("xxxx", a, b, c)
    stu, err := s.SelectList("xxxx", a, b, c)
}
```
