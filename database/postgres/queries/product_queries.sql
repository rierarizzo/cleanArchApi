-- name: GetProjects :many
select * from product;

-- name: GetProductCategoryById :one
select * from product_category where id = $1;

-- name: GetProductSizeByCode :one
select * from product_size where code = $1;

-- name: GetProductColorById :one
select * from product_color where id = $1;

-- name: GetProductSourceById :one
select * from product_source where id = $1;
