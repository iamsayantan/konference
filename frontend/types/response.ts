export interface SuccessResponse<T> {
  status_code: number,
  message: string,
  data? :T
}

export interface ErrorResponse {
  status_code: number,
  message: string,
  errors: any
}
