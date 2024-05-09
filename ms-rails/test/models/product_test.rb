require 'test_helper'

class ProductTest < ActiveSupport::TestCase
  test "product should be valid with all required attributes" do
    product = Product.new(
      name: "Example Product",
      price: 10.99,
      brand: "Example Brand",
      description: "Example Description",
      stock: 100
    )
    assert product.valid?, "Product should be valid with all required attributes"
  end

  test "name should be at least 4 characters long" do
    product = Product.new(name: "A"*3, price: 10.99, brand: "Example Brand", description: "Example Description", stock: 100)
    assert_not product.valid?, "Name should be at least 4 characters long"
  end

  test "price should be present and in the correct format" do
    product = Product.new(name: "Example Product", brand: "Example Brand", description: "Example Description", stock: 100)
    assert_not product.valid?, "Price should be present"

    invalid_prices = ["not a number", 0, 10.999]  # You can adjust the decimal precision test as needed
    invalid_prices.each do |price|
      product.price = price
      assert_not product.valid?, "Price should be invalid for #{price}"
    end
  end

  test "price should be greater than 0 and less than 1000000" do
    product = Product.new(name: "Example Product", price: 1000000, brand: "Example Brand", description: "Example Description", stock: 100)
    assert_not product.valid?, "Price should be less than 1,000,000"

    product.price = 0
    assert_not product.valid?, "Price should be greater than 0"

    product.price = 1000000
    assert_not product.valid?, "Price should be less than 1,000,000"
  end

  test "should validate presence of other attributes" do
    product = Product.new(name: "Example Product", price: 10.99, stock: 100)
    assert_not product.valid?
    assert_includes product.errors.full_messages, "Brand can't be blank"
    assert_includes product.errors.full_messages, "Description can't be blank"
  end

  test "stock should be present and a non-negative integer" do
    product = Product.new(name: "Example Product", price: 10.99, brand: "Example Brand", description: "Example Description")
    assert_not product.valid?, "Stock should be present"

    product.stock = -1
    assert_not product.valid?, "Stock should be non-negative"

    product.stock = 10.5
    assert_not product.valid?, "Stock should be an integer"

    product.stock = 10
    assert product.valid?, "Stock should be a valid integer"
  end
end
