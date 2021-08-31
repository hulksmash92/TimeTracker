import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { User } from 'src/app/models/user';
import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';
import { WindowService } from 'src/app/services/window/window.service';

@Component({
  selector: 'user-menu',
  templateUrl: './user-menu.component.html',
  styleUrls: ['./user-menu.component.scss']
})
export class UserMenuComponent implements OnInit {
  
  /**
   * Currently logged in user
   */
  get user(): User {
    return this.authService.user;
  }

  constructor(
    private readonly userService: UserService, 
    private readonly authService: AuthService,
    private readonly windowService: WindowService,
    private readonly router: Router,
  ) { }

  ngOnInit(): void {
    if (this.router.url !== 'auth') {
      this.userService.get().subscribe((res: any) => {
        this.authService.user = res;
      });
    }
  }

  /**
   * Handles the click event for the sign in with github button
   */
  handleGitHubLogin(): void {
    this.authService.gitHubUrl()
      .subscribe((loginUrl: string) => {
        if (!!loginUrl) {
          this.windowService.goExternal(loginUrl);
        }
      });
  }

}
