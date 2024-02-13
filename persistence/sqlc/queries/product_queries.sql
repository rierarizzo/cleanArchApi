-- name: GetProducts :many
select * from product_view;

-- name: GetProductById :one
select *
	from product_view
	where product_id = $1;

-- name: GetProductCategoryById :one
select *
	from product_category
	where
		id = $1;

-- name: GetProductSubcategoryById :one
select *
	from product_subcategory
	where
		id = $1;

-- name: GetProductSubcategoryByCategoryId :many
select *
	from product_subcategory
	where
		parent_category_id = $1;

-- name: GetProductSizeByCode :one
select *
	from product_size
	where
		code = $1;

-- name: GetProductColorById :one
select *
	from product_color
	where
		id = $1;

-- name: GetProductSourceById :one
select *
	from product_source
	where
		id = $1;

-- name: CreateProductCategory :one
insert
	into product_category
		( name, description )
	values
		( $1, $2 )
	returning id;

-- name: CreateProductSubcategory :one
insert
	into product_subcategory
		( parent_category_id, name, description )
	values
		( $1, $2, $3 )
	returning id;

-- name: CreateProductSource :one
insert
	into product_source
		( name )
	values
		( $1 )
	returning id;

-- name: CreateProduct :one
insert
	into product
		( subcategory_id, name, description, price, cost, quantity, size_code, color_id, brand, sku, upc, image_url,
		  source_id, source_url, is_offered, offer_percent )
	values
		( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16 )
	returning id;