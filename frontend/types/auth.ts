import {SuccessResponse} from "~/types/response";
import {User} from "~/types/user";

export interface LoginRequest {
  email: string,
  password: string
}

export interface RegistrationRequest {
  first_name: string,
  last_name: string,
  email: string,
  password: string
}

export interface LoginResponseData {
  user: User,
  access_token: string
}

export interface RegistrationResponse extends SuccessResponse <null> {

}

export interface LoginResponse extends SuccessResponse<LoginResponseData>{

}
