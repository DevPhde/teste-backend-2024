Rails.application.routes.draw do
  namespace :api do
    namespace :v1 do
      resources :products, only:[:index, :show, :create, :update] do
        collection do
          get 'index', to: 'products#paginatedIndex'
        end
      end
    end
  end
end
