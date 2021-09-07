import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';

const headers = new HttpHeaders({
  'Content-Type': 'application/json'
});

@Injectable({
  providedIn: 'root'
})
export class UserService {
  /**
   * Base URL of the user API endpoints
   */
  readonly API_URL = '/api/user';

  constructor(private http: HttpClient) { }

  /**
   * Gets the currently logged in user's details
   */
  get(): Observable<any> {
    return this.http.get(this.API_URL);
  }

  /**
   * Updates the a currently logged in user's details with those in the newValues object
   * 
   * @param newValues Fields to update
   * 
   * @returns true if updated else false
   */
  update(newValues: any): Observable<any> {
    return this.http.patch<any>(this.API_URL, JSON.stringify(newValues), { headers });
  }

  /**
   * Deletes the current user from the application and signs them out
   * 
   * @returns any object containing true if deleted successfully 
   */
  delete(): Observable<any> {
    return this.http.delete<any>(this.API_URL);
  }

}
