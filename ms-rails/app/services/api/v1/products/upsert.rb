module Services
  module Api
    module V1
      module Products
        class Upsert
          attr_accessor :params, :request, :operation

          def initialize(params, request, operation)
            @params = params
            @request = request
            @operation = operation
          end

          def execute
            ActiveRecord::Base.transaction do
              raise ActiveRecord::RecordNotFound, "Product Not Found" if product.blank?

              product.id        ||= params[:id]           if params[:id].present?
              product.name        = params[:name]         if params[:name].present?
              product.brand       = params[:brand]        if params[:brand].present?
              product.price       = params[:price]        if params[:price].present?
              product.description = params[:description]  if params[:description].present?
              product.stock       = params[:stock]        if params[:stock].present?

              product.save!

              produce_to_kafka(product) if !!params[:is_api]
            end

            product
          rescue StandardError => e
            raise ActiveRecord::Rollback, e.message
          end

          private
          def product
            @product ||= Product.find_by(id: params[:id]).presence || Product.new
          end

          require 'json'

          def produce_to_kafka(product)
            product_copy = operation == 'create' ? product.dup : product

            product_copy.id = 0 if operation == 'create'

            Karafka.producer.produce_sync(topic: 'rails-to-go', key: operation == 'create' ? 'create' : 'update', payload: product_copy.to_json)
          end

        end
      end
    end
  end
end
