import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router';

import { Subscription } from 'rxjs';
import { AuthService } from './services/auth/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit, OnDestroy {
  private routeQuerySub = new Subscription();
  details: any;
  authed: boolean;

  constructor(route: ActivatedRoute, private authService: AuthService) {
    this.routeQuerySub = route.queryParamMap.subscribe({
      next: (queryParams: ParamMap) => {
        if (queryParams.has('code')) {
          this.getAccessToken(queryParams.get('code'));
        }
      }
    });
  }

  ngOnInit(): void {
    this.authService.getUser()
      .subscribe((res: any) => {
        this.details = res;
        this.authed = !!res;
      });
  }

  ngOnDestroy(): void {
    if (!this.routeQuerySub.closed) {
      this.routeQuerySub.unsubscribe();
    }
  }

  handleGitHubLogin(): void {
    this.authService.gitHubUrl()
      .subscribe((loginUrl: string) => {
        if (loginUrl) {
          window.location.href = loginUrl;
        }
      });
  }

  private getAccessToken(sessionCode: string): void {
    this.authService.loginGitHub(sessionCode)
      .subscribe((res: any) => {
        this.details = res;
      });
  }

}
