module Services
  module Api
    module V1
      module Products
        class PaginatedIndex
          attr_accessor :params, :request

          def initialize(params, request)
            @params = params
            @request = request
          end

          def execute
            paginate_products
          end

          private

          def paginate_products
            page = params[:page] || 1
            limit = params[:itemsLimit] || 10

            offset = (page.to_i - 1) * limit.to_i
            products = Product.offset(offset).limit(limit)

            total_count = Product.count
            total_pages = (total_count.to_f / limit.to_i).ceil
            has_next_page = (page.to_i * limit.to_i) < total_count

            {
              data: products.present? ? products : [],
              has_next_page: has_next_page,
              total_pages: total_pages
            }
          end
        end
      end
    end
  end
end
