import {User} from "~/types/user";
import {SuccessResponse} from "~/types/response";

export interface Room {
  id: number
  created_by: User
}

export interface RoomDetailsResponse {
  room_details: Room
}

export interface RoomCreateResponse extends SuccessResponse<RoomDetailsResponse>{

}
