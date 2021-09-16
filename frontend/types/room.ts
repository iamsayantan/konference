import {User} from "~/types/user";

export interface Room {
  id: number
  invite_code: string
  created_by: User
}
