import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { User } from 'src/app/models/user';

import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';

@Component({
  selector: 'user-profile-form',
  templateUrl: './user-profile-form.component.html',
  styleUrls: ['./user-profile-form.component.scss']
})
export class UserProfileFormComponent implements OnInit {
  /**
   * Details of the form and all the controls
   */
  formGroup: FormGroup = new FormGroup({
    name: new FormControl(null),
    email: new FormControl(null, [Validators.email])
  });

  /**
   * Details of the user logged into the app
   * Used when resetting the form control values
   */
  get currentUser(): User {
    return this.authService.user;
  }

  constructor(private readonly userService: UserService, private authService: AuthService) { }

  ngOnInit(): void {
    this.reset();
  }

  /**
   * Resets the form
   */
  reset(): void {
    this.formGroup.reset();
    this.formGroup.patchValue({
      name: this.currentUser.name,
      email: this.currentUser.email,
    });
    this.formGroup.markAsUntouched();
  }

  /**
   * Submits the form details to be updated in the user profile
   */
  submit(): void {
    if (this.formGroup.valid) {
      const newValues = this.formGroup.value;

      this.userService.update(newValues)
        .subscribe((res: any) => {
          if (res.success) {
            this.authService.user.name = newValues.name;
            this.authService.user.email = newValues.email;
            this.reset();
          }
        });
    }
  }

}
