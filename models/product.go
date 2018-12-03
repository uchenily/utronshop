package models

//Product represent an product struct
type Product struct {
	ID     int    `scheme:"id"`
	Number int    `scheme:"number"` // 产品编号，四位数， 唯一
	Name   string `scheme:"name"`
	Type   string `scheme:"type"`
	Count  int    `scheme:"count"`
	Price  int    `scheme:"price"`
	Brief  string `scheme:"brief"`  // 产品描述,简述, 数据库中以文件路径形式存放
	Detail string `scheme:"detail"` // 产品描述,详述, 数据库中以文件路径形式存放
	Href   string `scheme:"href"`   // 链接路径 eg: /detail/3
	URL    string `scheme:"url"`    // 图片路径 eg: /static/img/item-3.jpg
}

/**
insert into products(name, type, count, price, brief, detail， href, url) values("macbook", "pro", 10, 3333, "/data/product/2/brief", "/data/product/2/detail", "/detail/3", "/static/img/item-3.jpg");


MariaDB [utronshop]> insert into products(name, type, count, price, brief, detail) values("macbook", "pro", 10, 3333, "/data/product/2/brief", "/data/product/2/detail");


insert into products(id, name, type, count, price, brief, detail, href, url) values(12, "dell", "i7 gtxg65m usb3.0 1080p", 20, 6666, "/data/product/8/brief", "/data/product/8/detail", "/detail/8", "/static/img/item-8.jpg");

*/
