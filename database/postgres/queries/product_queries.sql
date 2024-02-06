-- name: GetProjects :many
select pro.id                       as "product_id",
       pro_ca.id                    as "category_id",
       pro_ca.name                  as "category_name",
       pro_ca.description           as "category_description",
       pro_subca.id                 as "subcategory_id",
       pro_subca.parent_category_id as "subcategory_parent_category_id",
       pro_subca.name               as "subcategory_name",
       pro_subca.description        as "subcategory_description",
       pro.name                     as "product_name",
       pro.description              as "product_description",
       pro.price                    as "product_price",
       pro.cost                     as "product_cost",
       pro.quantity                 as "product_quantity",
       pro_si.code                  as "size_code",
       pro_si.description           as "size_description",
       pro_co.id                    as "color_id",
       pro_co.name                  as "color_name",
       pro_co.hex                   as "color_hex",
       pro.brand                    as "product_brand",
       pro.sku                      as "product_sku",
       pro.upc                      as "product_upc",
       pro.image_url                as "product_image_url",
       pro_so.id                    as "source_id",
       pro_so.name                  as "source_name",
       pro.source_url               as "product_source_url",
       pro.is_offered               as "product_is_offered",
       pro.offer_percent            as "product_offer_percent",
       pro.is_active                as "product_is_active",
       pro.created_at               as "product_created_at",
       pro.updated_at               as "product_updated_at"
from product as pro
         inner join product_subcategory as pro_subca on pro.subcategory_id = pro_subca.id
         inner join product_category as pro_ca on pro_subca.parent_category_id = pro_ca.id
         inner join product_size as pro_si on pro.size_code = pro_si.code
         inner join product_color as pro_co on pro.color_id = pro_co.id
         inner join product_source as pro_so on pro.source_id = pro_so.id;

-- name: GetProductCategoryById :one
select *
from product_category
where id = $1;

-- name: GetProductSubcategoryById :one
select *
from product_subcategory
where id = $1;

-- name: GetProductSubcategoryByCategoryId :many
select *
from product_subcategory
where parent_category_id = $1;

-- name: GetProductSizeByCode :one
select *
from product_size
where code = $1;

-- name: GetProductColorById :one
select *
from product_color
where id = $1;

-- name: GetProductSourceById :one
select *
from product_source
where id = $1;
