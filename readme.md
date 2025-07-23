1. Install golang
2. Export its path in zshrc / bashrc
2. Check `go -version`
4. Install sqlite3 : brew install sqlite, confirm after installing `sqlite3 --version`
5. Testing if sqlite 3 is installed properly or not
   1. one more way to test if sqlite is installed correctly by creating a db  : `sqlite3 ecommerce.db` : This opens a DB file ecommerce.db in the current folder. Exit with `.exit`
6. mkdir ecommerce-app
7. cd ecommerce-app
8. go mod init ecommerce-app
9. Installing Add GORM + SQLite driver
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```
GORM: ORM itself\
SQLite Driver: for connecting to SQLite from Go

#### Redefining project structure : 

```
ecommerce-app/
├── cmd/                # Entrypoints (main.go)
├── internal/
│   ├── db/             # DB connection logic
│   ├── models/         # GORM models
│   ├── handlers/       # HTTP handlers
│   ├── services/       # Business logic
│   ├── middleware/     # Auth, logging
│   ├── auth/           # JWT + OAuth
├── config/             # App config
├── go.mod
├── go.sum
```

Lets take one example : 
```go
type User struct {
    ID    uint   `gorm:"primaryKey"`
    Email string `gorm:"uniqueIndex"`
    Name  string
}
```
8. `primaryKey` → This column uniquely identifies each row in a table. it is not null, should be unique, Typically auto-incremented.\
__Note__ : GORM automatically uses `ID` as the primary key if you don’t specify

9. `uniqueIndex` → Enforces that the value in this column must be unique across all rows.

__Other useful GORM tags__ : 

| Tag             | Description                                                |
| :-------------- | :--------------------------------------------------------- |
| `not null`      | Cannot be null                                             |
| `default`       | Sets a default value if none provided                      |
| `unique`        | Marks the column as unique (shortcut for unique index)     |
| `index`         | Adds a non-unique index for faster lookups                 |
| `check`         | Adds a check constraint (example: age > 18)                |
| `size:255`      | Set column size (especially for strings in some databases) |
| `autoIncrement` | Auto-increments integer values                             |
| `embedded`      | Embed a struct as fields in the table                      |
| `foreignKey`    | Define foreign key relationships                           |

10. Meaning of `db.AutoMigrate(&User{})` : Look at the User struct, Check if a corresponding table exists in the connected database, if not, create it. else, compare the struct’s fields with existing columns. If new fields are added, try to add those to the table.

11. When New column is added, and we do AutoMigrate, the old rows wont be having values for this column and hence it will have value as `NULL`.\
If you want to set a value for existing records, you need to run a manual SQL query via GORM:
```go
func Migrate() {
    err := DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    // Example: add default true to existing nulls
    DB.Exec("UPDATE users SET is_active = true WHERE is_active IS NULL")
}
```
_usually this code will be part of migration files like migration.go which would be called from other file like `db.Migrate()`_

12. Assigning _multiple tags_ to go struct
```go
type Product struct {
    ID    uint   `gorm:"primaryKey;autoIncrement"`
    Code  string `gorm:"uniqueIndex;size:100"`
    Price int    `gorm:"not null;index"`
}
```
| Tag             | Effect                                       |
| :-------------- | :------------------------------------------- |
| `primaryKey`    | Primary key                                  |
| `autoIncrement` | Auto-incrementing integer (if integer type)  |
| `uniqueIndex`   | Creates a unique index                       |
| `size:100`      | String length constraint                     |
| `not null`      | Disallow null values                         |
| `index`         | Adds a normal non-unique index on the column |

__Note__ : Tags are separated by semicolons ; inside the backtick string.

***

Main models of our project : 
* User
* Product
* Category
* Cart
* Wishlist
* Order
* OrderItem
* Payment
* Address
* IssueReport
* OrderTrackingEvent

__Model relations__ : 
| Relation     | Example                            |
| :----------- | :--------------------------------- |
| One-to-Many  | User has many Orders               |
| One-to-Many  | Product belongs to one Category    |
| Many-to-Many | User has many Wishlist Products    |
| One-to-One   | Payment for an Order               |
| One-to-Many  | Order has many OrderTrackingEvents |

__Table relationships__ : 
| Type                    | Meaning                                                  | How It's Implemented            |
| :---------------------- | :------------------------------------------------------- | :------------------------------ |
| **One-to-One (1:1)**    | One record in table A relates to one record in table B   | Foreign key in either table     |
| **One-to-Many (1\:N)**  | One record in table A relates to many records in table B | Foreign key in "many" side      |
| **Many-to-One (N:1)**   | Many records in table A relate to one record in table B  | Same as One-to-Many             |
| **Many-to-Many (N\:N)** | Many records in table A relate to many in table B        | **Join Table (junction table)** |


1. __one to one__(1 : 1) mapping : 

```go
type User struct {
  ID      uint
  Profile Profile
}

// we can also say this is "belongs to" relationship.
type Profile struct {
  ID     uint
  UserID uint `gorm:"uniqueIndex"`
  Bio    string
}
```

2. __One to many__ : 
```go
type User struct {
  ID     uint
  Orders []Order
}

type Order struct {
  ID     uint
  UserID uint
}
```

3. __Many to Many relationship__ : 
```go
type User struct {
  ID        uint
  Wishlists []Product `gorm:"many2many:user_wishlists"`
}

type Product struct {
  ID uint
}
```
This will automatically create a table user_wishlists with columns:
* `user_id`
* `product_id`

*** 

__Summary of types created in our project at this point__ : 

__Note__ : GORM will automatically create a `user_wishlists` table for the many2many relation.


| Relation     | Example                    | Struct Field                                                 |
| :----------- | :------------------------- | :----------------------------------------------------------- |
| One-to-One   | Order → Payment            | `Payment Payment`                                            |
| One-to-Many  | User → Orders              | `Orders []Order`                                             |
| Many-to-Many | User ↔ Products (Wishlist) | `Wishlists []Product \`gorm:"many2many\:user\_wishlists"\`\` |
| Belongs To   | OrderItem → Order, Product | `OrderID uint`                                               |


__Why some fields will be of pointer type__ ? 

When to use pointer:

1. If the relation is optional (can be nil).
2. If you want to defer loading (e.g., via Preload) or avoid loading related objects unless needed.
3. To avoid zero-value struct initialization overhead.

example : 
```go
type Payment struct {
  OrderID uint
  Order   *Order // optional relation
}
```

Definition of foreign key in terms of SQL : 
```sql
user_id INTEGER,
FOREIGN KEY (user_id) REFERENCES users(id)
```
Means every user_id value in orders must exist in users.id.

__Note__ : 
If we dont want to follow convention of declaring field name as `<OtherStructFieldName>ID`, and want to declare our own field name which acts as foreign key : 
```go
type Product struct {
  CategoryRef uint `gorm:"column:category_ref"`
  Category    Category `gorm:"foreignKey:CategoryRef"`
}
```

__When not to prefer eager loading__ :
When we have a table/struct which has nested data like below, then just image loading 100 order would have to load rest of the unnecessary data as well. 

* Order
  * User
  * Products
    * Categories
    * Vendors
      * Addresses
      * Documents
        * ScanFiles

***

### The N+1 Query problem

The N+1 problem happens when your application runs one query to fetch a list of records, and then for each of those records, it runs an additional query to fetch related data.

Bad code : 
```go
var products []Product
db.Find(&products) // 1 query to fetch products

for _, p := range products {
    var category Category
    db.First(&category, p.CategoryID) // N queries (1 per product)
}
```

Total Queries: 1 + 1000 = 1001 queries\
Problem: Hugely inefficient — every product makes an additional DB round-trip.

How to __Fix N+1 Problem__? (Eager loading)

Use Preload to fetch everything in 2 queries:
```go
var products []Product
db.Preload("Category").Find(&products)
```

__Thumb rule__ : If you don’t prefetch, you’re running into the N+1 trap.

__Common scenarios__ : 
1. Fetching articles with their authors
2. Orders with their items
3. Comments with their users
4. In nested templates rendering: listing 100 users, then for each user loading 10 posts


