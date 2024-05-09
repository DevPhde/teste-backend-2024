require 'test_helper'

class ProductsDetailsTest < ActiveSupport::TestCase
  def setup
    # Create a valid product before each test
    @product = Product.create(name: "Test Product", price: 10.99, brand: "Test Brand", description: "Test Description", stock: 100)
  end

  test "should raise ActiveRecord::RecordNotFound if product is not found" do
    assert_raises(ActiveRecord::RecordNotFound) do
      Services::Api::V1::Products::Details.new({ id: -1 }, nil).execute
    end
  end

  test "should raise NoMethodError if product ID is 0" do
    assert_raises(NoMethodError) do
      Services::Api::V1::Products::Details.new({ id: 0 }, nil).execute
    rescue Exception => e
      assert_equal "Método indefinido `erros' para uma instância de String", e.mensagem
    end
  end

  test "should return product if found" do
    details_service = Services::Api::V1::Products::Details.new({ id: @product.id }, nil)
    assert_equal @product, details_service.execute
  end

  teardown do
    # Remove the created product after each test
    @product.destroy
  end
end
