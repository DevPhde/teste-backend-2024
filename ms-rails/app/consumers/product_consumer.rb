class ProductConsumer < Karafka::BaseConsumer
  def consume
    # Inclua manualmente um logger
    @logger ||= Karafka.logger

    messages.each do |message|
      @logger.info "Received message: #{message.payload.with_indifferent_access}"

      process_message(message)
    end
  end

  private

  def process_message(message)
    operation = message.key.to_sym

    product_params = message.payload.with_indifferent_access

    product_params[:is_api] = false

    product = Product.new(product_params.except(:is_api))

    request = nil
    Services::Api::V1::Products::Upsert.new(product_params, request, operation).execute
  end

end
