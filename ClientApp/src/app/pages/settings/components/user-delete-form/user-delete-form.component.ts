import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';

import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';

@Component({
  selector: 'user-delete-form',
  templateUrl: './user-delete-form.component.html',
  styleUrls: ['./user-delete-form.component.scss']
})
export class UserDeleteFormComponent {

  constructor(
    private readonly authService: AuthService,
    private readonly userService: UserService,
    private readonly matDialog: MatDialog,
  ) { }

  /**
   * 
   */
  handleBtnClick(): void {
    
  }

  /**
   * Calls the method to delete the current user from the application
   */
  deleteUser(): void {
    this.userService.delete()
      .subscribe((res: any) => {
        if (res?.success) {
          this.authService.resetUser();
        }
      });
  }

}
