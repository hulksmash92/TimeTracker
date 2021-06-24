import { Component, OnDestroy } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router';

import { Subscription } from 'rxjs';

import { AuthService } from 'src/app/services/auth/auth.service';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss']
})
export class AuthComponent implements OnDestroy {
  private routeQuerySub = new Subscription();

  constructor(route: ActivatedRoute, private readonly authService: AuthService) {
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
      this.routeQuerySub.unsubscribe();
    }
  }

  private getAccessToken(sessionCode: string): void {
    this.authService.loginGitHub(sessionCode)
      .subscribe((res: any) => {
        this.authService.user = res;
      });
  }

}
