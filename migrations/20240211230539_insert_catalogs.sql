-- +goose Up
-- +goose StatementBegin

insert
	into product_size
		( code, description )
	values
		( 'XXS', 'Extra Extra Small' ),
		( 'XS', 'Extra Small' ),
		( 'S', 'Small' ),
		( 'M', 'Medium' ),
		( 'L', 'Large' ),
		( 'XL', 'Extra Large' ),
		( 'XXL', 'Extra Extra Large' );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

delete from product_size where code in ('XXS', 'XS', 'S', 'M', 'L', 'XL', 'XXL');

-- +goose StatementEnd
