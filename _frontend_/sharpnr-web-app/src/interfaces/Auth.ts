/**
 * Login request payload
 *
 * Represents the data sent to the API when a user attempts to log in.
 *
 * @interface
 * @name LoginPayload
 * @kind interface
 * @exports
 */
export interface LoginPayload {
  /** User's email address used for login */
  email: string;

  /** User's password used for login */
  password: string;
}

/**
 * Login response payload
 *
 * Represents the structure of the response returned by the API after a successful login.
 *
 * @interface
 * @name LoginResponse
 * @kind interface
 * @exports
 */
export interface LoginResponse {
  /** Authentication tokens returned by the server */
  token: {
    /** Token used to refresh the access token */
    refresh_token: string;

    /** Token used to authenticate API requests */
    access_token: string;
  };
}
