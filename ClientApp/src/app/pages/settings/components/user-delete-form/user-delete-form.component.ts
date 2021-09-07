import { Component, OnDestroy } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';

import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';
import { UserDeleteConfirmComponent } from '../user-delete-confirm/user-delete-confirm.component';

@Component({
  selector: 'user-delete-form',
  templateUrl: './user-delete-form.component.html',
  styleUrls: ['./user-delete-form.component.scss']
})
export class UserDeleteFormComponent implements OnDestroy {

  constructor(
    private readonly authService: AuthService,
    private readonly userService: UserService,
    public matDialog: MatDialog,
  ) { }

  ngOnDestroy(): void {
    this.matDialog.closeAll();
  }

  /**
   * Opens a confirmation dialog for the user to confirm deletion
   */
  handleBtnClick(): void {
    const dialogRef = this.matDialog.open(UserDeleteConfirmComponent);

    // dialogRef.afterClosed().subscribe(result => {
    //   console.log('returned value:', result);
    // });
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
