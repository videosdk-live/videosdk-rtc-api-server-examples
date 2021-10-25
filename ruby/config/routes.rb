
Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
  get '/' ,to: "main#index"
  get 'get-token' , to:"main#generateToken"
  post 'create-meeting',to:"main#createMeeting"
  post 'validate-meeting/:meetingId',to:"main#validateMeeting"
end
