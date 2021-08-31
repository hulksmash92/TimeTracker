import { Component, OnDestroy } from '@angular/core';
import { ActivatedRoute, ParamMap, Router } from '@angular/router';

import { Subscription } from 'rxjs';

import { AuthService } from 'src/app/services/auth/auth.service';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss']
})
export class AuthComponent implements OnDestroy {
  /**
   * Holds the subscription for checking the query params
   * in the  current route if any exist.
   * Such as the session code returned in the URL by 
   * GitHub after auth redirect.
   */
  readonly routeQuerySub = new Subscription();

  constructor(
    route: ActivatedRoute, 
    private readonly authService: AuthService, 
    private readonly router: Router
  ) {
    this.routeQuerySub = route.queryParamMap.subscribe({
      next: (queryParams: ParamMap) => {
        if (queryParams.has('code')) {
          this.getAccessToken(queryParams.get('code'));
        }
      }
    });
  }

  ngOnDestroy(): void {
    if (!this.routeQuerySub.closed) {
      // Teardown the route query param subscription
      this.routeQuerySub.unsubscribe();
    }
  }

  /**
   * Sets the users details in the app and the access token using the session 
   * code returned by GitHub after successful authentication with OAuth
   * 
   * @param sessionCode the session code returned by GitHub 
   */
  getAccessToken(sessionCode: string): void {
    this.authService.loginGitHub(sessionCode)
      .subscribe((res: any) => {
        this.authService.user = res;
        if (!!res) {
          this.router.navigate(['/time']);
        }
      });
  }

}
