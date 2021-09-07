import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

import { Observable } from 'rxjs';

import { User } from 'src/app/models/user';
import { WindowService } from 'src/app/services/window/window.service';

const headers = new HttpHeaders({
  'Content-Type': 'application/json'
});

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  readonly API_URL = '/api/auth';
  readonly GH_API_URL = '/api/github';
  user: User;

  constructor(
    private readonly http: HttpClient, 
    private readonly router: Router,
    private readonly windowService: WindowService
  ) { }

  /**
   * Checks if the user is authenticated
   */
  async isAuthenticated(): Promise<boolean> {
    try {
      const res: any = await this.http.get<any>(`${this.API_URL}/isAuthenticated`).toPromise();
      const success: boolean = res?.success;

      if (!success) {
        this.resetUser();
      }
      return success;
    } catch (e) {
      return false;
    }
  }

  /**
   * Signs the user out from the application
   */
  async signOut(): Promise<void> {
    try {
      const res: any = await this.http.get<any>(`${this.API_URL}/signOut`).toPromise();
      const success: boolean = res?.success;

      if (success) {
        this.resetUser();
      }
    } catch (error) {
      console.error(error);
    }
  }

  /**
   * Redirects the user to the GitHub Login page
   */
  async gitHubLogin(): Promise<void> {
    try {
      const res: any = await this.http.get<any>(`${this.GH_API_URL}/url`).toPromise();
      const ghUrl: string = res?.data;
      this.windowService.goExternal(ghUrl);
    } catch (error) {
      console.error(error);
    }
  }

  /**
   * Passes the session code to the api to get the users GH access token 
   * and create a log in session for the user
   * 
   * @param sessionCode session code returned by the github oauth server
   * 
   * @returns details of the user
   */
  loginGitHub(sessionCode: string): Observable<any> {
    const url = `${this.GH_API_URL}/login`;
    const body = JSON.stringify({sessionCode});
    return this.http.post(url, body, { headers });
  }

  /**
   * Resets the user value to null and redirects the app user to the home page
   */
  resetUser(): void {
    this.user = null;
    this.router.navigate(['']);
  }

}
