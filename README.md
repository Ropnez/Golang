# Golang

### `paket kavramı`

Go için özel bir dosya ve paket adı bulunur: `main.go` ve `package main`. Eğer
bir go projesi altında `main.go` dosyası var ise mutlaka o dosyanın paket
adı da `main` olur. Bu uygulamaya giriş yeridir. Eğer orada bir go uygulaması
var ise o uygulamanın giriş kapısı `main.go` olur.

Dizin yapısına göre bir projede birden fazla `main.go` olabilir. Dosyanın
durduğu yere göre, **sadece bir tane** main paketi olur!

### `go run .`

`.` shell ortamında **current working directory** yani o an içinde buluduğumuz
dizin anlamındadır, bulunduğumuz yerdeki modül yapısına göre uygun `main.go`
dosyasını bulur ve çalıştırır. Genelde içinde `go.mod` olan bir proje dizininde
olmamız gerekir aksi halde;

### `go build`

Bulunduğumuz dizin içindeki `go.mod` yapısına göre ilgili `main.go` dosyasını
bulur ve derler. Ürettiği binary’i yine `.` yani **current working directory**
altına atar; sonra elle biz çalıştırırız:

Toplamda **25** tane anahtar kelime bulunur:

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var

```

Bunlara ek olarak;

**Öntanımlı sabitler ailesi**

- `true`
- `false`
- `iota`
- `nil`

**Gömülü gelen fonksiyonlar**

- `append`
- `cap`
- `close`
- `complex`
- `copy`
- `delete`,
- `imag`
- `len`
- `make`
- `new`
- `panic`
- `print`
- `println`
- `real`
- `recover`

**Diğer**

- `string`
- `error`

## Operatörler

`+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
     /    <<    /=    <<=    ++    =     :=    ,    ;
     %    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~`

## Built-in Veri Tipleri

Standart kütüphane bir kısım hazır veri tipi ile birlikte geliyor, kabaca;

- **Strings** : Metinsel tipler
- **Booleans** : `true` / `false` mantıksal veri tipleri
- **Numerics** : `int` / `float` ve `complex` familyası
- **Composite (Unnamed) Types** (Bileşik İsimsiz Tipler) : Array, Slice, Struct, Map

## Kod Stili

2 tür yorum (comment) yazma stili var;

1. **Line Comment** : `// bu bir yorum satırı` şeklinde
2. **General Comment** : `/* bu bir yorum satırı */` şeklinde

## Değişkenler

İçinde değer saklayan, depolayan ve değiştirilebilir olan şeylerdir. Veriye
ulaşmak için kullanılan referanstır. Go, değişkenin tipine göre içinde yazan
değeri algılar. 2 tür tanımla şekli vardır;

1. **Long Variable Declaration** : Uzun değişken tanımlama. `var` anahtar
   kelimesi ile kullanılır.
2. **Short Variable Declaration** : Kısa değişken tanımlama. `:=` ile kullanılır.

```go
var x = 5       // x’in değeri 5 ve tipi: dynamic type int
var y int = 5   // y’nin değeri 5 ve tipi: static type int
int := 5
```

Bazı tiplerin **zero-value** değerleri;

`var a int                // 0
var b float32            // 0
var d string             // ""
var e bool               // false
var f byte               // 0
var j func()             // nil`

### `fmt.Print`

```go
fmt.Print("hello", "world", 1, 2, []string{"foo"})
// helloworld1 2 [foo]
```

string dışındakilerin arasına bir boşluk karakteri (space) ekler çıktıya.
Satır sonuna otomatik olarak yeni satır `\n` (new line) karakteri **eklemez**

### `fmt.Println`

Neredeyse `Print` ile aynı fakat bu kez parametreler arasına otomatik boşluk
koyar ve satır sonuna otomatik `\n` ekler:

```go
fmt.Println("hello", "world", 1, 2, []string{"foo"})
// hello world 1 2 [foo]
```

### `fmt.Printf`

Çok sık kullanacağımız, metin formatlama, yer düzenleme (string interpolation)
gibi işlerde bize kolaylıklar sağlar. `%<VERB>` yani `%` işareti ve fiil alır,
satır sonuna otomatik olarak yeni satır `\n` (new line) karakteri **eklemez**;

```go
fmt.Printf("merhaba %s\n", "dünya")
// merhaba dünya
```

`%s` bir fiil’dir ve **the uninterpreted bytes of the string or slice** yani
yorumlanmamış string ya da slice (liste kesiti) byte’larını temsil eder.

```go
package main

import "fmt"

func main() {
	a := 5

	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)

	a := 8 // hata!! no new variables on left side of :=
    // Tekli atamalarda 2 kere tekrar edilemiyor

	fmt.Printf("%v\n", a)
}
```

Eğer kısa şekilde değişkeni tanımlamışsak artık değeri değiştirmek
istediğimizde `:=` yerine `=` kullanmamız gerekiyor. Çünkü artık `a` tipi
belli olan bir değişken:

Kısa değişken tanımlamanın bazı kısıtları var;

- Sadece **fonksiyon** içinde çalışıyor
- Tekli atamalarda **2 kere** tekrar edilemiyor
- Çoklu atamalarda tekrar oluyor ama her seferinde değeri değişiyor
- Kapsama göre tekrar olabiliyor

```go
package main

import "fmt"

a := 5 // error
       // non-declaration statement outside function body

func main() {
	fmt.Printf("a: %v\n", a)
}
```

Değişken isimlendirmesinde dikkat edeceğimiz kurallar;

1. Mutlaka harf ile başlamalı
2. İçinde harf, sayı ve `_` (*underscore*) olabilir ama olmasa iyi olur
3. `camelCase`, `BumpyCaps`, `mixedCase` şeklinde tanımlama yapılabilir
4. Anlaşılır olmalıdır

Örneğin veritabanından gelen kayıtların sayısı için bir değişken tanımlamak
gerekese; "NUMBER OF RECORDS" ya da "LENGTH OF RECORDS" ya da "RECORDS LENGTH"
kafamızda olsa;

```
var lengthOfRecords int // ya da
var recordsLength       // ya da
var numRecs             // çok tercih edilmememli
var recordsAmount       //
```

gibi varyasyonlar olabilir. Eğer imkan varsa tek bir kelime ile ifade etmek en
iyi yöntemdir. Tüm bu kurallar tüm **identifier**’lar için geçerlidir.

Nerede bir değişken kullanımı görürseniz mutlaka o değişkenin değerini yani
**Value of**’unu kullandığınızı **unutmayın**!