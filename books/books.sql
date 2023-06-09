create table books (
	title varchar(40) not null,
	author varchar(40) not null,
	pages_num int not null,
	review text not null
)

insert into books (
	title,
	author,
	pages_num,
	review
)
values (
	'Title',
	'Author',
	896,
	'NICE BOOK!'
)