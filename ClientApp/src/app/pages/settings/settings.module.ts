import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';

import { SettingsComponent } from './settings.component';
import { SettingsRoutingModule } from './settings-routing.module';
import { UserProfileFormComponent } from './components/user-profile-form/user-profile-form.component';
import { UserDeleteFormComponent } from './components/user-delete-form/user-delete-form.component';
import { UserDeleteConfirmComponent } from './components/user-delete-confirm/user-delete-confirm.component';


@NgModule({
  declarations: [
    SettingsComponent,
    UserProfileFormComponent,
    UserDeleteFormComponent,
    UserDeleteConfirmComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    SettingsRoutingModule,
    MatButtonModule,
    MatCardModule,
    MatDialogModule,
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
  ],
})
export class SettingsModule { }
