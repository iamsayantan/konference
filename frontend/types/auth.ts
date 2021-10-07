import {SuccessResponse} from "~/types/response";
import {User} from "~/types/user";

export interface LoginRequest {
  email: string,
  password: string
}

export interface RegistrationRequest {
  firstName: string,
  lastName: string,
  email: string,
  password: string
}

export interface LoginResponseData {
  user: User,
  accessToken: string
}

export interface RegistrationResponse extends SuccessResponse <null> {

}

export interface LoginResponse extends SuccessResponse<LoginResponseData>{

}
