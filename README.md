Hello gbatis

# Quick Star

## Config databaase information

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<configuration>
<dbs default="dev">
    <db id="dev">
        <property name="driver" value="mysql"/>
        <property name="dataSource" value="root:Changeme_123@tcp(127.0.0.1:3306)/orange?charset=utf8"/>
        <property name="maxOpenConns" value = "20"/>
        <property name="maxIdleConns" value = "10"/>
    </db>
    <db id="product">
        <property name="driver" value="mysql"/>
        <property name="dataSource" value="root:Changeme_123@tcp(127.0.0.1:3306)/orange?charset=utf8"/>
        <property name="maxOpenConns" value = "20"/>
        <property name="maxIdleConns" value = "10"/>
    </db>
</dbs>    
</configuration>
```

```go
err := gbatis.OpenDB("etc/conf/gbatis.xml")
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println("Open database success.")
```

## Named sql

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<mapper namespace="mapper.mall">
	<select id="selectMall" resultType="bean.Mall">
	  SELECT id as ID, name as Name, avatar as Avatar, createdAt as CreatedAt, 
			updatedAt as UpdatedAt, story as Story FROM t_mall WHERE id = ?
	</select>

	<select id="selectAllMalls" resultType="bean.Mall">
	  SELECT id as ID, name as Name, avatar as Avatar, story as Story FROM t_mall
	</select>

	<select id="selectShop" resultType="bean.Shop">
		SELECT s.id as id, name, avatar, story FROM t_shop s, t_brand b 
		where s.brandId = b.id and s.id = ?
	</select>

	<insert id="insertMall">
	  insert into t_mall(id, name, avatar, createdAt, updatedAt, story) 
			values (?, ?, ?, unix_timestamp(now()), unix_timestamp(now()), ?)
	</insert>

	<update id="updateMall">
		update t_mall set name = ?, updatedAt = unix_timestamp(now()) where id = ?
	</update>

	<delete id="deleteMall">
	  delete from t_mall where id = ?
	</delete>

	<delete id="deleteShop">
	  delete from t_shop where id = ?
	</delete>
</mapper>
```

## Code

```go
import (
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mallbook/gbatis"
	_ "github.com/mallbook/gbatis/bean"
)

func TestSelectRow(t *testing.T) {
	s, err := gbatis.OpenSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()

	row, err := s.SelectRow("mapper.mall.selectMall", "4")
	if err != nil {
		fmt.Println(err)
		return
	}

	var id string
	var name string
	var avatar string
	var story string
	err = row.Scan(&id, &name, &avatar, &story)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("id = %s, name=%s, avatar=%s, story=%s\n", id, name, avatar, story)
}

func TestSelectRow_2(t *testing.T) {

	row, err := gbatis.SelectRow("mapper.mall.selectMall", "4")
	if err != nil {
		fmt.Println(err)
		return
	}

	var id string
	var name string
	var avatar string
	var story string
	err = row.Scan(&id, &name, &avatar, &story)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("id = %s, name=%s, avatar=%s, story=%s\n", id, name, avatar, story)
}

```


